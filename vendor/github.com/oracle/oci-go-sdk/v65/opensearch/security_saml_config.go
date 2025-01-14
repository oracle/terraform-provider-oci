// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecuritySamlConfig SAML policy is optionally used for Opensearch cluster to config SAML authentication
type SecuritySamlConfig struct {

	// A flag determine whether SAML is enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The content of identity provider metadata
	IdpMetadataContent *string `mandatory:"true" json:"idpMetadataContent"`

	// The unique name for a identity provider entity
	IdpEntityId *string `mandatory:"true" json:"idpEntityId"`

	// The endpoint of opendashboard
	OpendashboardUrl *string `mandatory:"false" json:"opendashboardUrl"`

	// The backend role of admins who have all permissions like local master user
	AdminBackendRole *string `mandatory:"false" json:"adminBackendRole"`

	// The subject key is used to get username from SAML assertion. By default, it is NameID
	SubjectKey *string `mandatory:"false" json:"subjectKey"`

	// The roles key is sued to get backend roles from SAML assertion
	RolesKey *string `mandatory:"false" json:"rolesKey"`
}

func (m SecuritySamlConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecuritySamlConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
