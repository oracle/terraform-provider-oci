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

// BitbucketCloudConfigSource Metadata about the Bitbucket Cloud configuration source.
type BitbucketCloudConfigSource struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Bitbucket Cloud configuration source.
	ConfigurationSourceProviderId *string `mandatory:"true" json:"configurationSourceProviderId"`

	// The URL of the Bitbucket Cloud repository for the configuration source.
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// The id of the workspace in Bitbucket Cloud for the configuration source
	WorkspaceId *string `mandatory:"true" json:"workspaceId"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// Required when using a zip Terraform configuration (`configSourceType` value of `ZIP_UPLOAD`) that contains folders.
	// Ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	// For more information about required and recommended file structure, see
	// File Structure (Terraform Configurations for Resource Manager) (https://docs.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager.htm#filestructure).
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The name of the branch in the Bitbucket Cloud repository for the configuration source.
	BranchName *string `mandatory:"false" json:"branchName"`
}

// GetWorkingDirectory returns WorkingDirectory
func (m BitbucketCloudConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m BitbucketCloudConfigSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketCloudConfigSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketCloudConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketCloudConfigSource BitbucketCloudConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeBitbucketCloudConfigSource
	}{
		"BITBUCKET_CLOUD_CONFIG_SOURCE",
		(MarshalTypeBitbucketCloudConfigSource)(m),
	}

	return json.Marshal(&s)
}
