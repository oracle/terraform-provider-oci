// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeClusterCommandResult Compute Cluster Command execution Result Information.
type ComputeClusterCommandResult struct {

	// Summary of the error.
	Summary *string `mandatory:"false" json:"summary"`

	// Cause of the error.
	Cause *string `mandatory:"false" json:"cause"`

	Data ComputeClusterCommandOutputData `mandatory:"false" json:"data"`
}

func (m ComputeClusterCommandResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterCommandResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ComputeClusterCommandResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Summary *string                         `json:"summary"`
		Cause   *string                         `json:"cause"`
		Data    computeclustercommandoutputdata `json:"data"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Summary = model.Summary

	m.Cause = model.Cause

	nn, e = model.Data.UnmarshalPolymorphicJSON(model.Data.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Data = nn.(ComputeClusterCommandOutputData)
	} else {
		m.Data = nil
	}

	return
}
