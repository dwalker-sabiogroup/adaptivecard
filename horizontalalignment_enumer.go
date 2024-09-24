// Code generated by "enumer -type=HorizontalAlignment -json -transform=title-lower -trimprefix=HorizontalAlignment"; DO NOT EDIT.

package adaptivecard

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _HorizontalAlignmentName = "leftcenterright"

var _HorizontalAlignmentIndex = [...]uint8{0, 4, 10, 15}

const _HorizontalAlignmentLowerName = "leftcenterright"

func (i HorizontalAlignment) String() string {
	if i < 0 || i >= HorizontalAlignment(len(_HorizontalAlignmentIndex)-1) {
		return fmt.Sprintf("HorizontalAlignment(%d)", i)
	}
	return _HorizontalAlignmentName[_HorizontalAlignmentIndex[i]:_HorizontalAlignmentIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _HorizontalAlignmentNoOp() {
	var x [1]struct{}
	_ = x[HorizontalAlignmentLeft-(0)]
	_ = x[HorizontalAlignmentCenter-(1)]
	_ = x[HorizontalAlignmentRight-(2)]
}

var _HorizontalAlignmentValues = []HorizontalAlignment{HorizontalAlignmentLeft, HorizontalAlignmentCenter, HorizontalAlignmentRight}

var _HorizontalAlignmentNameToValueMap = map[string]HorizontalAlignment{
	_HorizontalAlignmentName[0:4]:        HorizontalAlignmentLeft,
	_HorizontalAlignmentLowerName[0:4]:   HorizontalAlignmentLeft,
	_HorizontalAlignmentName[4:10]:       HorizontalAlignmentCenter,
	_HorizontalAlignmentLowerName[4:10]:  HorizontalAlignmentCenter,
	_HorizontalAlignmentName[10:15]:      HorizontalAlignmentRight,
	_HorizontalAlignmentLowerName[10:15]: HorizontalAlignmentRight,
}

var _HorizontalAlignmentNames = []string{
	_HorizontalAlignmentName[0:4],
	_HorizontalAlignmentName[4:10],
	_HorizontalAlignmentName[10:15],
}

// HorizontalAlignmentString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func HorizontalAlignmentString(s string) (HorizontalAlignment, error) {
	if val, ok := _HorizontalAlignmentNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _HorizontalAlignmentNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to HorizontalAlignment values", s)
}

// HorizontalAlignmentValues returns all values of the enum
func HorizontalAlignmentValues() []HorizontalAlignment {
	return _HorizontalAlignmentValues
}

// HorizontalAlignmentStrings returns a slice of all String values of the enum
func HorizontalAlignmentStrings() []string {
	strs := make([]string, len(_HorizontalAlignmentNames))
	copy(strs, _HorizontalAlignmentNames)
	return strs
}

// IsAHorizontalAlignment returns "true" if the value is listed in the enum definition. "false" otherwise
func (i HorizontalAlignment) IsAHorizontalAlignment() bool {
	for _, v := range _HorizontalAlignmentValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for HorizontalAlignment
func (i HorizontalAlignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for HorizontalAlignment
func (i *HorizontalAlignment) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("HorizontalAlignment should be a string, got %s", data)
	}

	var err error
	*i, err = HorizontalAlignmentString(s)
	return err
}
