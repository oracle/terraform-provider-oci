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

// OlvmAgent Type representing a fence agent.
type OlvmAgent struct {

	// Fence agent address
	Address *string `mandatory:"false" json:"address"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// Specified whether the agent should be used concurrently or sequentially
	IsConcurrent *bool `mandatory:"false" json:"isConcurrent"`

	// Free text containing comments about this object.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether the options should be encrypted.
	IsEncryptOptions *bool `mandatory:"false" json:"isEncryptOptions"`

	// A unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// A human-readable name in plain text
	Name *string `mandatory:"false" json:"name"`

	// The order of this agent if used with other agents.
	Order *int `mandatory:"false" json:"order"`

	// Fence agent port.
	Port *int `mandatory:"false" json:"port"`

	// Fence agent type.
	Type *string `mandatory:"false" json:"type"`

	// Fence agent user name.
	Username *string `mandatory:"false" json:"username"`
}

func (m OlvmAgent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmAgent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
