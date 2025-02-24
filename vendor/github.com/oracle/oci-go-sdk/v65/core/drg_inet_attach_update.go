// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DrgInetAttachUpdate Data plane description for DRG internet attachment.
type DrgInetAttachUpdate struct {

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

	// The label given to the target drg attachment. If the target drg attachment is of
	// type VCN or Internet attachment, then this will be the label of the corresponding
	// VCN or Internet attachment. If the target drg attachment if of type Remote Peering
	// Connection, then this will be the label of the peered RPC attachment.
	DrgAttachmentLabel *int `mandatory:"false" json:"drgAttachmentLabel"`

	// The route data of internet attachment
	RouteCidr *string `mandatory:"false" json:"routeCidr"`

	// The drg “ingress” redirector vip (always IPv4) in the region of the peered RPC attachment
	// in case of RPC attachment. Only applies to RPC & Internet attachments.
	SubstrateNextHopIpAddress *string `mandatory:"false" json:"substrateNextHopIpAddress"`

	// The IPv4 address, in dot-decimal notation, used to encapsulate traffic routed to this internet gateway
	DrgVip *string `mandatory:"false" json:"drgVip"`
}

// GetId returns Id
func (m DrgInetAttachUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgInetAttachUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgInetAttachUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgInetAttachUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgInetAttachUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgInetAttachUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgInetAttachUpdate DrgInetAttachUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgInetAttachUpdate
	}{
		"DrgInetAttachUpdate",
		(MarshalTypeDrgInetAttachUpdate)(m),
	}

	return json.Marshal(&s)
}
