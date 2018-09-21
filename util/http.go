/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : http.go
 Time    : 2018/9/14 17:38
 Author  : yanue
 
 - http相关
 
------------------------------- go ---------------------------------*/

package util

import (
	"bytes"
	"compress/gzip"
	"github.com/yanue/go-esport-common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const httpGzipLen = 1024
const httpTimeout = time.Second * 10

type httpHelper struct {
}

var Http *httpHelper = new(httpHelper)

/*
 *@note Http远程调用
 *@param strUrl 请求的地址
 *@remark strUrl需要调用者做好urlencode
 */
func (this *httpHelper) RemoteCall(strUrl string) []byte {
	//r, err := http.NewRequest("GET", strUrl, nil)
	//if err != nil {
	//	fmt.Println("http.NewRequest: ", err.Error())
	//	return nil
	//}

	//DefaultClient := &http.Client{}
	//resp, err := DefaultClient.Do(r)
	resp, err := http.Get(strUrl)
	//resp, err := http.DefaultClient.Do(r)
	if err != nil {
		common.Logs.Warn("http.DefaultClient.Do: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
*@note 带重试的Http远程调用，如果失败会最大重试3次
*@param funName 功能模板名，只是会作为日志打印用
*@param strUrl 请求的地址
*@remark strUrl需要调用者做好urlencode
 */
func (this *httpHelper) RemoteCallWithTry(funName, strUrl string) bool {
	ok := this.RemoteCall(strUrl) != nil
	tryCount := 3
	for ; !ok && tryCount > 0; tryCount-- {
		// 失败重试3次
		common.Logs.Warn("%s err %d, url=%s\n", funName, 4-tryCount, strUrl)
		time.Sleep(1 * time.Second)
		ok = this.RemoteCall(strUrl) != nil
	}

	return ok
}

/*
 *@note 带超时Http远程调用
 *@param strUrl 请求的地址
 *@param timeout 超时时间：单位毫秒
 *@remark strUrl需要调用者做好urlencode
 */
func (this *httpHelper) RemoteCallWithTimeout(strUrl string, timeout time.Duration) []byte {
	var httpClient = &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		return nil
	}

	req.Close = true

	resp, err := httpClient.Do(req)
	if err != nil {
		common.Logs.Warn("httpClient.Do(req) fail, err=%s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("ioutil.ReadAll: " + err.Error())
		return nil
	}
	return data
}

/*
 *@note 带超时Http远程调用
 *@param strUrl 请求的地址
 *@param timeout 超时时间：单位毫秒
 *@remark strUrl需要调用者做好urlencode
 */
func (this *httpHelper) RemoteDeleteWithTimeout(strUrl string, timeout time.Duration) []byte {
	var httpClient = &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("DELETE", strUrl, nil)
	if err != nil {
		common.Logs.Warn("RemoteDeleteWithTimeout(%s) fail, err = %s", strUrl)
		return nil
	}

	req.Close = true

	resp, err := httpClient.Do(req)
	if err != nil {
		common.Logs.Warn("httpClient.Do(req) fail, err=%s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("ioutil.ReadAll: " + err.Error())
		return nil
	}
	return data
}

/*
 *@note Http post
 *@param strUrl 请求的地址
 */
func (this *httpHelper) RemotePost(url string, msg string) []byte {
	//body := bytes.NewBuffer([]byte(msg))
	common.Logs.Debug(url + "\n" + msg)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(msg))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Close = true

	resp, err := httpClient.Do(req)

	//body := strings.NewReader(msg)
	//resp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		common.Logs.Warn("http.Post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("ioutil.ReadAll: " + err.Error())
		return nil
	}

	common.Logs.Debug(string(data))
	return data
}

/*
 *@note Http post
 *@param uri 请求的地址
 *@param body post请求的body,二进制数据流，可用于传输序列化后的protobuf数据
 */
func (this *httpHelper) RemotePostOctStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Close = true

	resp, err := httpClient.Do(req)
	//resp, err := http.Post(uri, "application/octet-stream", bytes.NewBuffer(body))
	if err != nil {
		common.Logs.Warn("RemotePost_OctStream http.post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("RemotePost_OctStream resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("RemotePost_OctStream ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
 *@note Http put
 *@param uri 请求的地址
 *@param body put请求的body,二进制数据流，可用于传输序列化后的protobuf数据
 */
func (this *httpHelper) RemotePutOctStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Close = true

	resp, err := httpClient.Do(req)
	//resp, err := http.Put(uri, "application/octet-stream", bytes.NewBuffer(body))
	if err != nil {
		common.Logs.Warn("RemotePut_OctStream http.post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("RemotePut_OctStream resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("RemotePut_OctStream ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
 *@note Http post
 *@param uri 请求的地址
 *@param body post请求的body,二进制数据流，可用于传输序列化后的protobuf数据
 */
func (this *httpHelper) RemotePostProtoStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Close = true

	resp, err := httpClient.Do(req)
	//resp, err := http.Post(uri, "application/x-protobuf", bytes.NewBuffer(body))
	if err != nil {
		common.Logs.Warn("RemotePost_OctStream http.post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("RemotePost_OctStream resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("RemotePost_OctStream ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
 *@note Http post
 *@param uri 请求的地址
 *@param body post请求的body，浏览器的原生 form 表单
 */
func (this *httpHelper) RemotePostURLEncode(uri string, values url.Values) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("POST", uri, ioutil.NopCloser(strings.NewReader(values.Encode())))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Close = true

	resp, err := httpClient.Do(req)
	//resp, err := http.Post(uri, "application/x-www-form-urlencoded;charset=utf-8", ioutil.NopCloser(strings.NewReader(body.Encode())))
	if err != nil {
		common.Logs.Warn("RemotePost_URLEncode http.post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("RemotePost_URLEncode resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("RemotePost_URLEncode ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
 *@note Http post
 *@param uri 请求的地址
 *@param body post请求的body,Json数据
 */
func (this *httpHelper) RemotePostJson(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	resp, err := httpClient.Do(req)
	//resp, err := http.Post(uri, "application/json", bytes.NewBuffer(body))
	if err != nil {
		common.Logs.Warn("RemotePost_OctStream http.post: " + err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Logs.Warn("RemotePost_OctStream resp.StatusCode!=http.StatusOK: %d", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		common.Logs.Warn("RemotePost_OctStream ioutil.ReadAll: " + err.Error())
		return nil
	}

	return data
}

/*
 *@note Http head 操作，目前仅返回StatusCode
 *@param strUrl 请求的地址
 */
func (this *httpHelper) RemoteHead(strUrl string) int {
	resp, err := http.Head(strUrl)
	if err != nil {
		common.Logs.Warn("http.Get: " + err.Error())
		return 404
	}

	defer resp.Body.Close()
	return resp.StatusCode
}

/*
*@note HttpRespond 包装的http回包，支持gzip压缩
*@param w http.ResponseWriter
*@param r *http.Request
*@param d 要发送的数据
*@return 发送内容长度， 错误代码
 */
func (this *httpHelper) HttpRespond(w http.ResponseWriter, r *http.Request, d []byte) (int, error) {
	var n int
	var e error
	if len(d) > httpGzipLen && r.Header.Get("Accept-Encoding") == "gzip" {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		n, e = gz.Write(d)
		gz.Flush()
	} else {
		n, e = w.Write(d)
	}

	return n, e
}
