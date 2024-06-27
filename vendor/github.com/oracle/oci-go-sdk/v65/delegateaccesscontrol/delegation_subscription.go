// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DelegationSubscription Details of the Delegation Subscription.
type DelegationSubscription struct {

	// Unique identifier for the Delegation Subscription.
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier of the Service Provider.
	ServiceProviderId *string `mandatory:"true" json:"serviceProviderId"`

	// Subscribed Service Provider Service Type.
	SubscribedServiceType ServiceProviderServiceTypeEnum `mandatory:"true" json:"subscribedServiceType"`

	// The OCID of the compartment that contains the Delegation Subscription.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Delegation Subscription.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the Service Provider.
	LifecycleState DelegationSubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Service Provider was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Service Provider was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DelegationSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegationSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceProviderServiceTypeEnum(string(m.SubscribedServiceType)); !ok && m.SubscribedServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscribedServiceType: %s. Supported values are: %s.", m.SubscribedServiceType, strings.Join(GetServiceProviderServiceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDelegationSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDelegationSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DelegationSubscriptionLifecycleStateEnum Enum with underlying type: string
type DelegationSubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for DelegationSubscriptionLifecycleStateEnum
const (
	DelegationSubscriptionLifecycleStateCreating DelegationSubscriptionLifecycleStateEnum = "CREATING"
	DelegationSubscriptionLifecycleStateActive   DelegationSubscriptionLifecycleStateEnum = "ACTIVE"
	DelegationSubscriptionLifecycleStateUpdating DelegationSubscriptionLifecycleStateEnum = "UPDATING"
	DelegationSubscriptionLifecycleStateDeleting DelegationSubscriptionLifecycleStateEnum = "DELETING"
	DelegationSubscriptionLifecycleStateDeleted  DelegationSubscriptionLifecycleStateEnum = "DELETED"
	DelegationSubscriptionLifecycleStateFailed   DelegationSubscriptionLifecycleStateEnum = "FAILED"
)

var mappingDelegationSubscriptionLifecycleStateEnum = map[string]DelegationSubscriptionLifecycleStateEnum{
	"CREATING": DelegationSubscriptionLifecycleStateCreating,
	"ACTIVE":   DelegationSubscriptionLifecycleStateActive,
	"UPDATING": DelegationSubscriptionLifecycleStateUpdating,
	"DELETING": DelegationSubscriptionLifecycleStateDeleting,
	"DELETED":  DelegationSubscriptionLifecycleStateDeleted,
	"FAILED":   DelegationSubscriptionLifecycleStateFailed,
}

var mappingDelegationSubscriptionLifecycleStateEnumLowerCase = map[string]DelegationSubscriptionLifecycleStateEnum{
	"creating": DelegationSubscriptionLifecycleStateCreating,
	"active":   DelegationSubscriptionLifecycleStateActive,
	"updating": DelegationSubscriptionLifecycleStateUpdating,
	"deleting": DelegationSubscriptionLifecycleStateDeleting,
	"deleted":  DelegationSubscriptionLifecycleStateDeleted,
	"failed":   DelegationSubscriptionLifecycleStateFailed,
}

// GetDelegationSubscriptionLifecycleStateEnumValues Enumerates the set of values for DelegationSubscriptionLifecycleStateEnum
func GetDelegationSubscriptionLifecycleStateEnumValues() []DelegationSubscriptionLifecycleStateEnum {
	values := make([]DelegationSubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingDelegationSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for DelegationSubscriptionLifecycleStateEnum
func GetDelegationSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDelegationSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationSubscriptionLifecycleStateEnum(val string) (DelegationSubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingDelegationSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
