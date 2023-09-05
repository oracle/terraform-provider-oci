// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmwareSubscription A commitment term subscription added to a VMware billing link
type VmwareSubscription struct {

	// The compute shape name of the subscription.
	// ListSupportedHostShapes.
	Shape *string `mandatory:"true" json:"shape"`

	// The OCPU count of the subscription.
	OcpuCount *float32 `mandatory:"true" json:"ocpuCount"`

	// The commitment term selected for subscription.
	CommitmentTerm CommitmentEnum `mandatory:"true" json:"commitmentTerm"`

	// Subscription billing cycle expiration date. An empty string indicates subscription has not been assigned
	// to a compute instance yet.
	// Example: `2016-08-25T21:10:29.600Z`
	BillingExpirationDate *common.SDKTime `mandatory:"true" json:"billingExpirationDate"`

	// OCID of ESXi host that subscription is assigned to. An empty string indicates
	// subscription has not been assigned to a compute instance yet.  The `esxiHostId`
	// is the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of that ESXi host.
	EsxiHostId *string `mandatory:"true" json:"esxiHostId"`

	// The current state of the subscription
	SubscriptionState VmwareSubscriptionStatesEnum `mandatory:"true" json:"subscriptionState"`
}

func (m VmwareSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmwareSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCommitmentEnum(string(m.CommitmentTerm)); !ok && m.CommitmentTerm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommitmentTerm: %s. Supported values are: %s.", m.CommitmentTerm, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmwareSubscriptionStatesEnum(string(m.SubscriptionState)); !ok && m.SubscriptionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionState: %s. Supported values are: %s.", m.SubscriptionState, strings.Join(GetVmwareSubscriptionStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
