// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MlApplicationPackageArgumentDetails Represents single argument name value pair.
type MlApplicationPackageArgumentDetails struct {

	// Argument name
	Name *string `mandatory:"true" json:"name"`

	// Argument value
	Value *string `mandatory:"true" json:"value"`

	// short description of the argument
	Description *string `mandatory:"true" json:"description"`

	// type of the argument
	Type MlApplicationPackageArgumentDetailsTypeEnum `mandatory:"true" json:"type"`

	// argument is mandatory or not
	IsMandatory *bool `mandatory:"true" json:"isMandatory"`
}

func (m MlApplicationPackageArgumentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationPackageArgumentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationPackageArgumentDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMlApplicationPackageArgumentDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MlApplicationPackageArgumentDetailsTypeEnum Enum with underlying type: string
type MlApplicationPackageArgumentDetailsTypeEnum string

// Set of constants representing the allowable values for MlApplicationPackageArgumentDetailsTypeEnum
const (
	MlApplicationPackageArgumentDetailsTypeString MlApplicationPackageArgumentDetailsTypeEnum = "STRING"
	MlApplicationPackageArgumentDetailsTypeOcid   MlApplicationPackageArgumentDetailsTypeEnum = "OCID"
)

var mappingMlApplicationPackageArgumentDetailsTypeEnum = map[string]MlApplicationPackageArgumentDetailsTypeEnum{
	"STRING": MlApplicationPackageArgumentDetailsTypeString,
	"OCID":   MlApplicationPackageArgumentDetailsTypeOcid,
}

var mappingMlApplicationPackageArgumentDetailsTypeEnumLowerCase = map[string]MlApplicationPackageArgumentDetailsTypeEnum{
	"string": MlApplicationPackageArgumentDetailsTypeString,
	"ocid":   MlApplicationPackageArgumentDetailsTypeOcid,
}

// GetMlApplicationPackageArgumentDetailsTypeEnumValues Enumerates the set of values for MlApplicationPackageArgumentDetailsTypeEnum
func GetMlApplicationPackageArgumentDetailsTypeEnumValues() []MlApplicationPackageArgumentDetailsTypeEnum {
	values := make([]MlApplicationPackageArgumentDetailsTypeEnum, 0)
	for _, v := range mappingMlApplicationPackageArgumentDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationPackageArgumentDetailsTypeEnumStringValues Enumerates the set of values in String for MlApplicationPackageArgumentDetailsTypeEnum
func GetMlApplicationPackageArgumentDetailsTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"OCID",
	}
}

// GetMappingMlApplicationPackageArgumentDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationPackageArgumentDetailsTypeEnum(val string) (MlApplicationPackageArgumentDetailsTypeEnum, bool) {
	enum, ok := mappingMlApplicationPackageArgumentDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
