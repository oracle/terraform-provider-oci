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

// ModelSelect The information about the select object.
type ModelSelect struct {

	// Specifies whether the object is distinct.
	IsDistinct *bool `mandatory:"false" json:"isDistinct"`

	// An array of selected columns.
	SelectColumns []ShapeField `mandatory:"false" json:"selectColumns"`
}

func (m ModelSelect) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ModelSelect) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeModelSelect ModelSelect
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeModelSelect
	}{
		"SELECT",
		(MarshalTypeModelSelect)(m),
	}

	return json.Marshal(&s)
}
