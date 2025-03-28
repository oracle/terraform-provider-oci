// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OggDeployment Deployment Data for an OggDeployment
type OggDeployment struct {

	// The name given to the GoldenGate service deployment.
	// The name must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter.
	DeploymentName *string `mandatory:"true" json:"deploymentName"`

	// The GoldenGate deployment console username.
	AdminUsername *string `mandatory:"true" json:"adminUsername"`

	// Version of OGG
	OggVersion *string `mandatory:"false" json:"oggVersion"`

	// The base64 encoded content of the PEM file containing the SSL certificate.
	Certificate *string `mandatory:"false" json:"certificate"`

	// The type of credential store for OGG.
	CredentialStore CredentialStoreEnum `mandatory:"false" json:"credentialStore,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Identity Domain when IAM credential store is used.
	IdentityDomainId *string `mandatory:"false" json:"identityDomainId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the deployment password is stored.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	GroupToRolesMapping *GroupToRolesMappingDetails `mandatory:"false" json:"groupToRolesMapping"`
}

func (m OggDeployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OggDeployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCredentialStoreEnum(string(m.CredentialStore)); !ok && m.CredentialStore != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialStore: %s. Supported values are: %s.", m.CredentialStore, strings.Join(GetCredentialStoreEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
