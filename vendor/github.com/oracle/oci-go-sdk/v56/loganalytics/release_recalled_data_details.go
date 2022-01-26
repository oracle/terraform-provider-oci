// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ReleaseRecalledDataDetails This is the input used to release recalled data
type ReleaseRecalledDataDetails struct {

	// This is the compartment OCID for permission checking
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is the end of the time interval
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the start of the time interval
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the type of the recalled data to be released
	DataType StorageDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`
}

func (m ReleaseRecalledDataDetails) String() string {
	return common.PointerString(m)
}
