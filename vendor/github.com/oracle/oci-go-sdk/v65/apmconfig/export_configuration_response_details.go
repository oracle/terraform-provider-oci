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

// ExportConfigurationResponseDetails Array of configuration items with its dependencies to export.
type ExportConfigurationResponseDetails struct {

	// A list of Configurations Details .
	ConfigurationItems []ExportImportConfigSummary `mandatory:"true" json:"configurationItems"`
}

func (m ExportConfigurationResponseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportConfigurationResponseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExportConfigurationResponseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConfigurationItems []exportimportconfigsummary `json:"configurationItems"`
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
	return
}
