// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SetEncryptionStatusDetails This structure is used to update encryption status of a VCN or a region.
type SetEncryptionStatusDetails struct {

	// Indicates the encryption status of the entity.
	IsEncryptionEnabled *bool `mandatory:"false" json:"isEncryptionEnabled"`

	// OCID of a VCN or null for the region.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// Type of the entity.
	EntityType SetEncryptionStatusDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`

	// key rotation interval which must be provided in minutes.
	KeyRotationPeriodInMin *int64 `mandatory:"false" json:"keyRotationPeriodInMin"`

	// Indicates if existing keys must be rotated. Defaults to false.
	IsKeyRotationForced *bool `mandatory:"false" json:"isKeyRotationForced"`

	// key generation interval which is provided in minutes.
	KeyGenerationIntervalInMin *int64 `mandatory:"false" json:"keyGenerationIntervalInMin"`

	// max minutes of keys in future.
	MaxMinuteOfKeysInFuture *int64 `mandatory:"false" json:"maxMinuteOfKeysInFuture"`

	// initial key usage delay which is provided in minutes.
	InitialKeyUsageDelayInMin *int64 `mandatory:"false" json:"initialKeyUsageDelayInMin"`

	// maximum minutes to keep for force rotation.
	MaxMinuteToKeepForForceRotation *int64 `mandatory:"false" json:"maxMinuteToKeepForForceRotation"`
}

func (m SetEncryptionStatusDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SetEncryptionStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSetEncryptionStatusDetailsEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetSetEncryptionStatusDetailsEntityTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SetEncryptionStatusDetailsEntityTypeEnum Enum with underlying type: string
type SetEncryptionStatusDetailsEntityTypeEnum string

// Set of constants representing the allowable values for SetEncryptionStatusDetailsEntityTypeEnum
const (
	SetEncryptionStatusDetailsEntityTypeVcn    SetEncryptionStatusDetailsEntityTypeEnum = "VCN"
	SetEncryptionStatusDetailsEntityTypeRegion SetEncryptionStatusDetailsEntityTypeEnum = "REGION"
)

var mappingSetEncryptionStatusDetailsEntityTypeEnum = map[string]SetEncryptionStatusDetailsEntityTypeEnum{
	"VCN":    SetEncryptionStatusDetailsEntityTypeVcn,
	"REGION": SetEncryptionStatusDetailsEntityTypeRegion,
}

var mappingSetEncryptionStatusDetailsEntityTypeEnumLowerCase = map[string]SetEncryptionStatusDetailsEntityTypeEnum{
	"vcn":    SetEncryptionStatusDetailsEntityTypeVcn,
	"region": SetEncryptionStatusDetailsEntityTypeRegion,
}

// GetSetEncryptionStatusDetailsEntityTypeEnumValues Enumerates the set of values for SetEncryptionStatusDetailsEntityTypeEnum
func GetSetEncryptionStatusDetailsEntityTypeEnumValues() []SetEncryptionStatusDetailsEntityTypeEnum {
	values := make([]SetEncryptionStatusDetailsEntityTypeEnum, 0)
	for _, v := range mappingSetEncryptionStatusDetailsEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSetEncryptionStatusDetailsEntityTypeEnumStringValues Enumerates the set of values in String for SetEncryptionStatusDetailsEntityTypeEnum
func GetSetEncryptionStatusDetailsEntityTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"REGION",
	}
}

// GetMappingSetEncryptionStatusDetailsEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSetEncryptionStatusDetailsEntityTypeEnum(val string) (SetEncryptionStatusDetailsEntityTypeEnum, bool) {
	enum, ok := mappingSetEncryptionStatusDetailsEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
