/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : struct.go
 Time    : 2018/9/19 14:16
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanue/go-esport-common"
	"reflect"
	"strconv"
)

type structHelper struct {
}

var Struct *structHelper = new(structHelper)

/*
*@note 结构转[]byte
*@param v 结构体
*@return []byte
 */
func (this *structHelper) ToJsonByte(v interface{}) []byte {
	reply, err := json.Marshal(v)
	if err != nil {
		common.Logs.Debug("Struct2Byte(%v) fail, err=%s", v, err.Error())
		return nil
	}

	return reply
}

/*
*@note 结构转字符串
*@param v 结构体
*@return string
 */
func (this *structHelper) ToJsonString(v interface{}) string {
	reply, err := json.Marshal(v)
	if err != nil {
		common.Logs.Debug("Struct2String(%v) fail, err=%s", v, err.Error())
		return ""
	}

	return string(reply)
}

/**
*@note Map转struct
*@param obj 结构体
*@return map[string]interface{}
 */
func (this *structHelper) MapToStruct(data map[string]interface{}, val interface{}) error {
	var list [][]byte
	for key, value := range data {
		list = append(list, []byte(key), []byte(ToString(value)))
	}

	return this.toStruct(list, reflect.ValueOf(val))
}

/**
*@note 结构转Map
*@param obj 结构体
*@return map[string]interface{}
 */
func (this *structHelper) StructToMap(pStr interface{}) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	if pStr == nil {
		err = errors.New("pStr is nil")
		return
	}

	list := make([]string, 0)
	err = this.containerToString(reflect.ValueOf(pStr), &list)
	if err != nil {
		return
	}
	fmt.Println("list", list)
	for i, _ := range list {
		if i%2 == 1 {
			res[list[i-1]] = list[i]
		}
	}
	return
}

func (this *structHelper) toStruct(data [][]byte, val reflect.Value) error {
	switch v := val; v.Kind() {
	case reflect.Ptr:
		return this.toStruct(data, reflect.Indirect(v))
	case reflect.Interface:
		return this.toStruct(data, v.Elem())
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return errors.New("Invalid map type")
		}
		elemtype := v.Type().Elem()
		for i := 0; i < len(data)/2; i++ {
			mk := reflect.ValueOf(string(data[i*2]))
			mv := reflect.New(elemtype).Elem()
			this.writeTo(data[i*2+1], mv)
			v.SetMapIndex(mk, mv)
		}
	case reflect.Struct:
		for i := 0; i < len(data)/2; i++ {
			name := string(data[i*2])
			field := v.FieldByName(name)
			if !field.IsValid() {
				continue
			}
			this.writeTo(data[i*2+1], field)
		}
	default:
		return errors.New("Invalid container type")
	}
	return nil
}

func (this *structHelper) writeTo(data []byte, val reflect.Value) error {
	s := string(data)
	switch v := val; v.Kind() {
	// if we're writing to an interace value, just set the byte data
	// TODO: should we support writing to a pointer?
	case reflect.Interface:
		v.Set(reflect.ValueOf(data))
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ui, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		v.SetUint(ui)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		v.SetFloat(f)

	case reflect.String:
		v.SetString(s)
	case reflect.Slice:
		typ := v.Type()
		if typ.Elem().Kind() == reflect.Uint || typ.Elem().Kind() == reflect.Uint8 || typ.Elem().Kind() == reflect.Uint16 || typ.Elem().Kind() == reflect.Uint32 || typ.Elem().Kind() == reflect.Uint64 || typ.Elem().Kind() == reflect.Uintptr {
			v.Set(reflect.ValueOf(data))
		}
	}
	return nil
}

func (this *structHelper) containerToString(val reflect.Value, args *[]string) error {
	switch v := val; v.Kind() {
	case reflect.Ptr:
		return this.containerToString(reflect.Indirect(v), args)
	case reflect.Interface:
		return this.containerToString(v.Elem(), args)
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return errors.New("Unsupported type - map key must be a string:")
		}
		for _, k := range v.MapKeys() {
			*args = append(*args, k.String())
			s, err := this.valueToString(v.MapIndex(k))
			if err != nil {
				return err
			}
			*args = append(*args, s)
		}
	case reflect.Struct:
		st := v.Type()
		fmt.Println("st", st, st.NumField())
		for i := 0; i < st.NumField(); i++ {
			ft := st.FieldByIndex([]int{i})
			*args = append(*args, ft.Name)
			fmt.Println("ft", ft)
			s, err := this.valueToString(v.FieldByIndex([]int{i}))
			if err != nil {
				return err
			}
			*args = append(*args, s)
		}
	}
	return nil
}

//pretty much copy the json code from here.

func (this *structHelper) valueToString(v reflect.Value) (string, error) {
	if !v.IsValid() {
		return "null", nil
	}

	switch v.Kind() {
	case reflect.Ptr:
		return this.valueToString(reflect.Indirect(v))
	case reflect.Interface:
		return this.valueToString(v.Elem())
	case reflect.Bool:
		x := v.Bool()
		if x {
			return "true", nil
		} else {
			return "false", nil
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.UnsafePointer:
		return strconv.FormatUint(uint64(v.Pointer()), 10), nil

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64), nil

	case reflect.String:
		return v.String(), nil

		//This is kind of a rough hack to replace the old []byte
		//detection with reflect.Uint8Type, it doesn't catch
		//zero-length byte slices
	case reflect.Slice:
		typ := v.Type()
		if typ.Elem().Kind() == reflect.Uint || typ.Elem().Kind() == reflect.Uint8 || typ.Elem().Kind() == reflect.Uint16 || typ.Elem().Kind() == reflect.Uint32 || typ.Elem().Kind() == reflect.Uint64 || typ.Elem().Kind() == reflect.Uintptr {
			if v.Len() > 0 {
				if v.Index(0).OverflowUint(257) {
					return string(v.Interface().([]byte)), nil
				}
			}
		}
	}

	//fmt.Println("v.Kind()", v.Kind().String())

	return "", errors.New("Unsupported type -- " + v.Kind().String())
}
