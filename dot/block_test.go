package dot_test

import (
	"github.com/decoz/go_lib/dot"
	"log"
	"testing"
)

func TestBlock(t *testing.T) {


	d := dot.Make("a(1(x,y,z),this,is,game,end)")
	//d := dot.Make("a(b,c)")
	buff := dot.Block(d)
	d2	 := dot.Unblock(buff)		
	log.Println("TestBlock:", buff)
	log.Println("\t", d2)
	





}

