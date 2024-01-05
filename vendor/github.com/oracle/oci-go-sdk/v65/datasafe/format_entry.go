// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FormatEntry A format entry is part of a masking format and defines the logic to mask data. A format
// entry can be a basic masking format such as Random Number, or it can be a library masking
// format. If a masking format has more than one format entries, the combined output of all
// the format entries is used for masking.
type FormatEntry interface {

	// The description of the format entry.
	GetDescription() *string
}

type formatentry struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *formatentry) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerformatentry formatentry
	s := struct {
		Model Unmarshalerformatentry
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *formatentry) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RANDOM_STRING":
		mm := RandomStringFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETERMINISTIC_SUBSTITUTION":
		mm := DeterministicSubstitutionFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETERMINISTIC_ENCRYPTION":
		mm := DeterministicEncryptionFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_DECIMAL_NUMBER":
		mm := RandomDecimalNumberFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_SUBSTITUTION":
		mm := RandomSubstitutionFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POST_PROCESSING_FUNCTION":
		mm := PpfFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NULL_VALUE":
		mm := NullValueFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIXED_NUMBER":
		mm := FixedNumberFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REGULAR_EXPRESSION":
		mm := RegularExpressionFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_DEFINED_FUNCTION":
		mm := UdfFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHUFFLE":
		mm := ShuffleFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIXED_STRING":
		mm := FixedStringFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TRUNCATE_TABLE":
		mm := TruncateTableFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIBRARY_MASKING_FORMAT":
		mm := LibraryMaskingFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_EXPRESSION":
		mm := SqlExpressionFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETERMINISTIC_ENCRYPTION_DATE":
		mm := DeterministicEncryptionDateFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_DIGITS":
		mm := RandomDigitsFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELETE_ROWS":
		mm := DeleteRowsFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUBSTRING":
		mm := SubstringFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATTERN":
		mm := PatternFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_NUMBER":
		mm := RandomNumberFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRESERVE_ORIGINAL_DATA":
		mm := PreserveOriginalDataFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_DATE":
		mm := RandomDateFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RANDOM_LIST":
		mm := RandomListFormatEntry{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FormatEntry: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m formatentry) GetDescription() *string {
	return m.Description
}

func (m formatentry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m formatentry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
