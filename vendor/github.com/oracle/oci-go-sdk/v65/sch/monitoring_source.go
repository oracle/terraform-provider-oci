// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringSource A compartment-specific list of metric namespaces to retrieve data from.
type MonitoringSource struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric namespaces you want to use for the Monitoring source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	NamespaceDetails MonitoringSourceNamespaceDetails `mandatory:"true" json:"namespaceDetails"`
}

func (m MonitoringSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MonitoringSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId    *string                          `json:"compartmentId"`
		NamespaceDetails monitoringsourcenamespacedetails `json:"namespaceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	nn, e = model.NamespaceDetails.UnmarshalPolymorphicJSON(model.NamespaceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NamespaceDetails = nn.(MonitoringSourceNamespaceDetails)
	} else {
		m.NamespaceDetails = nil
	}

	return
}
