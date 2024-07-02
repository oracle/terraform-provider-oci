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

// UxConfigurationItemsCollection Collection of ux configuration item summary objects.
type UxConfigurationItemsCollection struct {

	// Array of configuration item summary objects.
	ConfigItems []ConfigurationItemSummary `mandatory:"false" json:"configItems"`
}

// GetConfigItems returns ConfigItems
func (m UxConfigurationItemsCollection) GetConfigItems() []ConfigurationItemSummary {
	return m.ConfigItems
}

func (m UxConfigurationItemsCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UxConfigurationItemsCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UxConfigurationItemsCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUxConfigurationItemsCollection UxConfigurationItemsCollection
	s := struct {
		DiscriminatorParam string `json:"opsiConfigType"`
		MarshalTypeUxConfigurationItemsCollection
	}{
		"UX_CONFIGURATION",
		(MarshalTypeUxConfigurationItemsCollection)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UxConfigurationItemsCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConfigItems []configurationitemsummary `json:"configItems"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConfigItems = make([]ConfigurationItemSummary, len(model.ConfigItems))
	for i, n := range model.ConfigItems {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ConfigItems[i] = nn.(ConfigurationItemSummary)
		} else {
			m.ConfigItems[i] = nil
		}
	}
	return
}
