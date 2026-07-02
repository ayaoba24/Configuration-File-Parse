package main


import (
	"fmt"
	"strings"
)

func main (){
	conifg := "host = local\nport=80080\ndebug=true"

	for _, line := range strings.Split(conifg, "\n"){
		part := strings.SplitN(line, "=", 2)
		key := strings.ToUpper(part[0][:1]) + part[0][1:]

		fmt.Println(key + ":" + part[1])

	}
}