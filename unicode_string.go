package gowin

import (
	"unicode/utf16"
	"unsafe"
)

func NewUnicodeString(str string) UnicodeString {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return UnicodeString{
		Length:        uint16(len(wchars)),
		MaximumLength: uint16(cap(wchars)),
		Buffer:        &wchars[0],
	}
}

func (u UnicodeString) String() string {
	sl := struct {
		addr uintptr
		len  int
		cap  int
	}{}
	sl.addr = uintptr(unsafe.Pointer(u.Buffer))
	sl.len = int(u.Length / 2)
	sl.cap = int(u.MaximumLength / 2)
	u16 := *(*[]uint16)(unsafe.Pointer(&sl))
	return string(utf16.Decode(u16))
}
