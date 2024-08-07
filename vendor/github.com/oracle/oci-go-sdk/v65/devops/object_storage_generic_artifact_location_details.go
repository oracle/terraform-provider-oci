// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageGenericArtifactLocationDetails ObjectStorage bucket details where artifact is located.
type ObjectStorageGenericArtifactLocationDetails struct {

	// Unique namespace in the tenancy which contains the bucket
	Namespace *string `mandatory:"true" json:"namespace"`

	// Specifies the Object Storage bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Specifies the name of object inside the bucket.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Region identifier.
	Region *string `mandatory:"true" json:"region"`
}

func (m ObjectStorageGenericArtifactLocationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageGenericArtifactLocationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageGenericArtifactLocationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageGenericArtifactLocationDetails ObjectStorageGenericArtifactLocationDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeObjectStorageGenericArtifactLocationDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageGenericArtifactLocationDetails)(m),
	}

	return json.Marshal(&s)
}
