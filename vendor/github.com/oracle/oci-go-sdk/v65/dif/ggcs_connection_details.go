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

// GgcsConnectionDetails Details of GGCS connections to be created.
type GgcsConnectionDetails struct {

	// Name of the connection to be created.
	ConnectionName *string `mandatory:"true" json:"connectionName"`

	// List of Service Dependency Details for connection creation.
	DifDependencies []DifDependencyDetails `mandatory:"false" json:"difDependencies"`

	// Vault secret OCID containing password that Oracle GoldenGate uses to connect the associated system of the given technology.
	GgAdminSecretId *string `mandatory:"false" json:"ggAdminSecretId"`

	// OCID of pre-created Oracle GoldenGate connection.
	ConnectionId *string `mandatory:"false" json:"connectionId"`
}

func (m GgcsConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GgcsConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
