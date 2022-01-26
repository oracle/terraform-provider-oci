// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EnrichedEntity This is used to specify runtime parameters for data entities such as files that need both the data entity and the format.
type EnrichedEntity struct {
	Entity DataEntity `mandatory:"false" json:"entity"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	// The model type for the entity which is referenced.
	ModelType *string `mandatory:"false" json:"modelType"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`
}

func (m EnrichedEntity) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *EnrichedEntity) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Entity     dataentity       `json:"entity"`
		DataFormat *DataFormat      `json:"dataFormat"`
		ModelType  *string          `json:"modelType"`
		ParentRef  *ParentReference `json:"parentRef"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Entity.UnmarshalPolymorphicJSON(model.Entity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Entity = nn.(DataEntity)
	} else {
		m.Entity = nil
	}

	m.DataFormat = model.DataFormat

	m.ModelType = model.ModelType

	m.ParentRef = model.ParentRef

	return
}
