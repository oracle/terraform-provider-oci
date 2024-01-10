// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// HydratedVolume Description of the hydration server volume.
type HydratedVolume struct {

	// ID of the vCenter disk obtained from Inventory.
	Uuid *string `mandatory:"true" json:"uuid"`

	// ID of the hydration server volume
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// The hydration server volume type
	VolumeType HydratedVolumeVolumeTypeEnum `mandatory:"true" json:"volumeType"`

	// ID of the unmodified volume
	UnmodifiedVolumeId *string `mandatory:"true" json:"unmodifiedVolumeId"`
}

func (m HydratedVolume) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HydratedVolume) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHydratedVolumeVolumeTypeEnum(string(m.VolumeType)); !ok && m.VolumeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VolumeType: %s. Supported values are: %s.", m.VolumeType, strings.Join(GetHydratedVolumeVolumeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HydratedVolumeVolumeTypeEnum Enum with underlying type: string
type HydratedVolumeVolumeTypeEnum string

// Set of constants representing the allowable values for HydratedVolumeVolumeTypeEnum
const (
	HydratedVolumeVolumeTypeBoot  HydratedVolumeVolumeTypeEnum = "BOOT"
	HydratedVolumeVolumeTypeBlock HydratedVolumeVolumeTypeEnum = "BLOCK"
)

var mappingHydratedVolumeVolumeTypeEnum = map[string]HydratedVolumeVolumeTypeEnum{
	"BOOT":  HydratedVolumeVolumeTypeBoot,
	"BLOCK": HydratedVolumeVolumeTypeBlock,
}

var mappingHydratedVolumeVolumeTypeEnumLowerCase = map[string]HydratedVolumeVolumeTypeEnum{
	"boot":  HydratedVolumeVolumeTypeBoot,
	"block": HydratedVolumeVolumeTypeBlock,
}

// GetHydratedVolumeVolumeTypeEnumValues Enumerates the set of values for HydratedVolumeVolumeTypeEnum
func GetHydratedVolumeVolumeTypeEnumValues() []HydratedVolumeVolumeTypeEnum {
	values := make([]HydratedVolumeVolumeTypeEnum, 0)
	for _, v := range mappingHydratedVolumeVolumeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHydratedVolumeVolumeTypeEnumStringValues Enumerates the set of values in String for HydratedVolumeVolumeTypeEnum
func GetHydratedVolumeVolumeTypeEnumStringValues() []string {
	return []string{
		"BOOT",
		"BLOCK",
	}
}

// GetMappingHydratedVolumeVolumeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHydratedVolumeVolumeTypeEnum(val string) (HydratedVolumeVolumeTypeEnum, bool) {
	enum, ok := mappingHydratedVolumeVolumeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
