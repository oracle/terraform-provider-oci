// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PiiEntityMasking Mask recognized PII entities with different modes.
type PiiEntityMasking interface {
}

type piientitymasking struct {
	JsonData []byte
	Mode     string `json:"mode"`
}

// UnmarshalJSON unmarshals json
func (m *piientitymasking) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpiientitymasking piientitymasking
	s := struct {
		Model Unmarshalerpiientitymasking
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Mode = s.Model.Mode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *piientitymasking) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Mode {
	case "REPLACE":
		mm := PiiEntityReplace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE":
		mm := PiiEntityRemove{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MASK":
		mm := PiiEntityMask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PiiEntityMasking: %s.", m.Mode)
		return *m, nil
	}
}

func (m piientitymasking) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m piientitymasking) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PiiEntityMaskingModeEnum Enum with underlying type: string
type PiiEntityMaskingModeEnum string

// Set of constants representing the allowable values for PiiEntityMaskingModeEnum
const (
	PiiEntityMaskingModeReplace PiiEntityMaskingModeEnum = "REPLACE"
	PiiEntityMaskingModeMask    PiiEntityMaskingModeEnum = "MASK"
	PiiEntityMaskingModeRemove  PiiEntityMaskingModeEnum = "REMOVE"
)

var mappingPiiEntityMaskingModeEnum = map[string]PiiEntityMaskingModeEnum{
	"REPLACE": PiiEntityMaskingModeReplace,
	"MASK":    PiiEntityMaskingModeMask,
	"REMOVE":  PiiEntityMaskingModeRemove,
}

var mappingPiiEntityMaskingModeEnumLowerCase = map[string]PiiEntityMaskingModeEnum{
	"replace": PiiEntityMaskingModeReplace,
	"mask":    PiiEntityMaskingModeMask,
	"remove":  PiiEntityMaskingModeRemove,
}

// GetPiiEntityMaskingModeEnumValues Enumerates the set of values for PiiEntityMaskingModeEnum
func GetPiiEntityMaskingModeEnumValues() []PiiEntityMaskingModeEnum {
	values := make([]PiiEntityMaskingModeEnum, 0)
	for _, v := range mappingPiiEntityMaskingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPiiEntityMaskingModeEnumStringValues Enumerates the set of values in String for PiiEntityMaskingModeEnum
func GetPiiEntityMaskingModeEnumStringValues() []string {
	return []string{
		"REPLACE",
		"MASK",
		"REMOVE",
	}
}

// GetMappingPiiEntityMaskingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPiiEntityMaskingModeEnum(val string) (PiiEntityMaskingModeEnum, bool) {
	enum, ok := mappingPiiEntityMaskingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
