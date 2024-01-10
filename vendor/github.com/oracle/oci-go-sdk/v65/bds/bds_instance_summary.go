// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BdsInstanceSummary Summary details of the Big Data Service cluster.
type BdsInstanceSummary struct {

	// The OCID of the Big Data Service resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the cluster.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the cluster.
	LifecycleState BdsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The number of nodes that form the cluster.
	NumberOfNodes *int `mandatory:"true" json:"numberOfNodes"`

	// Boolean flag specifying whether or not the cluster is highly available(HA).
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be set up as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// Boolean flag specifying whether Cloud SQL is configured or not.
	IsCloudSqlConfigured *bool `mandatory:"true" json:"isCloudSqlConfigured"`

	// Boolean flag specifying whether Kafka is configured or not.
	IsKafkaConfigured *bool `mandatory:"true" json:"isKafkaConfigured"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Number of nodes that require a maintenance reboot
	NumberOfNodesRequiringMaintenanceReboot *int `mandatory:"false" json:"numberOfNodesRequiringMaintenanceReboot"`

	// Version of the Hadoop distribution.
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"false" json:"clusterVersion,omitempty"`

	// Profile of the Big Data Service cluster.
	ClusterProfile BdsInstanceClusterProfileEnum `mandatory:"false" json:"clusterProfile,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BdsInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBdsInstanceClusterVersionEnum(string(m.ClusterVersion)); !ok && m.ClusterVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterVersion: %s. Supported values are: %s.", m.ClusterVersion, strings.Join(GetBdsInstanceClusterVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBdsInstanceClusterProfileEnum(string(m.ClusterProfile)); !ok && m.ClusterProfile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterProfile: %s. Supported values are: %s.", m.ClusterProfile, strings.Join(GetBdsInstanceClusterProfileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
