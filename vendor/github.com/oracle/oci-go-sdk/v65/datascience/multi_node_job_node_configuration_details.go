// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MultiNodeJobNodeConfigurationDetails MultiNodeJobNodeConfigurationDetails
type MultiNodeJobNodeConfigurationDetails struct {

	// A time bound for the execution of the job run. Timer starts when the job run is in progress.
	MaximumRuntimeInMinutes *int64 `mandatory:"false" json:"maximumRuntimeInMinutes"`

	JobNetworkConfiguration JobNetworkConfiguration `mandatory:"false" json:"jobNetworkConfiguration"`

	// List of JobNodeGroupConfigurationDetails
	JobNodeGroupConfigurationDetailsList []JobNodeGroupConfigurationDetails `mandatory:"false" json:"jobNodeGroupConfigurationDetailsList"`

	// The execution order of node groups
	StartupOrder MultiNodeJobNodeConfigurationDetailsStartupOrderEnum `mandatory:"false" json:"startupOrder,omitempty"`
}

func (m MultiNodeJobNodeConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MultiNodeJobNodeConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnum(string(m.StartupOrder)); !ok && m.StartupOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StartupOrder: %s. Supported values are: %s.", m.StartupOrder, strings.Join(GetMultiNodeJobNodeConfigurationDetailsStartupOrderEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MultiNodeJobNodeConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMultiNodeJobNodeConfigurationDetails MultiNodeJobNodeConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"jobNodeType"`
		MarshalTypeMultiNodeJobNodeConfigurationDetails
	}{
		"MULTI_NODE",
		(MarshalTypeMultiNodeJobNodeConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *MultiNodeJobNodeConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StartupOrder                         MultiNodeJobNodeConfigurationDetailsStartupOrderEnum `json:"startupOrder"`
		MaximumRuntimeInMinutes              *int64                                               `json:"maximumRuntimeInMinutes"`
		JobNetworkConfiguration              jobnetworkconfiguration                              `json:"jobNetworkConfiguration"`
		JobNodeGroupConfigurationDetailsList []JobNodeGroupConfigurationDetails                   `json:"jobNodeGroupConfigurationDetailsList"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StartupOrder = model.StartupOrder

	m.MaximumRuntimeInMinutes = model.MaximumRuntimeInMinutes

	nn, e = model.JobNetworkConfiguration.UnmarshalPolymorphicJSON(model.JobNetworkConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobNetworkConfiguration = nn.(JobNetworkConfiguration)
	} else {
		m.JobNetworkConfiguration = nil
	}

	m.JobNodeGroupConfigurationDetailsList = make([]JobNodeGroupConfigurationDetails, len(model.JobNodeGroupConfigurationDetailsList))
	copy(m.JobNodeGroupConfigurationDetailsList, model.JobNodeGroupConfigurationDetailsList)
	return
}

// MultiNodeJobNodeConfigurationDetailsStartupOrderEnum Enum with underlying type: string
type MultiNodeJobNodeConfigurationDetailsStartupOrderEnum string

// Set of constants representing the allowable values for MultiNodeJobNodeConfigurationDetailsStartupOrderEnum
const (
	MultiNodeJobNodeConfigurationDetailsStartupOrderOrder    MultiNodeJobNodeConfigurationDetailsStartupOrderEnum = "IN_ORDER"
	MultiNodeJobNodeConfigurationDetailsStartupOrderParallel MultiNodeJobNodeConfigurationDetailsStartupOrderEnum = "IN_PARALLEL"
)

var mappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnum = map[string]MultiNodeJobNodeConfigurationDetailsStartupOrderEnum{
	"IN_ORDER":    MultiNodeJobNodeConfigurationDetailsStartupOrderOrder,
	"IN_PARALLEL": MultiNodeJobNodeConfigurationDetailsStartupOrderParallel,
}

var mappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnumLowerCase = map[string]MultiNodeJobNodeConfigurationDetailsStartupOrderEnum{
	"in_order":    MultiNodeJobNodeConfigurationDetailsStartupOrderOrder,
	"in_parallel": MultiNodeJobNodeConfigurationDetailsStartupOrderParallel,
}

// GetMultiNodeJobNodeConfigurationDetailsStartupOrderEnumValues Enumerates the set of values for MultiNodeJobNodeConfigurationDetailsStartupOrderEnum
func GetMultiNodeJobNodeConfigurationDetailsStartupOrderEnumValues() []MultiNodeJobNodeConfigurationDetailsStartupOrderEnum {
	values := make([]MultiNodeJobNodeConfigurationDetailsStartupOrderEnum, 0)
	for _, v := range mappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetMultiNodeJobNodeConfigurationDetailsStartupOrderEnumStringValues Enumerates the set of values in String for MultiNodeJobNodeConfigurationDetailsStartupOrderEnum
func GetMultiNodeJobNodeConfigurationDetailsStartupOrderEnumStringValues() []string {
	return []string{
		"IN_ORDER",
		"IN_PARALLEL",
	}
}

// GetMappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnum(val string) (MultiNodeJobNodeConfigurationDetailsStartupOrderEnum, bool) {
	enum, ok := mappingMultiNodeJobNodeConfigurationDetailsStartupOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
