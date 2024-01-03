// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOnPremiseVantagePointDetails Details of the request body used to create a new On-premise vantage point.
type CreateOnPremiseVantagePointDetails struct {

	// Unique On-premise vantage point name that cannot be edited. The name should not contain any confidential information.
	Name *string `mandatory:"true" json:"name"`

	// Type of On-premise vantage point.
	Type CreateOnPremiseVantagePointDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A short description about the On-premise vantage point.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOnPremiseVantagePointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOnPremiseVantagePointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateOnPremiseVantagePointDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateOnPremiseVantagePointDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOnPremiseVantagePointDetailsTypeEnum Enum with underlying type: string
type CreateOnPremiseVantagePointDetailsTypeEnum string

// Set of constants representing the allowable values for CreateOnPremiseVantagePointDetailsTypeEnum
const (
	CreateOnPremiseVantagePointDetailsTypeOnPremiseDockerVantagePoint CreateOnPremiseVantagePointDetailsTypeEnum = "ON_PREMISE_DOCKER_VANTAGE_POINT"
)

var mappingCreateOnPremiseVantagePointDetailsTypeEnum = map[string]CreateOnPremiseVantagePointDetailsTypeEnum{
	"ON_PREMISE_DOCKER_VANTAGE_POINT": CreateOnPremiseVantagePointDetailsTypeOnPremiseDockerVantagePoint,
}

var mappingCreateOnPremiseVantagePointDetailsTypeEnumLowerCase = map[string]CreateOnPremiseVantagePointDetailsTypeEnum{
	"on_premise_docker_vantage_point": CreateOnPremiseVantagePointDetailsTypeOnPremiseDockerVantagePoint,
}

// GetCreateOnPremiseVantagePointDetailsTypeEnumValues Enumerates the set of values for CreateOnPremiseVantagePointDetailsTypeEnum
func GetCreateOnPremiseVantagePointDetailsTypeEnumValues() []CreateOnPremiseVantagePointDetailsTypeEnum {
	values := make([]CreateOnPremiseVantagePointDetailsTypeEnum, 0)
	for _, v := range mappingCreateOnPremiseVantagePointDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOnPremiseVantagePointDetailsTypeEnumStringValues Enumerates the set of values in String for CreateOnPremiseVantagePointDetailsTypeEnum
func GetCreateOnPremiseVantagePointDetailsTypeEnumStringValues() []string {
	return []string{
		"ON_PREMISE_DOCKER_VANTAGE_POINT",
	}
}

// GetMappingCreateOnPremiseVantagePointDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOnPremiseVantagePointDetailsTypeEnum(val string) (CreateOnPremiseVantagePointDetailsTypeEnum, bool) {
	enum, ok := mappingCreateOnPremiseVantagePointDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
