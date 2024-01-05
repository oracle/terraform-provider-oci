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

// BitbucketCloudConfigSourceRecord Metadata about the Bitbucket Cloud configuration source.
type BitbucketCloudConfigSourceRecord struct {

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm))
	// for the Bitbucket Cloud configuration source.
	ConfigurationSourceProviderId *string `mandatory:"true" json:"configurationSourceProviderId"`

	// The URL of the Bitbucket Cloud repository.
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// The id of the workspace in Bitbucket Cloud for the configuration source.
	WorkspaceId *string `mandatory:"true" json:"workspaceId"`

	// The name of the branch within the Bitbucket Cloud repository.
	BranchName *string `mandatory:"false" json:"branchName"`

	// The unique identifier (SHA-1 hash) of the individual change to the Bitbucket Cloud repository.
	CommitId *string `mandatory:"false" json:"commitId"`
}

func (m BitbucketCloudConfigSourceRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketCloudConfigSourceRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketCloudConfigSourceRecord) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketCloudConfigSourceRecord BitbucketCloudConfigSourceRecord
	s := struct {
		DiscriminatorParam string `json:"configSourceRecordType"`
		MarshalTypeBitbucketCloudConfigSourceRecord
	}{
		"BITBUCKET_CLOUD_CONFIG_SOURCE",
		(MarshalTypeBitbucketCloudConfigSourceRecord)(m),
	}

	return json.Marshal(&s)
}
