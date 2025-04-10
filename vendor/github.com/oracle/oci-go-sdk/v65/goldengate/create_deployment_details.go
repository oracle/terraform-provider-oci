// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDeploymentDetails The information about a new Deployment.
type CreateDeploymentDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint.
	// The subnet must be a private subnet. For backward compatibility, public subnets are allowed until May 31 2025,
	// after which the private subnet will be enforced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Specifies whether the deployment is used in a production or development/testing environment.
	EnvironmentType EnvironmentTypeEnum `mandatory:"false" json:"environmentType,omitempty"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	SourceDeploymentId *string `mandatory:"false" json:"sourceDeploymentId"`

	// The availability domain of a placement.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The fault domain of a placement.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// An array of local peers of deployment
	Placements []DeploymentPlacementDetails `mandatory:"false" json:"placements"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []AddResourceLockDetails `mandatory:"false" json:"locks"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	DeploymentBackupId *string `mandatory:"false" json:"deploymentBackupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy.
	// Can be provided only for public deployments. If provided, the loadbalancer will be created in this subnet instead of the service tenancy.
	// For backward compatibility, this is an optional property. It will become mandatory for public deployments after October 1, 2024.
	LoadBalancerSubnetId *string `mandatory:"false" json:"loadBalancerSubnetId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	OggData *CreateOggDeploymentDetails `mandatory:"false" json:"oggData"`

	MaintenanceWindow *CreateMaintenanceWindowDetails `mandatory:"false" json:"maintenanceWindow"`

	MaintenanceConfiguration *CreateMaintenanceConfigurationDetails `mandatory:"false" json:"maintenanceConfiguration"`

	BackupSchedule *CreateBackupScheduleDetails `mandatory:"false" json:"backupSchedule"`
}

func (m CreateDeploymentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDeploymentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetEnvironmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
