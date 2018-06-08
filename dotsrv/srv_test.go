package dotsrv_test

import (
	"github.com/decoz/go_lib/dotsrv"
	"github.com/decoz/go_lib/dot"
	"testing"
)

type simple_srv struct {

}

func (ss *simple_srv) Put(d *dot.Dot) *dot.Dot{
		return d
}


func TestRun(t *testing.T){
	s := new( simple_srv )
	dotsrv.Lunch(8515, s)
}
