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

// ScriptParameter Details of the script parameters, paramName must be from the script content and
// these details can be used to overwrite the default parameter present in the script content.
type ScriptParameter struct {

	// Name of the parameter.
	ParamName *string `mandatory:"true" json:"paramName"`

	// Value of the parameter.
	ParamValue *string `mandatory:"false" json:"paramValue"`

	// If the parameter value is secret and should be kept confidential, then set isSecret to true.
	IsSecret *bool `mandatory:"false" json:"isSecret"`
}

func (m ScriptParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScriptParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
