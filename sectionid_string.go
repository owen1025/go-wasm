// Code generated by "stringer -type SectionID -trimprefix Section"; DO NOT EDIT.

package wasm

import "strconv"

const _SectionID_name = "CustomTypeImportFunctionTableMemoryGlobalExportStartElementCodeData"

var _SectionID_index = [...]uint8{0, 6, 10, 16, 24, 29, 35, 41, 47, 52, 59, 63, 67}

func (i SectionID) String() string {
	if i >= SectionID(len(_SectionID_index)-1) {
		return "SectionID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SectionID_name[_SectionID_index[i]:_SectionID_index[i+1]]
}