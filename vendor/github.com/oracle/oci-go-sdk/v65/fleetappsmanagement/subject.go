// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Subject Subject.
type Subject interface {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	GetDisplayName() *string
}

type subject struct {
	JsonData    []byte
	DisplayName *string `mandatory:"false" json:"displayName"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *subject) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersubject subject
	s := struct {
		Model Unmarshalersubject
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *subject) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RECOMMENDED_PATCHES":
		mm := RecommendedPatchesBasedSubject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONE_OFF":
		mm := OneOffBasedSubject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATCH_GROUP":
		mm := PatchGroupBasedSubject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Subject: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m subject) GetDisplayName() *string {
	return m.DisplayName
}

func (m subject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m subject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
