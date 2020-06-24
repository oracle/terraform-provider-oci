// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ErrorDetails The details of an error that occured.
type ErrorDetails struct {

	// A short error code that defines the error, meant for programmatic parsing. See
	// API Errors (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm).
	Code *string `mandatory:"true" json:"code"`

	// A user-friendly error message.
	Message *string `mandatory:"true" json:"message"`
}

func (m ErrorDetails) String() string {
	return common.PointerString(m)
}
