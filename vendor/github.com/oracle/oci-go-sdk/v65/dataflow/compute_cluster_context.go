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

// ComputeClusterContext A compute cluster context object.
type ComputeClusterContext struct {

	// The provision identifier that is immutable on creation.
	ComputeClusterId *string `mandatory:"true" json:"computeClusterId"`

	// The unique identifier of a compute cluster context.
	Id *string `mandatory:"true" json:"id"`

	// Supported software coding language.
	DefaultLanguage LanguageEnum `mandatory:"true" json:"defaultLanguage"`

	// Compute Cluster Context State.
	LifecycleState ComputeClusterContextLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Compute Cluster Context was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m ComputeClusterContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLanguageEnum(string(m.DefaultLanguage)); !ok && m.DefaultLanguage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultLanguage: %s. Supported values are: %s.", m.DefaultLanguage, strings.Join(GetLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComputeClusterContextLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComputeClusterContextLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
