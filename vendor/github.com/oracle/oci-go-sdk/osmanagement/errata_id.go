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

// ErrataId Identifying information for an errata
type ErrataId struct {

	// errata identifier
	Id *string `mandatory:"true" json:"id"`

	// errata name
	Name *string `mandatory:"false" json:"name"`
}

func (m ErrataId) String() string {
	return common.PointerString(m)
}
