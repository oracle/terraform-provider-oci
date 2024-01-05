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

// CloudSqlDetails The information about added Cloud SQL capability
type CloudSqlDetails struct {

	// Shape of the node
	Shape *string `mandatory:"true" json:"shape"`

	// IP address of the Cloud SQL node.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The size of block volume in GB that needs to be attached to a given node.
	// All the necessary details needed for attachment are managed by service itself.
	BlockVolumeSizeInGBs *int64 `mandatory:"false" json:"blockVolumeSizeInGBs"`

	// Boolean flag specifying whether or not Kerberos principals are mapped
	// to database users.
	IsKerberosMappedToDatabaseUsers *bool `mandatory:"false" json:"isKerberosMappedToDatabaseUsers"`

	// Details about the Kerberos principals.
	KerberosDetails []KerberosDetails `mandatory:"false" json:"kerberosDetails"`
}

func (m CloudSqlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudSqlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
