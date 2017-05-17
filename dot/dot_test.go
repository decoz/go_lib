package dot_test

import (
	"fmt"
	"github.com/decoz/go_lib/dot"
	"log"
	"testing"
)

func TestPut(t *testing.T) {

	d := dot.Make("a.b")
	d.Put(dot.Make("(c,d,e)"))
	log.Println("dot.put:", d)

}

func TestGet(t *testing.T) {

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
		"c.d", "*.d.f", "*.b(d,e)"}

	qlist[1] = []string{
		"*.b", "?.b", "b.?.b"}

	for i, qstrs := range qlist {
		for _, qstr := range qstrs {

			fmt.Println("get ", qstr, " from ", dstrs[i])
			fmt.Println("result :", ds[i].Get(dot.Make(qstr)))

		}
	}

	d1 := dot.Make("a(b(d,e),c.d.f)")
	d2 := dot.Make("a(b.b.b,a.b.c.c)")

	log.Println("dot.get", d1)

	q1_0 := dot.Make("c.d")
	q1_1 := dot.Make("*.b(d,e)")
	q1 := dot.Make("*.d.f")
	q2 := dot.Make("*.b")
	q2_1 := dot.Make("?.b")
	q2_2 := dot.Make("b.?.b")

	fmt.Println("get ", d1, q1_0, " result :", d1.Get(q1_0))
	fmt.Println("get ", d1, q1_1, " result :", d1.Get(q1_1))
	fmt.Println("get ", d1, q1, " result :", d1.Get(q1))
	fmt.Println("get ", d2, q2, " result :", d2.Get(q2))
	fmt.Println("get ", d2, q2_1, " result :", d2.Get(q2_1))

	fmt.Println("get ", d2, q2_2, " result :", d2.Get(q2_2))

	/*
		if rid.String() != "(e,e)" {
			t.Error("N depth get fail")
		}
	*/
}

func TestParser(t *testing.T) {
	str := "x(a,b)"
	p := dot.CreateParser([]byte(str))
	n := p.Parse([]byte("root"), 0)
	log.Println(n)

}

func TestAttach(t *testing.T) {
	d := dot.Make("/(a,b)")
	d_attach := dot.Make("b.bb.bbb(1,2,3,4.x)")
	d.Attach(d_attach)
	log.Println(d)

}
