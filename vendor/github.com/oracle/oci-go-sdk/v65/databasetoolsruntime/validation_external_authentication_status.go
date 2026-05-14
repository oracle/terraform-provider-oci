// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidationExternalAuthenticationStatus Status details for an external authentication
type ValidationExternalAuthenticationStatus struct {

	// The Database Tools identity provider type.
	Type IdentityProviderTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The status of the external authentication. AVAILABLE if the external authentication can be used with this connection. UNAVAILABLE if not.
	// ENABLED if the external authentication is enabled.
	Status DatabaseToolsExternalAuthenticationStatusEnum `mandatory:"false" json:"status,omitempty"`

	// If the status is UNAVAILABLE this displays the cause.
	StatusDetails *string `mandatory:"false" json:"statusDetails"`
}

func (m ValidationExternalAuthenticationStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidationExternalAuthenticationStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdentityProviderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentityProviderTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsExternalAuthenticationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDatabaseToolsExternalAuthenticationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
