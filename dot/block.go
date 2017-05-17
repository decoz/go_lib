package dot

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
)

/*
	dot archtecture 를 serial 형태로 변환한다.
*/


type node struct{
	ccnt int16
	sz	 int16
}


func Block(dot *Dot ) []byte {

	
	//n := node{ccnt: int16(len(dot.child)), sz: int16(len(dot.value))}
	buff  := new(bytes.Buffer)
	//var cc, sz int16


	cl := int16( len(dot.child) ) 
	vl := len(dot.value)	
	
	binary.Write(buff, binary.LittleEndian, byte(vl))
	binary.Write(buff, binary.LittleEndian, cl)

	log.Println("dot.Block: ",len(buff.Bytes()) ,"/", buff.Bytes())
	return buff.Bytes()

}

func _block(dot *Dot, p []byte){

	

}

/*
	block 구조를 
*/
func Unblock(buff []byte ) *Dot {

	r:=  bytes.NewReader(buff)
	
	var val  int
	binary.Read(r, binary.LittleEndian, &val)	
		
	t := Make(strconv.Itoa(val))	
	return t 

}
