// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// DataSourceTypesEnum Enum with underlying type: string
type DataSourceTypesEnum string

// Set of constants representing the allowable values for DataSourceTypesEnum
const (
	DataSourceTypesKubernetesCluster DataSourceTypesEnum = "KUBERNETES_CLUSTER"
	DataSourceTypesPrometheusEmitter DataSourceTypesEnum = "PROMETHEUS_EMITTER"
)

var mappingDataSourceTypesEnum = map[string]DataSourceTypesEnum{
	"KUBERNETES_CLUSTER": DataSourceTypesKubernetesCluster,
	"PROMETHEUS_EMITTER": DataSourceTypesPrometheusEmitter,
}

var mappingDataSourceTypesEnumLowerCase = map[string]DataSourceTypesEnum{
	"kubernetes_cluster": DataSourceTypesKubernetesCluster,
	"prometheus_emitter": DataSourceTypesPrometheusEmitter,
}

// GetDataSourceTypesEnumValues Enumerates the set of values for DataSourceTypesEnum
func GetDataSourceTypesEnumValues() []DataSourceTypesEnum {
	values := make([]DataSourceTypesEnum, 0)
	for _, v := range mappingDataSourceTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSourceTypesEnumStringValues Enumerates the set of values in String for DataSourceTypesEnum
func GetDataSourceTypesEnumStringValues() []string {
	return []string{
		"KUBERNETES_CLUSTER",
		"PROMETHEUS_EMITTER",
	}
}

// GetMappingDataSourceTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSourceTypesEnum(val string) (DataSourceTypesEnum, bool) {
	enum, ok := mappingDataSourceTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
