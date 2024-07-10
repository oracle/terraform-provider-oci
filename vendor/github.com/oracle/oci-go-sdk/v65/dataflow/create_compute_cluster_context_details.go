// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateComputeClusterContextDetails Details about a new compute cluster Execution Context.
type CreateComputeClusterContextDetails struct {

	// Supported software coding language.
	DefaultLanguage LanguageEnum `mandatory:"true" json:"defaultLanguage"`
}

func (m CreateComputeClusterContextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeClusterContextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLanguageEnum(string(m.DefaultLanguage)); !ok && m.DefaultLanguage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultLanguage: %s. Supported values are: %s.", m.DefaultLanguage, strings.Join(GetLanguageEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
