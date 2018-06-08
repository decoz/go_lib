package dot

import (
  "strings"
	//"fmt"
	//"strconv"
)

func (dot *Dot) Set( s_query string) {

}

func set(d *Dot, query *Dot){
  qv := query.Val()
  el := strings.Split(qv, "/")
	pl := strings.Split(qv, "+")
	switch {
  case len(el) > 1 :
		if el[1] == "" {
			d.CClear()
			if el[0] != "" {
				nc := New(el[0])
				d.Append( nc )
				nc.Inherit(query)
			}
    }
  case len(pl) > 1 :
	default:
		child := d.ChildV(query.Val())
		if child  == nil  { return  }

		for _,qchild := range query.child {
			set( child, qchild )
		}
	}
}

func (dot *Dot) _put( query *Dot){

}
