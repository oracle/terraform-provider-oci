// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCrossConnectGroupDetails The representation of CreateCrossConnectGroupDetails
type CreateCrossConnectGroupDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the cross-connect group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A reference name or identifier for the physical fiber connection that this cross-connect
	// group uses.
	CustomerReferenceName *string `mandatory:"false" json:"customerReferenceName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	MacsecProperties *CreateMacsecProperties `mandatory:"false" json:"macsecProperties"`

	// (Optional) Minimum number of active cross-connects required for the cross-connect group to be considered
	// operational. During create cross-connect-group operation this value can only be set to 1 (If not specified,
	// this value defaults to 1) and can be edited using the update cross-connect group operation. Value must not
	// exceed the total number of cross-connects in the cross-connect group.
	MinimumLinks *int `mandatory:"false" json:"minimumLinks"`

	// The flag to enable or disable the down timer for the interface.
	IsInterfaceHoldTimerEnabled *bool `mandatory:"false" json:"isInterfaceHoldTimerEnabled"`

	// The duration of the interface down timer in milliseconds between 0 and 3000 in multiples of 500.
	InterfaceDownTimerValueInMilliseconds *int `mandatory:"false" json:"interfaceDownTimerValueInMilliseconds"`

	// (Optional) When true, restricts placement so cross-connects lands only on QoS-capable devices.
	// When false (default), placement may use any supported device. If no QoS-capable devices are available
	// in the selected location, the request fails.
	IsQosEnabled *bool `mandatory:"false" json:"isQosEnabled"`
}

func (m CreateCrossConnectGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCrossConnectGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
