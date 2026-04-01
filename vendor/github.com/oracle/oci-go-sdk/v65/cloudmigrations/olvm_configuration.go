// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmConfiguration Configuration of the OLVM virtual machine.
type OlvmConfiguration struct {

	// The document describing the virtual machine.
	Data *string `mandatory:"false" json:"data"`

	// Configuration format types.
	ConfigurationType OlvmConfigurationConfigurationTypeEnum `mandatory:"false" json:"configurationType,omitempty"`
}

func (m OlvmConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmConfigurationConfigurationTypeEnum(string(m.ConfigurationType)); !ok && m.ConfigurationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigurationType: %s. Supported values are: %s.", m.ConfigurationType, strings.Join(GetOlvmConfigurationConfigurationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmConfigurationConfigurationTypeEnum Enum with underlying type: string
type OlvmConfigurationConfigurationTypeEnum string

// Set of constants representing the allowable values for OlvmConfigurationConfigurationTypeEnum
const (
	OlvmConfigurationConfigurationTypeOva OlvmConfigurationConfigurationTypeEnum = "OVA"
	OlvmConfigurationConfigurationTypeOvf OlvmConfigurationConfigurationTypeEnum = "OVF"
)

var mappingOlvmConfigurationConfigurationTypeEnum = map[string]OlvmConfigurationConfigurationTypeEnum{
	"OVA": OlvmConfigurationConfigurationTypeOva,
	"OVF": OlvmConfigurationConfigurationTypeOvf,
}

var mappingOlvmConfigurationConfigurationTypeEnumLowerCase = map[string]OlvmConfigurationConfigurationTypeEnum{
	"ova": OlvmConfigurationConfigurationTypeOva,
	"ovf": OlvmConfigurationConfigurationTypeOvf,
}

// GetOlvmConfigurationConfigurationTypeEnumValues Enumerates the set of values for OlvmConfigurationConfigurationTypeEnum
func GetOlvmConfigurationConfigurationTypeEnumValues() []OlvmConfigurationConfigurationTypeEnum {
	values := make([]OlvmConfigurationConfigurationTypeEnum, 0)
	for _, v := range mappingOlvmConfigurationConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmConfigurationConfigurationTypeEnumStringValues Enumerates the set of values in String for OlvmConfigurationConfigurationTypeEnum
func GetOlvmConfigurationConfigurationTypeEnumStringValues() []string {
	return []string{
		"OVA",
		"OVF",
	}
}

// GetMappingOlvmConfigurationConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmConfigurationConfigurationTypeEnum(val string) (OlvmConfigurationConfigurationTypeEnum, bool) {
	enum, ok := mappingOlvmConfigurationConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
