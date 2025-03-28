// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GitHubActionsConfiguration Extends a Verify configuration with appropriate data to reach and use the build service provided by a GitHub Action.
type GitHubActionsConfiguration struct {

	// The location of the repository where the GitHub Actions is defined.
	// For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName]
	// For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName]
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// The Oracle Cloud Identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret.
	// The PAT provides the credentials to access the GitHub Action.
	PatSecretId *string `mandatory:"true" json:"patSecretId"`

	// The username that will trigger the GitHub Action.
	Username *string `mandatory:"true" json:"username"`

	// The name of the GitHub Actions workflow that defines the build pipeline.
	WorkflowName *string `mandatory:"true" json:"workflowName"`

	// Additional key-value pairs passed as parameters to the build service when running an experiment.
	AdditionalParameters map[string]string `mandatory:"false" json:"additionalParameters"`
}

func (m GitHubActionsConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitHubActionsConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitHubActionsConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitHubActionsConfiguration GitHubActionsConfiguration
	s := struct {
		DiscriminatorParam string `json:"buildServiceType"`
		MarshalTypeGitHubActionsConfiguration
	}{
		"GITHUB_ACTIONS",
		(MarshalTypeGitHubActionsConfiguration)(m),
	}

	return json.Marshal(&s)
}
