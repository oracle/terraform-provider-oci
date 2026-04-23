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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchSubnetReplaceInstruction Replaces the entire value of the selected subnet CIDR field with the specified final state.
// For IPv6 CIDR list selections (for example, `ipv6CidrBlocks`), the supplied array is treated
// as the authoritative set of CIDRs for that field:
//   - CIDRs present in both the existing list and the new list remain unchanged.
//   - CIDRs present in the existing list but not in the new list are removed.
//   - CIDRs present in the new list but not in the existing list are added.
type PatchSubnetReplaceInstruction struct {

	// The set of values to which the operation applies as a JMESPath expression (https://jmespath.org/specification.html) for evaluation
	// against the Subnet resource representation.
	// The PatchSubnet operation restricts supported selections (see PatchSubnet documentation).
	// Example: "ipv6CidrBlocks"
	Selection *string `mandatory:"true" json:"selection"`

	// The desired final IPv6 CIDR value(s) to apply to the selected field. This field must
	// always be a JSON object.
	// For fields that take a single CIDR (for example, `ipv6CidrBlock`), specify a single element.
	// For list fields (for example, `ipv6CidrBlocks`), specify the full desired list.
	// Examples:
	// - { "operation": "REPLACE", "selection": "ipv6CidrBlocks", "value": { "cidrs": [ "2001:db8:1234:1111::/64", "2001:db8:1234:2121::/64" ] } }
	// - { "operation": "REPLACE", "selection": "ipv6CidrBlock", "value": { "cidr": "2001:db8:1234:1111::/64" } }
	Value *interface{} `mandatory:"true" json:"value"`
}

// GetSelection returns Selection
func (m PatchSubnetReplaceInstruction) GetSelection() *string {
	return m.Selection
}

func (m PatchSubnetReplaceInstruction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchSubnetReplaceInstruction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchSubnetReplaceInstruction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchSubnetReplaceInstruction PatchSubnetReplaceInstruction
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchSubnetReplaceInstruction
	}{
		"REPLACE",
		(MarshalTypePatchSubnetReplaceInstruction)(m),
	}

	return json.Marshal(&s)
}
