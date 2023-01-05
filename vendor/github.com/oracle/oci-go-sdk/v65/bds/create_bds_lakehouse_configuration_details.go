// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBdsLakehouseConfigurationDetails The request body while creating a BDS lakehouse configuration.
type CreateBdsLakehouseConfigurationDetails struct {

	// The OCID of the lakehouse.
	LakehouseId *string `mandatory:"true" json:"lakehouseId"`

	// The OCID of BDS API key used for the Lakehouse integration.
	BdsApiKeyId *string `mandatory:"true" json:"bdsApiKeyId"`

	// The base-64 encoded passphrase of the BDS API key.
	BdsApiKeyPassphrase *string `mandatory:"true" json:"bdsApiKeyPassphrase"`

	// The base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// The display name of the lakehouse configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateBdsLakehouseConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBdsLakehouseConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
