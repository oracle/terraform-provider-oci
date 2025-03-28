// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePathAnalyzerTestDetails Details to update a `PathAnalyzerTest` resource.
type UpdatePathAnalyzerTestDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The IP protocol to use in the `PathAnalyzerTest` resource.
	Protocol *int `mandatory:"false" json:"protocol"`

	SourceEndpoint Endpoint `mandatory:"false" json:"sourceEndpoint"`

	DestinationEndpoint Endpoint `mandatory:"false" json:"destinationEndpoint"`

	ProtocolParameters ProtocolParameters `mandatory:"false" json:"protocolParameters"`

	QueryOptions *QueryOptions `mandatory:"false" json:"queryOptions"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdatePathAnalyzerTestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePathAnalyzerTestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdatePathAnalyzerTestDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                           `json:"displayName"`
		Protocol            *int                              `json:"protocol"`
		SourceEndpoint      endpoint                          `json:"sourceEndpoint"`
		DestinationEndpoint endpoint                          `json:"destinationEndpoint"`
		ProtocolParameters  protocolparameters                `json:"protocolParameters"`
		QueryOptions        *QueryOptions                     `json:"queryOptions"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Protocol = model.Protocol

	nn, e = model.SourceEndpoint.UnmarshalPolymorphicJSON(model.SourceEndpoint.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEndpoint = nn.(Endpoint)
	} else {
		m.SourceEndpoint = nil
	}

	nn, e = model.DestinationEndpoint.UnmarshalPolymorphicJSON(model.DestinationEndpoint.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DestinationEndpoint = nn.(Endpoint)
	} else {
		m.DestinationEndpoint = nil
	}

	nn, e = model.ProtocolParameters.UnmarshalPolymorphicJSON(model.ProtocolParameters.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProtocolParameters = nn.(ProtocolParameters)
	} else {
		m.ProtocolParameters = nil
	}

	m.QueryOptions = model.QueryOptions

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
