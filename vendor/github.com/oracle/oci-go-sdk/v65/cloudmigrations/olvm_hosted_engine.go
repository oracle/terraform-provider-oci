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

// OlvmHostedEngine The self-hosted engine status of this host.
type OlvmHostedEngine struct {

	// Indicates if this hosted engine is active.
	IsActive *bool `mandatory:"false" json:"isActive"`

	// Indicates if this hosted engine is configured.
	IsConfigured *bool `mandatory:"false" json:"isConfigured"`

	// Indicates if this hosted engine under global maintenance.
	IsGlobalMaintenance *bool `mandatory:"false" json:"isGlobalMaintenance"`

	// Indicates if this hosted engine under local maintenance.
	IsLocalMaintenance *bool `mandatory:"false" json:"isLocalMaintenance"`

	// A numerical value representing the health and status of the Manager virtual machine
	Score *int `mandatory:"false" json:"score"`
}

func (m OlvmHostedEngine) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmHostedEngine) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
