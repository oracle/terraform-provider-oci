// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRepositoryPrivateAccessDetails Information about a new Repository Private Access resource.
type CreateRepositoryPrivateAccessDetails struct {

	// The OCID of the compartment that will hold the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet where VNIC resources will be created for private endpoint.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Name of the reposiroy private access.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The textual description for the repository private access resource.
	Description *string `mandatory:"false" json:"description"`

	// An array of network security group OCIDs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// An array of IP CIDR ranges from where the access is allowed for REST API endpoints.
	AllowedApiCidrBlocks []string `mandatory:"false" json:"allowedApiCidrBlocks"`

	// The IP of the private endpoint provisioned.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The FQDN of the private endpoint provisioned.
	HostName *string `mandatory:"false" json:"hostName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRepositoryPrivateAccessDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRepositoryPrivateAccessDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
