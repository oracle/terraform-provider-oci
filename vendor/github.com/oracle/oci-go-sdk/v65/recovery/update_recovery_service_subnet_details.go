// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRecoveryServiceSubnetDetails Describes the parameters required to update a recovery service subnet.
type UpdateRecoveryServiceSubnetDetails struct {

	// A user-provided name for the recovery service subnet. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A list of OCIDs of the subnets associated with the recovery service subnet.
	Subnets []string `mandatory:"false" json:"subnets"`

	// A list of network security group (NSG) OCIDs that are associated with the Recovery Service subnet.
	// You can specify a maximum of 5 unique OCIDs, which implies that you can associate a maximum of 5 NSGs to each Recovery Service subnet.
	// Specify an empty array if you want to remove all the associated NSGs from a Recovery Service subnet.
	// See NetworkSecurityGroup for more information.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateRecoveryServiceSubnetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRecoveryServiceSubnetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
