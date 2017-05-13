package dot

import (
	"strings"
)




type Parser struct{
	pos int
	data []byte
	symbol []byte
	err chan error
}


const  (
	nt_value = 1 << iota
	nt_map
	nt_list
)



func CreateParser(idata []byte) *Parser{

	parser := new(Parser)
	parser.data = idata
	parser.err = make(chan error)
	parser.symbol = []byte{'.', '(', ')', ','}

	return parser

}

func (parser *Parser) SetData(input []byte){
	parser.data = input
	parser.pos = 0
}

func (parser *Parser) Parse(name []byte, depth int) *Dot{

	n := new(Dot)
	n.value = name
	n.child = make([]*Dot,5)[0:0]

	var (
		token []byte
		symbol byte
	)

	token,symbol = parser.next()


	for parser.pos < len(parser.data) {

		//log.Print("[",parser.pos,":",string(parser.data[parser.pos]),"]")

		switch symbol {

		case ',' :
					//if len(token) > 0 { n.add(New(token)) }
					n.add(newDot(token))
					if depth > 0 {
						//log.Println("back with",string(token))
						return n
					} else {
						parser.pos++
					}


		case '(' :  parser.pos++
					n.add(parser.Parse(token,0))

		case ')' :  n.add(newDot(token))
					if depth > 0 {
						return n
					} else {
						parser.pos++
						return n
					}
		case '.' : parser.pos++
				   n.add(parser.Parse(token,depth+1))

		}

		token,symbol = parser.next()
	}

	if len(token) > 0 { n.add( newDot(token) ) }
	return n
}

func (parser *Parser) next() ([]byte,byte){
/*
	특별 심볼까지 파싱을 진행한다.
*/
	buff := make([]byte,1000)[0:0]
	for parser.pos < len(parser.data) {
		c := parser.data[parser.pos]
		switch c {
			case '.', '(', ')', ',':
				//log.Println(string(buff),":",len(buff),"<-",string(c))
				return buff,c
			default:
				buff = append(buff,c)
		}
		parser.pos++
	}

	return buff,0
}




func TabParse(bdata []byte) *Dot{
/*
	tab 단위로 구분된 dot 구조를
	파싱하는 함수

	- 탭과 문자열

*/
	root := Make("")

	str := string(bdata)
	lines := strings.Split(str,"\r\n")


	depth := make( []int, 100 )

	path := make( []*Dot, 100)
	path[0] = root
	depth[0] = -1

	pos := 0

	for _,l := range(lines) {

		di := 0
		deep := 0

		for di<len(l)-1 && (l[di]=='\t' || l[di]==' ') {
			if l[di] == '\t' { deep += 5
			} else { deep++ }
			di++
		}

		er := len(l)-1

		for ; er >= 0 && ( l[er]==' ' || l[er]=='\t' ); er-- {}
		if di >= er+1 { continue }

		d := New( string(l[di:er+1]) )

		i:=pos
		for ;depth[i] >= deep; i-- {}

		path[i].Append(d)
		pos = i+1
		path[pos] = d
		depth[pos] = deep

		//log.Println(li, ":", deep, ">", pos, " ", d.Val())

	}

	return root

}

