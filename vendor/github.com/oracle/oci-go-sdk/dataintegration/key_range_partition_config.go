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

// KeyRangePartitionConfig The information about key range.
type KeyRangePartitionConfig struct {

	// The partition number for the key range.
	PartitionNumber *int `mandatory:"false" json:"partitionNumber"`

	KeyRange *KeyRange `mandatory:"false" json:"keyRange"`
}

func (m KeyRangePartitionConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m KeyRangePartitionConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKeyRangePartitionConfig KeyRangePartitionConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeKeyRangePartitionConfig
	}{
		"KEYRANGEPARTITIONCONFIG",
		(MarshalTypeKeyRangePartitionConfig)(m),
	}

	return json.Marshal(&s)
}
