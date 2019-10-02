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

// ClusterDetails Specific info about a Hadoop cluster
type ClusterDetails struct {

	// The time the cluster was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// BDA version installed in the cluster
	BdaVersion *string `mandatory:"false" json:"bdaVersion"`

	// BDM version installed in the cluster
	BdmVersion *string `mandatory:"false" json:"bdmVersion"`

	// The time the BDS instance was automatically, or manually refreshed.
	// An RFC3339 formatted datetime string
	TimeRefreshed *common.SDKTime `mandatory:"false" json:"timeRefreshed"`

	// The URL of a Cloudera Manager
	ClouderaManagerUrl *string `mandatory:"false" json:"clouderaManagerUrl"`

	// The URL of a Big Data Manager
	BigDataManagerUrl *string `mandatory:"false" json:"bigDataManagerUrl"`

	// The URL of a Hue Server
	HueServerUrl *string `mandatory:"false" json:"hueServerUrl"`
}

func (m ClusterDetails) String() string {
	return common.PointerString(m)
}
