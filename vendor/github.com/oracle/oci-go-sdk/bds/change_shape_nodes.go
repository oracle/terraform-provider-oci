// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

// ChangeShapeNodes The representation of ChangeShapeNodes
type ChangeShapeNodes struct {

	// worker nodes shape
	Worker *string `mandatory:"false" json:"worker"`

	// master nodes shape
	Master *string `mandatory:"false" json:"master"`

	// utility nodes shape
	Utility *string `mandatory:"false" json:"utility"`

	// cloudsql node shape
	Cloudsql *string `mandatory:"false" json:"cloudsql"`
}

func (m ChangeShapeNodes) String() string {
	return common.PointerString(m)
}
