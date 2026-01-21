// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbConnectionBundle Details of a database connection bundle.
type DbConnectionBundle struct {

	// The OCID of the database connection bundle.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the database connection bundle.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name for the connection bundle.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the database connection bundle was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the database connection bundle was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// True for the default, service-created Database Connection Bundle.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]string `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]string `mandatory:"true" json:"systemTags"`

	// The current lifecycle state of the database connection bundle.
	LifecycleState DbConnectionBundleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the database connection bundle.
	DbConnectionBundleType DbConnectionBundleDbConnectionBundleTypeEnum `mandatory:"true" json:"dbConnectionBundleType"`

	// The time the database connection bundle was last refreshed. An RFC3339 formatted datetime string.
	TimeLastRefreshed *common.SDKTime `mandatory:"false" json:"timeLastRefreshed"`

	// Details about the resources associated with the connection bundle.
	AssociatedResourceDetails []AssociatedResourceDetails `mandatory:"false" json:"associatedResourceDetails"`
}

func (m DbConnectionBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbConnectionBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbConnectionBundleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbConnectionBundleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbConnectionBundleDbConnectionBundleTypeEnum(string(m.DbConnectionBundleType)); !ok && m.DbConnectionBundleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbConnectionBundleType: %s. Supported values are: %s.", m.DbConnectionBundleType, strings.Join(GetDbConnectionBundleDbConnectionBundleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbConnectionBundleLifecycleStateEnum Enum with underlying type: string
type DbConnectionBundleLifecycleStateEnum string

// Set of constants representing the allowable values for DbConnectionBundleLifecycleStateEnum
const (
	DbConnectionBundleLifecycleStateCreating DbConnectionBundleLifecycleStateEnum = "CREATING"
	DbConnectionBundleLifecycleStateActive   DbConnectionBundleLifecycleStateEnum = "ACTIVE"
	DbConnectionBundleLifecycleStateInactive DbConnectionBundleLifecycleStateEnum = "INACTIVE"
	DbConnectionBundleLifecycleStateUpdating DbConnectionBundleLifecycleStateEnum = "UPDATING"
	DbConnectionBundleLifecycleStateDeleting DbConnectionBundleLifecycleStateEnum = "DELETING"
	DbConnectionBundleLifecycleStateDeleted  DbConnectionBundleLifecycleStateEnum = "DELETED"
	DbConnectionBundleLifecycleStateFailed   DbConnectionBundleLifecycleStateEnum = "FAILED"
)

var mappingDbConnectionBundleLifecycleStateEnum = map[string]DbConnectionBundleLifecycleStateEnum{
	"CREATING": DbConnectionBundleLifecycleStateCreating,
	"ACTIVE":   DbConnectionBundleLifecycleStateActive,
	"INACTIVE": DbConnectionBundleLifecycleStateInactive,
	"UPDATING": DbConnectionBundleLifecycleStateUpdating,
	"DELETING": DbConnectionBundleLifecycleStateDeleting,
	"DELETED":  DbConnectionBundleLifecycleStateDeleted,
	"FAILED":   DbConnectionBundleLifecycleStateFailed,
}

var mappingDbConnectionBundleLifecycleStateEnumLowerCase = map[string]DbConnectionBundleLifecycleStateEnum{
	"creating": DbConnectionBundleLifecycleStateCreating,
	"active":   DbConnectionBundleLifecycleStateActive,
	"inactive": DbConnectionBundleLifecycleStateInactive,
	"updating": DbConnectionBundleLifecycleStateUpdating,
	"deleting": DbConnectionBundleLifecycleStateDeleting,
	"deleted":  DbConnectionBundleLifecycleStateDeleted,
	"failed":   DbConnectionBundleLifecycleStateFailed,
}

// GetDbConnectionBundleLifecycleStateEnumValues Enumerates the set of values for DbConnectionBundleLifecycleStateEnum
func GetDbConnectionBundleLifecycleStateEnumValues() []DbConnectionBundleLifecycleStateEnum {
	values := make([]DbConnectionBundleLifecycleStateEnum, 0)
	for _, v := range mappingDbConnectionBundleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbConnectionBundleLifecycleStateEnumStringValues Enumerates the set of values in String for DbConnectionBundleLifecycleStateEnum
func GetDbConnectionBundleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbConnectionBundleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbConnectionBundleLifecycleStateEnum(val string) (DbConnectionBundleLifecycleStateEnum, bool) {
	enum, ok := mappingDbConnectionBundleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbConnectionBundleDbConnectionBundleTypeEnum Enum with underlying type: string
type DbConnectionBundleDbConnectionBundleTypeEnum string

// Set of constants representing the allowable values for DbConnectionBundleDbConnectionBundleTypeEnum
const (
	DbConnectionBundleDbConnectionBundleTypeTls  DbConnectionBundleDbConnectionBundleTypeEnum = "TLS"
	DbConnectionBundleDbConnectionBundleTypeMtls DbConnectionBundleDbConnectionBundleTypeEnum = "MTLS"
)

var mappingDbConnectionBundleDbConnectionBundleTypeEnum = map[string]DbConnectionBundleDbConnectionBundleTypeEnum{
	"TLS":  DbConnectionBundleDbConnectionBundleTypeTls,
	"MTLS": DbConnectionBundleDbConnectionBundleTypeMtls,
}

var mappingDbConnectionBundleDbConnectionBundleTypeEnumLowerCase = map[string]DbConnectionBundleDbConnectionBundleTypeEnum{
	"tls":  DbConnectionBundleDbConnectionBundleTypeTls,
	"mtls": DbConnectionBundleDbConnectionBundleTypeMtls,
}

// GetDbConnectionBundleDbConnectionBundleTypeEnumValues Enumerates the set of values for DbConnectionBundleDbConnectionBundleTypeEnum
func GetDbConnectionBundleDbConnectionBundleTypeEnumValues() []DbConnectionBundleDbConnectionBundleTypeEnum {
	values := make([]DbConnectionBundleDbConnectionBundleTypeEnum, 0)
	for _, v := range mappingDbConnectionBundleDbConnectionBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbConnectionBundleDbConnectionBundleTypeEnumStringValues Enumerates the set of values in String for DbConnectionBundleDbConnectionBundleTypeEnum
func GetDbConnectionBundleDbConnectionBundleTypeEnumStringValues() []string {
	return []string{
		"TLS",
		"MTLS",
	}
}

// GetMappingDbConnectionBundleDbConnectionBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbConnectionBundleDbConnectionBundleTypeEnum(val string) (DbConnectionBundleDbConnectionBundleTypeEnum, bool) {
	enum, ok := mappingDbConnectionBundleDbConnectionBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
