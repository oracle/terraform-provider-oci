// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ClusterDetails Specific info about a Hadoop cluster
type ClusterDetails struct {

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// BDA version installed in the cluster
	BdaVersion *string `mandatory:"false" json:"bdaVersion"`

	// Big Data Manager version installed in the cluster.
	BdmVersion *string `mandatory:"false" json:"bdmVersion"`

	// Big Data Service version installed in the cluster.
	BdsVersion *string `mandatory:"false" json:"bdsVersion"`

	// Oracle Linux version installed in the cluster.
	OsVersion *string `mandatory:"false" json:"osVersion"`

	// Cloud SQL query server database version.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Cloud SQL cell version.
	BdCellVersion *string `mandatory:"false" json:"bdCellVersion"`

	// Big Data SQL version.
	CsqlCellVersion *string `mandatory:"false" json:"csqlCellVersion"`

	// The time the cluster was automatically or manually refreshed, shown as an RFC 3339 formatted datetime string.
	TimeRefreshed *common.SDKTime `mandatory:"false" json:"timeRefreshed"`

	// The URL of Cloudera Manager
	ClouderaManagerUrl *string `mandatory:"false" json:"clouderaManagerUrl"`

	// The URL of Ambari
	AmbariUrl *string `mandatory:"false" json:"ambariUrl"`

	// The URL of Big Data Manager.
	BigDataManagerUrl *string `mandatory:"false" json:"bigDataManagerUrl"`

	// The URL of the Hue server.
	HueServerUrl *string `mandatory:"false" json:"hueServerUrl"`

	// Version of the ODH (Oracle Distribution including Apache Hadoop) installed on the cluster.
	OdhVersion *string `mandatory:"false" json:"odhVersion"`

	// The URL of the Jupyterhub.
	JupyterHubUrl *string `mandatory:"false" json:"jupyterHubUrl"`
}

func (m ClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
