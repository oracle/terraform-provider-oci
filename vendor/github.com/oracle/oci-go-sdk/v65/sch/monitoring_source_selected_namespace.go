// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringSourceSelectedNamespace A metric namespace for the compartment-specific list.
type MonitoringSourceSelectedNamespace struct {

	// The source service or application to use when querying for metric data points. Must begin with `oci_`.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"true" json:"namespace"`

	Metrics MonitoringSourceMetricDetails `mandatory:"true" json:"metrics"`
}

func (m MonitoringSourceSelectedNamespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringSourceSelectedNamespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MonitoringSourceSelectedNamespace) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Namespace *string                       `json:"namespace"`
		Metrics   monitoringsourcemetricdetails `json:"metrics"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Namespace = model.Namespace

	nn, e = model.Metrics.UnmarshalPolymorphicJSON(model.Metrics.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Metrics = nn.(MonitoringSourceMetricDetails)
	} else {
		m.Metrics = nil
	}

	return
}
