// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAutonomousDatabaseCloneDetails Details to create an Oracle Autonomous Database by cloning an existing Autonomous Database.
type CreateAutonomousDatabaseCloneDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of OCPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
	IsPreviewVersionWithServiceTermsAccepted *bool `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`

	// Indicates if auto scaling is enabled for the Autonomous Database OCPU core count. The default value is `FALSE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// True if the database is on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The client IP access control list (ACL). This feature is available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs.
	// For update operation, if you wish to delete all the existing whitelisted IP’s, use an array with a single empty string entry.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.0.0/16"]`
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Indicates whether the Autonomous Database has Data Guard enabled.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The Autonomous Database clone type.
	CloneType CreateAutonomousDatabaseCloneDetailsCloneTypeEnum `mandatory:"true" json:"cloneType"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	DbWorkload CreateAutonomousDatabaseBaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel CreateAutonomousDatabaseBaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

//GetCompartmentId returns CompartmentId
func (m CreateAutonomousDatabaseCloneDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDbName returns DbName
func (m CreateAutonomousDatabaseCloneDetails) GetDbName() *string {
	return m.DbName
}

//GetCpuCoreCount returns CpuCoreCount
func (m CreateAutonomousDatabaseCloneDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

//GetDbWorkload returns DbWorkload
func (m CreateAutonomousDatabaseCloneDetails) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

//GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m CreateAutonomousDatabaseCloneDetails) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

//GetIsFreeTier returns IsFreeTier
func (m CreateAutonomousDatabaseCloneDetails) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

//GetAdminPassword returns AdminPassword
func (m CreateAutonomousDatabaseCloneDetails) GetAdminPassword() *string {
	return m.AdminPassword
}

//GetDisplayName returns DisplayName
func (m CreateAutonomousDatabaseCloneDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetLicenseModel returns LicenseModel
func (m CreateAutonomousDatabaseCloneDetails) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

//GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m CreateAutonomousDatabaseCloneDetails) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

//GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

//GetIsDedicated returns IsDedicated
func (m CreateAutonomousDatabaseCloneDetails) GetIsDedicated() *bool {
	return m.IsDedicated
}

//GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m CreateAutonomousDatabaseCloneDetails) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

//GetWhitelistedIps returns WhitelistedIps
func (m CreateAutonomousDatabaseCloneDetails) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

//GetIsDataGuardEnabled returns IsDataGuardEnabled
func (m CreateAutonomousDatabaseCloneDetails) GetIsDataGuardEnabled() *bool {
	return m.IsDataGuardEnabled
}

//GetSubnetId returns SubnetId
func (m CreateAutonomousDatabaseCloneDetails) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m CreateAutonomousDatabaseCloneDetails) GetNsgIds() []string {
	return m.NsgIds
}

//GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m CreateAutonomousDatabaseCloneDetails) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

//GetFreeformTags returns FreeformTags
func (m CreateAutonomousDatabaseCloneDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateAutonomousDatabaseCloneDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDbVersion returns DbVersion
func (m CreateAutonomousDatabaseCloneDetails) GetDbVersion() *string {
	return m.DbVersion
}

func (m CreateAutonomousDatabaseCloneDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateAutonomousDatabaseCloneDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousDatabaseCloneDetails CreateAutonomousDatabaseCloneDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateAutonomousDatabaseCloneDetails
	}{
		"DATABASE",
		(MarshalTypeCreateAutonomousDatabaseCloneDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateAutonomousDatabaseCloneDetailsCloneTypeEnum Enum with underlying type: string
type CreateAutonomousDatabaseCloneDetailsCloneTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseCloneDetailsCloneTypeEnum
const (
	CreateAutonomousDatabaseCloneDetailsCloneTypeFull     CreateAutonomousDatabaseCloneDetailsCloneTypeEnum = "FULL"
	CreateAutonomousDatabaseCloneDetailsCloneTypeMetadata CreateAutonomousDatabaseCloneDetailsCloneTypeEnum = "METADATA"
)

var mappingCreateAutonomousDatabaseCloneDetailsCloneType = map[string]CreateAutonomousDatabaseCloneDetailsCloneTypeEnum{
	"FULL":     CreateAutonomousDatabaseCloneDetailsCloneTypeFull,
	"METADATA": CreateAutonomousDatabaseCloneDetailsCloneTypeMetadata,
}

// GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumValues Enumerates the set of values for CreateAutonomousDatabaseCloneDetailsCloneTypeEnum
func GetCreateAutonomousDatabaseCloneDetailsCloneTypeEnumValues() []CreateAutonomousDatabaseCloneDetailsCloneTypeEnum {
	values := make([]CreateAutonomousDatabaseCloneDetailsCloneTypeEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseCloneDetailsCloneType {
		values = append(values, v)
	}
	return values
}
