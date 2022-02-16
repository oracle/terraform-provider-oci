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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateManagedSshSessionTargetResourceDetails Details about a managed SSH session for a target resource.
type CreateManagedSshSessionTargetResourceDetails struct {

	// The name of the user on the target resource operating system that the session uses for the connection.
	TargetResourceOperatingSystemUserName *string `mandatory:"true" json:"targetResourceOperatingSystemUserName"`

	// The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// The port number to connect to on the target resource.
	TargetResourcePort *int `mandatory:"false" json:"targetResourcePort"`

	// The private IP address of the target resource that the session connects to.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

//GetTargetResourcePort returns TargetResourcePort
func (m CreateManagedSshSessionTargetResourceDetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m CreateManagedSshSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagedSshSessionTargetResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateManagedSshSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateManagedSshSessionTargetResourceDetails CreateManagedSshSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeCreateManagedSshSessionTargetResourceDetails
	}{
		"MANAGED_SSH",
		(MarshalTypeCreateManagedSshSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
