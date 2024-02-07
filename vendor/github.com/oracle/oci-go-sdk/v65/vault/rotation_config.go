// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RotationConfig Defines the frequency of the rotation and the information about the target system
type RotationConfig struct {
	TargetSystemDetails TargetSystemDetails `mandatory:"true" json:"targetSystemDetails"`

	// The time interval that indicates the frequency for rotating secret data, as described in ISO 8601 format.
	// The minimum value is 1 day and maximum value is 360 days.
	// For example, if you want to set the time interval for rotating a secret data as 30 days, the duration is expressed as "P30D."
	RotationInterval *string `mandatory:"false" json:"rotationInterval"`

	// Enables auto rotation, when set to true rotationInterval must be set.
	IsScheduledRotationEnabled *bool `mandatory:"false" json:"isScheduledRotationEnabled"`
}

func (m RotationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RotationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RotationInterval           *string             `json:"rotationInterval"`
		IsScheduledRotationEnabled *bool               `json:"isScheduledRotationEnabled"`
		TargetSystemDetails        targetsystemdetails `json:"targetSystemDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.RotationInterval = model.RotationInterval

	m.IsScheduledRotationEnabled = model.IsScheduledRotationEnabled

	nn, e = model.TargetSystemDetails.UnmarshalPolymorphicJSON(model.TargetSystemDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetSystemDetails = nn.(TargetSystemDetails)
	} else {
		m.TargetSystemDetails = nil
	}

	return
}
