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

// HostInsightResourceStatisticsAggregation Contains host details and resource statistics.
type HostInsightResourceStatisticsAggregation struct {
	HostDetails *HostDetails `mandatory:"true" json:"hostDetails"`

	CurrentStatistics HostResourceStatistics `mandatory:"true" json:"currentStatistics"`
}

func (m HostInsightResourceStatisticsAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostInsightResourceStatisticsAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *HostInsightResourceStatisticsAggregation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HostDetails       *HostDetails           `json:"hostDetails"`
		CurrentStatistics hostresourcestatistics `json:"currentStatistics"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HostDetails = model.HostDetails

	nn, e = model.CurrentStatistics.UnmarshalPolymorphicJSON(model.CurrentStatistics.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CurrentStatistics = nn.(HostResourceStatistics)
	} else {
		m.CurrentStatistics = nil
	}

	return
}
