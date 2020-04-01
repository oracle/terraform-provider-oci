// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeBdsInstanceCompartmentDetails Moves a BDS instance into a different compartment.
type ChangeBdsInstanceCompartmentDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeBdsInstanceCompartmentDetails) String() string {
	return common.PointerString(m)
}
