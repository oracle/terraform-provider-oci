// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VectorUserConfig Details for the Autonomous AI Vector Database user.
type VectorUserConfig struct {

	// Required when using the Autonomous AI Vector Database workload type. For all other workload types, it is an error to provide this field.
	Username *string `mandatory:"false" json:"username"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	// This cannot be used in conjunction with with OCI vault secrets (secretId).  When using the Autonomous AI Vector Database workload type, it is required to provide either this field or vectorSecretId. For all other workload types, it is an error to provide this field.
	Password *string `mandatory:"false" json:"password"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with password. When using the Autonomous AI Vector Database workload type, it is required to provide either this field or vectorPassword. For all other workload types, it is an error to provide this field.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used. Used in conjunction with vectorSecretId
	SecretVersionNumber *int `mandatory:"false" json:"secretVersionNumber"`
}

func (m VectorUserConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VectorUserConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
