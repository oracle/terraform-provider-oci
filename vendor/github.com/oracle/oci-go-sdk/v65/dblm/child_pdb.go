// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChildPdb Child PDB for a CDB.
type ChildPdb struct {

	// Identifier for the pluggable database.
	PdbId *string `mandatory:"true" json:"pdbId"`

	// Name fo the pluggable database.
	Name *string `mandatory:"true" json:"name"`

	// Open mode
	OpenMode *string `mandatory:"true" json:"openMode"`

	// Restricted
	Restricted *string `mandatory:"true" json:"restricted"`

	// Recovery status
	RecoveryStatus *string `mandatory:"true" json:"recoveryStatus"`

	// Last changed by
	LastChangedBy *string `mandatory:"true" json:"lastChangedBy"`
}

func (m ChildPdb) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChildPdb) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
