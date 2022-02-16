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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ObjectStorageSnapshotExportDetails Specifies where to output the export in Object Storage.
type ObjectStorageSnapshotExportDetails struct {

	// Bucket namespace name
	Namespace *string `mandatory:"true" json:"namespace"`

	// Bucket name
	Bucket *string `mandatory:"true" json:"bucket"`

	// Object path prefix to put snapshot file(s)
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m ObjectStorageSnapshotExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageSnapshotExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSnapshotExportDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSnapshotExportDetails ObjectStorageSnapshotExportDetails
	s := struct {
		DiscriminatorParam string `json:"exportType"`
		MarshalTypeObjectStorageSnapshotExportDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageSnapshotExportDetails)(m),
	}

	return json.Marshal(&s)
}
