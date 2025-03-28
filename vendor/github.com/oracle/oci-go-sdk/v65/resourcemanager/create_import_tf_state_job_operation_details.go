// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateImportTfStateJobOperationDetails Job details that are specific to import Terraform state operations.
type CreateImportTfStateJobOperationDetails struct {

	// Base64-encoded state file
	TfStateBase64Encoded []byte `mandatory:"true" json:"tfStateBase64Encoded"`

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`
}

// GetIsProviderUpgradeRequired returns IsProviderUpgradeRequired
func (m CreateImportTfStateJobOperationDetails) GetIsProviderUpgradeRequired() *bool {
	return m.IsProviderUpgradeRequired
}

func (m CreateImportTfStateJobOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateImportTfStateJobOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateImportTfStateJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateImportTfStateJobOperationDetails CreateImportTfStateJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeCreateImportTfStateJobOperationDetails
	}{
		"IMPORT_TF_STATE",
		(MarshalTypeCreateImportTfStateJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
