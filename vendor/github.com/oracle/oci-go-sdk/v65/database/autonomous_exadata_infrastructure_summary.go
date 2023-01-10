// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousExadataInfrastructureSummary **Deprecated** These APIs are deprecated with the introduction of the Autonomous Exadata VM Cluster resource and a shift to a common Exadata Infrastructure resource for all Exadata Cloud-based services, including Autonomous Database on dedicated Exadata infrastructure. For more details, see Latest Resource Model (https://docs.oracle.com/en/cloud/paas/autonomous-database/flddd/#articletitle).
// Infrastructure that enables the running of multiple Autonomous Databases within a dedicated DB system.
// For more information about Autonomous Exadata Infrastructure, see
// Oracle Autonomous Database (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see
// Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the ListAvailabilityDomains operation
// in the Identity service API.
type AutonomousExadataInfrastructureSummary struct {

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
	LifecycleState AutonomousExadataInfrastructureSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"true" json:"maintenanceWindow"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Additional information about the current lifecycle state of the Autonomous Exadata Infrastructure.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The Oracle license model that applies to all databases in the Autonomous Exadata Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel AutonomousExadataInfrastructureSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

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

func (m AutonomousExadataInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousExadataInfrastructureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousExadataInfrastructureSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousExadataInfrastructureSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousExadataInfrastructureSummaryLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousExadataInfrastructureSummaryLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousExadataInfrastructureSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousExadataInfrastructureSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousExadataInfrastructureSummaryLifecycleStateEnum
const (
	AutonomousExadataInfrastructureSummaryLifecycleStateProvisioning          AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousExadataInfrastructureSummaryLifecycleStateAvailable             AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousExadataInfrastructureSummaryLifecycleStateUpdating              AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "UPDATING"
	AutonomousExadataInfrastructureSummaryLifecycleStateTerminating           AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousExadataInfrastructureSummaryLifecycleStateTerminated            AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousExadataInfrastructureSummaryLifecycleStateFailed                AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "FAILED"
	AutonomousExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress AutonomousExadataInfrastructureSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousExadataInfrastructureSummaryLifecycleStateEnum = map[string]AutonomousExadataInfrastructureSummaryLifecycleStateEnum{
	"PROVISIONING":            AutonomousExadataInfrastructureSummaryLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousExadataInfrastructureSummaryLifecycleStateAvailable,
	"UPDATING":                AutonomousExadataInfrastructureSummaryLifecycleStateUpdating,
	"TERMINATING":             AutonomousExadataInfrastructureSummaryLifecycleStateTerminating,
	"TERMINATED":              AutonomousExadataInfrastructureSummaryLifecycleStateTerminated,
	"FAILED":                  AutonomousExadataInfrastructureSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress,
}

var mappingAutonomousExadataInfrastructureSummaryLifecycleStateEnumLowerCase = map[string]AutonomousExadataInfrastructureSummaryLifecycleStateEnum{
	"provisioning":            AutonomousExadataInfrastructureSummaryLifecycleStateProvisioning,
	"available":               AutonomousExadataInfrastructureSummaryLifecycleStateAvailable,
	"updating":                AutonomousExadataInfrastructureSummaryLifecycleStateUpdating,
	"terminating":             AutonomousExadataInfrastructureSummaryLifecycleStateTerminating,
	"terminated":              AutonomousExadataInfrastructureSummaryLifecycleStateTerminated,
	"failed":                  AutonomousExadataInfrastructureSummaryLifecycleStateFailed,
	"maintenance_in_progress": AutonomousExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress,
}

// GetAutonomousExadataInfrastructureSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousExadataInfrastructureSummaryLifecycleStateEnum
func GetAutonomousExadataInfrastructureSummaryLifecycleStateEnumValues() []AutonomousExadataInfrastructureSummaryLifecycleStateEnum {
	values := make([]AutonomousExadataInfrastructureSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousExadataInfrastructureSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousExadataInfrastructureSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousExadataInfrastructureSummaryLifecycleStateEnum
func GetAutonomousExadataInfrastructureSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousExadataInfrastructureSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousExadataInfrastructureSummaryLifecycleStateEnum(val string) (AutonomousExadataInfrastructureSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousExadataInfrastructureSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousExadataInfrastructureSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousExadataInfrastructureSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousExadataInfrastructureSummaryLicenseModelEnum
const (
	AutonomousExadataInfrastructureSummaryLicenseModelLicenseIncluded     AutonomousExadataInfrastructureSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousExadataInfrastructureSummaryLicenseModelBringYourOwnLicense AutonomousExadataInfrastructureSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousExadataInfrastructureSummaryLicenseModelEnum = map[string]AutonomousExadataInfrastructureSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

var mappingAutonomousExadataInfrastructureSummaryLicenseModelEnumLowerCase = map[string]AutonomousExadataInfrastructureSummaryLicenseModelEnum{
	"license_included":       AutonomousExadataInfrastructureSummaryLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousExadataInfrastructureSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousExadataInfrastructureSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousExadataInfrastructureSummaryLicenseModelEnum
func GetAutonomousExadataInfrastructureSummaryLicenseModelEnumValues() []AutonomousExadataInfrastructureSummaryLicenseModelEnum {
	values := make([]AutonomousExadataInfrastructureSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousExadataInfrastructureSummaryLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousExadataInfrastructureSummaryLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousExadataInfrastructureSummaryLicenseModelEnum
func GetAutonomousExadataInfrastructureSummaryLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousExadataInfrastructureSummaryLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousExadataInfrastructureSummaryLicenseModelEnum(val string) (AutonomousExadataInfrastructureSummaryLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousExadataInfrastructureSummaryLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
