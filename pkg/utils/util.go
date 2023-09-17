package utils

import (
	"bytes"
	"fmt"

	"github.com/KazuhiroIto0127/gachagacha-cli/resources"
)

func AaFromFile(fp string) string {
	file, err := resources.Aa.Open("aa/"+fp)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}
