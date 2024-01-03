// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StackTraceElement Stack trace element.
type StackTraceElement struct {

	// Name of the method containing the execution point.
	MethodName *string `mandatory:"false" json:"methodName"`

	// Name of the source file containing the execution point.
	FileName *string `mandatory:"false" json:"fileName"`

	// Line number of the source line containing the execution point.
	LineNumber *int `mandatory:"false" json:"lineNumber"`

	// Name of the class containing the execution point.
	ClassName *string `mandatory:"false" json:"className"`

	// The weight distribution that denotes the percentage occurrence of a method in the captured snapshots.
	Weightage *float32 `mandatory:"false" json:"weightage"`
}

func (m StackTraceElement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackTraceElement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
