// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ReferenceUsedBy Referenced object information.
type ReferenceUsedBy struct {

	// The key of the published object.
	Key *string `mandatory:"false" json:"key"`

	// The name of an published object.
	Name *string `mandatory:"false" json:"name"`

	// The name path of the published object.
	NamePath *string `mandatory:"false" json:"namePath"`
}

func (m ReferenceUsedBy) String() string {
	return common.PointerString(m)
}
