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

// OlvmBiosType Chipset and BIOS type combination.
type OlvmBiosType struct {

	// Chipset and BIOS type combination.
	Type OlvmBiosTypeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m OlvmBiosType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmBiosType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmBiosTypeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetOlvmBiosTypeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmBiosTypeTypeEnum Enum with underlying type: string
type OlvmBiosTypeTypeEnum string

// Set of constants representing the allowable values for OlvmBiosTypeTypeEnum
const (
	OlvmBiosTypeTypeClusterDefault OlvmBiosTypeTypeEnum = "CLUSTER_DEFAULT"
	OlvmBiosTypeTypeI440fxSeaBios  OlvmBiosTypeTypeEnum = "I440FX_SEA_BIOS"
	OlvmBiosTypeTypeQ35Ovmf        OlvmBiosTypeTypeEnum = "Q35_OVMF"
	OlvmBiosTypeTypeQ35SeaBios     OlvmBiosTypeTypeEnum = "Q35_SEA_BIOS"
	OlvmBiosTypeTypeQ35SecureBoot  OlvmBiosTypeTypeEnum = "Q35_SECURE_BOOT"
)

var mappingOlvmBiosTypeTypeEnum = map[string]OlvmBiosTypeTypeEnum{
	"CLUSTER_DEFAULT": OlvmBiosTypeTypeClusterDefault,
	"I440FX_SEA_BIOS": OlvmBiosTypeTypeI440fxSeaBios,
	"Q35_OVMF":        OlvmBiosTypeTypeQ35Ovmf,
	"Q35_SEA_BIOS":    OlvmBiosTypeTypeQ35SeaBios,
	"Q35_SECURE_BOOT": OlvmBiosTypeTypeQ35SecureBoot,
}

var mappingOlvmBiosTypeTypeEnumLowerCase = map[string]OlvmBiosTypeTypeEnum{
	"cluster_default": OlvmBiosTypeTypeClusterDefault,
	"i440fx_sea_bios": OlvmBiosTypeTypeI440fxSeaBios,
	"q35_ovmf":        OlvmBiosTypeTypeQ35Ovmf,
	"q35_sea_bios":    OlvmBiosTypeTypeQ35SeaBios,
	"q35_secure_boot": OlvmBiosTypeTypeQ35SecureBoot,
}

// GetOlvmBiosTypeTypeEnumValues Enumerates the set of values for OlvmBiosTypeTypeEnum
func GetOlvmBiosTypeTypeEnumValues() []OlvmBiosTypeTypeEnum {
	values := make([]OlvmBiosTypeTypeEnum, 0)
	for _, v := range mappingOlvmBiosTypeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmBiosTypeTypeEnumStringValues Enumerates the set of values in String for OlvmBiosTypeTypeEnum
func GetOlvmBiosTypeTypeEnumStringValues() []string {
	return []string{
		"CLUSTER_DEFAULT",
		"I440FX_SEA_BIOS",
		"Q35_OVMF",
		"Q35_SEA_BIOS",
		"Q35_SECURE_BOOT",
	}
}

// GetMappingOlvmBiosTypeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmBiosTypeTypeEnum(val string) (OlvmBiosTypeTypeEnum, bool) {
	enum, ok := mappingOlvmBiosTypeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
