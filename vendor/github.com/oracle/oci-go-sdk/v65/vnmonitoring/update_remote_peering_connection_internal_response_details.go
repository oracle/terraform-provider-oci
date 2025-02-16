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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRemotePeeringConnectionInternalResponseDetails Details on the existing peering connection to remote region.
type UpdateRemotePeeringConnectionInternalResponseDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG that has the remote peering connection.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The route target of the DRG that has the remote peering connection.
	DrgRouteTarget *string `mandatory:"true" json:"drgRouteTarget"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the tenant that contain the remote peering connection.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The route targets of the GFC attachments of the DRG that has the remote peering connection.
	GfcRouteTargets []GfcRouteTarget `mandatory:"false" json:"gfcRouteTargets"`

	// Ingress VIP of the acceptors DRG.
	IngressVip *string `mandatory:"false" json:"ingressVip"`
}

func (m UpdateRemotePeeringConnectionInternalResponseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRemotePeeringConnectionInternalResponseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
