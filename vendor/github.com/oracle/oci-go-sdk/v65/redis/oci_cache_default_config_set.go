// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCacheDefaultConfigSet Default configurations for OCI Cache to manage the behavior, performance, and functionality of the underlying cache engine.
type OciCacheDefaultConfigSet struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the OCI Cache Default Config Set.
	Id *string `mandatory:"true" json:"id"`

	// The engine version of the OCI Cache Default Config Set.
	SoftwareVersion OciCacheConfigSetSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// A user-friendly name of the OCI Cache Default Config Set.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the OCI Cache Default Config Set.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the OCI Cache Default Config Set was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the OCI Cache Default Config Set.
	LifecycleState OciCacheDefaultConfigSetLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	DefaultConfigurationDetails *DefaultConfigurationDetails `mandatory:"false" json:"defaultConfigurationDetails"`
}

func (m OciCacheDefaultConfigSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheDefaultConfigSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheConfigSetSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetOciCacheConfigSetSoftwareVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOciCacheDefaultConfigSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheDefaultConfigSetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciCacheDefaultConfigSetLifecycleStateEnum Enum with underlying type: string
type OciCacheDefaultConfigSetLifecycleStateEnum string

// Set of constants representing the allowable values for OciCacheDefaultConfigSetLifecycleStateEnum
const (
	OciCacheDefaultConfigSetLifecycleStateActive   OciCacheDefaultConfigSetLifecycleStateEnum = "ACTIVE"
	OciCacheDefaultConfigSetLifecycleStateInactive OciCacheDefaultConfigSetLifecycleStateEnum = "INACTIVE"
)

var mappingOciCacheDefaultConfigSetLifecycleStateEnum = map[string]OciCacheDefaultConfigSetLifecycleStateEnum{
	"ACTIVE":   OciCacheDefaultConfigSetLifecycleStateActive,
	"INACTIVE": OciCacheDefaultConfigSetLifecycleStateInactive,
}

var mappingOciCacheDefaultConfigSetLifecycleStateEnumLowerCase = map[string]OciCacheDefaultConfigSetLifecycleStateEnum{
	"active":   OciCacheDefaultConfigSetLifecycleStateActive,
	"inactive": OciCacheDefaultConfigSetLifecycleStateInactive,
}

// GetOciCacheDefaultConfigSetLifecycleStateEnumValues Enumerates the set of values for OciCacheDefaultConfigSetLifecycleStateEnum
func GetOciCacheDefaultConfigSetLifecycleStateEnumValues() []OciCacheDefaultConfigSetLifecycleStateEnum {
	values := make([]OciCacheDefaultConfigSetLifecycleStateEnum, 0)
	for _, v := range mappingOciCacheDefaultConfigSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheDefaultConfigSetLifecycleStateEnumStringValues Enumerates the set of values in String for OciCacheDefaultConfigSetLifecycleStateEnum
func GetOciCacheDefaultConfigSetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingOciCacheDefaultConfigSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheDefaultConfigSetLifecycleStateEnum(val string) (OciCacheDefaultConfigSetLifecycleStateEnum, bool) {
	enum, ok := mappingOciCacheDefaultConfigSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
