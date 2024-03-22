// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentMonitoringFilter Monitoring filter object.
type UnifiedAgentMonitoringFilter interface {

	// Unique name for the filter.
	GetName() *string
}

type unifiedagentmonitoringfilter struct {
	JsonData   []byte
	Name       *string `mandatory:"true" json:"name"`
	FilterType string  `json:"filterType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentmonitoringfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentmonitoringfilter unifiedagentmonitoringfilter
	s := struct {
		Model Unmarshalerunifiedagentmonitoringfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.FilterType = s.Model.FilterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentmonitoringfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FilterType {
	case "KUBERNETES_FILTER":
		mm := UnifiedAgentKubernetesFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "URL_FILTER":
		mm := UnifiedAgentUrlFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnifiedAgentMonitoringFilter: %s.", m.FilterType)
		return *m, nil
	}
}

// GetName returns Name
func (m unifiedagentmonitoringfilter) GetName() *string {
	return m.Name
}

func (m unifiedagentmonitoringfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m unifiedagentmonitoringfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAgentMonitoringFilterFilterTypeEnum Enum with underlying type: string
type UnifiedAgentMonitoringFilterFilterTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentMonitoringFilterFilterTypeEnum
const (
	UnifiedAgentMonitoringFilterFilterTypeKubernetesFilter UnifiedAgentMonitoringFilterFilterTypeEnum = "KUBERNETES_FILTER"
	UnifiedAgentMonitoringFilterFilterTypeUrlFilter        UnifiedAgentMonitoringFilterFilterTypeEnum = "URL_FILTER"
)

var mappingUnifiedAgentMonitoringFilterFilterTypeEnum = map[string]UnifiedAgentMonitoringFilterFilterTypeEnum{
	"KUBERNETES_FILTER": UnifiedAgentMonitoringFilterFilterTypeKubernetesFilter,
	"URL_FILTER":        UnifiedAgentMonitoringFilterFilterTypeUrlFilter,
}

var mappingUnifiedAgentMonitoringFilterFilterTypeEnumLowerCase = map[string]UnifiedAgentMonitoringFilterFilterTypeEnum{
	"kubernetes_filter": UnifiedAgentMonitoringFilterFilterTypeKubernetesFilter,
	"url_filter":        UnifiedAgentMonitoringFilterFilterTypeUrlFilter,
}

// GetUnifiedAgentMonitoringFilterFilterTypeEnumValues Enumerates the set of values for UnifiedAgentMonitoringFilterFilterTypeEnum
func GetUnifiedAgentMonitoringFilterFilterTypeEnumValues() []UnifiedAgentMonitoringFilterFilterTypeEnum {
	values := make([]UnifiedAgentMonitoringFilterFilterTypeEnum, 0)
	for _, v := range mappingUnifiedAgentMonitoringFilterFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentMonitoringFilterFilterTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentMonitoringFilterFilterTypeEnum
func GetUnifiedAgentMonitoringFilterFilterTypeEnumStringValues() []string {
	return []string{
		"KUBERNETES_FILTER",
		"URL_FILTER",
	}
}

// GetMappingUnifiedAgentMonitoringFilterFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentMonitoringFilterFilterTypeEnum(val string) (UnifiedAgentMonitoringFilterFilterTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentMonitoringFilterFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
