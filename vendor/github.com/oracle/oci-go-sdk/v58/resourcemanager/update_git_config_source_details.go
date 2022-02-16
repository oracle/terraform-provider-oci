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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateGitConfigSourceDetails Updates property details for the configuration git information.
type UpdateGitConfigSourceDetails struct {

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm))
	// for the Git configuration source.
	ConfigurationSourceProviderId *string `mandatory:"true" json:"configurationSourceProviderId"`

	// The path of the directory from which to run terraform. If not specified, the the root will be used. This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The URL of the Git repository.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// The name of the branch within the Git repository.
	BranchName *string `mandatory:"false" json:"branchName"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m UpdateGitConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m UpdateGitConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGitConfigSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateGitConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateGitConfigSourceDetails UpdateGitConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeUpdateGitConfigSourceDetails
	}{
		"GIT_CONFIG_SOURCE",
		(MarshalTypeUpdateGitConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
