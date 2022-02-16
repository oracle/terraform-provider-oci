// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PrivateApplication Full details of an application or a solution, which lives inside the tenancy and may be included into service catalogs.
type PrivateApplication struct {

	// The lifecycle state of the private application.
	LifecycleState PrivateApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the private application resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the private application in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// The name of the private application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of packages within this private application.
	PackageType PackageTypeEnumEnum `mandatory:"true" json:"packageType"`

	// The date and time the private application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-26T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the private application.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description of the private application.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	// The date and time the private application was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-12-10T05:10:29.721Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m PrivateApplication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateApplication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateApplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateApplicationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnumEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateApplicationLifecycleStateEnum Enum with underlying type: string
type PrivateApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateApplicationLifecycleStateEnum
const (
	PrivateApplicationLifecycleStateCreating PrivateApplicationLifecycleStateEnum = "CREATING"
	PrivateApplicationLifecycleStateUpdating PrivateApplicationLifecycleStateEnum = "UPDATING"
	PrivateApplicationLifecycleStateActive   PrivateApplicationLifecycleStateEnum = "ACTIVE"
	PrivateApplicationLifecycleStateDeleting PrivateApplicationLifecycleStateEnum = "DELETING"
	PrivateApplicationLifecycleStateDeleted  PrivateApplicationLifecycleStateEnum = "DELETED"
)

var mappingPrivateApplicationLifecycleStateEnum = map[string]PrivateApplicationLifecycleStateEnum{
	"CREATING": PrivateApplicationLifecycleStateCreating,
	"UPDATING": PrivateApplicationLifecycleStateUpdating,
	"ACTIVE":   PrivateApplicationLifecycleStateActive,
	"DELETING": PrivateApplicationLifecycleStateDeleting,
	"DELETED":  PrivateApplicationLifecycleStateDeleted,
}

// GetPrivateApplicationLifecycleStateEnumValues Enumerates the set of values for PrivateApplicationLifecycleStateEnum
func GetPrivateApplicationLifecycleStateEnumValues() []PrivateApplicationLifecycleStateEnum {
	values := make([]PrivateApplicationLifecycleStateEnum, 0)
	for _, v := range mappingPrivateApplicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateApplicationLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateApplicationLifecycleStateEnum
func GetPrivateApplicationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingPrivateApplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateApplicationLifecycleStateEnum(val string) (PrivateApplicationLifecycleStateEnum, bool) {
	mappingPrivateApplicationLifecycleStateEnumIgnoreCase := make(map[string]PrivateApplicationLifecycleStateEnum)
	for k, v := range mappingPrivateApplicationLifecycleStateEnum {
		mappingPrivateApplicationLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPrivateApplicationLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
