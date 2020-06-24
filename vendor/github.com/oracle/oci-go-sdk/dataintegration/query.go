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

// Query A query object.
type Query struct {

	// A query string.
	Query *string `mandatory:"false" json:"query"`
}

func (m Query) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Query) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeQuery Query
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeQuery
	}{
		"QUERY",
		(MarshalTypeQuery)(m),
	}

	return json.Marshal(&s)
}
