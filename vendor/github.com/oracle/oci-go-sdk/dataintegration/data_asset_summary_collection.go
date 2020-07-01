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
	"github.com/oracle/oci-go-sdk/common"
)

// DataAssetSummaryCollection This is the collection of data asset summaries, it may be a collection of lightweight details or full definitions.
type DataAssetSummaryCollection struct {

	// The array of DataAsset summaries
	Items []DataAssetSummary `mandatory:"true" json:"items"`
}

func (m DataAssetSummaryCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DataAssetSummaryCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []dataassetsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DataAssetSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DataAssetSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
