// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDatabaseInsightDetails The information to be updated.
type UpdateDatabaseInsightDetails interface {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatedatabaseinsightdetails struct {
	JsonData     []byte
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	EntitySource string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabaseinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabaseinsightdetails updatedatabaseinsightdetails
	s := struct {
		Model Unmarshalerupdatedatabaseinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabaseinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_EXTERNAL_DATABASE":
		mm := UpdateMacsManagedExternalDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EM_MANAGED_EXTERNAL_DATABASE":
		mm := UpdateEmManagedExternalDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE_COMANAGED_DATABASE":
		mm := UpdatePeComanagedDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := UpdateAutonomousDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MDS_MYSQL_DATABASE_SYSTEM":
		mm := UpdateMdsMySqlDatabaseInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateDatabaseInsightDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m updatedatabaseinsightdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatedatabaseinsightdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatedatabaseinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabaseinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
