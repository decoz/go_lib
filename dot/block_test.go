package dot_test

import (
	"github.com/decoz/go_lib/dot"
	"log"
	"testing"
)

func TestBlock(t *testing.T) {

	d := dot.Make("a(b,b)")
	buff := dot.Block(d)
	d2 := dot.Unblock(buff)		
	log.Println("TestBlock:", d2)
	




}

