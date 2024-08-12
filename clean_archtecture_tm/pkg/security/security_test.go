package security

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPasswordUtil(t *testing.T) {
	tests := []string{
		"Abc123",
		"A3!",
		"Abc123!!",
		"Abc123!!!",
		"1",
		"",
	}
	passwordUtil := PasswordUtil{}
	for _, sample := range tests {
		t.Run(sample, func(t *testing.T) {
			hashedPassword, _ := passwordUtil.EncryptPassword(sample)
			err := passwordUtil.ComparePassword(hashedPassword, sample)
			require.NoError(t, err)
			require.NotEqual(t, hashedPassword, sample)
			require.NotNil(t, hashedPassword)
		})
	}

}
