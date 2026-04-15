// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see Oracle Multicloud Hub (https://docs.oracle.com/iaas/Content/multicloud-hub/home.htm).
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MulticloudSubscriptionSummary A Multicloud subscription.
type MulticloudSubscriptionSummary struct {

	// Subscription ID for OCI and partner cloud in classic format.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// The partner cloud account ID.
	PartnerCloudAccountIdentifier *string `mandatory:"true" json:"partnerCloudAccountIdentifier"`

	// The date and time that the subscription was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// URL to the subscription details page.
	// Example: `https://{console-url}/org-mgmt/subscription/ocid1.organizationssubscription.oc1.iad.exampleuniqueid?region=us-ashburn-1`.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The partner cloud tenant ID.
	PartnerCloudTenantIdentifier *string `mandatory:"false" json:"partnerCloudTenantIdentifier"`

	// The cloud service provider.
	ServiceName SubscriptionTypeEnum `mandatory:"false" json:"serviceName,omitempty"`

	// The date and time that the Multicloud base compartment was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeLinkedDate *common.SDKTime `mandatory:"false" json:"timeLinkedDate"`

	// Payment plan for the subscription.
	PaymentPlan *string `mandatory:"false" json:"paymentPlan"`

	// Total value for the subscription.
	ActiveCommitment *string `mandatory:"false" json:"activeCommitment"`

	// The end date and time for the subscription, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeEndDate *common.SDKTime `mandatory:"false" json:"timeEndDate"`

	// The current state of the subscription.
	LifecycleState MulticloudSubscriptionSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Properties specific to the cloud service provider. For example, AzureSubnetId for Azure.
	CspAdditionalProperties map[string]string `mandatory:"false" json:"cspAdditionalProperties"`

	// The date and time that the subscription was updated, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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
}

func (m MulticloudSubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MulticloudSubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionTypeEnum(string(m.ServiceName)); !ok && m.ServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceName: %s. Supported values are: %s.", m.ServiceName, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMulticloudSubscriptionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMulticloudSubscriptionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MulticloudSubscriptionSummaryLifecycleStateEnum Enum with underlying type: string
type MulticloudSubscriptionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for MulticloudSubscriptionSummaryLifecycleStateEnum
const (
	MulticloudSubscriptionSummaryLifecycleStateActive   MulticloudSubscriptionSummaryLifecycleStateEnum = "ACTIVE"
	MulticloudSubscriptionSummaryLifecycleStateInactive MulticloudSubscriptionSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingMulticloudSubscriptionSummaryLifecycleStateEnum = map[string]MulticloudSubscriptionSummaryLifecycleStateEnum{
	"ACTIVE":   MulticloudSubscriptionSummaryLifecycleStateActive,
	"INACTIVE": MulticloudSubscriptionSummaryLifecycleStateInactive,
}

var mappingMulticloudSubscriptionSummaryLifecycleStateEnumLowerCase = map[string]MulticloudSubscriptionSummaryLifecycleStateEnum{
	"active":   MulticloudSubscriptionSummaryLifecycleStateActive,
	"inactive": MulticloudSubscriptionSummaryLifecycleStateInactive,
}

// GetMulticloudSubscriptionSummaryLifecycleStateEnumValues Enumerates the set of values for MulticloudSubscriptionSummaryLifecycleStateEnum
func GetMulticloudSubscriptionSummaryLifecycleStateEnumValues() []MulticloudSubscriptionSummaryLifecycleStateEnum {
	values := make([]MulticloudSubscriptionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingMulticloudSubscriptionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudSubscriptionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for MulticloudSubscriptionSummaryLifecycleStateEnum
func GetMulticloudSubscriptionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingMulticloudSubscriptionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudSubscriptionSummaryLifecycleStateEnum(val string) (MulticloudSubscriptionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingMulticloudSubscriptionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
