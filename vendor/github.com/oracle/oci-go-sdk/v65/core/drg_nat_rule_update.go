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

// DrgNatRuleUpdate Data plane information about the network address translated
// route rules associated with a DRG attachment.
type DrgNatRuleUpdate struct {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	Id *string `mandatory:"false" json:"id"`

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	IsDelete *bool `mandatory:"false" json:"isDelete"`

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The label which identifies this DRG NAT rule in encapsulated
	// traffic sent through the DRG attachment associated with the
	// DRG NAT policy.
	DrgNatRuleLabel *int64 `mandatory:"false" json:"drgNatRuleLabel"`

	// The label which identifies this DRG NAT policy to derive the DRG NAT
	// rules in traffic sent through the DRG attachment associated with
	// the DRG NAT policy.
	DrgNatPolicyLabel *int64 `mandatory:"false" json:"drgNatPolicyLabel"`

	// The priority associated with each DRG NAT rule.
	DrgNatRulePriority *int64 `mandatory:"false" json:"drgNatRulePriority"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Original CIDR range for Source NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	OriginalSource *string `mandatory:"false" json:"originalSource"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Translated CIDR range for Source NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	TranslatedSource *string `mandatory:"false" json:"translatedSource"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Original CIDR range for Destination NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	OriginalDestination *string `mandatory:"false" json:"originalDestination"`

	// Represents the range of IP addresses to match against when routing traffic.
	// Translated CIDR range for Destination NAT.
	// Potential values:
	//   * An IPv4 address range in CIDR notation. For example: `192.168.1.0/24`.
	TranslatedDestination *string `mandatory:"false" json:"translatedDestination"`
}

// GetId returns Id
func (m DrgNatRuleUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgNatRuleUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgNatRuleUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgNatRuleUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgNatRuleUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgNatRuleUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgNatRuleUpdate DrgNatRuleUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgNatRuleUpdate
	}{
		"DrgNatRuleUpdate",
		(MarshalTypeDrgNatRuleUpdate)(m),
	}

	return json.Marshal(&s)
}
