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
