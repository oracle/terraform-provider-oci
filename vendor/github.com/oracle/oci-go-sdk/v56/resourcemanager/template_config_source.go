// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TemplateConfigSource Information about the Template.
type TemplateConfigSource interface {
}

type templateconfigsource struct {
	JsonData                 []byte
	TemplateConfigSourceType string `json:"templateConfigSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *templateconfigsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertemplateconfigsource templateconfigsource
	s := struct {
		Model Unmarshalertemplateconfigsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TemplateConfigSourceType = s.Model.TemplateConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *templateconfigsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TemplateConfigSourceType {
	case "ZIP_UPLOAD":
		mm := TemplateZipUploadConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m templateconfigsource) String() string {
	return common.PointerString(m)
}

// TemplateConfigSourceTemplateConfigSourceTypeEnum Enum with underlying type: string
type TemplateConfigSourceTemplateConfigSourceTypeEnum string

// Set of constants representing the allowable values for TemplateConfigSourceTemplateConfigSourceTypeEnum
const (
	TemplateConfigSourceTemplateConfigSourceTypeZipUpload TemplateConfigSourceTemplateConfigSourceTypeEnum = "ZIP_UPLOAD"
)

var mappingTemplateConfigSourceTemplateConfigSourceType = map[string]TemplateConfigSourceTemplateConfigSourceTypeEnum{
	"ZIP_UPLOAD": TemplateConfigSourceTemplateConfigSourceTypeZipUpload,
}

// GetTemplateConfigSourceTemplateConfigSourceTypeEnumValues Enumerates the set of values for TemplateConfigSourceTemplateConfigSourceTypeEnum
func GetTemplateConfigSourceTemplateConfigSourceTypeEnumValues() []TemplateConfigSourceTemplateConfigSourceTypeEnum {
	values := make([]TemplateConfigSourceTemplateConfigSourceTypeEnum, 0)
	for _, v := range mappingTemplateConfigSourceTemplateConfigSourceType {
		values = append(values, v)
	}
	return values
}
