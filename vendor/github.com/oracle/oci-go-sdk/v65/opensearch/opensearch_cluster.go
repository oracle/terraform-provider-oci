// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpensearchCluster An OpenSearch cluster resource. An OpenSearch cluster is set of instances that provide OpenSearch functionality in OCI Search Service with OpenSearch.
// For more information, see About Search Service with OpenSearch (https://docs.cloud.oracle.com/iaas/Content/search-opensearch/Concepts/ociopensearch.htm).
type OpensearchCluster struct {

	// The OCID of the cluster.
	Id *string `mandatory:"true" json:"id"`

	// The name of the cluster. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment where the cluster is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the cluster.
	LifecycleState OpensearchClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The amount of time in milliseconds since the cluster was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The software version the cluster is running.
	SoftwareVersion *string `mandatory:"true" json:"softwareVersion"`

	// The size in GB of the cluster's total storage.
	TotalStorageGB *int `mandatory:"true" json:"totalStorageGB"`

	// The fully qualified domain name (FQDN) for the cluster's API endpoint.
	OpensearchFqdn *string `mandatory:"true" json:"opensearchFqdn"`

	// The cluster's private IP address.
	OpensearchPrivateIp *string `mandatory:"true" json:"opensearchPrivateIp"`

	// The fully qualified domain name (FQDN) for the cluster's OpenSearch Dashboard API endpoint.
	OpendashboardFqdn *string `mandatory:"true" json:"opendashboardFqdn"`

	// The private IP address for the cluster's OpenSearch Dashboard.
	OpendashboardPrivateIp *string `mandatory:"true" json:"opendashboardPrivateIp"`

	// The number of master nodes configured for the cluster.
	MasterNodeCount *int `mandatory:"true" json:"masterNodeCount"`

	// The instance type for the cluster's master nodes.
	MasterNodeHostType MasterNodeHostTypeEnum `mandatory:"true" json:"masterNodeHostType"`

	// The number of OCPUs configured for cluster's master nodes.
	MasterNodeHostOcpuCount *int `mandatory:"true" json:"masterNodeHostOcpuCount"`

	// The amount of memory in GB, for the cluster's master nodes.
	MasterNodeHostMemoryGB *int `mandatory:"true" json:"masterNodeHostMemoryGB"`

	// The number of data nodes configured for the cluster.
	DataNodeCount *int `mandatory:"true" json:"dataNodeCount"`

	// The instance type for the cluster's data nodes.
	DataNodeHostType DataNodeHostTypeEnum `mandatory:"true" json:"dataNodeHostType"`

	// The number of OCPUs configured for the cluster's data nodes.
	DataNodeHostOcpuCount *int `mandatory:"true" json:"dataNodeHostOcpuCount"`

	// The amount of memory in GB, for the cluster's data nodes.
	DataNodeHostMemoryGB *int `mandatory:"true" json:"dataNodeHostMemoryGB"`

	// The amount of storage in GB, to configure per node for the cluster's data nodes.
	DataNodeStorageGB *int `mandatory:"true" json:"dataNodeStorageGB"`

	// The number of OpenSearch Dashboard nodes configured for the cluster.
	OpendashboardNodeCount *int `mandatory:"true" json:"opendashboardNodeCount"`

	// The amount of memory in GB, for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostOcpuCount *int `mandatory:"true" json:"opendashboardNodeHostOcpuCount"`

	// The amount of memory in GB, for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostMemoryGB *int `mandatory:"true" json:"opendashboardNodeHostMemoryGB"`

	// The OCID of the cluster's VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the cluster's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID for the compartment where the cluster's VCN is located.
	VcnCompartmentId *string `mandatory:"true" json:"vcnCompartmentId"`

	// The OCID for the compartment where the cluster's subnet is located.
	SubnetCompartmentId *string `mandatory:"true" json:"subnetCompartmentId"`

	// The availability domains to distribute the cluser nodes across.
	AvailabilityDomains []string `mandatory:"true" json:"availabilityDomains"`

	// The amount of time in milliseconds since the cluster was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The amount of time in milliseconds since the cluster was updated.
	TimeDeleted *common.SDKTime `mandatory:"false" json:"timeDeleted"`

	// Additional information about the current lifecycle state of the cluster.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The bare metal shape for the cluster's master nodes.
	MasterNodeHostBareMetalShape *string `mandatory:"false" json:"masterNodeHostBareMetalShape"`

	// The bare metal shape for the cluster's data nodes.
	DataNodeHostBareMetalShape *string `mandatory:"false" json:"dataNodeHostBareMetalShape"`

	// The fully qualified domain name (FQDN) for the cluster's API endpoint.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The security mode of the cluster.
	SecurityMode SecurityModeEnum `mandatory:"false" json:"securityMode,omitempty"`

	// The name of the master user that are used to manage security config
	SecurityMasterUserName *string `mandatory:"false" json:"securityMasterUserName"`

	// The password hash of the master user that are used to manage security config
	SecurityMasterUserPasswordHash *string `mandatory:"false" json:"securityMasterUserPasswordHash"`
}

