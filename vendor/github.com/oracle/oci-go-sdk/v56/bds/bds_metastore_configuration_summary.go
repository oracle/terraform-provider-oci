// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BdsMetastoreConfigurationSummary The summary of metastore configuration information.
type BdsMetastoreConfigurationSummary struct {

	// The ID of the metastore configuration
	Id *string `mandatory:"true" json:"id"`

	// The display name of metastore configuration
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the metastore in the metastore configuration.
	MetastoreType BdsMetastoreConfigurationMetastoreTypeEnum `mandatory:"true" json:"metastoreType"`

	// the lifecycle state of the metastore configuration.
	LifecycleState BdsMetastoreConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the Data Catalog metastore. Set only if metastore's type is EXTERNAL.
	MetastoreId *string `mandatory:"false" json:"metastoreId"`

	// The ID of BDS API Key used for metastore configuration. Set only if metastore's type is EXTERNAL.
	BdsApiKeyId *string `mandatory:"false" json:"bdsApiKeyId"`

	// The time when the configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsMetastoreConfigurationSummary) String() string {
	return common.PointerString(m)
}
