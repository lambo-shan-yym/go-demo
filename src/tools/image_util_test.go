package tools

import "testing"

func TestGenerateImgName(t *testing.T) {
	name := GenerateImgName("filename.jpg")
	t.Log(name)
}