func (m OpensearchCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMasterNodeHostTypeEnum(string(m.MasterNodeHostType)); !ok && m.MasterNodeHostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MasterNodeHostType: %s. Supported values are: %s.", m.MasterNodeHostType, strings.Join(GetMasterNodeHostTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataNodeHostTypeEnum(string(m.DataNodeHostType)); !ok && m.DataNodeHostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataNodeHostType: %s. Supported values are: %s.", m.DataNodeHostType, strings.Join(GetDataNodeHostTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSecurityModeEnum(string(m.SecurityMode)); !ok && m.SecurityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityMode: %s. Supported values are: %s.", m.SecurityMode, strings.Join(GetSecurityModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpensearchClusterLifecycleStateEnum Enum with underlying type: string
type OpensearchClusterLifecycleStateEnum string

// Set of constants representing the allowable values for OpensearchClusterLifecycleStateEnum
const (
	OpensearchClusterLifecycleStateActive   OpensearchClusterLifecycleStateEnum = "ACTIVE"
	OpensearchClusterLifecycleStateCreating OpensearchClusterLifecycleStateEnum = "CREATING"
	OpensearchClusterLifecycleStateUpdating OpensearchClusterLifecycleStateEnum = "UPDATING"
	OpensearchClusterLifecycleStateDeleting OpensearchClusterLifecycleStateEnum = "DELETING"
	OpensearchClusterLifecycleStateDeleted  OpensearchClusterLifecycleStateEnum = "DELETED"
	OpensearchClusterLifecycleStateFailed   OpensearchClusterLifecycleStateEnum = "FAILED"
)

var mappingOpensearchClusterLifecycleStateEnum = map[string]OpensearchClusterLifecycleStateEnum{
	"ACTIVE":   OpensearchClusterLifecycleStateActive,
	"CREATING": OpensearchClusterLifecycleStateCreating,
	"UPDATING": OpensearchClusterLifecycleStateUpdating,
	"DELETING": OpensearchClusterLifecycleStateDeleting,
	"DELETED":  OpensearchClusterLifecycleStateDeleted,
	"FAILED":   OpensearchClusterLifecycleStateFailed,
}

var mappingOpensearchClusterLifecycleStateEnumLowerCase = map[string]OpensearchClusterLifecycleStateEnum{
	"active":   OpensearchClusterLifecycleStateActive,
	"creating": OpensearchClusterLifecycleStateCreating,
	"updating": OpensearchClusterLifecycleStateUpdating,
	"deleting": OpensearchClusterLifecycleStateDeleting,
	"deleted":  OpensearchClusterLifecycleStateDeleted,
	"failed":   OpensearchClusterLifecycleStateFailed,
}

// GetOpensearchClusterLifecycleStateEnumValues Enumerates the set of values for OpensearchClusterLifecycleStateEnum
func GetOpensearchClusterLifecycleStateEnumValues() []OpensearchClusterLifecycleStateEnum {
	values := make([]OpensearchClusterLifecycleStateEnum, 0)
	for _, v := range mappingOpensearchClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpensearchClusterLifecycleStateEnumStringValues Enumerates the set of values in String for OpensearchClusterLifecycleStateEnum
func GetOpensearchClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOpensearchClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpensearchClusterLifecycleStateEnum(val string) (OpensearchClusterLifecycleStateEnum, bool) {
	enum, ok := mappingOpensearchClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
