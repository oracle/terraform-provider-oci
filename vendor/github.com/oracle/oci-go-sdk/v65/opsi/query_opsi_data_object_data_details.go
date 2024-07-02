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

// QueryOpsiDataObjectDataDetails Information required to form and execute query on an OPSI data object.
type QueryOpsiDataObjectDataDetails struct {
	Query DataObjectQuery `mandatory:"true" json:"query"`

	// Unique OPSI data object identifier.
	DataObjectIdentifier *string `mandatory:"false" json:"dataObjectIdentifier"`

	// Details of OPSI data objects used in the query.
	DataObjects []OpsiDataObjectDetailsInQuery `mandatory:"false" json:"dataObjects"`

	ResourceFilters *ResourceFilters `mandatory:"false" json:"resourceFilters"`
}

func (m QueryOpsiDataObjectDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryOpsiDataObjectDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *QueryOpsiDataObjectDataDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataObjectIdentifier *string                        `json:"dataObjectIdentifier"`
		DataObjects          []opsidataobjectdetailsinquery `json:"dataObjects"`
		ResourceFilters      *ResourceFilters               `json:"resourceFilters"`
		Query                dataobjectquery                `json:"query"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DataObjectIdentifier = model.DataObjectIdentifier

	m.DataObjects = make([]OpsiDataObjectDetailsInQuery, len(model.DataObjects))
	for i, n := range model.DataObjects {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DataObjects[i] = nn.(OpsiDataObjectDetailsInQuery)
		} else {
			m.DataObjects[i] = nil
		}
	}
	m.ResourceFilters = model.ResourceFilters

	nn, e = model.Query.UnmarshalPolymorphicJSON(model.Query.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Query = nn.(DataObjectQuery)
	} else {
		m.Query = nil
	}

	return
}
