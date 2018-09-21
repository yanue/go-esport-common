/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : ip.go
 Time    : 2018/9/21 16:31
 Author  : yanue
 
 - ip相关
 
------------------------------- go ---------------------------------*/

package util

import (
	"fmt"
	"net"
	"strings"
)

type ipHelper struct {
}

var Ip *ipHelper

var privateBlocks []*net.IPNet
/*
 * @note 私有地址列表
 */
var privateIps = []string{
	"10.0.0.0/8",
	"127.0.0.0/8",
	"192.168.0.0/16",
}

func init() {
	privateBlocks = make([]*net.IPNet, len(privateIps))
	// Add each private block
	for i, ipCidr := range privateIps {
		_, block, err := net.ParseCIDR(ipCidr)
		if err != nil {
			panic(fmt.Sprintf("Bad cidr. Got %v", err))
		}
		privateBlocks[i] = block
	}
	Ip = new(ipHelper)
}

/*
 *@note 获取本地启用的网卡IP
 */
func (this *ipHelper) GetActiveIP() []string {
	return this.getIPv4Addresses(this.activeInterfaceAddresses(), false)
}

/*
 *@note 获取本地所有的私有地址
 */
func (this *ipHelper) GetPrivateIP() []string {
	return this.getIPv4Addresses(this.activeInterfaceAddresses(), true)
}

// 返回启用的网卡IP地址列表
func (this *ipHelper) activeInterfaceAddresses() []net.Addr {
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

func (this *ipHelper) getIPv4Addresses(addresses []net.Addr, private bool) []string {
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
		if private && !this.isPrivateIP(ip.String()) {
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
func (this *ipHelper) isPrivateIP(ipStr string) bool {
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
func (this *ipHelper) GetIP(RemoteAddr string) (ip string) {
	idx := strings.Index(RemoteAddr, ":")
	if -1 != idx {
		ip = RemoteAddr[:idx]
	} else {
		ip = RemoteAddr
	}
	return
}
