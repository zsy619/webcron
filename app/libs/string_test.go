package libs

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	pwd := Md5([]byte("admin888"))
	fmt.Println(pwd)
}
