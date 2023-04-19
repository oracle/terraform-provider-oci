// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HttpRateLimitConfigurationOverrides Overrides the existing rate limit values based on headers, queryParams or source type.
// Order matters as the overrides are processed sequentially.
type HttpRateLimitConfigurationOverrides struct {
	Limits *HttpRateLimitConfigurationOverridesLimit `mandatory:"true" json:"limits"`

	// Rate limits to be applied based on headers.
	Headers []StringMatch `mandatory:"false" json:"headers"`

	// Rate limits to be applied based on query params.
	QueryParams []StringMatch `mandatory:"false" json:"queryParams"`

	Source SourceMatch `mandatory:"false" json:"source"`
}

func (m HttpRateLimitConfigurationOverrides) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpRateLimitConfigurationOverrides) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *HttpRateLimitConfigurationOverrides) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Headers     []StringMatch                             `json:"headers"`
		QueryParams []StringMatch                             `json:"queryParams"`
		Source      sourcematch                               `json:"source"`
		Limits      *HttpRateLimitConfigurationOverridesLimit `json:"limits"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Headers = make([]StringMatch, len(model.Headers))
	for i, n := range model.Headers {
		m.Headers[i] = n
	}

	m.QueryParams = make([]StringMatch, len(model.QueryParams))
	for i, n := range model.QueryParams {
		m.QueryParams[i] = n
	}

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SourceMatch)
	} else {
		m.Source = nil
	}

	m.Limits = model.Limits

	return
}
