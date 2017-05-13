package dot

import (
	"strconv"
	"strings"
)

type Dots []*Dot

type Dot struct {
	value []byte
	child Dots
}

func New(str string) *Dot {
	/*
		문자열의 값을 지닌 단일 dot 을 생성
		문자열에 ., (, ) 등의 기호가 있으면
		encoding 함
	*/
	
	return newDot(Enc(str))
}


func newDot(value []byte) *Dot{
	n := new(Dot)
	n.value = value 	
	return n
}



func Make(str string) *Dot {
	/*
		Dot 을 파싱 / 생성
	*/

	p := CreateParser([]byte(str))

	dot := p.Parse([]byte(""), 0)
	if len(dot.child) == 1 {
		return dot.child[0]
	} else {
		return dot
	}

}

func (n *Dot) Attach(c *Dot) {
	/*
		c 를 n에 추가한다.
		Put과 중복 ( 둘중 하나 지울 계획)
	*/
	r := n.ChildV(c.Val())
	if r == nil {
		n.Append(c)
	} else {
		for _, child := range c.child {
			r.Attach(child)
		}
	}

}

func (dot *Dot) CClear() {
	dot.child = make([]*Dot, 5)[0:0]
}

func (dot *Dot) Print() string {
	// dot 의 내용을 시각적으로 보기 편하게 출력한다.
	return dot.print_dot(0)
}

func (dot *Dot) print_dot(lv int) string {

	str := ""
	for i := 0; i < lv; i++ {
		str += "\t"
	}

	sz := len(dot.value)

	if sz > 30 {
		str += string(dot.value[0:30])
		str += "... +" + strconv.Itoa(sz) + " byte"
	} else {
		str += string(dot.value)
	}

	str += "\n"

	for _, child := range dot.child {
		str += child.print_dot(lv + 1)
	}

	return str
}

func (n *Dot) add(c *Dot) {

	if n.child == nil {
		n.child = make([]*Dot, 5)[0:0]
	}
	if len(c.child) > 0 || len(c.value) > 0 {
		//log.Println("add ",string(c.value),"with", len(c.child),"children")
		n.child = append(n.child, c)
	}
}

func (dot Dot) String() string {
	return dot.str()
}

func (dot *Dot) str() string {

	str := ""
	ccnt := len(dot.child)

	if dot.value != nil {
		str += string(dot.value)
	}

	switch {

	case ccnt == 0:

	case ccnt == 1:
		str += "." + dot.child[0].str()

	case ccnt > 1:

		str += "("

		isfirst := true

		for _, child := range dot.child {
			if isfirst {
				isfirst = false
			} else {
				str += ","
			}

			str += child.str()
		}

		str += ")"
	}

	return str

}

func (dot *Dot) ChildN(i int) *Dot {

	if i < len(dot.child) {
		return dot.child[i]
	} else {
		return nil
	}
}

func (dot *Dot) CPath(path string) *Dot {

	arr := strings.Split(path, ".")

	cur := dot
	for _, p := range arr {
		cur = cur.ChildV(p)
		if cur == nil {
			return nil
		}
	}

	return cur
}


func (dot *Dot) CList() []*Dot {
	return dot.child

}

func (dot *Dot) Append(children ...*Dot) {
	for _, child := range children {
		dot.add(child)
	}
}

func (dot *Dot) Set(str string) {
	dot.value = Enc(str)

}

func (dot *Dot) Val() string {
	return Dec(dot.value)
}

func (dot *Dot) CVal() string {
	/*
		node 의 child 벨류 값을 리턴한다. 	현재는 퍼스트 노드의 값을 리턴하되
		차후에 하위 노드중 터미널 노드 의 값을 종합해서 리턴하는 것으로 업그레이드예정
	*/

	c := dot.ChildN(0)

	if c != nil {
		return c.Val()
	} else {
		return ""
	}

}
emove
func (dot *Dot) RemoveV(k string) bool {
	/*
		k값을 갖는 엘리먼트를 제거
		값이 발견된 경우 true 아니면 false 를 리턴
		(같은 값이 여러개일 경우 첫번째 경우만 지운다.)
	*/
	for i, d := range dot.child {
		if string(d.value) == k {
			dot.child = append(dot.child[:i], dot.child[i+1:]...)
			return true
		}
	}

	return false
}

func (dot *Dot) ChildV(k string) *Dot {
	for _, d := range dot.child {
		if string(d.value) == k {
			return d
	Remove
	}
	return nil
}


func Enc(str string) []byte {

	src := []byte(str)
	dst := make([]byte, len(str)+10)[0:0]

	for _, c := range src {
		switch c {
		case '.':
			dst = append(dst, '&', 'd')
		case ',':
			dst = append(dst, '&', 'c')
		case '(':
			dst = append(dst, '&', 's')
		case ')':
			dst = append(dst, '&', 'e')
		case '&':
			dst = append(dst, '&', '&')
		case ' ':
			dst = append(dst, '&', 'b')
		default:
			dst = append(dst, c)
		}
	}

	return dst
}

func Dec(src []byte) string {

	dst := make([]byte, len(src))[0:0]

	flag := false
	for _, c := range src {
		if flag {
			switch c {
			case 'd':
				dst = append(dst, '.')
			case 'c':
				dst = append(dst, ',')
			case 's':
				dst = append(dst, '(')
			case 'e':
				dst = append(dst, ')')
			case '&':
				dst = append(dst, '&')
			case 'b':
				dst = append(dst, ' ')
			}
			flag = false
		} else {
			if c == '&' {
				flag = true
			} else {
				dst = append(dst, c)
			}
		}

	}

	return string(dst)
}
