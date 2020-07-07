package tools

import (
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"
)

// generate upload image file name
func GenerateImgName(fname string) string {
	i := rand.Intn(9999)
	timestamp := time.Now().Unix()
	ext := path.Ext(fname)
	fileName := strings.TrimSuffix(fname, ext)
	fileName = GetMd5String(fmt.Sprintf("%s-%d-%d", fname, timestamp, i))

	return fileName + ext
}
