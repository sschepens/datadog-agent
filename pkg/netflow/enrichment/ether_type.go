package enrichment

var etherTypeMap = map[uint32]string{
	0x0800: "IPv4",
	0x0806: "ARP",
	0x86DD: "IPv6",
}

// MapEtherType map Ether Type number to human-readable Ether Type name
// https://www.iana.org/assignments/ieee-802-numbers/ieee-802-numbers.xhtml
func MapEtherType(etherTypeNumber uint32) string {
	protoStr, ok := etherTypeMap[etherTypeNumber]
	if !ok {
		return ""
	}
	return protoStr
}
