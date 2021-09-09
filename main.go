package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/janeczku/go-spinner"
)

func main() {

	c, err := net.Dial("tcp", fmt.Sprintf("%s:9100", os.Args[1]))
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(c, "\033*s2T\033*s2U\033*s4I\n")
	s := spinner.StartNew("Retrieving information from the printer. It will take a while...")
	line, err := bufio.NewReader(c).ReadString('\f')

	if err != nil {
		panic(err.Error())
	}

	s.Stop()
	fmt.Printf(">>> %s", line)
}
