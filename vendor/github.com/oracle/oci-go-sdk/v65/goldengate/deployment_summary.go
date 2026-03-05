// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeploymentSummary Summary of the Deployment.
type DeploymentSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint.
	// The subnet must be a private subnet. For backward compatibility, public subnets are allowed until May 31 2025,
	// after which the private subnet will be enforced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// The deployment category defines the broad separation of the deployment type into four categories.
	// Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS', 'DATA_TRANSFORMS' and 'DATA_VERIFICATION'.
	Category DeploymentCategoryEnum `mandatory:"true" json:"category"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states for a Deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Possible GGS lifecycle sub-states.
	LifecycleSubState LifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy used to host the public load balancer of the deployment.
	// Rules:
	// - Create: Mandatory when isPublic is true. Must be a public, regional subnet in the same VCN as subnetId.
	// - Update:
	//   - For public deployments, this property must be present and is immutable once set (cannot be changed to a different subnet).
	//   - Legacy exception: a public deployment created without this property may continue to be updated without providing it; once set, it becomes immutable.
	// Validation:
	// - Must reference a public subnet.
	// - Must be a regional subnet.
	// - Must be in the same VCN as subnetId.
	LoadBalancerSubnetId *string `mandatory:"false" json:"loadBalancerSubnetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the loadbalancer in the customer's subnet.
	// The loadbalancer of the public deployment created in the customer subnet.
	LoadBalancerId *string `mandatory:"false" json:"loadBalancerId"`

	// Flag to allow to configure the 'Bring Your Own License' (BYOL) license type CPU limit.
	// If enabled, the exact number of CPUs must be provided via byolCpuCoreCountLimit.
	IsByolCpuCoreCountLimitEnabled *bool `mandatory:"false" json:"isByolCpuCoreCountLimitEnabled"`

	// The maximum number of CPUs allowed with a 'Bring Your Own License' (BYOL) license type.
	// Any CPU usage above this limit is considered as License Included and billed.
	ByolCpuCoreCountLimit *int `mandatory:"false" json:"byolCpuCoreCountLimit"`

	// Specifies whether the deployment is used in a production or development/testing environment.
	EnvironmentType EnvironmentTypeEnum `mandatory:"false" json:"environmentType,omitempty"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// The public IP address representing the access point for the Deployment.
	PublicIpAddress *string `mandatory:"false" json:"publicIpAddress"`

	// The private IP address in the customer's VCN representing the access point for the
	// associated endpoint service in the GoldenGate service VCN.
	PrivateIpAddress *string `mandatory:"false" json:"privateIpAddress"`

	// The URL of a resource.
	DeploymentUrl *string `mandatory:"false" json:"deploymentUrl"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates if the resource is the the latest available version.
	IsLatestVersion *bool `mandatory:"false" json:"isLatestVersion"`

	// The amount of storage being utilized (in bytes)
	StorageUtilizationInBytes *int64 `mandatory:"false" json:"storageUtilizationInBytes"`

	// Deprecated: This field is not updated and will be removed in future versions. If storage utilization exceeds the limit, the respective warning message will appear in deployment messages, which can be accessed through /messages?deploymentId=.
	// Indicator will be true if the amount of storage being utilized exceeds the allowable storage utilization limit.  Exceeding the limit may be an indication of a misconfiguration of the deployment's GoldenGate service.
	IsStorageUtilizationLimitExceeded *bool `mandatory:"false" json:"isStorageUtilizationLimitExceeded"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The OCID(/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource.
	// Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud
	// subscription id is provided. Otherwise the cluster placement group must not be provided.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Indicates if disaster recovery is enabled for a deployment.
	// If not specified, disaster recovery is ENABLED when no clusterPlacementGroupId is provided, and DISABLED when a clusterPlacementGroupId is provided.
	DisasterRecoveryStatus DisasterRecoveryStatusEnum `mandatory:"false" json:"disasterRecoveryStatus,omitempty"`

	// The type of the deployment role.
	DeploymentRole DeploymentRoleEnum `mandatory:"false" json:"deploymentRole,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m DeploymentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetDeploymentCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeploymentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetLifecycleSubStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetEnvironmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryStatusEnum(string(m.DisasterRecoveryStatus)); !ok && m.DisasterRecoveryStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryStatus: %s. Supported values are: %s.", m.DisasterRecoveryStatus, strings.Join(GetDisasterRecoveryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentRoleEnum(string(m.DeploymentRole)); !ok && m.DeploymentRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentRole: %s. Supported values are: %s.", m.DeploymentRole, strings.Join(GetDeploymentRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
