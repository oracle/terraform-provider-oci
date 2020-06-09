// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ErrorBody Error Information.
type ErrorBody struct {

	// A short error code that defines the error, which is useful for programmatic parsing.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m ErrorBody) String() string {
	return common.PointerString(m)
}
