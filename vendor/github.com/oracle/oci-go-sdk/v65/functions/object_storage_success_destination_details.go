// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ObjectStorageSuccessDestinationDetails The destination bucket in the Object Storage service to which to send the response of the successful asynchronous function invocation.
// Example: `{"kind": "OBJECTSTORAGE", "namespace": "my-namespace" , "bucketName": "my-new-bucket1" }`
type ObjectStorageSuccessDestinationDetails struct {

	// The Object Storage namespace containing the bucket. For more information on Object Storage namespaces, see Understanding Object Storage Namespaces (https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/understandingnamespaces.htm).
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket. For more information on bucket names, see Managing Buckets (https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/managingbuckets.htm).
	// Example: my-new-bucket1
	BucketName *string `mandatory:"true" json:"bucketName"`
}

func (m ObjectStorageSuccessDestinationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageSuccessDestinationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSuccessDestinationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSuccessDestinationDetails ObjectStorageSuccessDestinationDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeObjectStorageSuccessDestinationDetails
	}{
		"OBJECTSTORAGE",
		(MarshalTypeObjectStorageSuccessDestinationDetails)(m),
	}

	return json.Marshal(&s)
}
