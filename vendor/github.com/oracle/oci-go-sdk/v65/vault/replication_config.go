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

// ReplicationConfig Defines the configuration that enables cross-region secret replication.
type ReplicationConfig struct {

	// List of the secret replication targets. By default, a maximum of 3 targets is allowed. To configure more than 3 targets, an override is required.
	ReplicationTargets []ReplicationTarget `mandatory:"true" json:"replicationTargets"`

	// (Optional) A Boolean value to enable forwarding of write requests from replicated secrets to the source secrets. The default value of false disables this option.
	IsWriteForwardEnabled *bool `mandatory:"false" json:"isWriteForwardEnabled"`
}

func (m ReplicationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
