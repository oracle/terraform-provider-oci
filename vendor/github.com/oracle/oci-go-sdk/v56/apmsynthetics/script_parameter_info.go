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

// ScriptParameterInfo Information about script parameters.
// isOverwritten specifies that the default parameter present in the script content is overwritten.
type ScriptParameterInfo struct {
	ScriptParameter *ScriptParameter `mandatory:"true" json:"scriptParameter"`

	// If parameter value is default or overwritten.
	IsOverwritten *bool `mandatory:"true" json:"isOverwritten"`
}

func (m ScriptParameterInfo) String() string {
	return common.PointerString(m)
}
