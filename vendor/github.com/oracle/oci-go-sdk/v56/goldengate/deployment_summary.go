// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DeploymentSummary Summary of the Deployment.
type DeploymentSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The time the resource was created. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Possible GGS lifecycle sub-states.
	LifecycleSubState LifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

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

	// The private IP address in the customer's VCN representing the access point for the associated endpoint service in the GoldenGate service VCN.
	PrivateIpAddress *string `mandatory:"false" json:"privateIpAddress"`

	// The URL of a resource.
	DeploymentUrl *string `mandatory:"false" json:"deploymentUrl"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates if the resource is the the latest available version.
	IsLatestVersion *bool `mandatory:"false" json:"isLatestVersion"`

	// The date the existing version in use will no longer be considered as usable and an upgrade will be required.  This date is typically 6 months after the version was released for use by GGS.  The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpgradeRequired *common.SDKTime `mandatory:"false" json:"timeUpgradeRequired"`

	// The type of deployment, the value determines the exact 'type' of service executed in the Deployment. NOTE: Use of the value OGG is maintained for backward compatibility purposes.  Its use is discouraged
	//       in favor of the equivalent DATABASE_ORACLE value.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`
}

func (m DeploymentSummary) String() string {
	return common.PointerString(m)
}
