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

// ObjectStorageWriteAttribute Properties to configure writing to Object Storage.
type ObjectStorageWriteAttribute struct {

	// Specifies whether to write output to single-file or not.
	WriteToSingleFile *bool `mandatory:"false" json:"writeToSingleFile"`
}

func (m ObjectStorageWriteAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageWriteAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageWriteAttribute ObjectStorageWriteAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeObjectStorageWriteAttribute
	}{
		"OBJECTSTORAGEWRITEATTRIBUTE",
		(MarshalTypeObjectStorageWriteAttribute)(m),
	}

	return json.Marshal(&s)
}
