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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentKubernetesScrapeTarget Monitoring scrape object.
type UnifiedAgentKubernetesScrapeTarget struct {

	// Type of resource to scrape metrics.
	ResourceType UnifiedAgentKubernetesScrapeTargetResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// K8s namespace of the resource.
	K8sNamespace *string `mandatory:"true" json:"k8sNamespace"`

	// Name of the service prepended to the endpoints.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Resource group in OCI monitoring.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`
}

func (m UnifiedAgentKubernetesScrapeTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentKubernetesScrapeTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetUnifiedAgentKubernetesScrapeTargetResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnifiedAgentKubernetesScrapeTargetResourceTypeEnum Enum with underlying type: string
type UnifiedAgentKubernetesScrapeTargetResourceTypeEnum string

// Set of constants representing the allowable values for UnifiedAgentKubernetesScrapeTargetResourceTypeEnum
const (
	UnifiedAgentKubernetesScrapeTargetResourceTypePods      UnifiedAgentKubernetesScrapeTargetResourceTypeEnum = "PODS"
	UnifiedAgentKubernetesScrapeTargetResourceTypeEndpoints UnifiedAgentKubernetesScrapeTargetResourceTypeEnum = "ENDPOINTS"
	UnifiedAgentKubernetesScrapeTargetResourceTypeNodes     UnifiedAgentKubernetesScrapeTargetResourceTypeEnum = "NODES"
	UnifiedAgentKubernetesScrapeTargetResourceTypeServices  UnifiedAgentKubernetesScrapeTargetResourceTypeEnum = "SERVICES"
)

var mappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnum = map[string]UnifiedAgentKubernetesScrapeTargetResourceTypeEnum{
	"PODS":      UnifiedAgentKubernetesScrapeTargetResourceTypePods,
	"ENDPOINTS": UnifiedAgentKubernetesScrapeTargetResourceTypeEndpoints,
	"NODES":     UnifiedAgentKubernetesScrapeTargetResourceTypeNodes,
	"SERVICES":  UnifiedAgentKubernetesScrapeTargetResourceTypeServices,
}

var mappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnumLowerCase = map[string]UnifiedAgentKubernetesScrapeTargetResourceTypeEnum{
	"pods":      UnifiedAgentKubernetesScrapeTargetResourceTypePods,
	"endpoints": UnifiedAgentKubernetesScrapeTargetResourceTypeEndpoints,
	"nodes":     UnifiedAgentKubernetesScrapeTargetResourceTypeNodes,
	"services":  UnifiedAgentKubernetesScrapeTargetResourceTypeServices,
}

// GetUnifiedAgentKubernetesScrapeTargetResourceTypeEnumValues Enumerates the set of values for UnifiedAgentKubernetesScrapeTargetResourceTypeEnum
func GetUnifiedAgentKubernetesScrapeTargetResourceTypeEnumValues() []UnifiedAgentKubernetesScrapeTargetResourceTypeEnum {
	values := make([]UnifiedAgentKubernetesScrapeTargetResourceTypeEnum, 0)
	for _, v := range mappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentKubernetesScrapeTargetResourceTypeEnumStringValues Enumerates the set of values in String for UnifiedAgentKubernetesScrapeTargetResourceTypeEnum
func GetUnifiedAgentKubernetesScrapeTargetResourceTypeEnumStringValues() []string {
	return []string{
		"PODS",
		"ENDPOINTS",
		"NODES",
		"SERVICES",
	}
}

// GetMappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnum(val string) (UnifiedAgentKubernetesScrapeTargetResourceTypeEnum, bool) {
	enum, ok := mappingUnifiedAgentKubernetesScrapeTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
