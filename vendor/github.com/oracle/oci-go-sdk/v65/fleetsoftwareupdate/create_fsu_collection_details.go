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

// CreateFsuCollectionDetails The information about new Exadata Fleet Update Collection.
type CreateFsuCollectionDetails interface {

	// Exadata service type for the target resource members.
	GetServiceType() CollectionServiceTypesEnum

	// Compartment Identifier
	GetCompartmentId() *string

	// Exadata Fleet Update Collection Identifier.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createfsucollectiondetails struct {
	JsonData      []byte
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ServiceType   CollectionServiceTypesEnum        `mandatory:"true" json:"serviceType"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	Type          string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createfsucollectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatefsucollectiondetails createfsucollectiondetails
	s := struct {
		Model Unmarshalercreatefsucollectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ServiceType = s.Model.ServiceType
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createfsucollectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DB":
		mm := CreateDbFsuCollectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GI":
		mm := CreateGiFsuCollectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateFsuCollectionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createfsucollectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createfsucollectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createfsucollectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetServiceType returns ServiceType
func (m createfsucollectiondetails) GetServiceType() CollectionServiceTypesEnum {
	return m.ServiceType
}

// GetCompartmentId returns CompartmentId
func (m createfsucollectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createfsucollectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createfsucollectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectionServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetCollectionServiceTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
