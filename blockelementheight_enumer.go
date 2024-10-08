// Code generated by "enumer -type=BlockElementHeight -json -transform=title-lower -trimprefix=BlockElementHeight"; DO NOT EDIT.

package adaptivecard

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _BlockElementHeightName = "autostretch"

var _BlockElementHeightIndex = [...]uint8{0, 4, 11}

const _BlockElementHeightLowerName = "autostretch"

func (i BlockElementHeight) String() string {
	if i < 0 || i >= BlockElementHeight(len(_BlockElementHeightIndex)-1) {
		return fmt.Sprintf("BlockElementHeight(%d)", i)
	}
	return _BlockElementHeightName[_BlockElementHeightIndex[i]:_BlockElementHeightIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _BlockElementHeightNoOp() {
	var x [1]struct{}
	_ = x[BlockElementHeightAuto-(0)]
	_ = x[BlockElementHeightStretch-(1)]
}

var _BlockElementHeightValues = []BlockElementHeight{BlockElementHeightAuto, BlockElementHeightStretch}

var _BlockElementHeightNameToValueMap = map[string]BlockElementHeight{
	_BlockElementHeightName[0:4]:       BlockElementHeightAuto,
	_BlockElementHeightLowerName[0:4]:  BlockElementHeightAuto,
	_BlockElementHeightName[4:11]:      BlockElementHeightStretch,
	_BlockElementHeightLowerName[4:11]: BlockElementHeightStretch,
}

var _BlockElementHeightNames = []string{
	_BlockElementHeightName[0:4],
	_BlockElementHeightName[4:11],
}

// BlockElementHeightString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func BlockElementHeightString(s string) (BlockElementHeight, error) {
	if val, ok := _BlockElementHeightNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _BlockElementHeightNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to BlockElementHeight values", s)
}

// BlockElementHeightValues returns all values of the enum
func BlockElementHeightValues() []BlockElementHeight {
	return _BlockElementHeightValues
}

// BlockElementHeightStrings returns a slice of all String values of the enum
func BlockElementHeightStrings() []string {
	strs := make([]string, len(_BlockElementHeightNames))
	copy(strs, _BlockElementHeightNames)
	return strs
}

// IsABlockElementHeight returns "true" if the value is listed in the enum definition. "false" otherwise
func (i BlockElementHeight) IsABlockElementHeight() bool {
	for _, v := range _BlockElementHeightValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for BlockElementHeight
func (i BlockElementHeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BlockElementHeight
func (i *BlockElementHeight) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BlockElementHeight should be a string, got %s", data)
	}

	var err error
	*i, err = BlockElementHeightString(s)
	return err
}
