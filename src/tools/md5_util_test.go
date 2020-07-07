package tools

import "testing"

func TestGetMd5String(t *testing.T) {
	password:="25d55ad283aa400af464c76d713c07ad"
	secretKey:="a1b2c3"
	md5String := GetMd5String(password + secretKey)
	t.Log(md5String)
}
