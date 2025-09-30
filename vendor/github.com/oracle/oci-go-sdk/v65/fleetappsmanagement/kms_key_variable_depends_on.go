// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KmsKeyVariableDependsOn Depends on object.
type KmsKeyVariableDependsOn struct {

	// OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the Vault.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The KMS crypto protection mode string.
	ProtectionMode *string `mandatory:"false" json:"protectionMode"`

	// Key cryptographic algorithm (RSA, AES, etc).
	Algorithm *string `mandatory:"false" json:"algorithm"`

	// Length of the cryptographic key in bits.
	Length *int `mandatory:"false" json:"length"`

	// Curve identifier for key creation.
	CurveId *string `mandatory:"false" json:"curveId"`
}

func (m KmsKeyVariableDependsOn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KmsKeyVariableDependsOn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
