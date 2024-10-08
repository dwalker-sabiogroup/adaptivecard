// Code generated by "enumer -type=FontType -json -transform=title-lower -trimprefix=FontType"; DO NOT EDIT.

package adaptivecard

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _FontTypeName = "defaultmonospace"

var _FontTypeIndex = [...]uint8{0, 7, 16}

const _FontTypeLowerName = "defaultmonospace"

func (i FontType) String() string {
	if i < 0 || i >= FontType(len(_FontTypeIndex)-1) {
		return fmt.Sprintf("FontType(%d)", i)
	}
	return _FontTypeName[_FontTypeIndex[i]:_FontTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FontTypeNoOp() {
	var x [1]struct{}
	_ = x[FontTypeDefault-(0)]
	_ = x[FontTypeMonospace-(1)]
}

var _FontTypeValues = []FontType{FontTypeDefault, FontTypeMonospace}

var _FontTypeNameToValueMap = map[string]FontType{
	_FontTypeName[0:7]:       FontTypeDefault,
	_FontTypeLowerName[0:7]:  FontTypeDefault,
	_FontTypeName[7:16]:      FontTypeMonospace,
	_FontTypeLowerName[7:16]: FontTypeMonospace,
}

var _FontTypeNames = []string{
	_FontTypeName[0:7],
	_FontTypeName[7:16],
}

// FontTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FontTypeString(s string) (FontType, error) {
	if val, ok := _FontTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FontTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FontType values", s)
}

// FontTypeValues returns all values of the enum
func FontTypeValues() []FontType {
	return _FontTypeValues
}

// FontTypeStrings returns a slice of all String values of the enum
func FontTypeStrings() []string {
	strs := make([]string, len(_FontTypeNames))
	copy(strs, _FontTypeNames)
	return strs
}

// IsAFontType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FontType) IsAFontType() bool {
	for _, v := range _FontTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for FontType
func (i FontType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FontType
func (i *FontType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FontType should be a string, got %s", data)
	}

	var err error
	*i, err = FontTypeString(s)
	return err
}
