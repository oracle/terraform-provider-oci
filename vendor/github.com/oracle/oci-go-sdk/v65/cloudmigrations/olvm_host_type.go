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

// OlvmHostType Indicates if the host contains a full installation of the operating system or a scaled-down version intended only to host virtual machines.
type OlvmHostType struct {

	// This enumerated type is used to determine which type of operating system is used by the host.
	HostType OlvmHostTypeHostTypeEnum `mandatory:"false" json:"hostType,omitempty"`
}

func (m OlvmHostType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmHostType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmHostTypeHostTypeEnum(string(m.HostType)); !ok && m.HostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HostType: %s. Supported values are: %s.", m.HostType, strings.Join(GetOlvmHostTypeHostTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmHostTypeHostTypeEnum Enum with underlying type: string
type OlvmHostTypeHostTypeEnum string

// Set of constants representing the allowable values for OlvmHostTypeHostTypeEnum
const (
	OlvmHostTypeHostTypeOvirtNode OlvmHostTypeHostTypeEnum = "OVIRT_NODE"
	OlvmHostTypeHostTypeRhel      OlvmHostTypeHostTypeEnum = "RHEL"
	OlvmHostTypeHostTypeRhevH     OlvmHostTypeHostTypeEnum = "RHEV_H"
)

var mappingOlvmHostTypeHostTypeEnum = map[string]OlvmHostTypeHostTypeEnum{
	"OVIRT_NODE": OlvmHostTypeHostTypeOvirtNode,
	"RHEL":       OlvmHostTypeHostTypeRhel,
	"RHEV_H":     OlvmHostTypeHostTypeRhevH,
}

var mappingOlvmHostTypeHostTypeEnumLowerCase = map[string]OlvmHostTypeHostTypeEnum{
	"ovirt_node": OlvmHostTypeHostTypeOvirtNode,
	"rhel":       OlvmHostTypeHostTypeRhel,
	"rhev_h":     OlvmHostTypeHostTypeRhevH,
}

// GetOlvmHostTypeHostTypeEnumValues Enumerates the set of values for OlvmHostTypeHostTypeEnum
func GetOlvmHostTypeHostTypeEnumValues() []OlvmHostTypeHostTypeEnum {
	values := make([]OlvmHostTypeHostTypeEnum, 0)
	for _, v := range mappingOlvmHostTypeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmHostTypeHostTypeEnumStringValues Enumerates the set of values in String for OlvmHostTypeHostTypeEnum
func GetOlvmHostTypeHostTypeEnumStringValues() []string {
	return []string{
		"OVIRT_NODE",
		"RHEL",
		"RHEV_H",
	}
}

// GetMappingOlvmHostTypeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmHostTypeHostTypeEnum(val string) (OlvmHostTypeHostTypeEnum, bool) {
	enum, ok := mappingOlvmHostTypeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
