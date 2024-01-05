// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// DataSciencePrivateEndpointLifecycleStateEnum Enum with underlying type: string
type DataSciencePrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for DataSciencePrivateEndpointLifecycleStateEnum
const (
	DataSciencePrivateEndpointLifecycleStateCreating       DataSciencePrivateEndpointLifecycleStateEnum = "CREATING"
	DataSciencePrivateEndpointLifecycleStateActive         DataSciencePrivateEndpointLifecycleStateEnum = "ACTIVE"
	DataSciencePrivateEndpointLifecycleStateUpdating       DataSciencePrivateEndpointLifecycleStateEnum = "UPDATING"
	DataSciencePrivateEndpointLifecycleStateDeleting       DataSciencePrivateEndpointLifecycleStateEnum = "DELETING"
	DataSciencePrivateEndpointLifecycleStateDeleted        DataSciencePrivateEndpointLifecycleStateEnum = "DELETED"
	DataSciencePrivateEndpointLifecycleStateFailed         DataSciencePrivateEndpointLifecycleStateEnum = "FAILED"
	DataSciencePrivateEndpointLifecycleStateNeedsAttention DataSciencePrivateEndpointLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDataSciencePrivateEndpointLifecycleStateEnum = map[string]DataSciencePrivateEndpointLifecycleStateEnum{
	"CREATING":        DataSciencePrivateEndpointLifecycleStateCreating,
	"ACTIVE":          DataSciencePrivateEndpointLifecycleStateActive,
	"UPDATING":        DataSciencePrivateEndpointLifecycleStateUpdating,
	"DELETING":        DataSciencePrivateEndpointLifecycleStateDeleting,
	"DELETED":         DataSciencePrivateEndpointLifecycleStateDeleted,
	"FAILED":          DataSciencePrivateEndpointLifecycleStateFailed,
	"NEEDS_ATTENTION": DataSciencePrivateEndpointLifecycleStateNeedsAttention,
}

var mappingDataSciencePrivateEndpointLifecycleStateEnumLowerCase = map[string]DataSciencePrivateEndpointLifecycleStateEnum{
	"creating":        DataSciencePrivateEndpointLifecycleStateCreating,
	"active":          DataSciencePrivateEndpointLifecycleStateActive,
	"updating":        DataSciencePrivateEndpointLifecycleStateUpdating,
	"deleting":        DataSciencePrivateEndpointLifecycleStateDeleting,
	"deleted":         DataSciencePrivateEndpointLifecycleStateDeleted,
	"failed":          DataSciencePrivateEndpointLifecycleStateFailed,
	"needs_attention": DataSciencePrivateEndpointLifecycleStateNeedsAttention,
}

// GetDataSciencePrivateEndpointLifecycleStateEnumValues Enumerates the set of values for DataSciencePrivateEndpointLifecycleStateEnum
func GetDataSciencePrivateEndpointLifecycleStateEnumValues() []DataSciencePrivateEndpointLifecycleStateEnum {
	values := make([]DataSciencePrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingDataSciencePrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSciencePrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for DataSciencePrivateEndpointLifecycleStateEnum
func GetDataSciencePrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDataSciencePrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSciencePrivateEndpointLifecycleStateEnum(val string) (DataSciencePrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingDataSciencePrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
