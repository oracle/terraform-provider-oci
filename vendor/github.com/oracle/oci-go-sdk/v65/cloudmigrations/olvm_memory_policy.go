// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmMemoryPolicy Logical grouping of memory-related properties of virtual machine-like entities.
type OlvmMemoryPolicy struct {

	// Indicates if ballooning is enabled
	IsBallooning *bool `mandatory:"false" json:"isBallooning"`

	// The amount of memory, in bytes, that is guaranteed to not be drained by the balloon mechanism
	GuaranteedMemoryInBytes *int64 `mandatory:"false" json:"guaranteedMemoryInBytes"`

	// Maximum virtual machine memory in Bytes
	MaxMemoryInBytes *int64 `mandatory:"false" json:"maxMemoryInBytes"`

	MemoryOverCommit *MemoryOverCommit `mandatory:"false" json:"memoryOverCommit"`

	TransparentHugePages *TransparentHugePages `mandatory:"false" json:"transparentHugePages"`
}

func (m OlvmMemoryPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmMemoryPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
