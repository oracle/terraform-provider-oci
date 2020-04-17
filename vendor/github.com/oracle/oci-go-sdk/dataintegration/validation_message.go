// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ValidationMessage The level, message key and validation message.
type ValidationMessage struct {

	// Total number of validation messages
	Level *string `mandatory:"false" json:"level"`

	// The key.
	MessageKey *string `mandatory:"false" json:"messageKey"`

	// The message itself.
	ValidationMessage *string `mandatory:"false" json:"validationMessage"`
}

func (m ValidationMessage) String() string {
	return common.PointerString(m)
}
