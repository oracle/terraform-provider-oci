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

// DbConnectionBundleSummary Summary of a database connection bundle.
type DbConnectionBundleSummary struct {

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
	LifecycleState DbConnectionBundleSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the database connection bundle.
	DbConnectionBundleType DbConnectionBundleSummaryDbConnectionBundleTypeEnum `mandatory:"true" json:"dbConnectionBundleType"`

	// The time the database connection bundle was last refreshed. An RFC3339 formatted datetime string.
	TimeLastRefreshed *common.SDKTime `mandatory:"false" json:"timeLastRefreshed"`

	// Details about the resources associated with the connection bundle.
	AssociatedResourceDetails []AssociatedResourceDetails `mandatory:"false" json:"associatedResourceDetails"`
}

func (m DbConnectionBundleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbConnectionBundleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbConnectionBundleSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbConnectionBundleSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbConnectionBundleSummaryDbConnectionBundleTypeEnum(string(m.DbConnectionBundleType)); !ok && m.DbConnectionBundleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbConnectionBundleType: %s. Supported values are: %s.", m.DbConnectionBundleType, strings.Join(GetDbConnectionBundleSummaryDbConnectionBundleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbConnectionBundleSummaryLifecycleStateEnum Enum with underlying type: string
type DbConnectionBundleSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbConnectionBundleSummaryLifecycleStateEnum
const (
	DbConnectionBundleSummaryLifecycleStateCreating DbConnectionBundleSummaryLifecycleStateEnum = "CREATING"
	DbConnectionBundleSummaryLifecycleStateActive   DbConnectionBundleSummaryLifecycleStateEnum = "ACTIVE"
	DbConnectionBundleSummaryLifecycleStateInactive DbConnectionBundleSummaryLifecycleStateEnum = "INACTIVE"
	DbConnectionBundleSummaryLifecycleStateUpdating DbConnectionBundleSummaryLifecycleStateEnum = "UPDATING"
	DbConnectionBundleSummaryLifecycleStateDeleting DbConnectionBundleSummaryLifecycleStateEnum = "DELETING"
	DbConnectionBundleSummaryLifecycleStateDeleted  DbConnectionBundleSummaryLifecycleStateEnum = "DELETED"
	DbConnectionBundleSummaryLifecycleStateFailed   DbConnectionBundleSummaryLifecycleStateEnum = "FAILED"
)

var mappingDbConnectionBundleSummaryLifecycleStateEnum = map[string]DbConnectionBundleSummaryLifecycleStateEnum{
	"CREATING": DbConnectionBundleSummaryLifecycleStateCreating,
	"ACTIVE":   DbConnectionBundleSummaryLifecycleStateActive,
	"INACTIVE": DbConnectionBundleSummaryLifecycleStateInactive,
	"UPDATING": DbConnectionBundleSummaryLifecycleStateUpdating,
	"DELETING": DbConnectionBundleSummaryLifecycleStateDeleting,
	"DELETED":  DbConnectionBundleSummaryLifecycleStateDeleted,
	"FAILED":   DbConnectionBundleSummaryLifecycleStateFailed,
}

var mappingDbConnectionBundleSummaryLifecycleStateEnumLowerCase = map[string]DbConnectionBundleSummaryLifecycleStateEnum{
	"creating": DbConnectionBundleSummaryLifecycleStateCreating,
	"active":   DbConnectionBundleSummaryLifecycleStateActive,
	"inactive": DbConnectionBundleSummaryLifecycleStateInactive,
	"updating": DbConnectionBundleSummaryLifecycleStateUpdating,
	"deleting": DbConnectionBundleSummaryLifecycleStateDeleting,
	"deleted":  DbConnectionBundleSummaryLifecycleStateDeleted,
	"failed":   DbConnectionBundleSummaryLifecycleStateFailed,
}

// GetDbConnectionBundleSummaryLifecycleStateEnumValues Enumerates the set of values for DbConnectionBundleSummaryLifecycleStateEnum
func GetDbConnectionBundleSummaryLifecycleStateEnumValues() []DbConnectionBundleSummaryLifecycleStateEnum {
	values := make([]DbConnectionBundleSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbConnectionBundleSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbConnectionBundleSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DbConnectionBundleSummaryLifecycleStateEnum
func GetDbConnectionBundleSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingDbConnectionBundleSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbConnectionBundleSummaryLifecycleStateEnum(val string) (DbConnectionBundleSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDbConnectionBundleSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbConnectionBundleSummaryDbConnectionBundleTypeEnum Enum with underlying type: string
type DbConnectionBundleSummaryDbConnectionBundleTypeEnum string

// Set of constants representing the allowable values for DbConnectionBundleSummaryDbConnectionBundleTypeEnum
const (
	DbConnectionBundleSummaryDbConnectionBundleTypeTls  DbConnectionBundleSummaryDbConnectionBundleTypeEnum = "TLS"
	DbConnectionBundleSummaryDbConnectionBundleTypeMtls DbConnectionBundleSummaryDbConnectionBundleTypeEnum = "MTLS"
)

var mappingDbConnectionBundleSummaryDbConnectionBundleTypeEnum = map[string]DbConnectionBundleSummaryDbConnectionBundleTypeEnum{
	"TLS":  DbConnectionBundleSummaryDbConnectionBundleTypeTls,
	"MTLS": DbConnectionBundleSummaryDbConnectionBundleTypeMtls,
}

var mappingDbConnectionBundleSummaryDbConnectionBundleTypeEnumLowerCase = map[string]DbConnectionBundleSummaryDbConnectionBundleTypeEnum{
	"tls":  DbConnectionBundleSummaryDbConnectionBundleTypeTls,
	"mtls": DbConnectionBundleSummaryDbConnectionBundleTypeMtls,
}

// GetDbConnectionBundleSummaryDbConnectionBundleTypeEnumValues Enumerates the set of values for DbConnectionBundleSummaryDbConnectionBundleTypeEnum
func GetDbConnectionBundleSummaryDbConnectionBundleTypeEnumValues() []DbConnectionBundleSummaryDbConnectionBundleTypeEnum {
	values := make([]DbConnectionBundleSummaryDbConnectionBundleTypeEnum, 0)
	for _, v := range mappingDbConnectionBundleSummaryDbConnectionBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbConnectionBundleSummaryDbConnectionBundleTypeEnumStringValues Enumerates the set of values in String for DbConnectionBundleSummaryDbConnectionBundleTypeEnum
func GetDbConnectionBundleSummaryDbConnectionBundleTypeEnumStringValues() []string {
	return []string{
		"TLS",
		"MTLS",
	}
}

// GetMappingDbConnectionBundleSummaryDbConnectionBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbConnectionBundleSummaryDbConnectionBundleTypeEnum(val string) (DbConnectionBundleSummaryDbConnectionBundleTypeEnum, bool) {
	enum, ok := mappingDbConnectionBundleSummaryDbConnectionBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
