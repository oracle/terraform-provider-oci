// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageBucketContentDetails Details about Object Storage Bucket source.
// Restricted to the region this is called
type ObjectStorageBucketContentDetails struct {

	// Namespace.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Bucket Name.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Object Name.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// SHA256 checksum of the artifact.
	Checksum *string `mandatory:"true" json:"checksum"`
}

func (m ObjectStorageBucketContentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageBucketContentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageBucketContentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageBucketContentDetails ObjectStorageBucketContentDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageBucketContentDetails
	}{
		"OBJECT_STORAGE_BUCKET",
		(MarshalTypeObjectStorageBucketContentDetails)(m),
	}

	return json.Marshal(&s)
}
