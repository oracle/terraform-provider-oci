// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LiteralArgument QueryString argument of type literal.
type LiteralArgument struct {

	// Data type of specified literal in queryString.
	DataType *string `mandatory:"false" json:"dataType"`

	// Literal value specified in queryString.
	Value *interface{} `mandatory:"false" json:"value"`
}

func (m LiteralArgument) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LiteralArgument) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLiteralArgument LiteralArgument
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeLiteralArgument
	}{
		"LITERAL",
		(MarshalTypeLiteralArgument)(m),
	}

	return json.Marshal(&s)
}
