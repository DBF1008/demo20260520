package test

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
	"testing"
)

func TestContainer(t *testing.T) {



	t.Run("ScanPath", scanPath)
}

func scanPath(t *testing.T) {
	var (
		packages     []string
		err          error
		packageNames []string
	)

	path := strings.Join([]string{"D:/goProject/p2022/gfast-v3/internal/app/demo/logic"}, "")

	packages, err = gfile.ScanDirFunc(path, "*", false, func(p string) string {
		if gfile.IsDir(p) {
			return p
		}
		return ""
	})
	if err != nil {
		panic(err)
	}
	packageNames = make([]string, len(packages))
	for k, v := range packages {
		fmt.Println("111111", v)
		v = gstr.Replace(v, "\\", "/")
		fmt.Println("2222222", v)
		packageNames[k] = gstr.SubStr(v, gstr.PosR(v, "/")+1)
	}
	fmt.Println("dddd", packageNames)
}
