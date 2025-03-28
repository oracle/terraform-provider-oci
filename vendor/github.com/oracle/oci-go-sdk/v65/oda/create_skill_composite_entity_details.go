// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSkillCompositeEntityDetails Properties that are required to create a skill composite entity.
type CreateSkillCompositeEntityDetails struct {

	// The entity name. This must be unique within the parent resource.
	Name *string `mandatory:"true" json:"name"`

	// List of entity attributes.
	Attributes []EntityAttribute `mandatory:"true" json:"attributes"`

	// List of entity actions.
	Actions []EntityAction `mandatory:"false" json:"actions"`

	NaturalLanguageMapping *EntityNaturalLanguageMapping `mandatory:"false" json:"naturalLanguageMapping"`
}

// GetName returns Name
func (m CreateSkillCompositeEntityDetails) GetName() *string {
	return m.Name
}

func (m CreateSkillCompositeEntityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSkillCompositeEntityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSkillCompositeEntityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSkillCompositeEntityDetails CreateSkillCompositeEntityDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateSkillCompositeEntityDetails
	}{
		"COMPOSITE",
		(MarshalTypeCreateSkillCompositeEntityDetails)(m),
	}

	return json.Marshal(&s)
}
