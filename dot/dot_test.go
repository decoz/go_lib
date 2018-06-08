package dot_test

import (
	"fmt"
	"github.com/decoz/go_lib/dot"
	"log"
	"testing"
)

func TestParser(t *testing.T) {
	fmt.Println("\nTestParser")

	str := "x(a, b)"
	p := dot.CreateParser([]byte(str))
	n := p.Parse([]byte("root"), 0)

	log.Println(p)
	log.Println(n)

}

func TestPut(t *testing.T) {

	fmt.Println("\nTestPut")

	dstrs := []string{
		"a(b.c,d.e)" }

	n := len(dstrs)
	ds := make([]*dot.Dot, n)
	qlist := make([][]string, n)

	for i, dstr := range dstrs {
		ds[i] = dot.Make(dstr)
		fmt.Println(ds[i])
	}
	qlist[0] = []string{
		"b.d", "d.e(f,g)", "d/x"}

	for i, qstrs := range qlist {
		for _, qstr := range qstrs {
			fmt.Println("\nget ", qstr, " from ", dstrs[i])
			ds[i].Set(qstr)
			fmt.Println("result :",ds[i])

		}
	}
}



func TestGet(t *testing.T) {

	fmt.Println("\nTestGet")
	dstrs := []string{
		"a(b(d,e),c.d.f)",
		"a(b.b.b,a.b.c.c)",
		"a.b.c.d.e.f.g.h.i"}

	n := len(dstrs)
	ds := make([]*dot.Dot, n)
	qlist := make([][]string, n)

	for i, dstr := range dstrs {
		ds[i] = dot.Make(dstr)
		fmt.Println(ds[i])
	}
	qlist[0] = []string{
		"c.d", "b(d,e)", "*.d.f", "*(d,e)"}

	qlist[1] = []string{
		"*.b", "?.b", "b.?.b", "*.?.?.c"}

	for i, qstrs := range qlist {
		for _, qstr := range qstrs {
			fmt.Println("\nget ", qstr, " from ", dstrs[i])
			fmt.Println("result :", ds[i].Get(qstr))

		}
	}
}


func TestAttach(t *testing.T) {
	fmt.Println("\nTestAttach")

	d := dot.Make("/(a,b)")
	d_attach := dot.Make("b.bb.bbb(1,2,3,4.x)")
	d.Attach(d_attach)
	log.Println(d)

}
