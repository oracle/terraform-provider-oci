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

// PatchVcnDetails The request to patch the VCN.
// Example:
//
//	{
//	  "patchVcnInstructions": [
//	    {
//	      "operation": "REPLACE",
//	      "selection": "ipv6CidrBlock",
//	      "value": {"cidr": "2001::/56"}
//	    },
//	    {
//	      "operation": "REPLACE",
//	      "selection": "ipv6PublicCidrBlock",
//	      "value": {"cidr": "2001:0db8:0123::/48"}
//	    },
//	    {
//	      "operation": "REPLACE",
//	      "selection": "byoipv6CidrDetails",
//	      "value": {
//	        "cidrs": [
//	          {
//	            "byoipv6RangeId": "ocid1.byoiprange.oc1.<unique_ID_1>",
//	            "ipv6CidrBlock": "2001:0db8:0123::/48"
//	          },
//	          {
//	            "byoipv6RangeId": "ocid1.byoiprange.oc1.<unique_ID_2>",
//	            "ipv6CidrBlock": "2001:0db8:0456::/48"
//	          }
//	        ]
//	      }
//	    },
//	    {
//	      "operation": "REPLACE",
//	      "selection": "ipv6PrivateCidrBlocks",
//	      "value": { "cidrs": ["fd00:1000:0:1::/64", "fd00:1000:0:2::/64"] }
//	    }
//	  ]
//	}
type PatchVcnDetails struct {

	// List of patch instructions for VCN.
	PatchVcnInstructions []PatchVcnInstruction `mandatory:"true" json:"patchVcnInstructions"`
}

func (m PatchVcnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchVcnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PatchVcnDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PatchVcnInstructions []patchvcninstruction `json:"patchVcnInstructions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PatchVcnInstructions = make([]PatchVcnInstruction, len(model.PatchVcnInstructions))
	for i, n := range model.PatchVcnInstructions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PatchVcnInstructions[i] = nn.(PatchVcnInstruction)
		} else {
			m.PatchVcnInstructions[i] = nil
		}
	}
	return
}
