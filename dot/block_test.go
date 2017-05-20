package dot_test

import (
	"github.com/decoz/go_lib/dot"
	"log"
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {

	fmt.Println("\nTestBlock")

	d := dot.Make("a(1(x,y,z),this,is,game,end)")
	//d := dot.Make("a(b,c)")
	buff := dot.Block(d)
	d2	 := dot.Unblock(buff)		
	log.Println( buff)
	log.Println( d2)

}

func TestTrim(t *testing.T) {
	fmt.Println("\nTestTrim")

	a :=  dot.Trim([]byte("\nx"))
	b :=  dot.Trim([]byte("\t  abc \t\n\t"))
	log.Println( a,  string(a))
	log.Println( b,  string(b))


}
