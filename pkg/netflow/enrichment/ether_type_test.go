package enrichment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapEtherType(t *testing.T) {
	assert.Equal(t, "", MapEtherType(0))
	assert.Equal(t, "IPv4", MapEtherType(0x0800))
	assert.Equal(t, "IPv6", MapEtherType(0x86DD))
	assert.Equal(t, "ARP", MapEtherType(0x0806))
	assert.Equal(t, "1000", MapEtherType(1000)) // invalid protocol number
}
