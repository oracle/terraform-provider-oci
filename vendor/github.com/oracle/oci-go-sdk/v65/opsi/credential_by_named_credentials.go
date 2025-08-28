// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CredentialByNamedCredentials Credentials by named credentials stored in management agent to connect database
type CredentialByNamedCredentials struct {

	// Credential source name that had been added in Management Agent wallet. This value is only required when Credential set by CREDENTIALS_BY_SOURCE and is optional properties for ther others.
	CredentialSourceName *string `mandatory:"false" json:"credentialSourceName"`

	// The credential OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) stored in management agent.
	NamedCredentialId *string `mandatory:"false" json:"namedCredentialId"`
}

// GetCredentialSourceName returns CredentialSourceName
func (m CredentialByNamedCredentials) GetCredentialSourceName() *string {
	return m.CredentialSourceName
}

func (m CredentialByNamedCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CredentialByNamedCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CredentialByNamedCredentials) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCredentialByNamedCredentials CredentialByNamedCredentials
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeCredentialByNamedCredentials
	}{
		"CREDENTIALS_BY_NAMED_CREDS",
		(MarshalTypeCredentialByNamedCredentials)(m),
	}

	return json.Marshal(&s)
}
