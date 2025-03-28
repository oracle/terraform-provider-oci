// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PortForwardingSessionTargetResourceDetails Details about a port forwarding session for a target resource.
type PortForwardingSessionTargetResourceDetails struct {

	// The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// The private IP address of the target resource that the session connects to.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`

	// The display name of the target Compute instance that the session connects to.
	TargetResourceDisplayName *string `mandatory:"false" json:"targetResourceDisplayName"`

	// The Fully Qualified Domain Name of the target resource that the session connects to.
	TargetResourceFqdn *string `mandatory:"false" json:"targetResourceFqdn"`

	// The port number to connect to on the target resource.
	TargetResourcePort *int `mandatory:"false" json:"targetResourcePort"`
}

func (m PortForwardingSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PortForwardingSessionTargetResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
