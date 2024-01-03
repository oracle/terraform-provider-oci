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

// OciCodeRepositoryConfiguration An OCI Code repository configuration extends a SCM Configuration with necessary data to reach and use the OCI DevOps Code Repository.
type OciCodeRepositoryConfiguration struct {

	// The branch used by ADM to patch vulnerabilities.
	Branch *string `mandatory:"true" json:"branch"`

	// If true, the Pull Request (PR) will be merged after the verify stage completes successfully
	// If false, the PR with the proposed changes must be reviewed and manually merged.
	IsAutomergeEnabled *bool `mandatory:"true" json:"isAutomergeEnabled"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the OCI DevOps repository.
	OciCodeRepositoryId *string `mandatory:"true" json:"ociCodeRepositoryId"`

	// The location of the build file relative to the root of the repository. Only Maven build files (POM) are currently supported.
	// If this property is not specified, ADM will use the build file located at the root of the repository.
	BuildFileLocation *string `mandatory:"false" json:"buildFileLocation"`
}

// GetBranch returns Branch
func (m OciCodeRepositoryConfiguration) GetBranch() *string {
	return m.Branch
}

// GetBuildFileLocation returns BuildFileLocation
func (m OciCodeRepositoryConfiguration) GetBuildFileLocation() *string {
	return m.BuildFileLocation
}

// GetIsAutomergeEnabled returns IsAutomergeEnabled
func (m OciCodeRepositoryConfiguration) GetIsAutomergeEnabled() *bool {
	return m.IsAutomergeEnabled
}

func (m OciCodeRepositoryConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCodeRepositoryConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciCodeRepositoryConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciCodeRepositoryConfiguration OciCodeRepositoryConfiguration
	s := struct {
		DiscriminatorParam string `json:"scmType"`
		MarshalTypeOciCodeRepositoryConfiguration
	}{
		"OCI_CODE_REPOSITORY",
		(MarshalTypeOciCodeRepositoryConfiguration)(m),
	}

	return json.Marshal(&s)
}
