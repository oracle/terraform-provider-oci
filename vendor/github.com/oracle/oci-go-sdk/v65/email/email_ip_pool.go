// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailIpPool A collection of IP addresses used by the Email Delivery service to submit an email.
type EmailIpPool struct {

	// The unique OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IpPool resource that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of the IpPool. The name must be unique within a region.
	// The name is case sensitive and supported characters include alphanumeric, hyphens ("-") and underscore ("_") characters.
	// Example: green_pool-1
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IpPool.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The description of the IpPool. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The current state of the IpPool.
	LifecycleState EmailIpPoolLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a
	// resource in 'INACTIVE' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Summary of outbound IPs assigned to the IpPool.
	OutboundIps []EmailOutboundIpSummary `mandatory:"false" json:"outboundIps"`

	// The time the IpPool was created.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time of the last change to the IpPool, due to a state change or
	// an update operation.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m EmailIpPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailIpPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEmailIpPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailIpPoolLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailIpPoolLifecycleStateEnum Enum with underlying type: string
type EmailIpPoolLifecycleStateEnum string

// Set of constants representing the allowable values for EmailIpPoolLifecycleStateEnum
const (
	EmailIpPoolLifecycleStateCreating EmailIpPoolLifecycleStateEnum = "CREATING"
	EmailIpPoolLifecycleStateUpdating EmailIpPoolLifecycleStateEnum = "UPDATING"
	EmailIpPoolLifecycleStateActive   EmailIpPoolLifecycleStateEnum = "ACTIVE"
	EmailIpPoolLifecycleStateDeleting EmailIpPoolLifecycleStateEnum = "DELETING"
	EmailIpPoolLifecycleStateDeleted  EmailIpPoolLifecycleStateEnum = "DELETED"
	EmailIpPoolLifecycleStateFailed   EmailIpPoolLifecycleStateEnum = "FAILED"
	EmailIpPoolLifecycleStateInactive EmailIpPoolLifecycleStateEnum = "INACTIVE"
)

var mappingEmailIpPoolLifecycleStateEnum = map[string]EmailIpPoolLifecycleStateEnum{
	"CREATING": EmailIpPoolLifecycleStateCreating,
	"UPDATING": EmailIpPoolLifecycleStateUpdating,
	"ACTIVE":   EmailIpPoolLifecycleStateActive,
	"DELETING": EmailIpPoolLifecycleStateDeleting,
	"DELETED":  EmailIpPoolLifecycleStateDeleted,
	"FAILED":   EmailIpPoolLifecycleStateFailed,
	"INACTIVE": EmailIpPoolLifecycleStateInactive,
}

var mappingEmailIpPoolLifecycleStateEnumLowerCase = map[string]EmailIpPoolLifecycleStateEnum{
	"creating": EmailIpPoolLifecycleStateCreating,
	"updating": EmailIpPoolLifecycleStateUpdating,
	"active":   EmailIpPoolLifecycleStateActive,
	"deleting": EmailIpPoolLifecycleStateDeleting,
	"deleted":  EmailIpPoolLifecycleStateDeleted,
	"failed":   EmailIpPoolLifecycleStateFailed,
	"inactive": EmailIpPoolLifecycleStateInactive,
}

// GetEmailIpPoolLifecycleStateEnumValues Enumerates the set of values for EmailIpPoolLifecycleStateEnum
func GetEmailIpPoolLifecycleStateEnumValues() []EmailIpPoolLifecycleStateEnum {
	values := make([]EmailIpPoolLifecycleStateEnum, 0)
	for _, v := range mappingEmailIpPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailIpPoolLifecycleStateEnumStringValues Enumerates the set of values in String for EmailIpPoolLifecycleStateEnum
func GetEmailIpPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingEmailIpPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailIpPoolLifecycleStateEnum(val string) (EmailIpPoolLifecycleStateEnum, bool) {
	enum, ok := mappingEmailIpPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
