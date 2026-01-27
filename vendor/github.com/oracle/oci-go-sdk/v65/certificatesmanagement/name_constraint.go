// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NameConstraint A constraint that specifies permitted and excluded namespaces for the hierarchical name forms in certificates that any CA in the certificate chain issues. You can define name constraints on a directory name, DNS address, or IP address. If you have a name constraint, you must define at least one permitted namespace or one excluded namespace.
type NameConstraint struct {

	// A list that contains permitted namespaces. If you have a name constraint with no excluded namespaces, you must specify at least one permitted namespace.
	PermittedSubtree []NameConstraintSubtreeNode `mandatory:"false" json:"permittedSubtree"`

	// A list that contains excluded (or prohibited) namespaces. If you have a name constraint with no permitted namespaces, you must specify at least one excluded namespace.
	ExcludedSubtree []NameConstraintSubtreeNode `mandatory:"false" json:"excludedSubtree"`
}

func (m NameConstraint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NameConstraint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
