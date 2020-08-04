// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateConfigSourceDetails Updates the property details for the configuration source.
type UpdateConfigSourceDetails interface {

	// The path of the directory from which to run terraform. If not specified, the the root will be used. This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	GetWorkingDirectory() *string
}

type updateconfigsourcedetails struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigsourcedetails updateconfigsourcedetails
	s := struct {
		Model Unmarshalerupdateconfigsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigSourceType = s.Model.ConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceType {
	case "GIT_CONFIG_SOURCE":
		mm := UpdateGitConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZIP_UPLOAD":
		mm := UpdateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetWorkingDirectory returns WorkingDirectory
func (m updateconfigsourcedetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m updateconfigsourcedetails) String() string {
	return common.PointerString(m)
}
