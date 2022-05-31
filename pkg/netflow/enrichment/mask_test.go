package enrichment

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestFormatMask(t *testing.T) {
	assert.Equal(t, "192.1.128.64/26", FormatMask(ipV4EtherType, []byte{192, 1, 128, 108}, 26))
	assert.Equal(t, "192.1.128.0/25", FormatMask(ipV4EtherType, []byte{192, 1, 128, 54}, 25))
	assert.Equal(t, "/50", FormatMask(ipV4EtherType, []byte{}, 50))
	assert.Equal(t, "/50", FormatMask(ipV4EtherType, []byte{192, 1, 128, 108}, 50))
	assert.Equal(t, "2001:db8:abcd:12::/112", FormatMask(ipV6EtherType, net.ParseIP("2001:0DB8:ABCD:0012:0000:0000:0000:0010"), 112))
	assert.Equal(t, "/300", FormatMask(ipV6EtherType, net.ParseIP("2001:0DB8:ABCD:0012:0000:0000:0000:0010"), 300))
}
