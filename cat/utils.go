package cat

import (
	"fmt"
	"net"
	"time"
)

var (
	ClientIp = ""
)

func getLocalhostIp() (ip net.IP, err error) {
	if ClientIp != "" {
		ip = net.ParseIP(ClientIp)
		return
	}
	ip = net.IPv4(127, 0, 0, 1)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				return
			}
		}
	}
	return
}

func ip2String(ip net.IP) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[12], ip[13], ip[14], ip[15])
}

func ip2HexString(ip net.IP) string {
	return fmt.Sprintf("%02x%02x%02x%02x", ip[12], ip[13], ip[14], ip[15])
}

func duration2Millis(duration time.Duration) int64 {
	return duration.Nanoseconds() / time.Millisecond.Nanoseconds()
}
