// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlexTunnelDrgAttachmentNetworkDetails Specifies the flex tunnel attached to the DRG.
type FlexTunnelDrgAttachmentNetworkDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the transport drg attachment of the flex tunnel.
	TransportAttachmentId *string `mandatory:"false" json:"transportAttachmentId"`

	// The BGP ASN to use for the Flex Tunnel connection's route target
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// Routes which may be imported from the attachment (subject to import policy) appear in the route reflectors
	// tagged with the attachment's import route target.
	ImportRouteTarget *string `mandatory:"false" json:"importRouteTarget"`

	// Routes which are exported to the attachment are exported to the route reflectors
	// with the route target set to the value of the attachment's export route target.
	ExportRouteTarget *string `mandatory:"false" json:"exportRouteTarget"`

	// The MPLS label of the DRG attachment.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// IPv4 address used to encapsulate ingress traffic to the DRG through this attachment
	IngressVip *string `mandatory:"false" json:"ingressVip"`
}

// GetId returns Id
func (m FlexTunnelDrgAttachmentNetworkDetails) GetId() *string {
	return m.Id
}

func (m FlexTunnelDrgAttachmentNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexTunnelDrgAttachmentNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FlexTunnelDrgAttachmentNetworkDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFlexTunnelDrgAttachmentNetworkDetails FlexTunnelDrgAttachmentNetworkDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFlexTunnelDrgAttachmentNetworkDetails
	}{
		"FLEX_TUNNEL",
		(MarshalTypeFlexTunnelDrgAttachmentNetworkDetails)(m),
	}

	return json.Marshal(&s)
}
