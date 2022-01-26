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

// MonitorScriptParameter Details of the script parameter that can be used to overwrite the parameter present in the script.
type MonitorScriptParameter struct {

	// Name of the parameter.
	ParamName *string `mandatory:"true" json:"paramName"`

	// Value of the parameter.
	ParamValue *string `mandatory:"true" json:"paramValue"`
}

func (m MonitorScriptParameter) String() string {
	return common.PointerString(m)
}
