// Code generated by "stringer -trimprefix sec -type sectionID"; DO NOT EDIT.

package wasm

import "strconv"

const _sectionID_name = "CustomTypeImportFunctionTableMemoryGlobalExportStartElementCodeData"

var _sectionID_index = [...]uint8{0, 6, 10, 16, 24, 29, 35, 41, 47, 52, 59, 63, 67}

func (i sectionID) String() string {
	if i >= sectionID(len(_sectionID_index)-1) {
		return "sectionID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _sectionID_name[_sectionID_index[i]:_sectionID_index[i+1]]
}
