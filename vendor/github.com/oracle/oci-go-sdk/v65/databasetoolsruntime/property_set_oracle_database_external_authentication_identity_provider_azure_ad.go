// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd External identity provider for AZURE_AD
type PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd struct {

	// External identity provider configuration parameters. Simple key-value pair
	// Example: { "tenant_id": "...", "application_id_uri": "...", ... }
	Configs map[string]string `mandatory:"true" json:"configs"`
}

func (m PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd
	}{
		"AZURE_AD",
		(MarshalTypePropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd)(m),
	}

	return json.Marshal(&s)
}
