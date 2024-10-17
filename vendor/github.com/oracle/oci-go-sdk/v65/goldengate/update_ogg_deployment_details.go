// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateOggDeploymentDetails Deployment Details for updating an OggDeployment
type UpdateOggDeploymentDetails struct {

	// The type of credential store for OGG.
	CredentialStore CredentialStoreEnum `mandatory:"false" json:"credentialStore,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Identity Domain when IAM credential store is used.
	IdentityDomainId *string `mandatory:"false" json:"identityDomainId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret where the deployment password is stored.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The GoldenGate deployment console username.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// The password associated with the GoldenGate deployment console username.
	// The password must be 8 to 30 characters long and must contain at least 1 uppercase, 1 lowercase, 1 numeric,
	// and 1 special character. Special characters such as '$', '^', or '?' are not allowed.
	// This field will be deprecated and replaced by "passwordSecretId".
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The base64 encoded content of the PEM file containing the SSL certificate.
	Certificate *string `mandatory:"false" json:"certificate"`

	// The base64 encoded content of the PEM file containing the private key.
	Key *string `mandatory:"false" json:"key"`

	GroupToRolesMapping *UpdateGroupToRolesMappingDetails `mandatory:"false" json:"groupToRolesMapping"`
}

func (m UpdateOggDeploymentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOggDeploymentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCredentialStoreEnum(string(m.CredentialStore)); !ok && m.CredentialStore != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialStore: %s. Supported values are: %s.", m.CredentialStore, strings.Join(GetCredentialStoreEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
