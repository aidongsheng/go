// generated by stringer -type=Class; DO NOT EDIT

package dwarf

import "fmt"

const _Class_name = "ClassUnknownClassAddressClassBlockClassConstantClassExprLocClassFlagClassLinePtrClassLocListPtrClassMacPtrClassRangeListPtrClassReferenceClassReferenceSigClassStringClassReferenceAltClassStringAlt"

var _Class_index = [...]uint8{0, 12, 24, 34, 47, 59, 68, 80, 95, 106, 123, 137, 154, 165, 182, 196}

func (i Class) String() string {
	if i < 0 || i+1 >= Class(len(_Class_index)) {
		return fmt.Sprintf("Class(%d)", i)
	}
	return _Class_name[_Class_index[i]:_Class_index[i+1]]
}
