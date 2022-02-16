// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BdsInstance Description of the cluster.
type BdsInstance struct {

	// The OCID of the Big Data Service resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the cluster.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the cluster.
	LifecycleState BdsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Boolean flag specifying whether or not the cluster is highly available (HA)
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be set up as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// Boolean flag specifying whether or not Cloud SQL should be configured.
	IsCloudSqlConfigured *bool `mandatory:"true" json:"isCloudSqlConfigured"`

	// The list of nodes in the cluster.
	Nodes []Node `mandatory:"true" json:"nodes"`

	// Number of nodes that forming the cluster
	NumberOfNodes *int `mandatory:"true" json:"numberOfNodes"`

	// Version of the Hadoop distribution.
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"false" json:"clusterVersion,omitempty"`

	NetworkConfig *NetworkConfig `mandatory:"false" json:"networkConfig"`

	ClusterDetails *ClusterDetails `mandatory:"false" json:"clusterDetails"`

	CloudSqlDetails *CloudSqlDetails `mandatory:"false" json:"cloudSqlDetails"`

	// The user who created the cluster.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BdsInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBdsInstanceClusterVersionEnum(string(m.ClusterVersion)); !ok && m.ClusterVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterVersion: %s. Supported values are: %s.", m.ClusterVersion, strings.Join(GetBdsInstanceClusterVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsInstanceLifecycleStateEnum Enum with underlying type: string
type BdsInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for BdsInstanceLifecycleStateEnum
const (
	BdsInstanceLifecycleStateCreating   BdsInstanceLifecycleStateEnum = "CREATING"
	BdsInstanceLifecycleStateActive     BdsInstanceLifecycleStateEnum = "ACTIVE"
	BdsInstanceLifecycleStateUpdating   BdsInstanceLifecycleStateEnum = "UPDATING"
	BdsInstanceLifecycleStateSuspending BdsInstanceLifecycleStateEnum = "SUSPENDING"
	BdsInstanceLifecycleStateSuspended  BdsInstanceLifecycleStateEnum = "SUSPENDED"
	BdsInstanceLifecycleStateResuming   BdsInstanceLifecycleStateEnum = "RESUMING"
	BdsInstanceLifecycleStateDeleting   BdsInstanceLifecycleStateEnum = "DELETING"
	BdsInstanceLifecycleStateDeleted    BdsInstanceLifecycleStateEnum = "DELETED"
	BdsInstanceLifecycleStateFailed     BdsInstanceLifecycleStateEnum = "FAILED"
)

var mappingBdsInstanceLifecycleStateEnum = map[string]BdsInstanceLifecycleStateEnum{
	"CREATING":   BdsInstanceLifecycleStateCreating,
	"ACTIVE":     BdsInstanceLifecycleStateActive,
	"UPDATING":   BdsInstanceLifecycleStateUpdating,
	"SUSPENDING": BdsInstanceLifecycleStateSuspending,
	"SUSPENDED":  BdsInstanceLifecycleStateSuspended,
	"RESUMING":   BdsInstanceLifecycleStateResuming,
	"DELETING":   BdsInstanceLifecycleStateDeleting,
	"DELETED":    BdsInstanceLifecycleStateDeleted,
	"FAILED":     BdsInstanceLifecycleStateFailed,
}

// GetBdsInstanceLifecycleStateEnumValues Enumerates the set of values for BdsInstanceLifecycleStateEnum
func GetBdsInstanceLifecycleStateEnumValues() []BdsInstanceLifecycleStateEnum {
	values := make([]BdsInstanceLifecycleStateEnum, 0)
	for _, v := range mappingBdsInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for BdsInstanceLifecycleStateEnum
func GetBdsInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"SUSPENDING",
		"SUSPENDED",
		"RESUMING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBdsInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceLifecycleStateEnum(val string) (BdsInstanceLifecycleStateEnum, bool) {
	mappingBdsInstanceLifecycleStateEnumIgnoreCase := make(map[string]BdsInstanceLifecycleStateEnum)
	for k, v := range mappingBdsInstanceLifecycleStateEnum {
		mappingBdsInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBdsInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// BdsInstanceClusterVersionEnum Enum with underlying type: string
type BdsInstanceClusterVersionEnum string

// Set of constants representing the allowable values for BdsInstanceClusterVersionEnum
const (
	BdsInstanceClusterVersionCdh5 BdsInstanceClusterVersionEnum = "CDH5"
	BdsInstanceClusterVersionCdh6 BdsInstanceClusterVersionEnum = "CDH6"
	BdsInstanceClusterVersionOdh1 BdsInstanceClusterVersionEnum = "ODH1"
)

var mappingBdsInstanceClusterVersionEnum = map[string]BdsInstanceClusterVersionEnum{
	"CDH5": BdsInstanceClusterVersionCdh5,
	"CDH6": BdsInstanceClusterVersionCdh6,
	"ODH1": BdsInstanceClusterVersionOdh1,
}

// GetBdsInstanceClusterVersionEnumValues Enumerates the set of values for BdsInstanceClusterVersionEnum
func GetBdsInstanceClusterVersionEnumValues() []BdsInstanceClusterVersionEnum {
	values := make([]BdsInstanceClusterVersionEnum, 0)
	for _, v := range mappingBdsInstanceClusterVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceClusterVersionEnumStringValues Enumerates the set of values in String for BdsInstanceClusterVersionEnum
func GetBdsInstanceClusterVersionEnumStringValues() []string {
	return []string{
		"CDH5",
		"CDH6",
		"ODH1",
	}
}

// GetMappingBdsInstanceClusterVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceClusterVersionEnum(val string) (BdsInstanceClusterVersionEnum, bool) {
	mappingBdsInstanceClusterVersionEnumIgnoreCase := make(map[string]BdsInstanceClusterVersionEnum)
	for k, v := range mappingBdsInstanceClusterVersionEnum {
		mappingBdsInstanceClusterVersionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBdsInstanceClusterVersionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
