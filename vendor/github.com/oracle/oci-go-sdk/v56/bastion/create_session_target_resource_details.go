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

// CreateSessionTargetResourceDetails Details about a bastion session's target resource.
type CreateSessionTargetResourceDetails interface {

	// The port number to connect to on the target resource.
	GetTargetResourcePort() *int
}

type createsessiontargetresourcedetails struct {
	JsonData           []byte
	TargetResourcePort *int   `mandatory:"false" json:"targetResourcePort"`
	SessionType        string `json:"sessionType"`
}

// UnmarshalJSON unmarshals json
func (m *createsessiontargetresourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatesessiontargetresourcedetails createsessiontargetresourcedetails
	s := struct {
		Model Unmarshalercreatesessiontargetresourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetResourcePort = s.Model.TargetResourcePort
	m.SessionType = s.Model.SessionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createsessiontargetresourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SessionType {
	case "MANAGED_SSH":
		mm := CreateManagedSshSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PORT_FORWARDING":
		mm := CreatePortForwardingSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTargetResourcePort returns TargetResourcePort
func (m createsessiontargetresourcedetails) GetTargetResourcePort() *int {
	return m.TargetResourcePort
}

func (m createsessiontargetresourcedetails) String() string {
	return common.PointerString(m)
}
