package dot

import (
	"fmt"
	"strconv"
)

/*
func (d *Dot) Get(query interface{}) *Dot {

	query := Make(qstr);
	d_root := d._get()(d, query)
	fmt.Println("get:", d_root.String())
	return d_root

}
*/
// Get list of dot *link by query
func (d *Dot) Get(query *Dot) *Dot {
	//	d_root := Make("")
	d_root := d._get()(d, query)
	//fmt.Println("get:", d_root.String())
	return d_root
}

type getfunc func(src, query *Dot) *Dot



func (d *Dot) _get() getfunc {

	d_root := New("")
	var f getfunc 

	f = func(src, query *Dot) *Dot {
		//fmt.Println("f() :", src.Val(), query.Val())

		var d_search *Dot
		switch query.value[0] {
		case '*':
			for _, qchild := range query.child {
				d_search = f(src, qchild)
				if d_search != nil {
					d_root.Append(src.ChildV(qchild.Val()))
				}
			}
			for _, dchild := range src.child {
				f(dchild, query)
			}

		case '?':
			for _, dchild := range src.child {
				for _, qchild := range query.child {
					d_search = f(dchild, qchild)
					if d_search != nil {
						d_root.Append(dchild)
					}
				}
			}
		case '+':
			//fmt.Println("+ operation")
			if len(query.Val()) < 2 {
				fmt.Println("copy all:", src.String())
				return src.Copy(-1)
			} else {
				num, err := strconv.Atoi(query.Val()[1:])
				fmt.Println("num:", num)
				if err != nil {
					fmt.Println("error: bad argument for +")
					return nil
				} else {
					return src.Copy(num)
				}

			}
		case '$':
		default:
			child := src.ChildV(query.Val())
			if child == nil { return nil } 
			switch len(query.child)  {
			case 0:
				return child
			case 1:
				return f(child, query.child[0])
			default:
				find := true
				for _, qchild := range query.child {
					d_r := f(child, qchild)
					if d_r == nil { find = false }
				}
				if find {
					return child
				} else {
					return nil
				}
			}
		}
		//fmt.Println("return//", d_root.String())
		return d_root
	}
	return f
}
