// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AnnotationFormatCollection Collection of annotation formats.
type AnnotationFormatCollection struct {

	// List of annotation formats.
	Items []AnnotationFormatSummary `mandatory:"true" json:"items"`
}

func (m AnnotationFormatCollection) String() string {
	return common.PointerString(m)
}
