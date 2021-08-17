// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ChildReferenceDetail References used in an application.
type ChildReferenceDetail struct {

	// The child reference key.
	Key *string `mandatory:"false" json:"key"`

	// The new reference object to use instead of the original reference. For example, this can be a connection reference.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`
}

func (m ChildReferenceDetail) String() string {
	return common.PointerString(m)
}
