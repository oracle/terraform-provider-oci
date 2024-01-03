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

// ClassificationType possible text classification modes
type ClassificationType interface {
}

type classificationtype struct {
	JsonData           []byte
	ClassificationMode string `json:"classificationMode"`
}

// UnmarshalJSON unmarshals json
func (m *classificationtype) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclassificationtype classificationtype
	s := struct {
		Model Unmarshalerclassificationtype
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ClassificationMode = s.Model.ClassificationMode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *classificationtype) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ClassificationMode {
	case "MULTI_CLASS":
		mm := ClassificationMultiClassModeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MULTI_LABEL":
		mm := ClassificationMultiLabelModeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ClassificationType: %s.", m.ClassificationMode)
		return *m, nil
	}
}

func (m classificationtype) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m classificationtype) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClassificationTypeClassificationModeEnum Enum with underlying type: string
type ClassificationTypeClassificationModeEnum string

// Set of constants representing the allowable values for ClassificationTypeClassificationModeEnum
const (
	ClassificationTypeClassificationModeClass ClassificationTypeClassificationModeEnum = "MULTI_CLASS"
	ClassificationTypeClassificationModeLabel ClassificationTypeClassificationModeEnum = "MULTI_LABEL"
)

var mappingClassificationTypeClassificationModeEnum = map[string]ClassificationTypeClassificationModeEnum{
	"MULTI_CLASS": ClassificationTypeClassificationModeClass,
	"MULTI_LABEL": ClassificationTypeClassificationModeLabel,
}

var mappingClassificationTypeClassificationModeEnumLowerCase = map[string]ClassificationTypeClassificationModeEnum{
	"multi_class": ClassificationTypeClassificationModeClass,
	"multi_label": ClassificationTypeClassificationModeLabel,
}

// GetClassificationTypeClassificationModeEnumValues Enumerates the set of values for ClassificationTypeClassificationModeEnum
func GetClassificationTypeClassificationModeEnumValues() []ClassificationTypeClassificationModeEnum {
	values := make([]ClassificationTypeClassificationModeEnum, 0)
	for _, v := range mappingClassificationTypeClassificationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetClassificationTypeClassificationModeEnumStringValues Enumerates the set of values in String for ClassificationTypeClassificationModeEnum
func GetClassificationTypeClassificationModeEnumStringValues() []string {
	return []string{
		"MULTI_CLASS",
		"MULTI_LABEL",
	}
}

// GetMappingClassificationTypeClassificationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassificationTypeClassificationModeEnum(val string) (ClassificationTypeClassificationModeEnum, bool) {
	enum, ok := mappingClassificationTypeClassificationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
