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

// BasicConfigurationItemSummary Basic configuration item summary.
// Value field contain the most preferred value for the specified scope (compartmentId), which could be from any of the ConfigurationItemValueSourceConfigurationType.
// Default value field contains the default value from Operations Insights.
type BasicConfigurationItemSummary struct {

	// Name of configuration item.
	Name *string `mandatory:"false" json:"name"`

	// Value of configuration item.
	Value *string `mandatory:"false" json:"value"`

	// Value of configuration item.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// List of contexts in Operations Insights where this configuration item is applicable.
	ApplicableContexts []string `mandatory:"false" json:"applicableContexts"`

	Metadata ConfigurationItemMetadata `mandatory:"false" json:"metadata"`

	// Source configuration from where the value is taken for a configuration item.
	ValueSourceConfig ConfigurationItemValueSourceConfigurationTypeEnum `mandatory:"false" json:"valueSourceConfig,omitempty"`
}

func (m BasicConfigurationItemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicConfigurationItemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigurationItemValueSourceConfigurationTypeEnum(string(m.ValueSourceConfig)); !ok && m.ValueSourceConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueSourceConfig: %s. Supported values are: %s.", m.ValueSourceConfig, strings.Join(GetConfigurationItemValueSourceConfigurationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicConfigurationItemSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicConfigurationItemSummary BasicConfigurationItemSummary
	s := struct {
		DiscriminatorParam string `json:"configItemType"`
		MarshalTypeBasicConfigurationItemSummary
	}{
		"BASIC",
		(MarshalTypeBasicConfigurationItemSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *BasicConfigurationItemSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name               *string                                           `json:"name"`
		Value              *string                                           `json:"value"`
		ValueSourceConfig  ConfigurationItemValueSourceConfigurationTypeEnum `json:"valueSourceConfig"`
		DefaultValue       *string                                           `json:"defaultValue"`
		ApplicableContexts []string                                          `json:"applicableContexts"`
		Metadata           configurationitemmetadata                         `json:"metadata"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Value = model.Value

	m.ValueSourceConfig = model.ValueSourceConfig

	m.DefaultValue = model.DefaultValue

	m.ApplicableContexts = make([]string, len(model.ApplicableContexts))
	copy(m.ApplicableContexts, model.ApplicableContexts)
	nn, e = model.Metadata.UnmarshalPolymorphicJSON(model.Metadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Metadata = nn.(ConfigurationItemMetadata)
	} else {
		m.Metadata = nil
	}

	return
}
