// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BdsMetastoreConfiguration The metastore configuration information.
type BdsMetastoreConfiguration struct {

	// The ID of the metastore configuration
	Id *string `mandatory:"true" json:"id"`

	// The display name of metastore configuration
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the metastore in the metastore configuration.
	MetastoreType BdsMetastoreConfigurationMetastoreTypeEnum `mandatory:"true" json:"metastoreType"`

	// the lifecycle state of the metastore configuration.
	LifecycleState BdsMetastoreConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the Data Catalog metastore. Set only if metastore's type is EXTERNAL.
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// The ID of BDS API Key used for metastore configuration. Set only if metastore's type is EXTERNAL.
	BdsApiKeyId *string `mandatory:"false" json:"bdsApiKeyId"`

	// The time when the configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsMetastoreConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsMetastoreConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsMetastoreConfigurationMetastoreTypeEnum(string(m.MetastoreType)); !ok && m.MetastoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetastoreType: %s. Supported values are: %s.", m.MetastoreType, strings.Join(GetBdsMetastoreConfigurationMetastoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBdsMetastoreConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsMetastoreConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsMetastoreConfigurationMetastoreTypeEnum Enum with underlying type: string
type BdsMetastoreConfigurationMetastoreTypeEnum string

// Set of constants representing the allowable values for BdsMetastoreConfigurationMetastoreTypeEnum
const (
	BdsMetastoreConfigurationMetastoreTypeLocal    BdsMetastoreConfigurationMetastoreTypeEnum = "LOCAL"
	BdsMetastoreConfigurationMetastoreTypeExternal BdsMetastoreConfigurationMetastoreTypeEnum = "EXTERNAL"
)

var mappingBdsMetastoreConfigurationMetastoreTypeEnum = map[string]BdsMetastoreConfigurationMetastoreTypeEnum{
	"LOCAL":    BdsMetastoreConfigurationMetastoreTypeLocal,
	"EXTERNAL": BdsMetastoreConfigurationMetastoreTypeExternal,
}

var mappingBdsMetastoreConfigurationMetastoreTypeEnumLowerCase = map[string]BdsMetastoreConfigurationMetastoreTypeEnum{
	"local":    BdsMetastoreConfigurationMetastoreTypeLocal,
	"external": BdsMetastoreConfigurationMetastoreTypeExternal,
}

// GetBdsMetastoreConfigurationMetastoreTypeEnumValues Enumerates the set of values for BdsMetastoreConfigurationMetastoreTypeEnum
func GetBdsMetastoreConfigurationMetastoreTypeEnumValues() []BdsMetastoreConfigurationMetastoreTypeEnum {
	values := make([]BdsMetastoreConfigurationMetastoreTypeEnum, 0)
	for _, v := range mappingBdsMetastoreConfigurationMetastoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsMetastoreConfigurationMetastoreTypeEnumStringValues Enumerates the set of values in String for BdsMetastoreConfigurationMetastoreTypeEnum
func GetBdsMetastoreConfigurationMetastoreTypeEnumStringValues() []string {
	return []string{
		"LOCAL",
		"EXTERNAL",
	}
}

// GetMappingBdsMetastoreConfigurationMetastoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsMetastoreConfigurationMetastoreTypeEnum(val string) (BdsMetastoreConfigurationMetastoreTypeEnum, bool) {
	enum, ok := mappingBdsMetastoreConfigurationMetastoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BdsMetastoreConfigurationLifecycleStateEnum Enum with underlying type: string
type BdsMetastoreConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for BdsMetastoreConfigurationLifecycleStateEnum
const (
	BdsMetastoreConfigurationLifecycleStateCreating   BdsMetastoreConfigurationLifecycleStateEnum = "CREATING"
	BdsMetastoreConfigurationLifecycleStateActivating BdsMetastoreConfigurationLifecycleStateEnum = "ACTIVATING"
	BdsMetastoreConfigurationLifecycleStateActive     BdsMetastoreConfigurationLifecycleStateEnum = "ACTIVE"
	BdsMetastoreConfigurationLifecycleStateInactive   BdsMetastoreConfigurationLifecycleStateEnum = "INACTIVE"
	BdsMetastoreConfigurationLifecycleStateUpdating   BdsMetastoreConfigurationLifecycleStateEnum = "UPDATING"
	BdsMetastoreConfigurationLifecycleStateFailed     BdsMetastoreConfigurationLifecycleStateEnum = "FAILED"
	BdsMetastoreConfigurationLifecycleStateDeleting   BdsMetastoreConfigurationLifecycleStateEnum = "DELETING"
	BdsMetastoreConfigurationLifecycleStateDeleted    BdsMetastoreConfigurationLifecycleStateEnum = "DELETED"
)

var mappingBdsMetastoreConfigurationLifecycleStateEnum = map[string]BdsMetastoreConfigurationLifecycleStateEnum{
	"CREATING":   BdsMetastoreConfigurationLifecycleStateCreating,
	"ACTIVATING": BdsMetastoreConfigurationLifecycleStateActivating,
	"ACTIVE":     BdsMetastoreConfigurationLifecycleStateActive,
	"INACTIVE":   BdsMetastoreConfigurationLifecycleStateInactive,
	"UPDATING":   BdsMetastoreConfigurationLifecycleStateUpdating,
	"FAILED":     BdsMetastoreConfigurationLifecycleStateFailed,
	"DELETING":   BdsMetastoreConfigurationLifecycleStateDeleting,
	"DELETED":    BdsMetastoreConfigurationLifecycleStateDeleted,
}

var mappingBdsMetastoreConfigurationLifecycleStateEnumLowerCase = map[string]BdsMetastoreConfigurationLifecycleStateEnum{
	"creating":   BdsMetastoreConfigurationLifecycleStateCreating,
	"activating": BdsMetastoreConfigurationLifecycleStateActivating,
	"active":     BdsMetastoreConfigurationLifecycleStateActive,
	"inactive":   BdsMetastoreConfigurationLifecycleStateInactive,
	"updating":   BdsMetastoreConfigurationLifecycleStateUpdating,
	"failed":     BdsMetastoreConfigurationLifecycleStateFailed,
	"deleting":   BdsMetastoreConfigurationLifecycleStateDeleting,
	"deleted":    BdsMetastoreConfigurationLifecycleStateDeleted,
}

// GetBdsMetastoreConfigurationLifecycleStateEnumValues Enumerates the set of values for BdsMetastoreConfigurationLifecycleStateEnum
func GetBdsMetastoreConfigurationLifecycleStateEnumValues() []BdsMetastoreConfigurationLifecycleStateEnum {
	values := make([]BdsMetastoreConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingBdsMetastoreConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsMetastoreConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for BdsMetastoreConfigurationLifecycleStateEnum
func GetBdsMetastoreConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBdsMetastoreConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsMetastoreConfigurationLifecycleStateEnum(val string) (BdsMetastoreConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingBdsMetastoreConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
