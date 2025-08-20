// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCccFlexNetworkDetails Updates Compute Cloud@Customer flexNetwork configuration details.
type UpdateCccFlexNetworkDetails struct {

	// The name that will be used to display the Compute Cloud@Customer FlexNetwork
	// in the Oracle Cloud Console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer flexNetwork.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// An array of switch ports to be used to establish the network connection.  The ports are represented
	// using a string to support breakout mode.  As an example if switch port 7 is a 100Gbps port, it can
	// be configured in breakout mode as four (4) smaller 25Gbps ports, 7/1, 7/2, 7/3 and 7/4.  The minimum
	// entry will be a single port X, where the maximum entry is XX/X.  This is a required field. The
	// minimum number of entries is 1 while the maximum number of entries is 4.
	Ports []string `mandatory:"false" json:"ports"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCccFlexNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCccFlexNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
