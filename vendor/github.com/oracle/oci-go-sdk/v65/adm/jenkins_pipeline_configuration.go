// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// JenkinsPipelineConfiguration Extends a Verify configuration with appropriate data to reach and use the build service provided by a Jenkins Pipeline.
type JenkinsPipelineConfiguration struct {

	// The username that will be used to authenticate with Jenkins.
	Username *string `mandatory:"true" json:"username"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret.
	// The PAT provides the credentials to access the Jenkins Pipeline.
	PatSecretId *string `mandatory:"true" json:"patSecretId"`

	// The URL that locates the Jenkins pipeline.
	JenkinsUrl *string `mandatory:"true" json:"jenkinsUrl"`

	// The name of the Jenkins pipeline job that identifies the build pipeline.
	JobName *string `mandatory:"true" json:"jobName"`

	// Additional key-value pairs passed as parameters to the build service when running an experiment.
	AdditionalParameters map[string]string `mandatory:"false" json:"additionalParameters"`
}

func (m JenkinsPipelineConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JenkinsPipelineConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JenkinsPipelineConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJenkinsPipelineConfiguration JenkinsPipelineConfiguration
	s := struct {
		DiscriminatorParam string `json:"buildServiceType"`
		MarshalTypeJenkinsPipelineConfiguration
	}{
		"JENKINS_PIPELINE",
		(MarshalTypeJenkinsPipelineConfiguration)(m),
	}

	return json.Marshal(&s)
}
