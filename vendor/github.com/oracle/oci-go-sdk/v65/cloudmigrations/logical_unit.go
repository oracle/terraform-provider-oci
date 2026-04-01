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

// LogicalUnit Logical unit of the host storage.
type LogicalUnit struct {

	// Address of the logical unit.
	Address *string `mandatory:"false" json:"address"`

	// The maximum number of bytes that can be discarded by the logical unit???s underlying storage in a single operation.
	DiscardMaxSizeInBytes *int64 `mandatory:"false" json:"discardMaxSizeInBytes"`

	// True, if previously discarded blocks in the logical unit???s underlying storage are read back as zeros.
	IsDiscardZeroesData *bool `mandatory:"false" json:"isDiscardZeroesData"`

	// Disk ID of the logical unit.
	DiskId *string `mandatory:"false" json:"diskId"`

	// a unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// LUN Mapping
	LunMapping *int `mandatory:"false" json:"lunMapping"`

	// LUN Status
	LunStatus LogicalUnitLunStatusEnum `mandatory:"false" json:"lunStatus,omitempty"`

	// Password of the logical unit.
	Password *string `mandatory:"false" json:"password"`

	// Paths of the logical unit.
	Paths *int `mandatory:"false" json:"paths"`

	// Port of the logical unit.
	Port *int `mandatory:"false" json:"port"`

	// Portal of the logical unit.
	Portal *string `mandatory:"false" json:"portal"`

	// Product ID of the logical unit.
	ProductId *string `mandatory:"false" json:"productId"`

	// Serial of the logical unit.
	Serial *string `mandatory:"false" json:"serial"`

	// Size of the logical unit.
	Size *int `mandatory:"false" json:"size"`

	// Storage Domain ID of the logical unit.
	StorageDomainId *string `mandatory:"false" json:"storageDomainId"`

	// Target of the logical unit.
	Target *string `mandatory:"false" json:"target"`

	// Username of the logical unit.
	Username *string `mandatory:"false" json:"username"`

	// Vendor ID of the logical unit.
	VendorId *string `mandatory:"false" json:"vendorId"`

	// Volume Group ID of the logical unit.
	VolumeGroupId *string `mandatory:"false" json:"volumeGroupId"`
}

func (m LogicalUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogicalUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogicalUnitLunStatusEnum(string(m.LunStatus)); !ok && m.LunStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LunStatus: %s. Supported values are: %s.", m.LunStatus, strings.Join(GetLogicalUnitLunStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogicalUnitLunStatusEnum Enum with underlying type: string
type LogicalUnitLunStatusEnum string

// Set of constants representing the allowable values for LogicalUnitLunStatusEnum
const (
	LogicalUnitLunStatusFree     LogicalUnitLunStatusEnum = "FREE"
	LogicalUnitLunStatusUnusable LogicalUnitLunStatusEnum = "UNUSABLE"
	LogicalUnitLunStatusUsed     LogicalUnitLunStatusEnum = "USED"
)

var mappingLogicalUnitLunStatusEnum = map[string]LogicalUnitLunStatusEnum{
	"FREE":     LogicalUnitLunStatusFree,
	"UNUSABLE": LogicalUnitLunStatusUnusable,
	"USED":     LogicalUnitLunStatusUsed,
}

var mappingLogicalUnitLunStatusEnumLowerCase = map[string]LogicalUnitLunStatusEnum{
	"free":     LogicalUnitLunStatusFree,
	"unusable": LogicalUnitLunStatusUnusable,
	"used":     LogicalUnitLunStatusUsed,
}

// GetLogicalUnitLunStatusEnumValues Enumerates the set of values for LogicalUnitLunStatusEnum
func GetLogicalUnitLunStatusEnumValues() []LogicalUnitLunStatusEnum {
	values := make([]LogicalUnitLunStatusEnum, 0)
	for _, v := range mappingLogicalUnitLunStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogicalUnitLunStatusEnumStringValues Enumerates the set of values in String for LogicalUnitLunStatusEnum
func GetLogicalUnitLunStatusEnumStringValues() []string {
	return []string{
		"FREE",
		"UNUSABLE",
		"USED",
	}
}

// GetMappingLogicalUnitLunStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogicalUnitLunStatusEnum(val string) (LogicalUnitLunStatusEnum, bool) {
	enum, ok := mappingLogicalUnitLunStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
