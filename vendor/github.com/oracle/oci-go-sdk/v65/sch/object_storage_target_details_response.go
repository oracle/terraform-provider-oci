// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageTargetDetailsResponse The destination bucket for data transferred from the source.
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type ObjectStorageTargetDetailsResponse struct {

	// The name of the bucket. Valid characters are letters (upper or lower case), numbers, hyphens (-),
	// underscores(_), and periods (.). Bucket names must be unique within an Object Storage namespace.
	// Avoid entering confidential information. Example: my-new-bucket1
	BucketName *string `mandatory:"true" json:"bucketName"`

	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`

	// The namespace.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The prefix of the objects. Avoid entering confidential information.
	ObjectNamePrefix *string `mandatory:"false" json:"objectNamePrefix"`

	// The batch rollover size in megabytes.
	BatchRolloverSizeInMBs *int `mandatory:"false" json:"batchRolloverSizeInMBs"`

	// The batch rollover time in milliseconds.
	BatchRolloverTimeInMs *int `mandatory:"false" json:"batchRolloverTimeInMs"`
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m ObjectStorageTargetDetailsResponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m ObjectStorageTargetDetailsResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageTargetDetailsResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageTargetDetailsResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageTargetDetailsResponse ObjectStorageTargetDetailsResponse
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeObjectStorageTargetDetailsResponse
	}{
		"objectStorage",
		(MarshalTypeObjectStorageTargetDetailsResponse)(m),
	}

	return json.Marshal(&s)
}
