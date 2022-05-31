package enrichment

import (
	"net"
	"strconv"
)

const ipV4EtherType = 0x800
const ipV6EtherType = 0x86dd

func FormatMask(etherType uint32, ipAddr []byte, maskRawValue uint32) string {
	// TODO: check for `.` or `:` for v4 vs v6
	maskSuffix := "/" + strconv.Itoa(int(maskRawValue))
	var maskBits int
	switch etherType {
	case ipV4EtherType:
		maskBits = 32
	case ipV6EtherType:
		maskBits = 128
	default:
		if maskRawValue != 0 {
			return maskSuffix
		}
		return ""
	}

	ip := net.IP(ipAddr)
	if ip == nil {
		return maskSuffix
	}

	mask := net.CIDRMask(int(maskRawValue), maskBits)
	if mask == nil {
		return maskSuffix
	}
	return ip.Mask(mask).String() + maskSuffix
}
