// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseTagDefinitionValidator Validates a definedTag value. Each validator performs validation steps in addition to the standard
// validation for definedTag values. For more information, see
// Limits on Tags (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm#limits).
// If you define a validator after a value has been set for a defined tag, then any updates that
// attempt to change the value must pass the additional validation defined by the current rule.
// Previously set values (even those that would fail the current validation) are not updated. You can
// still update other attributes to resources that contain a non-valid defined tag.
// To clear the validator call UpdateTag with
// DefaultTagDefinitionValidator (https://docs.oracle.com/iaas/api/#/en/identity/latest/datatypes/DefaultTagDefinitionValidator).
type BaseTagDefinitionValidator interface {
}

type basetagdefinitionvalidator struct {
	JsonData      []byte
	ValidatorType string `json:"validatorType"`
}

// UnmarshalJSON unmarshals json
func (m *basetagdefinitionvalidator) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasetagdefinitionvalidator basetagdefinitionvalidator
	s := struct {
		Model Unmarshalerbasetagdefinitionvalidator
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValidatorType = s.Model.ValidatorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basetagdefinitionvalidator) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValidatorType {
	case "DEFAULT":
		mm := DefaultTagDefinitionValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENUM":
		mm := EnumTagDefinitionValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BaseTagDefinitionValidator: %s.", m.ValidatorType)
		return *m, nil
	}
}

func (m basetagdefinitionvalidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basetagdefinitionvalidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseTagDefinitionValidatorValidatorTypeEnum Enum with underlying type: string
type BaseTagDefinitionValidatorValidatorTypeEnum string

// Set of constants representing the allowable values for BaseTagDefinitionValidatorValidatorTypeEnum
const (
	BaseTagDefinitionValidatorValidatorTypeEnumvalue BaseTagDefinitionValidatorValidatorTypeEnum = "ENUM"
	BaseTagDefinitionValidatorValidatorTypeDefault   BaseTagDefinitionValidatorValidatorTypeEnum = "DEFAULT"
)

var mappingBaseTagDefinitionValidatorValidatorTypeEnum = map[string]BaseTagDefinitionValidatorValidatorTypeEnum{
	"ENUM":    BaseTagDefinitionValidatorValidatorTypeEnumvalue,
	"DEFAULT": BaseTagDefinitionValidatorValidatorTypeDefault,
}

var mappingBaseTagDefinitionValidatorValidatorTypeEnumLowerCase = map[string]BaseTagDefinitionValidatorValidatorTypeEnum{
	"enum":    BaseTagDefinitionValidatorValidatorTypeEnumvalue,
	"default": BaseTagDefinitionValidatorValidatorTypeDefault,
}

// GetBaseTagDefinitionValidatorValidatorTypeEnumValues Enumerates the set of values for BaseTagDefinitionValidatorValidatorTypeEnum
func GetBaseTagDefinitionValidatorValidatorTypeEnumValues() []BaseTagDefinitionValidatorValidatorTypeEnum {
	values := make([]BaseTagDefinitionValidatorValidatorTypeEnum, 0)
	for _, v := range mappingBaseTagDefinitionValidatorValidatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseTagDefinitionValidatorValidatorTypeEnumStringValues Enumerates the set of values in String for BaseTagDefinitionValidatorValidatorTypeEnum
func GetBaseTagDefinitionValidatorValidatorTypeEnumStringValues() []string {
	return []string{
		"ENUM",
		"DEFAULT",
	}
}

// GetMappingBaseTagDefinitionValidatorValidatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseTagDefinitionValidatorValidatorTypeEnum(val string) (BaseTagDefinitionValidatorValidatorTypeEnum, bool) {
	enum, ok := mappingBaseTagDefinitionValidatorValidatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
