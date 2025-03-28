// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateHostInsightDetails The information about the host to be analyzed.
type CreateHostInsightDetails interface {

	// Compartment Identifier of host
	GetCompartmentId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createhostinsightdetails struct {
	JsonData      []byte
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	EntitySource  string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *createhostinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatehostinsightdetails createhostinsightdetails
	s := struct {
		Model Unmarshalercreatehostinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createhostinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_CLOUD_HOST":
		mm := CreateMacsManagedCloudHostInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := CreateMacsManagedExternalHostInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EM_MANAGED_EXTERNAL_HOST":
		mm := CreateEmManagedExternalHostInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateHostInsightDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m createhostinsightdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createhostinsightdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createhostinsightdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createhostinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createhostinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
