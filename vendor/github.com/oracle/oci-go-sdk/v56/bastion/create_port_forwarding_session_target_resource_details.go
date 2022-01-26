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

// CreatePortForwardingSessionTargetResourceDetails Details about a port forwarding session for a target resource.
type CreatePortForwardingSessionTargetResourceDetails struct {

	// The port number to connect to on the target resource.
	TargetResourcePort *int `mandatory:"false" json:"targetResourcePort"`

	// The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// The private IP address of the target resource that the session connects to.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m CreatePortForwardingSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m CreatePortForwardingSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreatePortForwardingSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePortForwardingSessionTargetResourceDetails CreatePortForwardingSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeCreatePortForwardingSessionTargetResourceDetails
	}{
		"PORT_FORWARDING",
		(MarshalTypeCreatePortForwardingSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
