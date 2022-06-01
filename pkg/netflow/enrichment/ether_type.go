package enrichment

import "strconv"

var etherTypeMap = map[uint32]string{
	0x0800: "IPv4",
	0x0806: "ARP",
	0x86DD: "IPv6",
}

// MapEtherType map Ether Type number to human-readable Ether Type name
// https://www.iana.org/assignments/ieee-802-numbers/ieee-802-numbers.xhtml
func MapEtherType(etherTypeNumber uint32) string {
	if etherTypeNumber == 0 {
		return ""
	}
	protoStr, ok := etherTypeMap[etherTypeNumber]
	if !ok {
		return strconv.Itoa(int(etherTypeNumber))
	}
	return protoStr
}
