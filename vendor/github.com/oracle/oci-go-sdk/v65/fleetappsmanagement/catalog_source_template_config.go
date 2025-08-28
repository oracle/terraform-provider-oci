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

// CatalogSourceTemplateConfig Catalog source template config.
type CatalogSourceTemplateConfig struct {

	// File path to the directory to use for running Terraform. If not specified, the root directory is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The Base64 encoded template. This payload will trigger CreateTemplate API, where the parameter will be passed.
	ZipFileBase64Encoded *string `mandatory:"false" json:"zipFileBase64Encoded"`

	// Template Description
	Description *string `mandatory:"false" json:"description"`

	// Template Long Description
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Template Display Name
	TemplateDisplayName *string `mandatory:"false" json:"templateDisplayName"`
}

// GetWorkingDirectory returns WorkingDirectory
func (m CatalogSourceTemplateConfig) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CatalogSourceTemplateConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogSourceTemplateConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CatalogSourceTemplateConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCatalogSourceTemplateConfig CatalogSourceTemplateConfig
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCatalogSourceTemplateConfig
	}{
		"STACK_TEMPLATE_CATALOG_SOURCE",
		(MarshalTypeCatalogSourceTemplateConfig)(m),
	}

	return json.Marshal(&s)
}
