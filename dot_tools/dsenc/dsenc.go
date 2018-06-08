package main

import (
	"flag"
	"github.com/decoz/go_lib/dot"
	"fmt"
	"os"
	"bytes"
)

func main() {

	flag.Parse()
	args := flag.Args()

	loop := 1
	f := os.Stdin
	if len(args) > 0  { loop = len(args) }

	var err error

	for i:=0; i<loop; i++{
		if len(args) > 0 {
			f,err = os.OpenFile(args[i], os.O_RDONLY, 0755)
		}

		if err != nil {
			fmt.Println("fail to open file", args[i])
			continue
		}
		denc(f)

	}
}

func denc(f *os.File ){

	bufw := new( bytes.Buffer  )
	buff := make([]byte, 100)
	for {
		n,err := f.Read(buff)
		if err != nil { break }
		if n>0 { bufw.Write(buff[0:n]) }
	} 

	f.Close()

	b := bufw.Bytes()	
	fmt.Println(b, string(b))
	d := dot.Make(string(b))
	//d := dot.Make( string(bufw.Bytes()) )
	fmt.Println(d)
}


