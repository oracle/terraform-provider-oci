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

// GitConfigSource Metadata about the Git configuration source.
type GitConfigSource struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Git configuration source.
	ConfigurationSourceProviderId *string `mandatory:"true" json:"configurationSourceProviderId"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The URL of the Git repository for the configuration source.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// The name of the branch in the Git repository for the configuration source.
	BranchName *string `mandatory:"false" json:"branchName"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m GitConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m GitConfigSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitConfigSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitConfigSource GitConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeGitConfigSource
	}{
		"GIT_CONFIG_SOURCE",
		(MarshalTypeGitConfigSource)(m),
	}

	return json.Marshal(&s)
}
