/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : http.go
 Time    : 2018/9/14 17:38
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/yanue/go-esport-common"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
 * @note 私有地址列表
 * 10.0.0.0/8
 * 100.64.0.0/10
 * 127.0.0.0/8
 * 169.254.0.0/16
 * 172.16.0.0/12
 * 192.168.0.0/16
 */
var privateBlocks []*net.IPNet

func init() {
	// Add each private block
	privateBlocks = make([]*net.IPNet, 6)

	_, block, err := net.ParseCIDR("10.0.0.0/8")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}
	privateBlocks[0] = block

	_, block, err = net.ParseCIDR("100.64.0.0/10")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}
	privateBlocks[1] = block

	_, block, err = net.ParseCIDR("127.0.0.0/8")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}
	privateBlocks[2] = block

	_, block, err = net.ParseCIDR("169.254.0.0/16")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}
	privateBlocks[3] = block

	_, block, err = net.ParseCIDR("172.16.0.0/12")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}
	privateBlocks[4] = block

	_, block, err = net.ParseCIDR("192.168.0.0/16")
	if err != nil {
		panic(fmt.Sprintf("Bad cidr. Got %v", err))
	}

	privateBlocks[5] = block
}

/*
 *@note Http远程调用
 *@param strUrl 请求的地址
 *@remark strUrl需要调用者做好urlencode
 */
func RemoteCall(strUrl string) []byte {
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
 *@note 带超时Http远程调用
 *@param strUrl 请求的地址
 *@param timeout 超时时间：单位毫秒
 *@remark strUrl需要调用者做好urlencode
 */
func RemoteCallWithTimeout(strUrl string, timeout time.Duration) []byte {
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
func RemoteDeleteWithTimeout(strUrl string, timeout time.Duration) []byte {
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
func RemotePost(url string, msg string) []byte {
	//body := bytes.NewBuffer([]byte(msg))
	common.Logs.Debug(url + "\n" + msg)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemotePost_OctStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemotePut_OctStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemotePost_ProtoStream(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemotePost_URLEncode(uri string, values url.Values) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemotePost_Json(uri string, body []byte) []byte {
	common.Logs.Debug(uri)

	var httpClient = &http.Client{
		Timeout: DEF_HTTP_TIMEOUT,
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
func RemoteHead(strUrl string) int {
	resp, err := http.Head(strUrl)
	if err != nil {
		common.Logs.Warn("http.Get: " + err.Error())
		return 404
	}

	defer resp.Body.Close()
	return resp.StatusCode
}

//const smUrl = "http://222.73.117.158:80/msg/HttpBatchSendSM?account=wansha&pswd=Wans2531&mobile=%s&msg=%s"

/*
*@note 带重试的Http远程调用，如果失败会最大重试3次
*@param funName 功能模板名，只是会作为日志打印用
*@param strUrl 请求的地址
*@remark strUrl需要调用者做好urlencode
 */
func RemoteCallEx(funName, strUrl string) bool {
	ok := RemoteCall(strUrl) != nil
	tryCount := 3
	for ; !ok && tryCount > 0; tryCount-- {
		// 失败重试3次
		common.Logs.Warn("%s err %d, url=%s\n", funName, 4-tryCount, strUrl)
		time.Sleep(1 * time.Second)
		ok = RemoteCall(strUrl) != nil
	}

	return ok
}

/*
*@note HttpRespond 包装的http回包，支持gzip压缩
*@param w http.ResponseWriter
*@param r *http.Request
*@param d 要发送的数据
*@return 发送内容长度， 错误代码
 */
func HttpRespond(w http.ResponseWriter, r *http.Request, d []byte) (int, error) {
	var n int
	var e error
	if len(d) > HTTP_GZIP_LEN && r.Header.Get("Accept-Encoding") == "gzip" {
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

/*
 *@note 获取本地启用的网卡IP
 */
func GetActiveIP() []string {
	return getIPv4Addresses(activeInterfaceAddresses(), false)
}

/*
 *@note 获取本地所有的私有地址
 */
func GetPrivateIP() []string {
	return getIPv4Addresses(activeInterfaceAddresses(), true)
}

// 返回启用的网卡IP地址列表
func activeInterfaceAddresses() []net.Addr {
	var upAddrs []net.Addr

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, iface := range ifaces {
		// 已经启用
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// 忽略本地回路地址
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}
		upAddrs = append(upAddrs, addresses...)
	}

	return upAddrs
}

func getIPv4Addresses(addresses []net.Addr, private bool) []string {
	var candidates []string

	for _, rawAddr := range addresses {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if ip.To4() == nil {
			continue
		}
		if private && !isPrivateIP(ip.String()) {
			continue
		}
		candidates = append(candidates, ip.String())
	}
	return candidates
}

/*
 *@note 是否是私有地址
 *ipStr ipv4地址
 */
func isPrivateIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}

/*
*@note 从ip:port这样的字符串取出ip
*@param RemoteAddr ip:port的字符串，比如127.0.0.1:80
*@return ip 127.0.0.1
 */
func GetIP(RemoteAddr string) (ip string) {
	idx := strings.Index(RemoteAddr, ":")
	if -1 != idx {
		ip = RemoteAddr[:idx]
	} else {
		ip = RemoteAddr
	}
	return
}
