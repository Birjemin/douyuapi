package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("900150983cd24fb0d6963f7d28e17f72", Md5ByByte([]byte("abc")))
	ast.Equal("900150983cd24fb0d6963f7d28e17f72", GetMD5String("abc"))
	ast.Equal(12, len(Hex(6)))
}
