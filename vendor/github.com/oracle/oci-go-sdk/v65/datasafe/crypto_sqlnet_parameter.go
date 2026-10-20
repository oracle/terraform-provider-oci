// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoSqlnetParameter One SQLNET parameter and its evaluated quantum readiness.
type CryptoSqlnetParameter struct {

	// SQLNET parameter name.
	Name *string `mandatory:"true" json:"name"`

	Value *CryptoSqlnetParameterValue `mandatory:"true" json:"value"`

	// Quantum-readiness classification of this parameter.
	QuantumReadiness CryptoQuantumReadinessEnum `mandatory:"true" json:"quantumReadiness"`
}

func (m CryptoSqlnetParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoSqlnetParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoQuantumReadinessEnum(string(m.QuantumReadiness)); !ok && m.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", m.QuantumReadiness, strings.Join(GetCryptoQuantumReadinessEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
