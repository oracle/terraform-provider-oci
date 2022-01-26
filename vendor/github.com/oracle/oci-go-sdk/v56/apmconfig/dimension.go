// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Dimension A dimension is a label that can be used to describe or group metrics.
type Dimension struct {

	// The dimension name
	Name *string `mandatory:"true" json:"name"`

	// The source to populate the dimension. Must be NULL at the moment.
	ValueSource *string `mandatory:"false" json:"valueSource"`
}

func (m Dimension) String() string {
	return common.PointerString(m)
}
