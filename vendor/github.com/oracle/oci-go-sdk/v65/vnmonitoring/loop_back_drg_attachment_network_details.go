// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoopBackDrgAttachmentNetworkDetails Specifies the loopback attachment on the DRG. A loopback attachment can be used to terminate a virtual circuit that is carrying an IPSec tunnel, routing traffic directly to the IPSec tunnel attachment where the tunnel can terminate.
type LoopBackDrgAttachmentNetworkDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target IPSec tunnel attachment.
	Ids []string `mandatory:"false" json:"ids"`

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
