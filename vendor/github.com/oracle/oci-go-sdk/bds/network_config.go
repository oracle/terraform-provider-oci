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

// NetworkConfig Additional configuration of customer's network.
type NetworkConfig struct {

	// A boolean flag whether to configure a NAT gateway.
	IsNatGatewayRequired *bool `mandatory:"false" json:"isNatGatewayRequired"`

	// The CIDR IP address block of the VCN.
	CidrBlock *string `mandatory:"false" json:"cidrBlock"`
}

func (m NetworkConfig) String() string {
	return common.PointerString(m)
}
