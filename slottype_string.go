// Code generated by "stringer -type=SlotType"; DO NOT EDIT

package mtx

import "fmt"

const _SlotType_name = "DataTransferStorageSlotMailSlot"

var _SlotType_index = [...]uint8{0, 12, 23, 31}

func (i SlotType) String() string {
	if i < 0 || i >= SlotType(len(_SlotType_index)-1) {
		return fmt.Sprintf("SlotType(%d)", i)
	}
	return _SlotType_name[_SlotType_index[i]:_SlotType_index[i+1]]
}
