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

// ConfigurationItemsCollection Collection of configuration item summary objects.
type ConfigurationItemsCollection interface {

	// Array of configuration item summary objects.
	GetConfigItems() []ConfigurationItemSummary
}

type configurationitemscollection struct {
	JsonData       []byte
	ConfigItems    json.RawMessage `mandatory:"false" json:"configItems"`
	OpsiConfigType string          `json:"opsiConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *configurationitemscollection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigurationitemscollection configurationitemscollection
	s := struct {
		Model Unmarshalerconfigurationitemscollection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigItems = s.Model.ConfigItems
	m.OpsiConfigType = s.Model.OpsiConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configurationitemscollection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OpsiConfigType {
	case "UX_CONFIGURATION":
		mm := UxConfigurationItemsCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConfigurationItemsCollection: %s.", m.OpsiConfigType)
		return *m, nil
	}
}

// GetConfigItems returns ConfigItems
func (m configurationitemscollection) GetConfigItems() json.RawMessage {
	return m.ConfigItems
}

func (m configurationitemscollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configurationitemscollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
