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

// MonitoringSource A compartment-specific list of metric namespaces to retrieve data from.
type MonitoringSource struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment containing metric namespaces you want to use for the Monitoring source.
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
