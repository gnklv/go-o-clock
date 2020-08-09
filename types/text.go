package types

type Text []Symbol

func (t Text) Width() int {
	w := 0
	for _, s := range t {
		w += s.Width()
	}
	return w
}

func (t Text) Height() int {
	return len(t[0])
}
