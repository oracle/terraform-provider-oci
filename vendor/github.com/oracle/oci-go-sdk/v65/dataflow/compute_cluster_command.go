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

// ComputeClusterCommand A compute cluster execution command.
type ComputeClusterCommand struct {

	// The unique identifier of a compute cluster command.
	Id *string `mandatory:"true" json:"id"`

	// The time the Compute Cluster Command was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Supported software coding language.
	Language LanguageEnum `mandatory:"true" json:"language"`

	// Compute Cluster Command Lifecycle state.
	Status ComputeClusterCommandLifecycleStateEnum `mandatory:"true" json:"status"`

	Result *ComputeClusterCommandResult `mandatory:"true" json:"result"`

	// Command to be executed.
	Command *string `mandatory:"false" json:"command"`
}

func (m ComputeClusterCommand) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterCommand) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComputeClusterCommandLifecycleStateEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetComputeClusterCommandLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
