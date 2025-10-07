// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IotDomainGroup An IoT domain group is an Oracle Cloud Infrastructure resource that provides a managed environment for organizing
// and managing IoT domains, within a compartment.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type IotDomainGroup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of an IoT Domain Group.
	LifecycleState IotDomainGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the resource was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The host name of the database corresponding to the IoT Domain group.
	DataHost *string `mandatory:"false" json:"dataHost"`

	// This is an array of VCN OCID (virtual cloud network Oracle Cloud ID) that is allowed to connect the data host.
	DbAllowListedVcnIds []string `mandatory:"false" json:"dbAllowListedVcnIds"`

	// The connection string used to connect to the data host associated with the IoT domain group.
	DbConnectionString *string `mandatory:"false" json:"dbConnectionString"`

	// The token scope used to connect to the data host associated with the IoT domain group.
	DbTokenScope *string `mandatory:"false" json:"dbTokenScope"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time when the resource was last updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m IotDomainGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IotDomainGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotDomainGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIotDomainGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IotDomainGroupLifecycleStateEnum Enum with underlying type: string
type IotDomainGroupLifecycleStateEnum string

// Set of constants representing the allowable values for IotDomainGroupLifecycleStateEnum
const (
	IotDomainGroupLifecycleStateCreating IotDomainGroupLifecycleStateEnum = "CREATING"
	IotDomainGroupLifecycleStateUpdating IotDomainGroupLifecycleStateEnum = "UPDATING"
	IotDomainGroupLifecycleStateActive   IotDomainGroupLifecycleStateEnum = "ACTIVE"
	IotDomainGroupLifecycleStateDeleting IotDomainGroupLifecycleStateEnum = "DELETING"
	IotDomainGroupLifecycleStateDeleted  IotDomainGroupLifecycleStateEnum = "DELETED"
	IotDomainGroupLifecycleStateFailed   IotDomainGroupLifecycleStateEnum = "FAILED"
)

var mappingIotDomainGroupLifecycleStateEnum = map[string]IotDomainGroupLifecycleStateEnum{
	"CREATING": IotDomainGroupLifecycleStateCreating,
	"UPDATING": IotDomainGroupLifecycleStateUpdating,
	"ACTIVE":   IotDomainGroupLifecycleStateActive,
	"DELETING": IotDomainGroupLifecycleStateDeleting,
	"DELETED":  IotDomainGroupLifecycleStateDeleted,
	"FAILED":   IotDomainGroupLifecycleStateFailed,
}

var mappingIotDomainGroupLifecycleStateEnumLowerCase = map[string]IotDomainGroupLifecycleStateEnum{
	"creating": IotDomainGroupLifecycleStateCreating,
	"updating": IotDomainGroupLifecycleStateUpdating,
	"active":   IotDomainGroupLifecycleStateActive,
	"deleting": IotDomainGroupLifecycleStateDeleting,
	"deleted":  IotDomainGroupLifecycleStateDeleted,
	"failed":   IotDomainGroupLifecycleStateFailed,
}

// GetIotDomainGroupLifecycleStateEnumValues Enumerates the set of values for IotDomainGroupLifecycleStateEnum
func GetIotDomainGroupLifecycleStateEnumValues() []IotDomainGroupLifecycleStateEnum {
	values := make([]IotDomainGroupLifecycleStateEnum, 0)
	for _, v := range mappingIotDomainGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIotDomainGroupLifecycleStateEnumStringValues Enumerates the set of values in String for IotDomainGroupLifecycleStateEnum
func GetIotDomainGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIotDomainGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIotDomainGroupLifecycleStateEnum(val string) (IotDomainGroupLifecycleStateEnum, bool) {
	enum, ok := mappingIotDomainGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
