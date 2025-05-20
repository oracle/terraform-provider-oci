// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// KubeClusterEntitySourceEnum Enum with underlying type: string
type KubeClusterEntitySourceEnum string

// Set of constants representing the allowable values for KubeClusterEntitySourceEnum
const (
	KubeClusterEntitySourceOkeCluster KubeClusterEntitySourceEnum = "OKE_CLUSTER"
)

var mappingKubeClusterEntitySourceEnum = map[string]KubeClusterEntitySourceEnum{
	"OKE_CLUSTER": KubeClusterEntitySourceOkeCluster,
}

var mappingKubeClusterEntitySourceEnumLowerCase = map[string]KubeClusterEntitySourceEnum{
	"oke_cluster": KubeClusterEntitySourceOkeCluster,
}

// GetKubeClusterEntitySourceEnumValues Enumerates the set of values for KubeClusterEntitySourceEnum
func GetKubeClusterEntitySourceEnumValues() []KubeClusterEntitySourceEnum {
	values := make([]KubeClusterEntitySourceEnum, 0)
	for _, v := range mappingKubeClusterEntitySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetKubeClusterEntitySourceEnumStringValues Enumerates the set of values in String for KubeClusterEntitySourceEnum
func GetKubeClusterEntitySourceEnumStringValues() []string {
	return []string{
		"OKE_CLUSTER",
	}
}

// GetMappingKubeClusterEntitySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKubeClusterEntitySourceEnum(val string) (KubeClusterEntitySourceEnum, bool) {
	enum, ok := mappingKubeClusterEntitySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
