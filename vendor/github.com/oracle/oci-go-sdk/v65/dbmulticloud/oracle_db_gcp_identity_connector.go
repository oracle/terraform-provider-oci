// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// <b>Microsoft Azure</b>:<br>
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
// <b>Google Cloud</b>:<br>
// 1. Oracle Google Cloud Connector Resource: This is for installing Google Identity in ExaCS VM Cluster.<br>
// 2. Discover Google Key-Rings and Keys Resource: This is for to discover Google Key-Rings.<br>
// 3. Google Key-Rings Resource: This is for to maintain Google Key-Rings in Oracle Cloud.<br>
// 4. Google Key Resource: This is for to maintain Google Key in Oracle Cloud for a Google Key-Ring.<br>
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbGcpIdentityConnector Oracle DB GCP Connector Details, this resource is for to creates GCP Identity Connector.
type OracleDbGcpIdentityConnector struct {

	// The ID of the Oracle DB GCP Identity Connector resource.
	Id *string `mandatory:"true" json:"id"`

	// The ID of the compartment that contains Oracle DB GCP Identity  Configuration resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OIDC token issuer Url
	IssuerUrl *string `mandatory:"true" json:"issuerUrl"`

	// Project id of the customer project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The ID of the GCP VM Cluster resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// GCP Location.
	GcpLocation *string `mandatory:"true" json:"gcpLocation"`

	// The ID of the cloud GCP Workload Identity Pool.
	GcpWorkloadIdentityPoolId *string `mandatory:"true" json:"gcpWorkloadIdentityPoolId"`

	// The ID of the GCP Workload Identity Provider.
	GcpWorkloadIdentityProviderId *string `mandatory:"true" json:"gcpWorkloadIdentityProviderId"`

	// The ID of the GCP Resource Service Agent.
	GcpResourceServiceAgentId *string `mandatory:"true" json:"gcpResourceServiceAgentId"`

	// Oracle DB GCP Identity Connector resource name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of All VMs where GCP Identity Connector is configured for this VMCluster.
	GcpNodes []GcpNodes `mandatory:"false" json:"gcpNodes"`

	// The current Connectivity status of GCP Identity Connector Resource.
	GcpIdentityConnectivityStatus OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum `mandatory:"false" json:"gcpIdentityConnectivityStatus,omitempty"`

	// The current lifecycle state of the GCP Identity Connector Resource.
	LifecycleState OracleDbGcpIdentityConnectorLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Oracle DB GCP Identity Connector Resource was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Oracle DB GCP Identity Connector Resource was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OracleDbGcpIdentityConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbGcpIdentityConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum(string(m.GcpIdentityConnectivityStatus)); !ok && m.GcpIdentityConnectivityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GcpIdentityConnectivityStatus: %s. Supported values are: %s.", m.GcpIdentityConnectivityStatus, strings.Join(GetOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleDbGcpIdentityConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbGcpIdentityConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum Enum with underlying type: string
type OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum string

// Set of constants representing the allowable values for OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum
const (
	OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusConnected          OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum = "CONNECTED"
	OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusDisconnected       OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum = "DISCONNECTED"
	OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusPartiallyConnected OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum = "PARTIALLY_CONNECTED"
	OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusUnknown            OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum = "UNKNOWN"
)

var mappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum = map[string]OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum{
	"CONNECTED":           OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusConnected,
	"DISCONNECTED":        OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusDisconnected,
	"PARTIALLY_CONNECTED": OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusPartiallyConnected,
	"UNKNOWN":             OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusUnknown,
}

var mappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumLowerCase = map[string]OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum{
	"connected":           OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusConnected,
	"disconnected":        OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusDisconnected,
	"partially_connected": OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusPartiallyConnected,
	"unknown":             OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusUnknown,
}

// GetOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumValues Enumerates the set of values for OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum
func GetOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumValues() []OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum {
	values := make([]OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum, 0)
	for _, v := range mappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumStringValues Enumerates the set of values in String for OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum
func GetOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"PARTIALLY_CONNECTED",
		"UNKNOWN",
	}
}

// GetMappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum(val string) (OracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnum, bool) {
	enum, ok := mappingOracleDbGcpIdentityConnectorGcpIdentityConnectivityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleDbGcpIdentityConnectorLifecycleStateEnum Enum with underlying type: string
type OracleDbGcpIdentityConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbGcpIdentityConnectorLifecycleStateEnum
const (
	OracleDbGcpIdentityConnectorLifecycleStateCreating OracleDbGcpIdentityConnectorLifecycleStateEnum = "CREATING"
	OracleDbGcpIdentityConnectorLifecycleStateActive   OracleDbGcpIdentityConnectorLifecycleStateEnum = "ACTIVE"
	OracleDbGcpIdentityConnectorLifecycleStateUpdating OracleDbGcpIdentityConnectorLifecycleStateEnum = "UPDATING"
	OracleDbGcpIdentityConnectorLifecycleStateDeleting OracleDbGcpIdentityConnectorLifecycleStateEnum = "DELETING"
	OracleDbGcpIdentityConnectorLifecycleStateDeleted  OracleDbGcpIdentityConnectorLifecycleStateEnum = "DELETED"
	OracleDbGcpIdentityConnectorLifecycleStateFailed   OracleDbGcpIdentityConnectorLifecycleStateEnum = "FAILED"
)

var mappingOracleDbGcpIdentityConnectorLifecycleStateEnum = map[string]OracleDbGcpIdentityConnectorLifecycleStateEnum{
	"CREATING": OracleDbGcpIdentityConnectorLifecycleStateCreating,
	"ACTIVE":   OracleDbGcpIdentityConnectorLifecycleStateActive,
	"UPDATING": OracleDbGcpIdentityConnectorLifecycleStateUpdating,
	"DELETING": OracleDbGcpIdentityConnectorLifecycleStateDeleting,
	"DELETED":  OracleDbGcpIdentityConnectorLifecycleStateDeleted,
	"FAILED":   OracleDbGcpIdentityConnectorLifecycleStateFailed,
}

var mappingOracleDbGcpIdentityConnectorLifecycleStateEnumLowerCase = map[string]OracleDbGcpIdentityConnectorLifecycleStateEnum{
	"creating": OracleDbGcpIdentityConnectorLifecycleStateCreating,
	"active":   OracleDbGcpIdentityConnectorLifecycleStateActive,
	"updating": OracleDbGcpIdentityConnectorLifecycleStateUpdating,
	"deleting": OracleDbGcpIdentityConnectorLifecycleStateDeleting,
	"deleted":  OracleDbGcpIdentityConnectorLifecycleStateDeleted,
	"failed":   OracleDbGcpIdentityConnectorLifecycleStateFailed,
}

// GetOracleDbGcpIdentityConnectorLifecycleStateEnumValues Enumerates the set of values for OracleDbGcpIdentityConnectorLifecycleStateEnum
func GetOracleDbGcpIdentityConnectorLifecycleStateEnumValues() []OracleDbGcpIdentityConnectorLifecycleStateEnum {
	values := make([]OracleDbGcpIdentityConnectorLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbGcpIdentityConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbGcpIdentityConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbGcpIdentityConnectorLifecycleStateEnum
func GetOracleDbGcpIdentityConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbGcpIdentityConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbGcpIdentityConnectorLifecycleStateEnum(val string) (OracleDbGcpIdentityConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbGcpIdentityConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
