// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API.
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateZipUploadConfigSourceDetails Property details for uploading the configuration zip file.
type CreateZipUploadConfigSourceDetails struct {
	ZipFileBase64Encoded *string `mandatory:"true" json:"zipFileBase64Encoded"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CreateZipUploadConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CreateZipUploadConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateZipUploadConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZipUploadConfigSourceDetails CreateZipUploadConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCreateZipUploadConfigSourceDetails
	}{
		"ZIP_UPLOAD",
		(MarshalTypeCreateZipUploadConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
