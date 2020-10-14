// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v27/common"
)

// SortKeyRule A rule to define the set of fields to sort by, and whether to sort by ascending or descending values.
type SortKeyRule struct {
	WrappedRule ProjectionRule `mandatory:"false" json:"wrappedRule"`

	// Specifies if the sort key has ascending order.
	IsAscending *bool `mandatory:"false" json:"isAscending"`
}

func (m SortKeyRule) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *SortKeyRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		WrappedRule projectionrule `json:"wrappedRule"`
		IsAscending *bool          `json:"isAscending"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.WrappedRule.UnmarshalPolymorphicJSON(model.WrappedRule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WrappedRule = nn.(ProjectionRule)
	} else {
		m.WrappedRule = nil
	}

	m.IsAscending = model.IsAscending

	return
}
