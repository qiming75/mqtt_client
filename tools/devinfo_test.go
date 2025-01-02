package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMacAddrs(t *testing.T) {
	macAddrs := getMacAddrs3()
	fmt.Print(macAddrs)
	assert.NotEmpty(t, macAddrs)
}
