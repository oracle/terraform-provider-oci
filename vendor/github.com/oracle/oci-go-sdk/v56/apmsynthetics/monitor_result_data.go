// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MonitorResultData Details of the monitor result data.
type MonitorResultData struct {

	// Name of the data.
	Name *string `mandatory:"false" json:"name"`

	// Data content in byte format.
	// Example: Zip or Screenshot.
	ByteContent []byte `mandatory:"false" json:"byteContent"`

	// Data content in string format.
	// Example: HAR.
	StringContent *string `mandatory:"false" json:"stringContent"`

	// The time when the data was generated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m MonitorResultData) String() string {
	return common.PointerString(m)
}
