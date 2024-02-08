// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportableComputeEntitySummary A compute entity that can be imported into Operations Insights.
type ImportableComputeEntitySummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
	GetComputeId() *string

	// The Display Name (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Compute Instance
	GetComputeDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string
}

type importablecomputeentitysummary struct {
	JsonData           []byte
	ComputeId          *string `mandatory:"true" json:"computeId"`
	ComputeDisplayName *string `mandatory:"true" json:"computeDisplayName"`
	CompartmentId      *string `mandatory:"true" json:"compartmentId"`
	EntitySource       string  `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *importablecomputeentitysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimportablecomputeentitysummary importablecomputeentitysummary
	s := struct {
		Model Unmarshalerimportablecomputeentitysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComputeId = s.Model.ComputeId
	m.ComputeDisplayName = s.Model.ComputeDisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *importablecomputeentitysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_CLOUD_HOST":
		mm := CloudImportableComputeEntitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ImportableComputeEntitySummary: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetComputeId returns ComputeId
func (m importablecomputeentitysummary) GetComputeId() *string {
	return m.ComputeId
}

// GetComputeDisplayName returns ComputeDisplayName
func (m importablecomputeentitysummary) GetComputeDisplayName() *string {
	return m.ComputeDisplayName
}

// GetCompartmentId returns CompartmentId
func (m importablecomputeentitysummary) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m importablecomputeentitysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m importablecomputeentitysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
