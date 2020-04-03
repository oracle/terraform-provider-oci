// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeApplicationCompartmentDetails The change application compartment details.
type ChangeApplicationCompartmentDetails struct {

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeApplicationCompartmentDetails) String() string {
	return common.PointerString(m)
}
