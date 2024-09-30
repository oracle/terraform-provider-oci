// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ZprPolicy A ZprPolicy is a rule that governs the communication between specific endpoints identified by their security attributes.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ZprPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ZprPolicy.
	Id *string `mandatory:"true" json:"id"`

	// The name you assign to the ZprPolicy during creation. The name must be unique across all ZPL policies in the tenancy.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the ZprPolicy during creation. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An array of ZprPolicy statements (up to 25 statements per ZprPolicy) written in the Zero Trust Packet Routing Policy Language.
	Statements []string `mandatory:"true" json:"statements"`

	// The current state of the ZprPolicy.
	LifecycleState ZprPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the ZprPolicy was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A message that describes the current state of the ZprPolicy in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the ZprPolicy was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ZprPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingZprPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZprPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ZprPolicyLifecycleStateEnum Enum with underlying type: string
type ZprPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for ZprPolicyLifecycleStateEnum
const (
	ZprPolicyLifecycleStateActive         ZprPolicyLifecycleStateEnum = "ACTIVE"
	ZprPolicyLifecycleStateCreating       ZprPolicyLifecycleStateEnum = "CREATING"
	ZprPolicyLifecycleStateFailed         ZprPolicyLifecycleStateEnum = "FAILED"
	ZprPolicyLifecycleStateUpdating       ZprPolicyLifecycleStateEnum = "UPDATING"
	ZprPolicyLifecycleStateDeleting       ZprPolicyLifecycleStateEnum = "DELETING"
	ZprPolicyLifecycleStateDeleted        ZprPolicyLifecycleStateEnum = "DELETED"
	ZprPolicyLifecycleStateNeedsAttention ZprPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingZprPolicyLifecycleStateEnum = map[string]ZprPolicyLifecycleStateEnum{
	"ACTIVE":          ZprPolicyLifecycleStateActive,
	"CREATING":        ZprPolicyLifecycleStateCreating,
	"FAILED":          ZprPolicyLifecycleStateFailed,
	"UPDATING":        ZprPolicyLifecycleStateUpdating,
	"DELETING":        ZprPolicyLifecycleStateDeleting,
	"DELETED":         ZprPolicyLifecycleStateDeleted,
	"NEEDS_ATTENTION": ZprPolicyLifecycleStateNeedsAttention,
}

var mappingZprPolicyLifecycleStateEnumLowerCase = map[string]ZprPolicyLifecycleStateEnum{
	"active":          ZprPolicyLifecycleStateActive,
	"creating":        ZprPolicyLifecycleStateCreating,
	"failed":          ZprPolicyLifecycleStateFailed,
	"updating":        ZprPolicyLifecycleStateUpdating,
	"deleting":        ZprPolicyLifecycleStateDeleting,
	"deleted":         ZprPolicyLifecycleStateDeleted,
	"needs_attention": ZprPolicyLifecycleStateNeedsAttention,
}

// GetZprPolicyLifecycleStateEnumValues Enumerates the set of values for ZprPolicyLifecycleStateEnum
func GetZprPolicyLifecycleStateEnumValues() []ZprPolicyLifecycleStateEnum {
	values := make([]ZprPolicyLifecycleStateEnum, 0)
	for _, v := range mappingZprPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZprPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for ZprPolicyLifecycleStateEnum
func GetZprPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingZprPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZprPolicyLifecycleStateEnum(val string) (ZprPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingZprPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
