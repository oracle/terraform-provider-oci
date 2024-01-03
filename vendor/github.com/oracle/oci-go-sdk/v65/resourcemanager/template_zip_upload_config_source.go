// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TemplateZipUploadConfigSource Metadata about the zip file containing the Terraform configuration for the template.
type TemplateZipUploadConfigSource struct {
}

func (m TemplateZipUploadConfigSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateZipUploadConfigSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TemplateZipUploadConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTemplateZipUploadConfigSource TemplateZipUploadConfigSource
	s := struct {
		DiscriminatorParam string `json:"templateConfigSourceType"`
		MarshalTypeTemplateZipUploadConfigSource
	}{
		"ZIP_UPLOAD",
		(MarshalTypeTemplateZipUploadConfigSource)(m),
	}

	return json.Marshal(&s)
}
