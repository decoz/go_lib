package dot

import (
	"bytes"
	"encoding/binary"
	//"log"
	//"strconv"
)

/*
	dot archtecture 를 serial 형태로 변환한다.
*/


type node struct{
	ccnt int16
	sz	 int16
}



func d_block(dot *Dot, nw *bytes.Buffer, vw *bytes.Buffer) {

	cnt := uint16( len(dot.child) )
	sz := uint16( len(dot.value) )

	//log.Println(string(dot.value), ":", cnt , "/", sz )

	binary.Write(nw, binary.LittleEndian, cnt)
	binary.Write(nw, binary.LittleEndian, sz)

	vw.Write(dot.value)

	for _,cdot := range(dot.child) {
		d_block(cdot, nw, vw)
	}

}

func Block(dot *Dot ) []byte {

	ncnt,_ :=  _get_size(dot)

	nw := new( bytes.Buffer )
	vw := new( bytes.Buffer )

	binary.Write(nw, binary.LittleEndian, uint16(ncnt))
	d_block(dot, nw, vw)
	
	nw.Write( vw.Bytes())
	return nw.Bytes()

}


func Unblock(buf []byte ) *Dot {

	nr:= bytes.NewReader(buf)

	var ncnt  uint16
	binary.Read(nr, binary.LittleEndian, &ncnt)

	vr:= bytes.NewReader( buf[ ncnt * 4 + 2 :] )
	dot := d_unblock(nr, vr)

	return dot
}

func d_unblock( nr *bytes.Reader, vr *bytes.Reader) *Dot{

	var ccnt, sz uint16
	binary.Read(nr, binary.LittleEndian,&ccnt )
	binary.Read(nr, binary.LittleEndian,&sz )

	val := make([]byte, int(sz) )
	vr.Read(val)
	d := newDot(val)

	for i:=0; i< int(ccnt); i++ {
		d.Append(d_unblock(nr, vr))
	}


	return d

}


func _get_size(dot *Dot) (int, int){

	c := 1
	l := len(dot.value)

	for _,ch := range(dot.child){
		cc,cl := _get_size(ch)
		c += cc
		l += cl
	}

	return c,l

}
