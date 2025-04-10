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

// UpdateDeploymentDetails The information to use to update a Deployment.
type UpdateDeploymentDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Specifies whether the deployment is used in a production or development/testing environment.
	EnvironmentType EnvironmentTypeEnum `mandatory:"false" json:"environmentType,omitempty"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet of the deployment's private endpoint.
	// The subnet must be a private subnet. For backward compatibility, public subnets are allowed until May 31 2025,
	// after which the private subnet will be enforced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a public subnet in the customer tenancy.
	// Can be provided only for public deployments. If provided, the loadbalancer will be created in this subnet instead of the service tenancy.
	// For backward compatibility, this is an optional property. It will become mandatory for public deployments after October 1, 2024.
	LoadBalancerSubnetId *string `mandatory:"false" json:"loadBalancerSubnetId"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// An array of local peers of deployment
	Placements []DeploymentPlacementDetails `mandatory:"false" json:"placements"`

	OggData *UpdateOggDeploymentDetails `mandatory:"false" json:"oggData"`

	MaintenanceWindow *UpdateMaintenanceWindowDetails `mandatory:"false" json:"maintenanceWindow"`

	MaintenanceConfiguration *UpdateMaintenanceConfigurationDetails `mandatory:"false" json:"maintenanceConfiguration"`

	BackupSchedule *UpdateBackupScheduleDetails `mandatory:"false" json:"backupSchedule"`
}

func (m UpdateDeploymentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDeploymentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetEnvironmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
