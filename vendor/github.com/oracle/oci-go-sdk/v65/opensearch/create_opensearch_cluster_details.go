// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateOpensearchClusterDetails The configuration details for a new OpenSearch cluster.
type CreateOpensearchClusterDetails struct {

	// The name of the cluster. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment to create the cluster in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The version of the software the cluster is running.
	SoftwareVersion *string `mandatory:"true" json:"softwareVersion"`

	// The number of master nodes to configure for the cluster.
	MasterNodeCount *int `mandatory:"true" json:"masterNodeCount"`

	// The instance type for the cluster's master nodes.
	MasterNodeHostType MasterNodeHostTypeEnum `mandatory:"true" json:"masterNodeHostType"`

	// The number of OCPUs to configure for the cluser's master nodes.
	MasterNodeHostOcpuCount *int `mandatory:"true" json:"masterNodeHostOcpuCount"`

	// The amount of memory in GB, to configure per node for the cluster's master nodes.
	MasterNodeHostMemoryGB *int `mandatory:"true" json:"masterNodeHostMemoryGB"`

	// The number of data nodes to configure for the cluster.
	DataNodeCount *int `mandatory:"true" json:"dataNodeCount"`

	// TThe instance type for the cluster's data nodes.
	DataNodeHostType DataNodeHostTypeEnum `mandatory:"true" json:"dataNodeHostType"`

	// The number of OCPUs to configure for the cluster's data nodes.
	DataNodeHostOcpuCount *int `mandatory:"true" json:"dataNodeHostOcpuCount"`

	// The amount of memory in GB, to configure per node for the cluster's data nodes.
	DataNodeHostMemoryGB *int `mandatory:"true" json:"dataNodeHostMemoryGB"`

	// The amount of storage in GB, to configure per node for the cluster's data nodes.
	DataNodeStorageGB *int `mandatory:"true" json:"dataNodeStorageGB"`

	// The number of OpenSearch Dashboard nodes to configure for the cluster.
	OpendashboardNodeCount *int `mandatory:"true" json:"opendashboardNodeCount"`

	// The number of OCPUs to configure for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostOcpuCount *int `mandatory:"true" json:"opendashboardNodeHostOcpuCount"`

	// The amount of memory in GB, to configure for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostMemoryGB *int `mandatory:"true" json:"opendashboardNodeHostMemoryGB"`

	// The OCID of the cluster's VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the cluster's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID for the compartment where the cluster's VCN is located.
	VcnCompartmentId *string `mandatory:"true" json:"vcnCompartmentId"`

	// The OCID for the compartment where the cluster's subnet is located.
	SubnetCompartmentId *string `mandatory:"true" json:"subnetCompartmentId"`

	// The bare metal shape for the cluster's master nodes.
	MasterNodeHostBareMetalShape *string `mandatory:"false" json:"masterNodeHostBareMetalShape"`

	// The node shape for the cluster's master nodes.
	MasterNodeHostShape *string `mandatory:"false" json:"masterNodeHostShape"`

	// The bare metal shape for the cluster's data nodes.
	DataNodeHostBareMetalShape *string `mandatory:"false" json:"dataNodeHostBareMetalShape"`

	// The node shape for the cluster's data nodes.
	DataNodeHostShape *string `mandatory:"false" json:"dataNodeHostShape"`

	// The node shape for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostShape *string `mandatory:"false" json:"opendashboardNodeHostShape"`

	// The number of search nodes configured for the cluster.
	SearchNodeCount *int `mandatory:"false" json:"searchNodeCount"`

	// The instance type for the cluster's search nodes.
	SearchNodeHostType SearchNodeHostTypeEnum `mandatory:"false" json:"searchNodeHostType,omitempty"`

	// The node shape for the cluster's search nodes.
	SearchNodeHostShape *string `mandatory:"false" json:"searchNodeHostShape"`

	// The number of OCPUs configured for the cluster's search nodes.
	SearchNodeHostOcpuCount *int `mandatory:"false" json:"searchNodeHostOcpuCount"`

	// The amount of memory in GB, for the cluster's search nodes.
	SearchNodeHostMemoryGB *int `mandatory:"false" json:"searchNodeHostMemoryGB"`

	// The amount of storage in GB, to configure per node for the cluster's search nodes.
	SearchNodeStorageGB *int `mandatory:"false" json:"searchNodeStorageGB"`

	// The security mode of the cluster.
	SecurityMode SecurityModeEnum `mandatory:"false" json:"securityMode,omitempty"`

	// The name of the master user that are used to manage security config
	SecurityMasterUserName *string `mandatory:"false" json:"securityMasterUserName"`

	// The password hash of the master user that are used to manage security config
	SecurityMasterUserPasswordHash *string `mandatory:"false" json:"securityMasterUserPasswordHash"`

	SecuritySamlConfig *SecuritySamlConfig `mandatory:"false" json:"securitySamlConfig"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	// The OCID of the NSG where the private endpoint vnic will be attached.
	NsgId *string `mandatory:"false" json:"nsgId"`

	// The customer IP addresses of the endpoint in customer VCN
	ReverseConnectionEndpointCustomerIps []string `mandatory:"false" json:"reverseConnectionEndpointCustomerIps"`

	// List of inbound clusters that will be queried using cross cluster search
	InboundClusterIds []string `mandatory:"false" json:"inboundClusterIds"`

	OutboundClusterConfig *OutboundClusterConfig `mandatory:"false" json:"outboundClusterConfig"`

	MaintenanceDetails *CreateMaintenanceDetails `mandatory:"false" json:"maintenanceDetails"`

	CertificateConfig *CertificateConfig `mandatory:"false" json:"certificateConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`
}

func (m CreateOpensearchClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOpensearchClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMasterNodeHostTypeEnum(string(m.MasterNodeHostType)); !ok && m.MasterNodeHostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MasterNodeHostType: %s. Supported values are: %s.", m.MasterNodeHostType, strings.Join(GetMasterNodeHostTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDataNodeHostTypeEnum(string(m.DataNodeHostType)); !ok && m.DataNodeHostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataNodeHostType: %s. Supported values are: %s.", m.DataNodeHostType, strings.Join(GetDataNodeHostTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSearchNodeHostTypeEnum(string(m.SearchNodeHostType)); !ok && m.SearchNodeHostType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SearchNodeHostType: %s. Supported values are: %s.", m.SearchNodeHostType, strings.Join(GetSearchNodeHostTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityModeEnum(string(m.SecurityMode)); !ok && m.SecurityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityMode: %s. Supported values are: %s.", m.SecurityMode, strings.Join(GetSecurityModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
