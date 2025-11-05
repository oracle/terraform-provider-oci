// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbCredentialsDetail DB credentials for ADB connection.
type DbCredentialsDetail struct {

	// Username for ADB to be created or updated.
	UserName *string `mandatory:"true" json:"userName"`

	// Vault secret OCID containing the corresponding user password.
	SecretId *string `mandatory:"true" json:"secretId"`

	// Type of the user. Allowed values are "ADMIN" or "CUSTOM" or "GGCS".
	UserType *string `mandatory:"true" json:"userType"`
}

func (m DbCredentialsDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbCredentialsDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
