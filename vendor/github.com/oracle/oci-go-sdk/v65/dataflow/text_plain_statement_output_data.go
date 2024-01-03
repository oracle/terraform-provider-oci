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

// TextPlainStatementOutputData The statement output data in text format.
type TextPlainStatementOutputData struct {

	// The statement code execution output in text format.
	Value *string `mandatory:"true" json:"value"`
}

func (m TextPlainStatementOutputData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextPlainStatementOutputData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextPlainStatementOutputData) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextPlainStatementOutputData TextPlainStatementOutputData
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTextPlainStatementOutputData
	}{
		"TEXT_PLAIN",
		(MarshalTypeTextPlainStatementOutputData)(m),
	}

	return json.Marshal(&s)
}
