// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousExadataInfrastructure The representation of AutonomousExadataInfrastructure
type AutonomousExadataInfrastructure struct {

	// The OCID of the Autonomous Exadata Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous Exadata Infrastructure.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the availability domain that the Autonomous Exadata Infrastructure is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the subnet the Autonomous Exadata Infrastructure is associated with.
	// **Subnet Restrictions:**
	// - For Autonomous Databases with Autonomous Exadata Infrastructure, do not use a subnet that overlaps with 192.168.128.0/20
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the Autonomous Exadata Infrastructure. The shape determines resources to allocate to the Autonomous Exadata Infrastructure (CPU cores, memory and storage).
	Shape *string `mandatory:"true" json:"shape"`

	// The host name for the Autonomous Exadata Infrastructure node.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The domain name for the Autonomous Exadata Infrastructure.
	Domain *string `mandatory:"true" json:"domain"`

	// The current lifecycle state of the Autonomous Exadata Infrastructure.
	LifecycleState AutonomousExadataInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"true" json:"maintenanceWindow"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Additional information about the current lifecycle state of the Autonomous Exadata Infrastructure.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The Oracle license model that applies to all databases in the Autonomous Exadata Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousExadataInfrastructureLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The date and time the Autonomous Exadata Infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The FQDN of the DNS record for the SCAN IP addresses that are associated with the Autonomous Exadata Infrastructure.
	ScanDnsName *string `mandatory:"false" json:"scanDnsName"`

	// The OCID of the zone the Autonomous Exadata Infrastructure is associated with.
	ZoneId *string `mandatory:"false" json:"zoneId"`
}

func (m AutonomousExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousExadataInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousExadataInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousExadataInfrastructureLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousExadataInfrastructureLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousExadataInfrastructureLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousExadataInfrastructureLifecycleStateEnum Enum with underlying type: string
type AutonomousExadataInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousExadataInfrastructureLifecycleStateEnum
const (
	AutonomousExadataInfrastructureLifecycleStateProvisioning          AutonomousExadataInfrastructureLifecycleStateEnum = "PROVISIONING"
	AutonomousExadataInfrastructureLifecycleStateAvailable             AutonomousExadataInfrastructureLifecycleStateEnum = "AVAILABLE"
	AutonomousExadataInfrastructureLifecycleStateUpdating              AutonomousExadataInfrastructureLifecycleStateEnum = "UPDATING"
	AutonomousExadataInfrastructureLifecycleStateTerminating           AutonomousExadataInfrastructureLifecycleStateEnum = "TERMINATING"
	AutonomousExadataInfrastructureLifecycleStateTerminated            AutonomousExadataInfrastructureLifecycleStateEnum = "TERMINATED"
	AutonomousExadataInfrastructureLifecycleStateFailed                AutonomousExadataInfrastructureLifecycleStateEnum = "FAILED"
	AutonomousExadataInfrastructureLifecycleStateMaintenanceInProgress AutonomousExadataInfrastructureLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousExadataInfrastructureLifecycleStateEnum = map[string]AutonomousExadataInfrastructureLifecycleStateEnum{
	"PROVISIONING":            AutonomousExadataInfrastructureLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousExadataInfrastructureLifecycleStateAvailable,
	"UPDATING":                AutonomousExadataInfrastructureLifecycleStateUpdating,
	"TERMINATING":             AutonomousExadataInfrastructureLifecycleStateTerminating,
	"TERMINATED":              AutonomousExadataInfrastructureLifecycleStateTerminated,
	"FAILED":                  AutonomousExadataInfrastructureLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousExadataInfrastructureLifecycleStateMaintenanceInProgress,
}

var mappingAutonomousExadataInfrastructureLifecycleStateEnumLowerCase = map[string]AutonomousExadataInfrastructureLifecycleStateEnum{
	"provisioning":            AutonomousExadataInfrastructureLifecycleStateProvisioning,
	"available":               AutonomousExadataInfrastructureLifecycleStateAvailable,
	"updating":                AutonomousExadataInfrastructureLifecycleStateUpdating,
	"terminating":             AutonomousExadataInfrastructureLifecycleStateTerminating,
	"terminated":              AutonomousExadataInfrastructureLifecycleStateTerminated,
	"failed":                  AutonomousExadataInfrastructureLifecycleStateFailed,
	"maintenance_in_progress": AutonomousExadataInfrastructureLifecycleStateMaintenanceInProgress,
}

// GetAutonomousExadataInfrastructureLifecycleStateEnumValues Enumerates the set of values for AutonomousExadataInfrastructureLifecycleStateEnum
func GetAutonomousExadataInfrastructureLifecycleStateEnumValues() []AutonomousExadataInfrastructureLifecycleStateEnum {
	values := make([]AutonomousExadataInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousExadataInfrastructureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousExadataInfrastructureLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousExadataInfrastructureLifecycleStateEnum
func GetAutonomousExadataInfrastructureLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingAutonomousExadataInfrastructureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousExadataInfrastructureLifecycleStateEnum(val string) (AutonomousExadataInfrastructureLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousExadataInfrastructureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousExadataInfrastructureLicenseModelEnum Enum with underlying type: string
type AutonomousExadataInfrastructureLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousExadataInfrastructureLicenseModelEnum
const (
	AutonomousExadataInfrastructureLicenseModelLicenseIncluded     AutonomousExadataInfrastructureLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousExadataInfrastructureLicenseModelBringYourOwnLicense AutonomousExadataInfrastructureLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousExadataInfrastructureLicenseModelEnum = map[string]AutonomousExadataInfrastructureLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousExadataInfrastructureLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousExadataInfrastructureLicenseModelBringYourOwnLicense,
}

var mappingAutonomousExadataInfrastructureLicenseModelEnumLowerCase = map[string]AutonomousExadataInfrastructureLicenseModelEnum{
	"license_included":       AutonomousExadataInfrastructureLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousExadataInfrastructureLicenseModelBringYourOwnLicense,
}

// GetAutonomousExadataInfrastructureLicenseModelEnumValues Enumerates the set of values for AutonomousExadataInfrastructureLicenseModelEnum
func GetAutonomousExadataInfrastructureLicenseModelEnumValues() []AutonomousExadataInfrastructureLicenseModelEnum {
	values := make([]AutonomousExadataInfrastructureLicenseModelEnum, 0)
	for _, v := range mappingAutonomousExadataInfrastructureLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousExadataInfrastructureLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousExadataInfrastructureLicenseModelEnum
func GetAutonomousExadataInfrastructureLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousExadataInfrastructureLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousExadataInfrastructureLicenseModelEnum(val string) (AutonomousExadataInfrastructureLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousExadataInfrastructureLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
