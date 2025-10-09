// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClusterNamespaceStatus Represents the state of a ClusterNamespace within a specific OKE cluster
type ClusterNamespaceStatus struct {

	// The OCID of the OKE cluster where the namespace is provisioned
	ClusterId *string `mandatory:"false" json:"clusterId"`

	// The lifecycle state of ClusterNamespace resources in this cluster
	Status ClusterNamespaceStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Indicates whether new ClusterNamespace can be scheduled to Cluster
	SchedulingReadiness ClusterNamespaceStatusSchedulingReadinessEnum `mandatory:"false" json:"schedulingReadiness,omitempty"`
}

func (m ClusterNamespaceStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterNamespaceStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClusterNamespaceStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetClusterNamespaceStatusStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClusterNamespaceStatusSchedulingReadinessEnum(string(m.SchedulingReadiness)); !ok && m.SchedulingReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchedulingReadiness: %s. Supported values are: %s.", m.SchedulingReadiness, strings.Join(GetClusterNamespaceStatusSchedulingReadinessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterNamespaceStatusStatusEnum Enum with underlying type: string
type ClusterNamespaceStatusStatusEnum string

// Set of constants representing the allowable values for ClusterNamespaceStatusStatusEnum
const (
	ClusterNamespaceStatusStatusCreating ClusterNamespaceStatusStatusEnum = "CREATING"
	ClusterNamespaceStatusStatusUpdating ClusterNamespaceStatusStatusEnum = "UPDATING"
	ClusterNamespaceStatusStatusActive   ClusterNamespaceStatusStatusEnum = "ACTIVE"
	ClusterNamespaceStatusStatusDeleting ClusterNamespaceStatusStatusEnum = "DELETING"
	ClusterNamespaceStatusStatusDeleted  ClusterNamespaceStatusStatusEnum = "DELETED"
	ClusterNamespaceStatusStatusFailed   ClusterNamespaceStatusStatusEnum = "FAILED"
)

var mappingClusterNamespaceStatusStatusEnum = map[string]ClusterNamespaceStatusStatusEnum{
	"CREATING": ClusterNamespaceStatusStatusCreating,
	"UPDATING": ClusterNamespaceStatusStatusUpdating,
	"ACTIVE":   ClusterNamespaceStatusStatusActive,
	"DELETING": ClusterNamespaceStatusStatusDeleting,
	"DELETED":  ClusterNamespaceStatusStatusDeleted,
	"FAILED":   ClusterNamespaceStatusStatusFailed,
}

var mappingClusterNamespaceStatusStatusEnumLowerCase = map[string]ClusterNamespaceStatusStatusEnum{
	"creating": ClusterNamespaceStatusStatusCreating,
	"updating": ClusterNamespaceStatusStatusUpdating,
	"active":   ClusterNamespaceStatusStatusActive,
	"deleting": ClusterNamespaceStatusStatusDeleting,
	"deleted":  ClusterNamespaceStatusStatusDeleted,
	"failed":   ClusterNamespaceStatusStatusFailed,
}

// GetClusterNamespaceStatusStatusEnumValues Enumerates the set of values for ClusterNamespaceStatusStatusEnum
func GetClusterNamespaceStatusStatusEnumValues() []ClusterNamespaceStatusStatusEnum {
	values := make([]ClusterNamespaceStatusStatusEnum, 0)
	for _, v := range mappingClusterNamespaceStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceStatusStatusEnumStringValues Enumerates the set of values in String for ClusterNamespaceStatusStatusEnum
func GetClusterNamespaceStatusStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterNamespaceStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceStatusStatusEnum(val string) (ClusterNamespaceStatusStatusEnum, bool) {
	enum, ok := mappingClusterNamespaceStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ClusterNamespaceStatusSchedulingReadinessEnum Enum with underlying type: string
type ClusterNamespaceStatusSchedulingReadinessEnum string

// Set of constants representing the allowable values for ClusterNamespaceStatusSchedulingReadinessEnum
const (
	ClusterNamespaceStatusSchedulingReadinessReady         ClusterNamespaceStatusSchedulingReadinessEnum = "READY"
	ClusterNamespaceStatusSchedulingReadinessNotReady      ClusterNamespaceStatusSchedulingReadinessEnum = "NOT_READY"
	ClusterNamespaceStatusSchedulingReadinessNotApplicable ClusterNamespaceStatusSchedulingReadinessEnum = "NOT_APPLICABLE"
)

var mappingClusterNamespaceStatusSchedulingReadinessEnum = map[string]ClusterNamespaceStatusSchedulingReadinessEnum{
	"READY":          ClusterNamespaceStatusSchedulingReadinessReady,
	"NOT_READY":      ClusterNamespaceStatusSchedulingReadinessNotReady,
	"NOT_APPLICABLE": ClusterNamespaceStatusSchedulingReadinessNotApplicable,
}

var mappingClusterNamespaceStatusSchedulingReadinessEnumLowerCase = map[string]ClusterNamespaceStatusSchedulingReadinessEnum{
	"ready":          ClusterNamespaceStatusSchedulingReadinessReady,
	"not_ready":      ClusterNamespaceStatusSchedulingReadinessNotReady,
	"not_applicable": ClusterNamespaceStatusSchedulingReadinessNotApplicable,
}

// GetClusterNamespaceStatusSchedulingReadinessEnumValues Enumerates the set of values for ClusterNamespaceStatusSchedulingReadinessEnum
func GetClusterNamespaceStatusSchedulingReadinessEnumValues() []ClusterNamespaceStatusSchedulingReadinessEnum {
	values := make([]ClusterNamespaceStatusSchedulingReadinessEnum, 0)
	for _, v := range mappingClusterNamespaceStatusSchedulingReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceStatusSchedulingReadinessEnumStringValues Enumerates the set of values in String for ClusterNamespaceStatusSchedulingReadinessEnum
func GetClusterNamespaceStatusSchedulingReadinessEnumStringValues() []string {
	return []string{
		"READY",
		"NOT_READY",
		"NOT_APPLICABLE",
	}
}

// GetMappingClusterNamespaceStatusSchedulingReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceStatusSchedulingReadinessEnum(val string) (ClusterNamespaceStatusSchedulingReadinessEnum, bool) {
	enum, ok := mappingClusterNamespaceStatusSchedulingReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
