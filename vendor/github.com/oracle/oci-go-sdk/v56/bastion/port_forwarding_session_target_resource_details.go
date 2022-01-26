// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PortForwardingSessionTargetResourceDetails Details about a port forwarding session for a target resource.
type PortForwardingSessionTargetResourceDetails struct {

	// The port number to connect to on the target resource.
	TargetResourcePort *int `mandatory:"true" json:"targetResourcePort"`

	// The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// The private IP address of the target resource that the session connects to.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`

	// The display name of the target Compute instance that the session connects to.
	TargetResourceDisplayName *string `mandatory:"false" json:"targetResourceDisplayName"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m PortForwardingSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m PortForwardingSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PortForwardingSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePortForwardingSessionTargetResourceDetails PortForwardingSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypePortForwardingSessionTargetResourceDetails
	}{
		"PORT_FORWARDING",
		(MarshalTypePortForwardingSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
