// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TcpService TCP Service used on the firewall policy rules.
type TcpService struct {

	// Name of the service.
	Name *string `mandatory:"true" json:"name"`

	// OCID of the Network Firewall Policy this service belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`

	// List of port-ranges used.
	PortRanges []PortRange `mandatory:"true" json:"portRanges"`
}

// GetName returns Name
func (m TcpService) GetName() *string {
	return m.Name
}

// GetParentResourceId returns ParentResourceId
func (m TcpService) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m TcpService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TcpService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TcpService) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTcpService TcpService
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTcpService
	}{
		"TCP_SERVICE",
		(MarshalTypeTcpService)(m),
	}

	return json.Marshal(&s)
}
