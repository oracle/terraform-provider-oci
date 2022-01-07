// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// SnapshotDatasetDetails Allows outputting the latest records paired with annotations and write them to object storage.
type SnapshotDatasetDetails struct {

	// Whether annotations are to be included in the export dataset digest.
	AreAnnotationsIncluded *bool `mandatory:"true" json:"areAnnotationsIncluded"`

	// Whether to include records that have yet to be annotated in the export dataset digest.
	AreUnannotatedRecordsIncluded *bool `mandatory:"true" json:"areUnannotatedRecordsIncluded"`

	ExportDetails *ObjectStorageSnapshotExportDetails `mandatory:"true" json:"exportDetails"`
}

func (m SnapshotDatasetDetails) String() string {
	return common.PointerString(m)
}
