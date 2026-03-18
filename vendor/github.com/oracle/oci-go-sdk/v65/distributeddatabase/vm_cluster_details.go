// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmClusterDetails Details of the request to create exadb vm cluster for shard or catalog of the distributed database.
type VmClusterDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure.
	BackupSubnetId *string `mandatory:"true" json:"backupSubnetId"`

	// The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure.
	EnabledECpuCount *int `mandatory:"true" json:"enabledECpuCount"`

	// The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure.
	TotalECpuCount *int `mandatory:"false" json:"totalECpuCount"`

	// File System Storage Size in GBs for Exadata VM cluster.
	VmFileSystemStorageSize *int `mandatory:"false" json:"vmFileSystemStorageSize"`

	// The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel VmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// A domain name used for the Exadata VM cluster on Exascale Infrastructure.
	// If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used
	// (do not provide one). Otherwise, provide a valid DNS domain name.
	// Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only.
	Domain *string `mandatory:"false" json:"domain"`

	// The private zone ID in which you want DNS records to be created.
	PrivateZoneId *string `mandatory:"false" json:"privateZoneId"`

	// Indicates whether diagnostic collection is enabled for the VM cluster.
	// Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.
	// Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.
	// You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API.
	IsDiagnosticsEventsEnabled *bool `mandatory:"false" json:"isDiagnosticsEventsEnabled"`

	// Indicates whether health monitoring is enabled for the VM cluster.
	// Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.
	// You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system.
	// Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API.
	IsHealthMonitoringEnabled *bool `mandatory:"false" json:"isHealthMonitoringEnabled"`

	// Indicates whether incident logs and trace collection are enabled for the VM cluster.
	// Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them.
	// Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API.
	IsIncidentLogsEnabled *bool `mandatory:"false" json:"isIncidentLogsEnabled"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.
	// Setting this to an empty list removes all resources from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.
	// Setting this to an empty array after the list is created removes the resource from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`
}

func (m VmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterDetailsLicenseModelEnum Enum with underlying type: string
type VmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for VmClusterDetailsLicenseModelEnum
const (
	VmClusterDetailsLicenseModelLicenseIncluded     VmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	VmClusterDetailsLicenseModelBringYourOwnLicense VmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingVmClusterDetailsLicenseModelEnum = map[string]VmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       VmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": VmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingVmClusterDetailsLicenseModelEnumLowerCase = map[string]VmClusterDetailsLicenseModelEnum{
	"license_included":       VmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": VmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for VmClusterDetailsLicenseModelEnum
func GetVmClusterDetailsLicenseModelEnumValues() []VmClusterDetailsLicenseModelEnum {
	values := make([]VmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for VmClusterDetailsLicenseModelEnum
func GetVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterDetailsLicenseModelEnum(val string) (VmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
