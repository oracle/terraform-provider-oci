// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageSourceDetails Specifies the dataset location in object storage. This requires that all records are in this bucket, and under this prefix. We do not support a dataset with objects in arbitrary locations across buckets or prefixes.
type ObjectStorageSourceDetails struct {

	// The namespace of the bucket that contains the dataset data source.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The object storage bucket that contains the dataset data source.
	Bucket *string `mandatory:"true" json:"bucket"`

	// A common path prefix shared by the objects that make up the dataset. Except for the CSV file type, records are not generated for the objects whose names exactly match with the prefix.
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m ObjectStorageSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSourceDetails ObjectStorageSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageSourceDetails)(m),
	}

	return json.Marshal(&s)
}
