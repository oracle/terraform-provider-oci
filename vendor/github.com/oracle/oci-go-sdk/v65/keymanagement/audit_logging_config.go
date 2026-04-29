// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditLoggingConfig Configuration for management audit logging for the cluster.
type AuditLoggingConfig struct {

	// Indicates if management audit logging is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Name of the object store bucket for audit logs.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Object storage namespace for the bucket.
	Namespace *string `mandatory:"false" json:"namespace"`

	// OCID of the compartment for the bucket.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m AuditLoggingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditLoggingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
