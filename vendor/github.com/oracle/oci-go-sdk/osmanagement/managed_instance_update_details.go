// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
