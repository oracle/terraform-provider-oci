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

// FileStorageMountConfigurationDetails The File Storage Mount Configuration Details.
type FileStorageMountConfigurationDetails struct {

	// The local directory name to be mounted
	DestinationDirectoryName *string `mandatory:"true" json:"destinationDirectoryName"`

	// OCID of the mount target
	MountTargetId *string `mandatory:"true" json:"mountTargetId"`

	// OCID of the export
	ExportId *string `mandatory:"true" json:"exportId"`

	// The local path of the mounted directory, excluding directory name.
	DestinationPath *string `mandatory:"false" json:"destinationPath"`
}

// GetDestinationDirectoryName returns DestinationDirectoryName
func (m FileStorageMountConfigurationDetails) GetDestinationDirectoryName() *string {
	return m.DestinationDirectoryName
}

// GetDestinationPath returns DestinationPath
func (m FileStorageMountConfigurationDetails) GetDestinationPath() *string {
	return m.DestinationPath
}

func (m FileStorageMountConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileStorageMountConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FileStorageMountConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFileStorageMountConfigurationDetails FileStorageMountConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeFileStorageMountConfigurationDetails
	}{
		"FILE_STORAGE",
		(MarshalTypeFileStorageMountConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
