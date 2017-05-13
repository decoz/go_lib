package dot

// check it is matched
func (d *Dot) _match(query *Dot) bool {

	return true
}

// used when '*' query is parent
func (d *Dot) _find(query *Dot) *Dot {

	return nil
}

// Put : target 구조를 추가
func (d *Dot) Put(target *Dot) {

	var child *Dot

	if target.Val() == "" {
		child = d
	} else {
		child = d.ChildV(target.Val())
	}

	if child == nil {
		d.Append(target)
		return
	}

	for _, t_child := range target.child {
		child.Put(t_child)
	}

}

// Copy : 구조를 lv 레벨까지 카피한다.
func (d *Dot) Copy(lv int) *Dot {
	/*
		lv 가 -1 이면 전체 구조를 카피한다.
	*/

	cp_d := newDot(d.value)
	if lv > 0 || lv < 0 {
		for _, t_child := range d.child {
			cp_d.add(t_child.Copy(lv - 1))
		}
	}

	return cp_d
}
