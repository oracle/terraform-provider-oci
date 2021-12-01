// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// API to manage Service manager proxy.
//

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// ServiceEnvironmentEndPointOverview Model describing the properties of service environment endPoint overview.
type ServiceEnvironmentEndPointOverview struct {

	// Service Environemnt EndPoint type.
	EnvironmentType ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum `mandatory:"true" json:"environmentType"`

	// Service Environemnt Instance EndPoint url.
	Url *string `mandatory:"true" json:"url"`
}

func (m ServiceEnvironmentEndPointOverview) String() string {
	return common.PointerString(m)
}

// ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum Enum with underlying type: string
type ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum string

// Set of constants representing the allowable values for ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum
const (
	ServiceEnvironmentEndPointOverviewEnvironmentTypeProd ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_PROD"
	ServiceEnvironmentEndPointOverviewEnvironmentTypeTest ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_TEST"
	ServiceEnvironmentEndPointOverviewEnvironmentTypeDev  ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_DEV"
)

var mappingServiceEnvironmentEndPointOverviewEnvironmentType = map[string]ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum{
	"INSTANCE_URL_PROD": ServiceEnvironmentEndPointOverviewEnvironmentTypeProd,
	"INSTANCE_URL_TEST": ServiceEnvironmentEndPointOverviewEnvironmentTypeTest,
	"INSTANCE_URL_DEV":  ServiceEnvironmentEndPointOverviewEnvironmentTypeDev,
}

// GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumValues Enumerates the set of values for ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum
func GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumValues() []ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum {
	values := make([]ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum, 0)
	for _, v := range mappingServiceEnvironmentEndPointOverviewEnvironmentType {
		values = append(values, v)
	}
	return values
}
