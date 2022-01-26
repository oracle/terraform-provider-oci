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

// ImageDatasetFormatDetails Indicates the dataset is comprised of images.
type ImageDatasetFormatDetails struct {
}

func (m ImageDatasetFormatDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImageDatasetFormatDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageDatasetFormatDetails ImageDatasetFormatDetails
	s := struct {
		DiscriminatorParam string `json:"formatType"`
		MarshalTypeImageDatasetFormatDetails
	}{
		"IMAGE",
		(MarshalTypeImageDatasetFormatDetails)(m),
	}

	return json.Marshal(&s)
}
