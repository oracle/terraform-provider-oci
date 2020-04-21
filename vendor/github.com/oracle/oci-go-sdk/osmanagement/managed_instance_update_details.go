// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagedInstanceUpdateDetails Updated information for the managed instance
type ManagedInstanceUpdateDetails struct {

	// Managed Instance identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Information specified by the user about the managed instance
	Description *string `mandatory:"false" json:"description"`
}

func (m ManagedInstanceUpdateDetails) String() string {
	return common.PointerString(m)
}
