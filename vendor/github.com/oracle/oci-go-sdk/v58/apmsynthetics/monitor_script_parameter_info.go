// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MonitorScriptParameterInfo Details of the script parameters in the monitor.
// isOverwritten specifies that the script parameters are overwritten in the monitor.
// If the user overwrites the parameter value in the monitor, then the overwritten values will be used to run the monitor.
type MonitorScriptParameterInfo struct {
	MonitorScriptParameter *MonitorScriptParameter `mandatory:"true" json:"monitorScriptParameter"`

	// Describes if  the parameter value is secret and should be kept confidential.
	// isSecret is specified in either CreateScript or UpdateScript API.
	IsSecret *bool `mandatory:"true" json:"isSecret"`

	// If parameter value is default or overwritten.
	IsOverwritten *bool `mandatory:"true" json:"isOverwritten"`
}

func (m MonitorScriptParameterInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitorScriptParameterInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
