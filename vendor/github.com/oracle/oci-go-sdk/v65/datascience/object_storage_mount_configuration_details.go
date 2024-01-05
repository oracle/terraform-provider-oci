// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageMountConfigurationDetails The Object Storage Configuration Details.
type ObjectStorageMountConfigurationDetails struct {

	// The local directory name to be mounted
	DestinationDirectoryName *string `mandatory:"true" json:"destinationDirectoryName"`

	// The object storage namespace
	Namespace *string `mandatory:"true" json:"namespace"`

	// The object storage bucket
	Bucket *string `mandatory:"true" json:"bucket"`

	// The local path of the mounted directory, excluding directory name.
	DestinationPath *string `mandatory:"false" json:"destinationPath"`

	// Prefix in the bucket to mount
	Prefix *string `mandatory:"false" json:"prefix"`
}

// GetDestinationDirectoryName returns DestinationDirectoryName
func (m ObjectStorageMountConfigurationDetails) GetDestinationDirectoryName() *string {
	return m.DestinationDirectoryName
}

// GetDestinationPath returns DestinationPath
func (m ObjectStorageMountConfigurationDetails) GetDestinationPath() *string {
	return m.DestinationPath
}

func (m ObjectStorageMountConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageMountConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageMountConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageMountConfigurationDetails ObjectStorageMountConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeObjectStorageMountConfigurationDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageMountConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
