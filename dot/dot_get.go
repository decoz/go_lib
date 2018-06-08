package dot

import (
	//"fmt"
	//"strconv"
)

// Get list of dot *link by query
func (d *Dot) Get(s_query string) *Dot {
	//	d_root := Make("")
	query := Make(s_query)
	d_root := New("")
	for _,child := range d.child {
		d_r := child._qget(query)
		if d_r != nil {
			if d_r.Val() == ""  {
				d_r.Inherit( d_root )
			} else {
				d_root.Append(d_r)
			}
		}
	}
	return d_root
}

func (dot *Dot) Inherit(d_scc *Dot){

	for _,child := range dot.child {
		d_scc.Append( child )
	}

}


func (d *Dot) _qget(query *Dot) *Dot {

	//fmt.Println(" q() :", d.Val(), query.Val())
	rr := New("") // rr

	take := func(d *Dot, td *Dot ){
		if td.Val() == ""  {
			td.Inherit(d)
		} else {
			d.Append(td)
		}
	}

	match := func(d *Dot, query *Dot) bool{
		for _, qchild := range query.child {
			find := false
			for _, child := range d.child {
				d_r := child._qget(qchild)
				if d_r != nil { find = true }
			}
			if !find  { return false }
		}
		return true
	}

	switch query.value[0] {
	case '*':

		if match(d, query)  { rr.Append(d) }

		for _, dchild := range d.child {
			d_r := dchild._qget(query)
			if d_r != nil {
				take(rr, d_r )
			}
		}

	case '?':
		if match(d, query) { rr.Append(d) }

	case '+': // 미정
	case '$':

	default:
		if 	d.Val() != query.Val() ||
			len(d.child) < len(query.child){
			return nil
		}

		switch len(query.child) {
		case 0:
			return d
		case 1:
			 for _, child := range d.child {
				d_r :=  child._qget(query.child[0])
				if d_r != nil {      // only return first search
					take(rr,  d_r )
				}
			}
		default:
			if match(d, query)	{ rr.Append(d) }
		}
	}
	//fmt.Println(" return ://", rr)

	switch len(rr.child) {
	case 0 :
		return nil
	case 1 :
		return rr.ChildN(0)
	default:
		return rr
	}


}
