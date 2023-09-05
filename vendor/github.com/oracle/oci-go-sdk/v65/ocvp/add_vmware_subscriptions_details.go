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

// AddVmwareSubscriptionsDetails Commitment term subscriptions added to a VMware billing link.  An will be added for every host in the
// request.
type AddVmwareSubscriptionsDetails struct {

	// The compute shape name of the subscription.
	// ListSupportedHostShapes.
	Shape *string `mandatory:"true" json:"shape"`

	// The number of hosts in subscription.
	HostCount *float32 `mandatory:"true" json:"hostCount"`

	// The per host OCPU count of the subscription.
	OcpuCount *float32 `mandatory:"true" json:"ocpuCount"`

	// The commitment term selected for subscription.
	CommitmentTerm CommitmentEnum `mandatory:"true" json:"commitmentTerm"`
}

func (m AddVmwareSubscriptionsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddVmwareSubscriptionsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCommitmentEnum(string(m.CommitmentTerm)); !ok && m.CommitmentTerm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommitmentTerm: %s. Supported values are: %s.", m.CommitmentTerm, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
