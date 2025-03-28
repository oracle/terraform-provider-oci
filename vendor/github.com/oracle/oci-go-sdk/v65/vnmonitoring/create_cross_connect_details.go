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

// CreateCrossConnectDetails The representation of CreateCrossConnectDetails
type CreateCrossConnectDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the cross-connect.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the FastConnect location where this cross-connect will be installed.
	// To get a list of the available locations, see
	// ListCrossConnectLocations.
	// Example: `CyrusOne, Chandler, AZ`
	LocationName *string `mandatory:"true" json:"locationName"`

	// The port speed for this cross-connect. To get a list of the available port speeds, see
	// ListCrossconnectPortSpeedShapes.
	// Example: `10 Gbps`
	PortSpeedShapeName *string `mandatory:"true" json:"portSpeedShapeName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId *string `mandatory:"false" json:"crossConnectGroupId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If you already have an existing cross-connect or cross-connect group at this FastConnect
	// location, and you want this new cross-connect to be on a different router (for the
	// purposes of redundancy), provide the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that existing cross-connect or
	// cross-connect group.
	FarCrossConnectOrCrossConnectGroupId *string `mandatory:"false" json:"farCrossConnectOrCrossConnectGroupId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// If you already have an existing cross-connect or cross-connect group at this FastConnect
	// location, and you want this new cross-connect to be on the same router, provide the
	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId *string `mandatory:"false" json:"nearCrossConnectOrCrossConnectGroupId"`

	// A reference name or identifier for the physical fiber connection that this cross-connect
	// uses.
	CustomerReferenceName *string `mandatory:"false" json:"customerReferenceName"`

	MacsecProperties *CreateMacsecProperties `mandatory:"false" json:"macsecProperties"`
}

func (m CreateCrossConnectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCrossConnectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
