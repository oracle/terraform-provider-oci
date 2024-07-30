// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateFsuActionDetails Exadata Fleet Update Action resource details to update.
type UpdateFsuActionDetails interface {

	// Exadata Fleet Update Action display name.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatefsuactiondetails struct {
	JsonData     []byte
	DisplayName  *string                           `mandatory:"false" json:"displayName"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Type         string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatefsuactiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatefsuactiondetails updatefsuactiondetails
	s := struct {
		Model Unmarshalerupdatefsuactiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatefsuactiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STAGE":
		mm := UpdateStageActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := UpdateApplyActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROLLBACK_AND_REMOVE_TARGET":
		mm := UpdateRollbackActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRECHECK":
		mm := UpdatePrecheckActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLEANUP":
		mm := UpdateCleanupActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateFsuActionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatefsuactiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m updatefsuactiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatefsuactiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatefsuactiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatefsuactiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
