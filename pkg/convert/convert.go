package convert

import "strconv"

type StrTO string

func (s StrTO) string() string {
	return string(s)
}
func (s StrTO) Int() (int, error) {
	v, err := strconv.Atoi(s.string())
	return v, err
}
func (s StrTO) MustInt() int {
	v, _ := s.Int()
	return v
}
func (s StrTO) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.string())
	return uint32(v), err
}
func (s StrTO) MustInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
