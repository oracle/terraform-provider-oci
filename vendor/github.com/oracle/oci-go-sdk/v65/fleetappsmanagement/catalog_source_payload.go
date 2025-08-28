// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CatalogSourcePayload Catalog source payload.
type CatalogSourcePayload interface {

	// File path to the directory to use for running Terraform. If not specified, the root directory is used.
	GetWorkingDirectory() *string
}

type catalogsourcepayload struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *catalogsourcepayload) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercatalogsourcepayload catalogsourcepayload
	s := struct {
		Model Unmarshalercatalogsourcepayload
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
func (m *catalogsourcepayload) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceType {
	case "GIT_CATALOG_SOURCE":
		mm := CatalogGitSourceConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STACK_TEMPLATE_CATALOG_SOURCE":
		mm := CatalogSourceTemplateConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PAR_CATALOG_SOURCE":
		mm := CatalogParSourceConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MARKETPLACE_CATALOG_SOURCE":
		mm := CatalogMarketplaceSourceConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CatalogSourcePayload: %s.", m.ConfigSourceType)
		return *m, nil
	}
}

// GetWorkingDirectory returns WorkingDirectory
func (m catalogsourcepayload) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m catalogsourcepayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m catalogsourcepayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
