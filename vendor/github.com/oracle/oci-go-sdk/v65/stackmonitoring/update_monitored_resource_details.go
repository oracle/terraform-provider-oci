// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMonitoredResourceDetails The information about updating a monitored resource.
type UpdateMonitoredResourceDetails struct {

	// Monitored resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Host name of the monitored resource
	HostName *string `mandatory:"false" json:"hostName"`

	// Time zone in the form of tz database canonical zone ID.
	ResourceTimeZone *string `mandatory:"false" json:"resourceTimeZone"`

	// List of monitored resource properties
	Properties []MonitoredResourceProperty `mandatory:"false" json:"properties"`

	DatabaseConnectionDetails *ConnectionDetails `mandatory:"false" json:"databaseConnectionDetails"`

	Credentials MonitoredResourceCredential `mandatory:"false" json:"credentials"`

	Aliases *MonitoredResourceAliasCredential `mandatory:"false" json:"aliases"`
}

func (m UpdateMonitoredResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMonitoredResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMonitoredResourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                           `json:"displayName"`
		HostName                  *string                           `json:"hostName"`
		ResourceTimeZone          *string                           `json:"resourceTimeZone"`
		Properties                []MonitoredResourceProperty       `json:"properties"`
		DatabaseConnectionDetails *ConnectionDetails                `json:"databaseConnectionDetails"`
		Credentials               monitoredresourcecredential       `json:"credentials"`
		Aliases                   *MonitoredResourceAliasCredential `json:"aliases"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.HostName = model.HostName

	m.ResourceTimeZone = model.ResourceTimeZone

	m.Properties = make([]MonitoredResourceProperty, len(model.Properties))
	for i, n := range model.Properties {
		m.Properties[i] = n
	}

	m.DatabaseConnectionDetails = model.DatabaseConnectionDetails

	nn, e = model.Credentials.UnmarshalPolymorphicJSON(model.Credentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Credentials = nn.(MonitoredResourceCredential)
	} else {
		m.Credentials = nil
	}

	m.Aliases = model.Aliases

	return
}
