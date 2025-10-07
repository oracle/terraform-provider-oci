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

// IotDomain An IoT domain is an Oracle Cloud Infrastructure resource that provides a managed environment for organizing and managing
// digital twin resources, such as models, adapters, instances and relationships, within a compartment and IoT domain group.
// To use any API operations, you must be authorized in an IAM policy. If you are not authorized, contact an administrator.
// If you are an administrator who needs to create policies to grant users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type IotDomain struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain group.
	IotDomainGroupId *string `mandatory:"true" json:"iotDomainGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the IoT domain.
	LifecycleState IotDomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the resource was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// Host name of an IoT domain, where IoT devices can connect to.
	DeviceHost *string `mandatory:"false" json:"deviceHost"`

	// List of IAM groups of form described in here (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/mnqmn/#GUID-3634D6C9-A7F1-4875-9925-BAEA2D3C5197) that are allowed to directly connect to the data host.
	DbAllowListedIdentityGroupNames []string `mandatory:"false" json:"dbAllowListedIdentityGroupNames"`

	// Host name of identity domain that is used for authenticating connect to data host via ORDS.
	DbAllowedIdentityDomainHost *string `mandatory:"false" json:"dbAllowedIdentityDomainHost"`

	DataRetentionPeriodsInDays *DataRetentionPeriodsInDays `mandatory:"false" json:"dataRetentionPeriodsInDays"`

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

func (m IotDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IotDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIotDomainLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IotDomainLifecycleStateEnum Enum with underlying type: string
type IotDomainLifecycleStateEnum string

// Set of constants representing the allowable values for IotDomainLifecycleStateEnum
const (
	IotDomainLifecycleStateCreating IotDomainLifecycleStateEnum = "CREATING"
	IotDomainLifecycleStateUpdating IotDomainLifecycleStateEnum = "UPDATING"
	IotDomainLifecycleStateActive   IotDomainLifecycleStateEnum = "ACTIVE"
	IotDomainLifecycleStateDeleting IotDomainLifecycleStateEnum = "DELETING"
	IotDomainLifecycleStateDeleted  IotDomainLifecycleStateEnum = "DELETED"
	IotDomainLifecycleStateFailed   IotDomainLifecycleStateEnum = "FAILED"
)

var mappingIotDomainLifecycleStateEnum = map[string]IotDomainLifecycleStateEnum{
	"CREATING": IotDomainLifecycleStateCreating,
	"UPDATING": IotDomainLifecycleStateUpdating,
	"ACTIVE":   IotDomainLifecycleStateActive,
	"DELETING": IotDomainLifecycleStateDeleting,
	"DELETED":  IotDomainLifecycleStateDeleted,
	"FAILED":   IotDomainLifecycleStateFailed,
}

var mappingIotDomainLifecycleStateEnumLowerCase = map[string]IotDomainLifecycleStateEnum{
	"creating": IotDomainLifecycleStateCreating,
	"updating": IotDomainLifecycleStateUpdating,
	"active":   IotDomainLifecycleStateActive,
	"deleting": IotDomainLifecycleStateDeleting,
	"deleted":  IotDomainLifecycleStateDeleted,
	"failed":   IotDomainLifecycleStateFailed,
}

// GetIotDomainLifecycleStateEnumValues Enumerates the set of values for IotDomainLifecycleStateEnum
func GetIotDomainLifecycleStateEnumValues() []IotDomainLifecycleStateEnum {
	values := make([]IotDomainLifecycleStateEnum, 0)
	for _, v := range mappingIotDomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIotDomainLifecycleStateEnumStringValues Enumerates the set of values in String for IotDomainLifecycleStateEnum
func GetIotDomainLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIotDomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIotDomainLifecycleStateEnum(val string) (IotDomainLifecycleStateEnum, bool) {
	enum, ok := mappingIotDomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
