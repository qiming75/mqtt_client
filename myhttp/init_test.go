package myhttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMqttAccount(t *testing.T) {
	// ip, userId, pwd, err := getMqttAccount("17fb36d68d751")
	// assert.Nil(t, err)
	// assert.Equal(t, "124.192.140.240", ip)
	// assert.Equal(t, "e0c1ff56", userId)
	// assert.Equal(t, "ebd7937b238a29edf", pwd)

	ip, userId, pwd, err := getMqttAccount("17fb36d68d751")
	assert.Nil(t, err)
	assert.Equal(t, "124.192.140.240", ip)
	assert.Equal(t, "e0c1ff56", userId)
	assert.Equal(t, "ebd7937b238a29edf", pwd)
}
