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

// OciCacheConfigSet Configurations for OCI Cache to manage the behavior, performance, and functionality of the underlying cache engine.
type OciCacheConfigSet struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the OCI Cache Config Set.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the OCI Cache Config Set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the OCI Cache Config Set.
	LifecycleState OciCacheConfigSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCI Cache engine version that the cluster is running.
	SoftwareVersion OciCacheConfigSetSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the default OCI Cache Config Set which the custom OCI Cache Config Set is based upon.
	DefaultConfigSetId *string `mandatory:"false" json:"defaultConfigSetId"`

	// A description of the OCI Cache Config Set.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the OCI Cache Config Set was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the OCI Cache Config Set was updated. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"false" json:"configurationDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OciCacheConfigSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheConfigSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheConfigSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheConfigSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheConfigSetSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetOciCacheConfigSetSoftwareVersionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciCacheConfigSetLifecycleStateEnum Enum with underlying type: string
type OciCacheConfigSetLifecycleStateEnum string

// Set of constants representing the allowable values for OciCacheConfigSetLifecycleStateEnum
const (
	OciCacheConfigSetLifecycleStateCreating OciCacheConfigSetLifecycleStateEnum = "CREATING"
	OciCacheConfigSetLifecycleStateUpdating OciCacheConfigSetLifecycleStateEnum = "UPDATING"
	OciCacheConfigSetLifecycleStateActive   OciCacheConfigSetLifecycleStateEnum = "ACTIVE"
	OciCacheConfigSetLifecycleStateDeleting OciCacheConfigSetLifecycleStateEnum = "DELETING"
	OciCacheConfigSetLifecycleStateDeleted  OciCacheConfigSetLifecycleStateEnum = "DELETED"
	OciCacheConfigSetLifecycleStateFailed   OciCacheConfigSetLifecycleStateEnum = "FAILED"
)

var mappingOciCacheConfigSetLifecycleStateEnum = map[string]OciCacheConfigSetLifecycleStateEnum{
	"CREATING": OciCacheConfigSetLifecycleStateCreating,
	"UPDATING": OciCacheConfigSetLifecycleStateUpdating,
	"ACTIVE":   OciCacheConfigSetLifecycleStateActive,
	"DELETING": OciCacheConfigSetLifecycleStateDeleting,
	"DELETED":  OciCacheConfigSetLifecycleStateDeleted,
	"FAILED":   OciCacheConfigSetLifecycleStateFailed,
}

var mappingOciCacheConfigSetLifecycleStateEnumLowerCase = map[string]OciCacheConfigSetLifecycleStateEnum{
	"creating": OciCacheConfigSetLifecycleStateCreating,
	"updating": OciCacheConfigSetLifecycleStateUpdating,
	"active":   OciCacheConfigSetLifecycleStateActive,
	"deleting": OciCacheConfigSetLifecycleStateDeleting,
	"deleted":  OciCacheConfigSetLifecycleStateDeleted,
	"failed":   OciCacheConfigSetLifecycleStateFailed,
}

// GetOciCacheConfigSetLifecycleStateEnumValues Enumerates the set of values for OciCacheConfigSetLifecycleStateEnum
func GetOciCacheConfigSetLifecycleStateEnumValues() []OciCacheConfigSetLifecycleStateEnum {
	values := make([]OciCacheConfigSetLifecycleStateEnum, 0)
	for _, v := range mappingOciCacheConfigSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheConfigSetLifecycleStateEnumStringValues Enumerates the set of values in String for OciCacheConfigSetLifecycleStateEnum
func GetOciCacheConfigSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOciCacheConfigSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheConfigSetLifecycleStateEnum(val string) (OciCacheConfigSetLifecycleStateEnum, bool) {
	enum, ok := mappingOciCacheConfigSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OciCacheConfigSetSoftwareVersionEnum Enum with underlying type: string
type OciCacheConfigSetSoftwareVersionEnum string

// Set of constants representing the allowable values for OciCacheConfigSetSoftwareVersionEnum
const (
	OciCacheConfigSetSoftwareVersionV705     OciCacheConfigSetSoftwareVersionEnum = "V7_0_5"
	OciCacheConfigSetSoftwareVersionRedis70  OciCacheConfigSetSoftwareVersionEnum = "REDIS_7_0"
	OciCacheConfigSetSoftwareVersionValkey72 OciCacheConfigSetSoftwareVersionEnum = "VALKEY_7_2"
)

var mappingOciCacheConfigSetSoftwareVersionEnum = map[string]OciCacheConfigSetSoftwareVersionEnum{
	"V7_0_5":     OciCacheConfigSetSoftwareVersionV705,
	"REDIS_7_0":  OciCacheConfigSetSoftwareVersionRedis70,
	"VALKEY_7_2": OciCacheConfigSetSoftwareVersionValkey72,
}

var mappingOciCacheConfigSetSoftwareVersionEnumLowerCase = map[string]OciCacheConfigSetSoftwareVersionEnum{
	"v7_0_5":     OciCacheConfigSetSoftwareVersionV705,
	"redis_7_0":  OciCacheConfigSetSoftwareVersionRedis70,
	"valkey_7_2": OciCacheConfigSetSoftwareVersionValkey72,
}

// GetOciCacheConfigSetSoftwareVersionEnumValues Enumerates the set of values for OciCacheConfigSetSoftwareVersionEnum
func GetOciCacheConfigSetSoftwareVersionEnumValues() []OciCacheConfigSetSoftwareVersionEnum {
	values := make([]OciCacheConfigSetSoftwareVersionEnum, 0)
	for _, v := range mappingOciCacheConfigSetSoftwareVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheConfigSetSoftwareVersionEnumStringValues Enumerates the set of values in String for OciCacheConfigSetSoftwareVersionEnum
func GetOciCacheConfigSetSoftwareVersionEnumStringValues() []string {
	return []string{
		"V7_0_5",
		"REDIS_7_0",
		"VALKEY_7_2",
	}
}

// GetMappingOciCacheConfigSetSoftwareVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheConfigSetSoftwareVersionEnum(val string) (OciCacheConfigSetSoftwareVersionEnum, bool) {
	enum, ok := mappingOciCacheConfigSetSoftwareVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
