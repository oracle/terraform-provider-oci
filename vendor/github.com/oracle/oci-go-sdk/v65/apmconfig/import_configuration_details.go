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

// ImportConfigurationDetails Array of configuration items with dependencies to import.
type ImportConfigurationDetails struct {

	// A list of Configurations Details .
	ConfigurationItems []ExportImportConfigSummary `mandatory:"true" json:"configurationItems"`

	// Simple key-value pair that has parameters related to the import process (EnableOcidSubstitution, Skip, …) and more.
	// Example: `{"parameter-key": "parameter-value"}`
	// Supported parameters:
	// — Enable the OCIDs in instructions to be replaced, if set to "true" The Config Service replace any OCIDs it finds
	// in the instructions, if set to true the Config Service will replace.
	// — Compartment Id we want to import the configuration Items,
	// if the compartment Id is not provided it will be the default destination domain compartmentId.
	ConfigurationMap map[string]string `mandatory:"true" json:"configurationMap"`
}

func (m ImportConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImportConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConfigurationItems []exportimportconfigsummary `json:"configurationItems"`
		ConfigurationMap   map[string]string           `json:"configurationMap"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConfigurationItems = make([]ExportImportConfigSummary, len(model.ConfigurationItems))
	for i, n := range model.ConfigurationItems {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ConfigurationItems[i] = nn.(ExportImportConfigSummary)
		} else {
			m.ConfigurationItems[i] = nil
		}
	}
	m.ConfigurationMap = model.ConfigurationMap

	return
}
