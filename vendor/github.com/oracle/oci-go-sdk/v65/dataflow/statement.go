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

// Statement A statement object.
type Statement struct {

	// The statement ID.
	Id *int64 `mandatory:"true" json:"id"`

	// The statement code to execute.
	// Example: `println(sc.version)`
	Code *string `mandatory:"true" json:"code"`

	// The current state of this statement.
	LifecycleState StatementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	Output *StatementOutput `mandatory:"false" json:"output"`

	// The execution progress.
	Progress *float64 `mandatory:"false" json:"progress"`

	// The ID of a run.
	RunId *string `mandatory:"false" json:"runId"`

	// The date and time a statement execution was completed, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2022-05-31T21:10:29.600Z`
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m Statement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Statement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStatementLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStatementLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
