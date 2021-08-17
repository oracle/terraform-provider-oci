// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v46/common"
)

// OracleWriteAttributes Properties to configure when writing to an Oracle Database.
type OracleWriteAttributes struct {

	// The batch size for writing.
	BatchSize *int `mandatory:"false" json:"batchSize"`

	// Specifies whether to truncate.
	IsTruncate *bool `mandatory:"false" json:"isTruncate"`

	// Specifies the isolation level.
	IsolationLevel *string `mandatory:"false" json:"isolationLevel"`
}

func (m OracleWriteAttributes) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleWriteAttributes) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleWriteAttributes OracleWriteAttributes
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleWriteAttributes
	}{
		"ORACLE_WRITE_ATTRIBUTE",
		(MarshalTypeOracleWriteAttributes)(m),
	}

	return json.Marshal(&s)
}
