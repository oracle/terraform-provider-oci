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
	"github.com/oracle/oci-go-sdk/v44/common"
)

// CreateStackTemplateConfigSourceDetails The template to use as the source of the Terraform configuration.
type CreateStackTemplateConfigSourceDetails struct {
	TemplateId *string `mandatory:"true" json:"templateId"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CreateStackTemplateConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CreateStackTemplateConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateStackTemplateConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateStackTemplateConfigSourceDetails CreateStackTemplateConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCreateStackTemplateConfigSourceDetails
	}{
		"TEMPLATE_CONFIG_SOURCE",
		(MarshalTypeCreateStackTemplateConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
