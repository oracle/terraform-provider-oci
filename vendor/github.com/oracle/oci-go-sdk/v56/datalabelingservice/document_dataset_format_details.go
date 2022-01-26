// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DocumentDatasetFormatDetails Allows the user to specify that the dataset is comprised of document files (e.g. PDFs, DOCs, etc.).  It is open for further configurability.
type DocumentDatasetFormatDetails struct {
}

func (m DocumentDatasetFormatDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DocumentDatasetFormatDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentDatasetFormatDetails DocumentDatasetFormatDetails
	s := struct {
		DiscriminatorParam string `json:"formatType"`
		MarshalTypeDocumentDatasetFormatDetails
	}{
		"DOCUMENT",
		(MarshalTypeDocumentDatasetFormatDetails)(m),
	}

	return json.Marshal(&s)
}
