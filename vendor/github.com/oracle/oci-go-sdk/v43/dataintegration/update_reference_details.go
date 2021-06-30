// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v43/common"
)

// UpdateReferenceDetails Application references that need to be updated.
type UpdateReferenceDetails struct {

	// A list of options such as `ignoreObjectOnError`.
	Options map[string]string `mandatory:"false" json:"options"`

	// The new target object to reference. This should be of type `DataAsset`. The child references can be of type `Connection`.
	TargetObject *interface{} `mandatory:"false" json:"targetObject"`

	// The list of child references that also need to be updated.
	ChildReferences []ChildReferenceDetail `mandatory:"false" json:"childReferences"`
}

func (m UpdateReferenceDetails) String() string {
	return common.PointerString(m)
}
