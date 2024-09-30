// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Zero Trust Packet Routing Control Plane API
//
// Use the Zero Trust Packet Routing Control Plane API to manage ZPR configuration and policy. See the Zero Trust Packet Routing (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/home.htm) documentation for more information.
//

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateZprPolicyDetails The data to create a ZprPolicy.
type CreateZprPolicyDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the ZprPolicy in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The description you assign to the ZprPolicy during creation. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description"`

	// The name you assign to the ZprPolicy during creation. The name must be unique across all ZPL policies in the tenancy.
	Name *string `mandatory:"true" json:"name"`

	// An array of ZprPolicy statements(up to 25 statements per ZprPolicy) written in the Zero Trust Packet Routing Policy Language.
	Statements []string `mandatory:"true" json:"statements"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateZprPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateZprPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
