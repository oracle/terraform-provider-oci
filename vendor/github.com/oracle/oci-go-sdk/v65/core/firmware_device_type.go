// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"strings"
)

// FirmwareDeviceTypeEnum Enum with underlying type: string
type FirmwareDeviceTypeEnum string

// Set of constants representing the allowable values for FirmwareDeviceTypeEnum
const (
	FirmwareDeviceTypeNic    FirmwareDeviceTypeEnum = "NIC"
	FirmwareDeviceTypeBios   FirmwareDeviceTypeEnum = "BIOS"
	FirmwareDeviceTypeNvme   FirmwareDeviceTypeEnum = "NVME"
	FirmwareDeviceTypeHba    FirmwareDeviceTypeEnum = "HBA"
	FirmwareDeviceTypeGpu    FirmwareDeviceTypeEnum = "GPU"
	FirmwareDeviceTypeSas    FirmwareDeviceTypeEnum = "SAS"
	FirmwareDeviceTypeSata   FirmwareDeviceTypeEnum = "SATA"
	FirmwareDeviceTypeJbod   FirmwareDeviceTypeEnum = "JBOD"
	FirmwareDeviceTypeHddSsd FirmwareDeviceTypeEnum = "HDD_SSD"
)

var mappingFirmwareDeviceTypeEnum = map[string]FirmwareDeviceTypeEnum{
	"NIC":     FirmwareDeviceTypeNic,
	"BIOS":    FirmwareDeviceTypeBios,
	"NVME":    FirmwareDeviceTypeNvme,
	"HBA":     FirmwareDeviceTypeHba,
	"GPU":     FirmwareDeviceTypeGpu,
	"SAS":     FirmwareDeviceTypeSas,
	"SATA":    FirmwareDeviceTypeSata,
	"JBOD":    FirmwareDeviceTypeJbod,
	"HDD_SSD": FirmwareDeviceTypeHddSsd,
}

var mappingFirmwareDeviceTypeEnumLowerCase = map[string]FirmwareDeviceTypeEnum{
	"nic":     FirmwareDeviceTypeNic,
	"bios":    FirmwareDeviceTypeBios,
	"nvme":    FirmwareDeviceTypeNvme,
	"hba":     FirmwareDeviceTypeHba,
	"gpu":     FirmwareDeviceTypeGpu,
	"sas":     FirmwareDeviceTypeSas,
	"sata":    FirmwareDeviceTypeSata,
	"jbod":    FirmwareDeviceTypeJbod,
	"hdd_ssd": FirmwareDeviceTypeHddSsd,
}

// GetFirmwareDeviceTypeEnumValues Enumerates the set of values for FirmwareDeviceTypeEnum
func GetFirmwareDeviceTypeEnumValues() []FirmwareDeviceTypeEnum {
	values := make([]FirmwareDeviceTypeEnum, 0)
	for _, v := range mappingFirmwareDeviceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFirmwareDeviceTypeEnumStringValues Enumerates the set of values in String for FirmwareDeviceTypeEnum
func GetFirmwareDeviceTypeEnumStringValues() []string {
	return []string{
		"NIC",
		"BIOS",
		"NVME",
		"HBA",
		"GPU",
		"SAS",
		"SATA",
		"JBOD",
		"HDD_SSD",
	}
}

// GetMappingFirmwareDeviceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirmwareDeviceTypeEnum(val string) (FirmwareDeviceTypeEnum, bool) {
	enum, ok := mappingFirmwareDeviceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
