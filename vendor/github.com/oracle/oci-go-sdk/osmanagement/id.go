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

// Id An id along with a name to simplify display for a user
type Id struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// User friendly name
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m Id) String() string {
	return common.PointerString(m)
}
