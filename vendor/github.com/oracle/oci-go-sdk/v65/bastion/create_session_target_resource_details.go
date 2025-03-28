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

// CreateSessionTargetResourceDetails Details about a bastion session's target resource.
type CreateSessionTargetResourceDetails interface {
}

type createsessiontargetresourcedetails struct {
	JsonData    []byte
	SessionType string `json:"sessionType"`
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
	case "DYNAMIC_PORT_FORWARDING":
		mm := CreateDynamicPortForwardingSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PORT_FORWARDING":
		mm := CreatePortForwardingSessionTargetResourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateSessionTargetResourceDetails: %s.", m.SessionType)
		return *m, nil
	}
}

func (m createsessiontargetresourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createsessiontargetresourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
