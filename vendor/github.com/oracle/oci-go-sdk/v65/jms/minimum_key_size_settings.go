// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MinimumKeySizeSettings test
type MinimumKeySizeSettings struct {

	// Updates the minimum key size for the specified encryption algorithm.
	// The JDK property jdk.tls.disabledAlgorithms will be updated with the following supported actions:
	// - Changing minimum key length for Diffie-Hellman
	Tls []KeySizeAlgorithm `mandatory:"false" json:"tls"`

	// Updates the minimum key size for the specified encryption algorithm.
	// The JDK property jdk.jar.disabledAlgorithms will be updated with the following supported actions:
	// - Changing minimum key length for RSA signed jars
	// - Changing minimum key length for EC
	// - Changing minimum key length for DSA
	Jar []KeySizeAlgorithm `mandatory:"false" json:"jar"`

	// Updates the minimum key size for the specified encryption algorithm.
	// The JDK property jdk.certpath.disabledAlgorithms will be updated with the following supported actions:
	// - Changing minimum key length for RSA signed jars
	// - Changing minimum key length for EC
	// - Changing minimum key length for DSA
	Certpath []KeySizeAlgorithm `mandatory:"false" json:"certpath"`
}

func (m MinimumKeySizeSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MinimumKeySizeSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
