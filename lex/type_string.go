// generated by stringer -type=Type; DO NOT EDIT

package lex

import "fmt"

const _Type_name = "EOFIgnoreNewLineIncTapeDecTapeIncByteDecByteOutputByteStoreByteLoopStartLoopEnd"

var _Type_index = [...]uint8{0, 3, 9, 16, 23, 30, 37, 44, 54, 63, 72, 79}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}