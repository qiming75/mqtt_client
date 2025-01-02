package tools

import (
	"fmt"
	"sort"
	"strings"

	"net"

	cpu "github.com/shirou/gopsutil/v3/cpu"
	disk "github.com/shirou/gopsutil/v3/disk"
	mem "github.com/shirou/gopsutil/v3/mem"
	net2 "github.com/shirou/gopsutil/v3/net"
)

type DEVInfo struct {
	MacAddress string
	CPUInfo    any
	MEMInfo    any
	DISKInfo   any
	NETInfo    any
}

var keyMAC string = ""

func GetKeyMAC() string {
	if keyMAC == "" {
		keyMAC = getMacAddrs3()
	}
	return keyMAC
}

func GetDEVInfo() (info DEVInfo) {
	info = DEVInfo{
		MacAddress: getMacAddrs3(),
		CPUInfo:    getCPUInfo(),
		MEMInfo:    getMEMInfo(),
		DISKInfo:   getDiskInfo(),
		NETInfo:    getNetInfo(),
	}
	return
}

func getMacAddrs3() (macAddrs string) {
	macAddrs = getMacAddrs2()
	macAddrs = strings.ToUpper(macAddrs)
	return
}

func getMacAddrs2() (macAddrs string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return
	}

	networkInterface := GetConf().NetworkInterface

	macs := make([]string, 0)
	for _, netInterface := range netInterfaces {
		mac := netInterface.HardwareAddr.String()
		if len(mac) == 0 {
			continue
		}
		fmt.Printf("discover network interface: %s, mac: %s\n", netInterface.Name, mac)
		if networkInterface != "" && networkInterface == netInterface.Name {
			fmt.Printf("select network interface: %s, mac: %s\n", netInterface.Name, mac)
			return mac
		}
		if len(mac) > 3 {
			macs = append(macs, mac)
		}
	}
	sort.Strings(macs)
	if len(macs) > 0 {
		macAddrs = macs[0]
	}
	fmt.Printf("select default network interface mac: %s\n", macAddrs)
	// for _, netInterface := range netInterfaces {
	// 	if netInterface.Flags&net.FlagUp != 0 && netInterface.Flags&net.FlagRunning != 0 && netInterface.Flags&net.FlagBroadcast != 0 {
	// 		macAddrs = netInterface.HardwareAddr.String()
	// 		return
	// 	}
	// }

	return
}

// // lint
// func getMacAddrs() (macAddrs []string) {
// 	netInterfaces, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Printf("fail to get net interfaces: %v", err)
// 		return
// 	}

// 	fmt.Println("22222222222222")
// 	for _, netInterface := range netInterfaces {
// 		if netInterface.Flags&net.FlagUp != 0 && netInterface.Flags&net.FlagRunning != 0 && netInterface.Flags&net.FlagBroadcast != 0 {
// 			fmt.Println("555", netInterface.HardwareAddr)
// 		}
// 		fmt.Println("111111111111111")
// 		fmt.Println(netInterface.Flags, netInterface.MTU, netInterface.Name, netInterface.Index, netInterface.HardwareAddr)
// 		macAddr := netInterface.HardwareAddr.String()
// 		if len(macAddr) == 0 {
// 			continue
// 		}

// 		macAddrs = append(macAddrs, macAddr)
// 	}
// 	return macAddrs
// }

func getMEMInfo() (info any) {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	info = v
	return
}

func getCPUInfo() any {
	type CPUInfo struct {
		BaseInfo    any
		UsedPercent any
	}

	cpuInfo := CPUInfo{}

	baseInfo, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	usedPercentInfo, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Println(err)
	}

	cpuInfo.BaseInfo = baseInfo
	cpuInfo.UsedPercent = usedPercentInfo

	return cpuInfo
}

func getDiskInfo() (info any) {
	v, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	info = v
	return
}

func getNetInfo() any {
	type NETInfo struct {
		LocalIPList any
		BaseInfo    any
	}

	netInfo := NETInfo{}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ips := make([]string, 0)
	for _, address := range addrs {
		// 检查ipnet结构体是否nil，如果不是nil，则转换为ipNet类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	netInfo.LocalIPList = ips
	counters, _ := net2.IOCounters(false)
	netInfo.BaseInfo = counters

	return netInfo
}
