package douyuapi

import (
	"fmt"
	"github.com/birjemin/douyuapi/utils"
)

// GetSign ...
func GetSign(secret, uri string, params map[string]string) string {
	str := fmt.Sprintf("%s?%s%s", uri, utils.HttpQueryBuild(params), secret)
	return utils.GetMD5String(str)
}
