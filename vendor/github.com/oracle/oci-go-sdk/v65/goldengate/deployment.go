// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Deployment A container for your OCI GoldenGate resources, such as the OCI GoldenGate deployment console.
type Deployment struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	DeploymentBackupId *string `mandatory:"false" json:"deploymentBackupId"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

	// True if all of the aggregate resources are working correctly.
	IsHealthy *bool `mandatory:"false" json:"isHealthy"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy.
	// Can be provided only for public deployments. If provided, the loadbalancer will be created in this subnet instead of the service tenancy.
	// For backward compatiblity this is an optional property for now, but it will become mandatory (for public deployments only) after October 1, 2024.
	LoadBalancerSubnetId *string `mandatory:"false" json:"loadBalancerSubnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the loadbalancer in the customer's subnet.
	// The loadbalancer of the public deployment created in the customer subnet.
	LoadBalancerId *string `mandatory:"false" json:"loadBalancerId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

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
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates if the resource is the the latest available version.
	IsLatestVersion *bool `mandatory:"false" json:"isLatestVersion"`

	// Note: Deprecated: Use timeOfNextMaintenance instead, or related upgrade records
	// to check, when deployment will be forced to upgrade to a newer version.
	// Old description:
	// The date the existing version in use will no longer be considered as usable
	// and an upgrade will be required.  This date is typically 6 months after the
	// version was released for use by GGS.  The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpgradeRequired *common.SDKTime `mandatory:"false" json:"timeUpgradeRequired"`

	// The amount of storage being utilized (in bytes)
	StorageUtilizationInBytes *int64 `mandatory:"false" json:"storageUtilizationInBytes"`

	// Indicator will be true if the amount of storage being utilized exceeds the allowable storage utilization limit.  Exceeding the limit may be an indication of a misconfiguration of the deployment's GoldenGate service.
	IsStorageUtilizationLimitExceeded *bool `mandatory:"false" json:"isStorageUtilizationLimitExceeded"`

	OggData *OggDeployment `mandatory:"false" json:"oggData"`

	DeploymentDiagnosticData *DeploymentDiagnosticData `mandatory:"false" json:"deploymentDiagnosticData"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The time of next maintenance schedule. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeOfNextMaintenance *common.SDKTime `mandatory:"false" json:"timeOfNextMaintenance"`

	// Type of the next maintenance.
	NextMaintenanceActionType MaintenanceActionTypeEnum `mandatory:"false" json:"nextMaintenanceActionType,omitempty"`

	// Description of the next maintenance.
	NextMaintenanceDescription *string `mandatory:"false" json:"nextMaintenanceDescription"`

	MaintenanceConfiguration *MaintenanceConfiguration `mandatory:"false" json:"maintenanceConfiguration"`

	// The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeOggVersionSupportedUntil *common.SDKTime `mandatory:"false" json:"timeOggVersionSupportedUntil"`

	// List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.
	// Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`
}

func (m Deployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Deployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetLifecycleSubStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceActionTypeEnum(string(m.NextMaintenanceActionType)); !ok && m.NextMaintenanceActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextMaintenanceActionType: %s. Supported values are: %s.", m.NextMaintenanceActionType, strings.Join(GetMaintenanceActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
