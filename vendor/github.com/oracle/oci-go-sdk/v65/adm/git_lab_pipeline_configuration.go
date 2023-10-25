// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GitLabPipelineConfiguration Extends a Verify configuration with appropriate data to reach and use the build service provided by a GitLab Pipeline.
type GitLabPipelineConfiguration struct {

	// The location of the Repository where the GitLab Pipeline will be run.
	// The expected format is https://gitlab.com/[groupName]/[repoName]
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// The username that will trigger the GitLab Pipeline.
	Username *string `mandatory:"true" json:"username"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret.
	// The PAT provides the credentials to access the GitLab pipeline.
	PatSecretId *string `mandatory:"true" json:"patSecretId"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the trigger Secret.
	// The Secret provides access to the trigger for a GitLab pipeline.
	TriggerSecretId *string `mandatory:"true" json:"triggerSecretId"`

	// Additional key-value pairs passed as parameters to the build service when running an experiment.
	AdditionalParameters map[string]string `mandatory:"false" json:"additionalParameters"`
}

func (m GitLabPipelineConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitLabPipelineConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitLabPipelineConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitLabPipelineConfiguration GitLabPipelineConfiguration
	s := struct {
		DiscriminatorParam string `json:"buildServiceType"`
		MarshalTypeGitLabPipelineConfiguration
	}{
		"GITLAB_PIPELINE",
		(MarshalTypeGitLabPipelineConfiguration)(m),
	}

	return json.Marshal(&s)
}
