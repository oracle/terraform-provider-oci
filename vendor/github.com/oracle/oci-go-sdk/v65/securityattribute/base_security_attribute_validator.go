// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseSecurityAttributeValidator Validates a security attribute value. Each validator performs validation steps in addition to the standard
// validation for security attribute values. For more information, see
// Limits on Security Attributes (https://docs.cloud.oracle.com/Content/zero-trust-packet-routing/overview.htm).
// If you define a validator after a value has been set for a security attribute, then any updates that
// attempt to change the value must pass the additional validation defined by the current rule.
// Previously set values (even those that would fail the current validation) are not updated. You can
// still update other attributes to resources that contain a non-valid security attribute.
// To clear the validator call UpdateSecurityAttribute with
// DefaultSecuirtyAttributeValidator (https://docs.cloud.oracle.com/api/#/en/securityattribute/latest/datatypes/DefaultTagDefinitionValidator).
type BaseSecurityAttributeValidator interface {
}

type basesecurityattributevalidator struct {
	JsonData      []byte
	ValidatorType string `json:"validatorType"`
}

// UnmarshalJSON unmarshals json
func (m *basesecurityattributevalidator) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasesecurityattributevalidator basesecurityattributevalidator
	s := struct {
		Model Unmarshalerbasesecurityattributevalidator
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValidatorType = s.Model.ValidatorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basesecurityattributevalidator) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValidatorType {
	case "ENUM":
		mm := EnumSecurityAttributeValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFAULT":
		mm := DefaultSecurityAttributeValidator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BaseSecurityAttributeValidator: %s.", m.ValidatorType)
		return *m, nil
	}
}

func (m basesecurityattributevalidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basesecurityattributevalidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseSecurityAttributeValidatorValidatorTypeEnum Enum with underlying type: string
type BaseSecurityAttributeValidatorValidatorTypeEnum string

// Set of constants representing the allowable values for BaseSecurityAttributeValidatorValidatorTypeEnum
const (
	BaseSecurityAttributeValidatorValidatorTypeEnumvalue BaseSecurityAttributeValidatorValidatorTypeEnum = "ENUM"
	BaseSecurityAttributeValidatorValidatorTypeDefault   BaseSecurityAttributeValidatorValidatorTypeEnum = "DEFAULT"
)

var mappingBaseSecurityAttributeValidatorValidatorTypeEnum = map[string]BaseSecurityAttributeValidatorValidatorTypeEnum{
	"ENUM":    BaseSecurityAttributeValidatorValidatorTypeEnumvalue,
	"DEFAULT": BaseSecurityAttributeValidatorValidatorTypeDefault,
}

var mappingBaseSecurityAttributeValidatorValidatorTypeEnumLowerCase = map[string]BaseSecurityAttributeValidatorValidatorTypeEnum{
	"enum":    BaseSecurityAttributeValidatorValidatorTypeEnumvalue,
	"default": BaseSecurityAttributeValidatorValidatorTypeDefault,
}

// GetBaseSecurityAttributeValidatorValidatorTypeEnumValues Enumerates the set of values for BaseSecurityAttributeValidatorValidatorTypeEnum
func GetBaseSecurityAttributeValidatorValidatorTypeEnumValues() []BaseSecurityAttributeValidatorValidatorTypeEnum {
	values := make([]BaseSecurityAttributeValidatorValidatorTypeEnum, 0)
	for _, v := range mappingBaseSecurityAttributeValidatorValidatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseSecurityAttributeValidatorValidatorTypeEnumStringValues Enumerates the set of values in String for BaseSecurityAttributeValidatorValidatorTypeEnum
func GetBaseSecurityAttributeValidatorValidatorTypeEnumStringValues() []string {
	return []string{
		"ENUM",
		"DEFAULT",
	}
}

// GetMappingBaseSecurityAttributeValidatorValidatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseSecurityAttributeValidatorValidatorTypeEnum(val string) (BaseSecurityAttributeValidatorValidatorTypeEnum, bool) {
	enum, ok := mappingBaseSecurityAttributeValidatorValidatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
