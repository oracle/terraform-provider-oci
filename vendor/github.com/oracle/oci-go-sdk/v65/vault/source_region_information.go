// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceRegionInformation Details for the source that the source secret has.
type SourceRegionInformation struct {

	// The OCID of the source region KMS key.
	SourceKeyId *string `mandatory:"true" json:"sourceKeyId"`

	// The name of the source's region.
	SourceRegion *string `mandatory:"true" json:"sourceRegion"`

	// The OCID of the source region's Vault.
	SourceVaultId *string `mandatory:"true" json:"sourceVaultId"`
}

func (m SourceRegionInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceRegionInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
