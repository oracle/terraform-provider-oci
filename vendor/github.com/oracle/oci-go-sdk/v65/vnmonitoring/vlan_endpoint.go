// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// VlanEndpoint Defines the details required for a VLAN-type `Endpoint`.
type VlanEndpoint struct {

	// The IPv4 address of the `Endpoint`.
	Address *string `mandatory:"true" json:"address"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN containing the IP address.
	// This can be used to disambiguate which VLAN is queried, in case the endpoint IP
	// address belongs to more than one VLAN (when there are VLANs with overlapping IP ranges).
	VlanId *string `mandatory:"true" json:"vlanId"`
}

func (m VlanEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VlanEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VlanEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVlanEndpoint VlanEndpoint
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVlanEndpoint
	}{
		"VLAN",
		(MarshalTypeVlanEndpoint)(m),
	}

	return json.Marshal(&s)
}
