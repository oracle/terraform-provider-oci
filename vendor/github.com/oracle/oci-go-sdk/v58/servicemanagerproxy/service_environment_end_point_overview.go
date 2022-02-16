// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ServiceEnvironmentEndPointOverview An overview of service environment endpoints.
type ServiceEnvironmentEndPointOverview struct {

	// Service environment endpoint type.
	EnvironmentType ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum `mandatory:"true" json:"environmentType"`

	// Service environment instance URL.
	Url *string `mandatory:"true" json:"url"`

	// Description of the environment link
	Description *string `mandatory:"false" json:"description"`
}

func (m ServiceEnvironmentEndPointOverview) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceEnvironmentEndPointOverview) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum Enum with underlying type: string
type ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum string

// Set of constants representing the allowable values for ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum
const (
	ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlProd ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_PROD"
	ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlTest ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_TEST"
	ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlDev  ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "INSTANCE_URL_DEV"
	ServiceEnvironmentEndPointOverviewEnvironmentTypeOther           ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = "OTHER"
)

var mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum = map[string]ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum{
	"INSTANCE_URL_PROD": ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlProd,
	"INSTANCE_URL_TEST": ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlTest,
	"INSTANCE_URL_DEV":  ServiceEnvironmentEndPointOverviewEnvironmentTypeInstanceUrlDev,
	"OTHER":             ServiceEnvironmentEndPointOverviewEnvironmentTypeOther,
}

// GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumValues Enumerates the set of values for ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum
func GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumValues() []ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum {
	values := make([]ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum, 0)
	for _, v := range mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumStringValues Enumerates the set of values in String for ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum
func GetServiceEnvironmentEndPointOverviewEnvironmentTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_URL_PROD",
		"INSTANCE_URL_TEST",
		"INSTANCE_URL_DEV",
		"OTHER",
	}
}

// GetMappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum(val string) (ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum, bool) {
	mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnumIgnoreCase := make(map[string]ServiceEnvironmentEndPointOverviewEnvironmentTypeEnum)
	for k, v := range mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnum {
		mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingServiceEnvironmentEndPointOverviewEnvironmentTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
