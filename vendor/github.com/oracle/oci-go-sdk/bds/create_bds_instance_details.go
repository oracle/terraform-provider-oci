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

// CreateBdsInstanceDetails The information about new BDS instance
type CreateBdsInstanceDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the BDS instance
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Version of the Hadoop distribution
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"true" json:"clusterVersion"`

	// The SSH public key used to authenticate the cluster connection.
	ClusterPublicKey *string `mandatory:"true" json:"clusterPublicKey"`

	// Base-64 encoded password for Cloudera Manager admin user
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// Boolean flag specifying whether or not the cluster is HA
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be setup as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// The list of nodes in the BDS instance
	Nodes []CreateNodeDetails `mandatory:"true" json:"nodes"`

	// Additional configuration of customer's network.
	NetworkConfig *NetworkConfig `mandatory:"false" json:"networkConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBdsInstanceDetails) String() string {
	return common.PointerString(m)
}
