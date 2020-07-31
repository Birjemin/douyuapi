package douyuapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSign
func TestSign(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(
		"5638f3e24319fc150be904c93182cd45",
		GetSign("VFCS4u6e6kpev7uO!t", "/api/thirdPart/token", map[string]string{
			"time": "1521791007",
			"aid":  "ttlive",
		}),
	)
}
