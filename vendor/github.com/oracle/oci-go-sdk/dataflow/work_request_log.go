// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestLog A Data Flow work request log object.
type WorkRequestLog struct {

	// A human readable log message.
	Message *string `mandatory:"true" json:"message"`

	// The time the log message was written. An RFC3339 formatted datetime string.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// The id of a work request log.
	Id *int64 `mandatory:"false" json:"id"`

	// The OCID of a work request.
	WorkRequestid *string `mandatory:"false" json:"workRequestid"`
}

func (m WorkRequestLog) String() string {
	return common.PointerString(m)
}
