// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivilegedApiControl A PrivilegedApiControl is a kind of Policy definition which provides details about which operations needs to be secure; who can approve a privilegedApiRequest requesting for a particular operation, whether the operations needs to be approved by customer or is it preApproved.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type PrivilegedApiControl struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiControl.
	Id *string `mandatory:"true" json:"id"`

	// Name of the privilegedApi control. The name must be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the PrivilegedApiControl was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the PrivilegedApiControl.
	State *string `mandatory:"true" json:"state"`

	// The current state of the PrivilegedApiControl.
	LifecycleState PrivilegedApiControlLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Description of privilegedApi control.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the OCI Notification topic to publish messages related to this Privileged Api Control.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// List of IAM user group ids who can approve an privilegedApi request associated with a target resource under the governance of this operator control.
	ApproverGroupIdList []string `mandatory:"false" json:"approverGroupIdList"`

	// resourceType for which the PrivilegedApiControl is applicable
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// contains Resource details
	Resources []string `mandatory:"false" json:"resources"`

	// List of privileged operations/apis. These operations/apis will be treaated as secured, once enabled by the Privileged API Managment for a resource. Any of these operations, if needs to be executed, needs to be raised as a PrivilegedApi Request which needs to be approved by customers or it can be pre-approved.
	PrivilegedOperationList []PrivilegedApiDetails `mandatory:"false" json:"privilegedOperationList"`

	// Number of approvers required to approve an privilegedApi request.
	NumberOfApprovers *int `mandatory:"false" json:"numberOfApprovers"`

	// The date and time the PrivilegedApiControl was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the PrivilegedApiControl was marked for delete, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeDeleted *common.SDKTime `mandatory:"false" json:"timeDeleted"`

	// A message that describes the current state of the PrivilegedApiControl in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	StateDetails *string `mandatory:"false" json:"stateDetails"`

	// A message that describes the current state of the PrivilegedApiControl in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PrivilegedApiControl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivilegedApiControl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivilegedApiControlLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivilegedApiControlLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivilegedApiControlLifecycleStateEnum Enum with underlying type: string
type PrivilegedApiControlLifecycleStateEnum string

// Set of constants representing the allowable values for PrivilegedApiControlLifecycleStateEnum
const (
	PrivilegedApiControlLifecycleStateCreating       PrivilegedApiControlLifecycleStateEnum = "CREATING"
	PrivilegedApiControlLifecycleStateUpdating       PrivilegedApiControlLifecycleStateEnum = "UPDATING"
	PrivilegedApiControlLifecycleStateActive         PrivilegedApiControlLifecycleStateEnum = "ACTIVE"
	PrivilegedApiControlLifecycleStateDeleting       PrivilegedApiControlLifecycleStateEnum = "DELETING"
	PrivilegedApiControlLifecycleStateDeleted        PrivilegedApiControlLifecycleStateEnum = "DELETED"
	PrivilegedApiControlLifecycleStateFailed         PrivilegedApiControlLifecycleStateEnum = "FAILED"
	PrivilegedApiControlLifecycleStateNeedsAttention PrivilegedApiControlLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingPrivilegedApiControlLifecycleStateEnum = map[string]PrivilegedApiControlLifecycleStateEnum{
	"CREATING":        PrivilegedApiControlLifecycleStateCreating,
	"UPDATING":        PrivilegedApiControlLifecycleStateUpdating,
	"ACTIVE":          PrivilegedApiControlLifecycleStateActive,
	"DELETING":        PrivilegedApiControlLifecycleStateDeleting,
	"DELETED":         PrivilegedApiControlLifecycleStateDeleted,
	"FAILED":          PrivilegedApiControlLifecycleStateFailed,
	"NEEDS_ATTENTION": PrivilegedApiControlLifecycleStateNeedsAttention,
}

var mappingPrivilegedApiControlLifecycleStateEnumLowerCase = map[string]PrivilegedApiControlLifecycleStateEnum{
	"creating":        PrivilegedApiControlLifecycleStateCreating,
	"updating":        PrivilegedApiControlLifecycleStateUpdating,
	"active":          PrivilegedApiControlLifecycleStateActive,
	"deleting":        PrivilegedApiControlLifecycleStateDeleting,
	"deleted":         PrivilegedApiControlLifecycleStateDeleted,
	"failed":          PrivilegedApiControlLifecycleStateFailed,
	"needs_attention": PrivilegedApiControlLifecycleStateNeedsAttention,
}

// GetPrivilegedApiControlLifecycleStateEnumValues Enumerates the set of values for PrivilegedApiControlLifecycleStateEnum
func GetPrivilegedApiControlLifecycleStateEnumValues() []PrivilegedApiControlLifecycleStateEnum {
	values := make([]PrivilegedApiControlLifecycleStateEnum, 0)
	for _, v := range mappingPrivilegedApiControlLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivilegedApiControlLifecycleStateEnumStringValues Enumerates the set of values in String for PrivilegedApiControlLifecycleStateEnum
func GetPrivilegedApiControlLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingPrivilegedApiControlLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivilegedApiControlLifecycleStateEnum(val string) (PrivilegedApiControlLifecycleStateEnum, bool) {
	enum, ok := mappingPrivilegedApiControlLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
