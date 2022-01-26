// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// IngestDatabaseConfigurationDetails Database Configuration Metrics details.
type IngestDatabaseConfigurationDetails struct {

	// Array of one or more database configuration metrics objects.
	Items []DatabaseConfigurationMetricGroup `mandatory:"true" json:"items"`
}

func (m IngestDatabaseConfigurationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *IngestDatabaseConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []databaseconfigurationmetricgroup `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DatabaseConfigurationMetricGroup, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DatabaseConfigurationMetricGroup)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
