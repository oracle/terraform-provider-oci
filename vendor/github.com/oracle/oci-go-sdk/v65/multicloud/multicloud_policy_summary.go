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

// MulticloudPolicySummary Summary of missing IAM policies for a multicloud subscription.
type MulticloudPolicySummary struct {

	// Oracle Cloud Infrastructure Subscription Type.
	SubscriptionType SubscriptionTypeEnum `mandatory:"true" json:"subscriptionType"`

	// Compartment The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Subscription
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// groups required for the particular subscriptionType IAM policy statements required.
	Groups []string `mandatory:"false" json:"groups"`

	// Missing policy definitions.
	Policies []MulticloudPolicy `mandatory:"false" json:"policies"`

	// The current state of the Multicloud Network Alert.
	LifecycleState MulticloudPolicyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m MulticloudPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MulticloudPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubscriptionTypeEnum(string(m.SubscriptionType)); !ok && m.SubscriptionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionType: %s. Supported values are: %s.", m.SubscriptionType, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMulticloudPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMulticloudPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
