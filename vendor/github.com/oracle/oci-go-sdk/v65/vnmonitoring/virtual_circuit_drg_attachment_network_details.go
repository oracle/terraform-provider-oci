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

// VirtualCircuitDrgAttachmentNetworkDetails Specifies the virtual circuit attached to the DRG.
type VirtualCircuitDrgAttachmentNetworkDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"false" json:"id"`

	// Whether the Fast Connect is an FFAB VirtualCircuit.
	// Example: `true`
	IsFFAB *bool `mandatory:"false" json:"isFFAB"`

	// This indicates whether FastConnect extends through an edge POP region.
	// Example: `true`
	IsEdgePop *bool `mandatory:"false" json:"isEdgePop"`

	// Routes which may be imported from the attachment (subject to import policy) appear in the route reflectors
	// tagged with the attachment's import route target.
	ImportRouteTarget *string `mandatory:"false" json:"importRouteTarget"`

	// Routes which are exported to the attachment are exported to the route reflectors
	// with the route target set to the value of the attachment's export route target.
	ExportRouteTarget *string `mandatory:"false" json:"exportRouteTarget"`

	// The MPLS label of the DRG attachment.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// The BGP ASN to use for the IPSec connection's route target.
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// The Oracle Cloud Infrastructure region name.
	RegionName *string `mandatory:"false" json:"regionName"`

	// Common Export route target to use for the DRG Attachment instead of per-attachment route target.
	// This is applicable to DRG attachments that are assigned to a DRG route table which is whitelisted for high
	// throughput mode.
	CommonExportRT *string `mandatory:"false" json:"commonExportRT"`

	// Common Import route target to use for the DRG attachment instead of per-attachment import route target.
	// This is applicable to DRG attachments that are assigned to a DRG route table which is whitelisted for high
	// throughput mode.
	CommonImportRT *string `mandatory:"false" json:"commonImportRT"`

	// This indicates whether the DRG route table associated with the DRG attachment is whitelisted for High Throughput mode.
	// Example: `true`
	IsIngressDisintermediationEnabled *bool `mandatory:"false" json:"isIngressDisintermediationEnabled"`

	// This indicates whether the DRG route table associated with the DRG attachment is whitelisted for route unification.
	// Example: `true`
	IsRouteUnificationEnabled *bool `mandatory:"false" json:"isRouteUnificationEnabled"`

	// Boolean flag that determines wether all traffic over the virtual circuits is encrypted.
	// Example: `true`
	TransportOnlyMode *bool `mandatory:"false" json:"transportOnlyMode"`
}

// GetId returns Id
func (m VirtualCircuitDrgAttachmentNetworkDetails) GetId() *string {
	return m.Id
}

func (m VirtualCircuitDrgAttachmentNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualCircuitDrgAttachmentNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VirtualCircuitDrgAttachmentNetworkDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVirtualCircuitDrgAttachmentNetworkDetails VirtualCircuitDrgAttachmentNetworkDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVirtualCircuitDrgAttachmentNetworkDetails
	}{
		"VIRTUAL_CIRCUIT",
		(MarshalTypeVirtualCircuitDrgAttachmentNetworkDetails)(m),
	}

	return json.Marshal(&s)
}
