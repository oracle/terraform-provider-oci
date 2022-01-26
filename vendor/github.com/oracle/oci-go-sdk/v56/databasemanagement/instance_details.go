// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstanceDetails The details of the Oracle Real Application Clusters (Oracle RAC) database instance.
type InstanceDetails struct {

	// The ID of the Oracle RAC database instance.
	Id *int `mandatory:"true" json:"id"`

	// The name of the Oracle RAC database instance.
	Name *string `mandatory:"true" json:"name"`

	// The name of the host of the Oracle RAC database instance.
	HostName *string `mandatory:"true" json:"hostName"`

	// The status of the Oracle RAC database instance.
	Status InstanceDetailsStatusEnum `mandatory:"true" json:"status"`
}

func (m InstanceDetails) String() string {
	return common.PointerString(m)
}

// InstanceDetailsStatusEnum Enum with underlying type: string
type InstanceDetailsStatusEnum string

// Set of constants representing the allowable values for InstanceDetailsStatusEnum
const (
	InstanceDetailsStatusUp      InstanceDetailsStatusEnum = "UP"
	InstanceDetailsStatusDown    InstanceDetailsStatusEnum = "DOWN"
	InstanceDetailsStatusUnknown InstanceDetailsStatusEnum = "UNKNOWN"
)

var mappingInstanceDetailsStatus = map[string]InstanceDetailsStatusEnum{
	"UP":      InstanceDetailsStatusUp,
	"DOWN":    InstanceDetailsStatusDown,
	"UNKNOWN": InstanceDetailsStatusUnknown,
}

// GetInstanceDetailsStatusEnumValues Enumerates the set of values for InstanceDetailsStatusEnum
func GetInstanceDetailsStatusEnumValues() []InstanceDetailsStatusEnum {
	values := make([]InstanceDetailsStatusEnum, 0)
	for _, v := range mappingInstanceDetailsStatus {
		values = append(values, v)
	}
	return values
}
