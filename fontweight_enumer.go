// Code generated by "enumer -type=FontWeight -json -transform=title-lower -trimprefix=FontWeight"; DO NOT EDIT.

package adaptivecard

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _FontWeightName = "defaultlighterbolder"

var _FontWeightIndex = [...]uint8{0, 7, 14, 20}

const _FontWeightLowerName = "defaultlighterbolder"

func (i FontWeight) String() string {
	if i < 0 || i >= FontWeight(len(_FontWeightIndex)-1) {
		return fmt.Sprintf("FontWeight(%d)", i)
	}
	return _FontWeightName[_FontWeightIndex[i]:_FontWeightIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FontWeightNoOp() {
	var x [1]struct{}
	_ = x[FontWeightDefault-(0)]
	_ = x[FontWeightLighter-(1)]
	_ = x[FontWeightBolder-(2)]
}

var _FontWeightValues = []FontWeight{FontWeightDefault, FontWeightLighter, FontWeightBolder}

var _FontWeightNameToValueMap = map[string]FontWeight{
	_FontWeightName[0:7]:        FontWeightDefault,
	_FontWeightLowerName[0:7]:   FontWeightDefault,
	_FontWeightName[7:14]:       FontWeightLighter,
	_FontWeightLowerName[7:14]:  FontWeightLighter,
	_FontWeightName[14:20]:      FontWeightBolder,
	_FontWeightLowerName[14:20]: FontWeightBolder,
}

var _FontWeightNames = []string{
	_FontWeightName[0:7],
	_FontWeightName[7:14],
	_FontWeightName[14:20],
}

// FontWeightString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FontWeightString(s string) (FontWeight, error) {
	if val, ok := _FontWeightNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FontWeightNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FontWeight values", s)
}

// FontWeightValues returns all values of the enum
func FontWeightValues() []FontWeight {
	return _FontWeightValues
}

// FontWeightStrings returns a slice of all String values of the enum
func FontWeightStrings() []string {
	strs := make([]string, len(_FontWeightNames))
	copy(strs, _FontWeightNames)
	return strs
}

// IsAFontWeight returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FontWeight) IsAFontWeight() bool {
	for _, v := range _FontWeightValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for FontWeight
func (i FontWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FontWeight
func (i *FontWeight) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FontWeight should be a string, got %s", data)
	}

	var err error
	*i, err = FontWeightString(s)
	return err
}