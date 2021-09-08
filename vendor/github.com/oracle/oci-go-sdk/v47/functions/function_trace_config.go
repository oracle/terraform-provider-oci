// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// FunctionTraceConfig Define the tracing configuration for a function.
type FunctionTraceConfig struct {

	// Define if tracing is enabled for the resource.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m FunctionTraceConfig) String() string {
	return common.PointerString(m)
}
