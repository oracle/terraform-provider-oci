// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageArchiveSourceDetails The details for the Archive source of the Function from Object Storage.
// This is used when the function code is stored in Object Storage.
// It is suitable for scenarios where the function code is uploaded as an archive file to a specific bucket and namespace
type ObjectStorageArchiveSourceDetails struct {

	// The name of the Object Storage bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the Object Storage object.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// VersionId used to identify a particular version of the object. If not specified, the latest version of the object is used.
	ObjectVersionId *string `mandatory:"false" json:"objectVersionId"`
}

func (m ObjectStorageArchiveSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageArchiveSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageArchiveSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageArchiveSourceDetails ObjectStorageArchiveSourceDetails
	s := struct {
		DiscriminatorParam string `json:"archiveSourceType"`
		MarshalTypeObjectStorageArchiveSourceDetails
	}{
		"OBJECT_STORAGE_ARCHIVE",
		(MarshalTypeObjectStorageArchiveSourceDetails)(m),
	}

	return json.Marshal(&s)
}
