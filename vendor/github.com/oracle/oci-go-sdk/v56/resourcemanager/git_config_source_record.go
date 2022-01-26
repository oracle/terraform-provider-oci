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

// GitConfigSourceRecord Metadata about the Git configuration source.
type GitConfigSourceRecord struct {

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm))
	// for the Git configuration source.
	ConfigurationSourceProviderId *string `mandatory:"true" json:"configurationSourceProviderId"`

	// The URL of the Git repository.
	RepositoryUrl *string `mandatory:"false" json:"repositoryUrl"`

	// The name of the branch within the Git repository.
	BranchName *string `mandatory:"false" json:"branchName"`

	// The unique identifier (SHA-1 hash) of the individual change to the Git repository.
	CommitId *string `mandatory:"false" json:"commitId"`
}

func (m GitConfigSourceRecord) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GitConfigSourceRecord) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitConfigSourceRecord GitConfigSourceRecord
	s := struct {
		DiscriminatorParam string `json:"configSourceRecordType"`
		MarshalTypeGitConfigSourceRecord
	}{
		"GIT_CONFIG_SOURCE",
		(MarshalTypeGitConfigSourceRecord)(m),
	}

	return json.Marshal(&s)
}
