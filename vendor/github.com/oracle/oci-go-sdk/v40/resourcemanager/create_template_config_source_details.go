// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v40/common"
)

// CreateTemplateConfigSourceDetails Property details for the configuration source used for the template.
type CreateTemplateConfigSourceDetails interface {
}

type createtemplateconfigsourcedetails struct {
	JsonData                 []byte
	TemplateConfigSourceType string `json:"templateConfigSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createtemplateconfigsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetemplateconfigsourcedetails createtemplateconfigsourcedetails
	s := struct {
		Model Unmarshalercreatetemplateconfigsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TemplateConfigSourceType = s.Model.TemplateConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtemplateconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TemplateConfigSourceType {
	case "ZIP_UPLOAD":
		mm := CreateTemplateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createtemplateconfigsourcedetails) String() string {
	return common.PointerString(m)
}
