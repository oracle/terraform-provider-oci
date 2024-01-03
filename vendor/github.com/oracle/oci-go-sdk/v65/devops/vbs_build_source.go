// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VbsBuildSource VBS Server Build Source for Build Stage
type VbsBuildSource struct {

	// Name of the build source. This must be unique within a build source collection. The name can be used by customers to locate the working directory pertinent to this repository.
	Name *string `mandatory:"true" json:"name"`

	// URL for the repository.
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// Branch name.
	Branch *string `mandatory:"true" json:"branch"`

	// Connection identifier pertinent to VBS Server source provider
	ConnectionId *string `mandatory:"true" json:"connectionId"`
}

// GetName returns Name
func (m VbsBuildSource) GetName() *string {
	return m.Name
}

// GetRepositoryUrl returns RepositoryUrl
func (m VbsBuildSource) GetRepositoryUrl() *string {
	return m.RepositoryUrl
}

// GetBranch returns Branch
func (m VbsBuildSource) GetBranch() *string {
	return m.Branch
}

func (m VbsBuildSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbsBuildSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VbsBuildSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVbsBuildSource VbsBuildSource
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeVbsBuildSource
	}{
		"VBS",
		(MarshalTypeVbsBuildSource)(m),
	}

	return json.Marshal(&s)
}
