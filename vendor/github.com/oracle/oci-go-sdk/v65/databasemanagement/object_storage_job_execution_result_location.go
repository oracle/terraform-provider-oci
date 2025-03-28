// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageJobExecutionResultLocation The details about Object Storage job execution result location type.
type ObjectStorageJobExecutionResultLocation struct {

	// The Object Storage namespace used for job execution result storage.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The name of the bucket used for job execution result storage.
	BucketName *string `mandatory:"false" json:"bucketName"`
}

func (m ObjectStorageJobExecutionResultLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageJobExecutionResultLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageJobExecutionResultLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageJobExecutionResultLocation ObjectStorageJobExecutionResultLocation
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeObjectStorageJobExecutionResultLocation
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageJobExecutionResultLocation)(m),
	}

	return json.Marshal(&s)
}
