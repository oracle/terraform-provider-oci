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

// UpdatePropertySetApexFaIntegrationDetails Contains the update details of an APEX FA Integration property set
type UpdatePropertySetApexFaIntegrationDetails struct {

	// APEX FA Integration key-value pairs.
	AuthenticationSubstitutions map[string]string `mandatory:"false" json:"authenticationSubstitutions"`

	// Specifies whether database credentials can be used in all workspaces on the APEX instance. Supported values include: "Y", "N" and empty string.
	InstanceDbmsCredentialEnabled *string `mandatory:"false" json:"instanceDbmsCredentialEnabled"`
}

func (m UpdatePropertySetApexFaIntegrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePropertySetApexFaIntegrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePropertySetApexFaIntegrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePropertySetApexFaIntegrationDetails UpdatePropertySetApexFaIntegrationDetails
	s := struct {
		DiscriminatorParam string `json:"key"`
		MarshalTypeUpdatePropertySetApexFaIntegrationDetails
	}{
		"APEX_FA_INTEGRATION",
		(MarshalTypeUpdatePropertySetApexFaIntegrationDetails)(m),
	}

	return json.Marshal(&s)
}
