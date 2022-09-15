// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoopBackDrgAttachmentNetworkDetails Specifies the Loopback attached to the DRG.
type LoopBackDrgAttachmentNetworkDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Ipsec tunnel attachment.
	FrontEndAttachmentId []string `mandatory:"false" json:"frontEndAttachmentId"`

	// The IP address for loopback.
	LoopBackIp *string `mandatory:"false" json:"loopBackIp"`

	// The MPLS label of the DRG attachment.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// Routes which may be imported from the attachment (subject to import policy) appear in the route reflectors
	// tagged with the attachment's import route target.
	ImportRouteTarget *string `mandatory:"false" json:"importRouteTarget"`

	// Routes which are exported to the attachment are exported to the route reflectors
	// with the route target set to the value of the attachment's export route target.
	ExportRouteTarget *string `mandatory:"false" json:"exportRouteTarget"`
}

//GetId returns Id
func (m LoopBackDrgAttachmentNetworkDetails) GetId() *string {
	return m.Id
}

func (m LoopBackDrgAttachmentNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoopBackDrgAttachmentNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LoopBackDrgAttachmentNetworkDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLoopBackDrgAttachmentNetworkDetails LoopBackDrgAttachmentNetworkDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeLoopBackDrgAttachmentNetworkDetails
	}{
		"LOOPBACK",
		(MarshalTypeLoopBackDrgAttachmentNetworkDetails)(m),
	}

	return json.Marshal(&s)
}
