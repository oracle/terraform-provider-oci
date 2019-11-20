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

// SoftwareSourceId Identifying information for the specified software source
type SoftwareSourceId struct {

	// software source identifier
	Id *string `mandatory:"true" json:"id"`

	// software source name
	Name *string `mandatory:"false" json:"name"`
}

func (m SoftwareSourceId) String() string {
	return common.PointerString(m)
}
