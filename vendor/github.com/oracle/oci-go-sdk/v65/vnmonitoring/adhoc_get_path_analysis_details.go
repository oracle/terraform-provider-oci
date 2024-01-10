// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdhocGetPathAnalysisDetails Defines the configuration for getting an ad-hoc path analysis.
type AdhocGetPathAnalysisDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The IP protocol to used for the path analysis.
	Protocol *int `mandatory:"true" json:"protocol"`

	SourceEndpoint Endpoint `mandatory:"true" json:"sourceEndpoint"`

	DestinationEndpoint Endpoint `mandatory:"true" json:"destinationEndpoint"`

	ProtocolParameters ProtocolParameters `mandatory:"false" json:"protocolParameters"`

	QueryOptions *QueryOptions `mandatory:"false" json:"queryOptions"`
}

func (m AdhocGetPathAnalysisDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdhocGetPathAnalysisDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdhocGetPathAnalysisDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdhocGetPathAnalysisDetails AdhocGetPathAnalysisDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAdhocGetPathAnalysisDetails
	}{
		"ADHOC_QUERY",
		(MarshalTypeAdhocGetPathAnalysisDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AdhocGetPathAnalysisDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ProtocolParameters  protocolparameters `json:"protocolParameters"`
		QueryOptions        *QueryOptions      `json:"queryOptions"`
		CompartmentId       *string            `json:"compartmentId"`
		Protocol            *int               `json:"protocol"`
		SourceEndpoint      endpoint           `json:"sourceEndpoint"`
		DestinationEndpoint endpoint           `json:"destinationEndpoint"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.CompartmentId = model.CompartmentId

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

	return
}
