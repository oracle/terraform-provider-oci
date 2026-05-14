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

// PatchSubnetDetails The request to patch the subnet.
// Example:
//
//	{
//	  "patchSubnetInstructions": [
//	    {
//	      "operation": "REPLACE",
//	      "selection": "ipv6CidrBlock",
//	      "value": {"cidr": "2001::/56"}
//	    },
//	    {
//	      "operation": "REPLACE",
//	      "selection": "ipv6CidrBlocks",
//	      "value": { "cidrs": [ "2001:db8:1234:1111::/64", "2001:db8:1234:2121::/64" ] }
//	    }
//	  ]
//	}
type PatchSubnetDetails struct {

	// List of patch instructions for Subnet.
	PatchSubnetInstructions []PatchSubnetInstruction `mandatory:"true" json:"patchSubnetInstructions"`
}

func (m PatchSubnetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchSubnetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PatchSubnetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PatchSubnetInstructions []patchsubnetinstruction `json:"patchSubnetInstructions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PatchSubnetInstructions = make([]PatchSubnetInstruction, len(model.PatchSubnetInstructions))
	for i, n := range model.PatchSubnetInstructions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PatchSubnetInstructions[i] = nn.(PatchSubnetInstruction)
		} else {
			m.PatchSubnetInstructions[i] = nil
		}
	}
	return
}
