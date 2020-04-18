package main

import (
	"awesomeProject/shakespeare"
	"fmt"
	"strings"
)

func main(){
	x:=shakespeare.DidShakespeareSayThat("help")
	fmt.Println(strings.Join(x, "\n\n"))


}
