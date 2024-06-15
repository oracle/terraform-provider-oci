// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ZprTagNamespace A managed container for defined ZPR tags. A ZPR tag namespace is unique in a tenancy. For more information,
// see Managing Tags and Tag Namespaces (https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type ZprTagNamespace struct {

	// The OCID of the ZPR tag namespace.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the tag namespace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the ZPR tag namespace.
	Description *string `mandatory:"true" json:"description"`

	// Indeicates whether the tag namespace is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"true" json:"isRetired"`

	// Date and time the ZPR tag namespace was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// ZPR tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"oracle-zpr": {"td": {"value": "42", "mode": "audit"}}}`
	ZprTags map[string]map[string]interface{} `mandatory:"false" json:"zprTags"`

	// Usage of ZPR system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-internal-zpr": {"administrator": {"value": "true", "mode": "enforce"}}}`
	ZprSystemTags map[string]map[string]interface{} `mandatory:"false" json:"zprSystemTags"`

	// The ZPR tagnamespace's current state. After creating a ZPR tagnamespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a ZPR tag namespace, make sure its `lifecycleState` is INACTIVE.
	LifecycleState ZprTagNamespaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ZprTagNamespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprTagNamespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingZprTagNamespaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZprTagNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ZprTagNamespaceLifecycleStateEnum Enum with underlying type: string
type ZprTagNamespaceLifecycleStateEnum string

// Set of constants representing the allowable values for ZprTagNamespaceLifecycleStateEnum
const (
	ZprTagNamespaceLifecycleStateActive   ZprTagNamespaceLifecycleStateEnum = "ACTIVE"
	ZprTagNamespaceLifecycleStateInactive ZprTagNamespaceLifecycleStateEnum = "INACTIVE"
	ZprTagNamespaceLifecycleStateDeleting ZprTagNamespaceLifecycleStateEnum = "DELETING"
	ZprTagNamespaceLifecycleStateDeleted  ZprTagNamespaceLifecycleStateEnum = "DELETED"
)

var mappingZprTagNamespaceLifecycleStateEnum = map[string]ZprTagNamespaceLifecycleStateEnum{
	"ACTIVE":   ZprTagNamespaceLifecycleStateActive,
	"INACTIVE": ZprTagNamespaceLifecycleStateInactive,
	"DELETING": ZprTagNamespaceLifecycleStateDeleting,
	"DELETED":  ZprTagNamespaceLifecycleStateDeleted,
}

var mappingZprTagNamespaceLifecycleStateEnumLowerCase = map[string]ZprTagNamespaceLifecycleStateEnum{
	"active":   ZprTagNamespaceLifecycleStateActive,
	"inactive": ZprTagNamespaceLifecycleStateInactive,
	"deleting": ZprTagNamespaceLifecycleStateDeleting,
	"deleted":  ZprTagNamespaceLifecycleStateDeleted,
}

// GetZprTagNamespaceLifecycleStateEnumValues Enumerates the set of values for ZprTagNamespaceLifecycleStateEnum
func GetZprTagNamespaceLifecycleStateEnumValues() []ZprTagNamespaceLifecycleStateEnum {
	values := make([]ZprTagNamespaceLifecycleStateEnum, 0)
	for _, v := range mappingZprTagNamespaceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZprTagNamespaceLifecycleStateEnumStringValues Enumerates the set of values in String for ZprTagNamespaceLifecycleStateEnum
func GetZprTagNamespaceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingZprTagNamespaceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZprTagNamespaceLifecycleStateEnum(val string) (ZprTagNamespaceLifecycleStateEnum, bool) {
	enum, ok := mappingZprTagNamespaceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
