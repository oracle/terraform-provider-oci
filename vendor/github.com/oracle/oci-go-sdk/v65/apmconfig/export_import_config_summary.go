// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportImportConfigSummary A description of a configuration item or dependency. It specifies all the properties that define the configuration item or dependency that will be exported.
type ExportImportConfigSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
	// when the item is created.
	GetId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type exportimportconfigsummary struct {
	JsonData     []byte
	Id           *string                           `mandatory:"false" json:"id"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigType   string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *exportimportconfigsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexportimportconfigsummary exportimportconfigsummary
	s := struct {
		Model Unmarshalerexportimportconfigsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *exportimportconfigsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "OPTIONS":
		mm := ExportImportOptionsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APDEX":
		mm := ExportImportApdexRulesSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPAN_FILTER":
		mm := ExportImportSpanFilterSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "METRIC_GROUP":
		mm := ExportImportMetricGroupSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExportImportConfigSummary: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetId returns Id
func (m exportimportconfigsummary) GetId() *string {
	return m.Id
}

// GetFreeformTags returns FreeformTags
func (m exportimportconfigsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m exportimportconfigsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m exportimportconfigsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m exportimportconfigsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
