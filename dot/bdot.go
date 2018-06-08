/*
	goalang Base Dot Object

	DB, File 등의 Stroage 와 결합하기 위해 Block 화와 Ublock을 지원하는
	Dot Object

*/
package dot

import (

)

type DoBase interface {

	Write(key uint32,data []byte ) error
	Read(key uint32) ( []byte, error )

}

type Bdot struct {
	Dot
	closed bool
	key	uint32
	db	DoBase
}

func (bdot *Bdot) Val() string{

	if len(bdot.value) == 0 && bdot.value[0] == 0 {
		bdot.Load()
	}

	return Dec(bdot.value)
}


func (bdot *Bdot) Load(){

	data,_  := bdot.db.Read(bdot.key)
	nd := dot.Block(data)
	bdot.value = nd.value
	bdot.child = nd.child

}
