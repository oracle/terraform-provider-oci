// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSecurityContextDetails Security context for container.
type CreateSecurityContextDetails interface {
}

type createsecuritycontextdetails struct {
	JsonData            []byte
	SecurityContextType string `json:"securityContextType"`
}

// UnmarshalJSON unmarshals json
func (m *createsecuritycontextdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatesecuritycontextdetails createsecuritycontextdetails
	s := struct {
		Model Unmarshalercreatesecuritycontextdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SecurityContextType = s.Model.SecurityContextType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createsecuritycontextdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SecurityContextType {
	case "LINUX":
		mm := CreateLinuxSecurityContextDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateSecurityContextDetails: %s.", m.SecurityContextType)
		return *m, nil
	}
}

func (m createsecuritycontextdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createsecuritycontextdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
