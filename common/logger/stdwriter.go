package logger

import (
	"fmt"
	"os"
)

type StdWriter struct {
}

func (sw *StdWriter) Write(data interface{}) {
	str := fmt.Sprintf("%v\n", data)

	if _, err := os.Stdout.WriteString(str); err != nil {
		panic(err)
	}
}