// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryJobSummary The Summary of DiscoveryJob details.
type DiscoveryJobSummary struct {

	// The OCID of Discovery job
	Id *string `mandatory:"true" json:"id"`

	// Resource Type
	ResourceType DiscoveryJobSummaryResourceTypeEnum `mandatory:"false" json:"resourceType,omitempty"`

	// The name of resource type
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// License edition of the monitored resource.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	// The OCID of the Compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Add option submits new discovery Job. Add with retry option to re-submit failed discovery job. Refresh option refreshes the existing discovered resources.
	DiscoveryType DiscoveryJobSummaryDiscoveryTypeEnum `mandatory:"false" json:"discoveryType,omitempty"`

	// Specifies the status of the discovery job
	Status DiscoveryJobSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The short summary of the status of the discovery job
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The OCID of Tenant
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The OCID of user in which the job is submitted
	UserId *string `mandatory:"false" json:"userId"`

	// The time the discovery Job was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the DiscoveryJob Resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DiscoveryJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveryJobSummaryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDiscoveryJobSummaryResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobSummaryDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobSummaryDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveryJobSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobSummaryResourceTypeEnum Enum with underlying type: string
type DiscoveryJobSummaryResourceTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobSummaryResourceTypeEnum
const (
	DiscoveryJobSummaryResourceTypeWeblogicDomain   DiscoveryJobSummaryResourceTypeEnum = "WEBLOGIC_DOMAIN"
	DiscoveryJobSummaryResourceTypeEbsInstance      DiscoveryJobSummaryResourceTypeEnum = "EBS_INSTANCE"
	DiscoveryJobSummaryResourceTypeSqlServer        DiscoveryJobSummaryResourceTypeEnum = "SQL_SERVER"
	DiscoveryJobSummaryResourceTypeApacheTomcat     DiscoveryJobSummaryResourceTypeEnum = "APACHE_TOMCAT"
	DiscoveryJobSummaryResourceTypeOracleDatabase   DiscoveryJobSummaryResourceTypeEnum = "ORACLE_DATABASE"
	DiscoveryJobSummaryResourceTypeOciOracleDb      DiscoveryJobSummaryResourceTypeEnum = "OCI_ORACLE_DB"
	DiscoveryJobSummaryResourceTypeOciOracleCdb     DiscoveryJobSummaryResourceTypeEnum = "OCI_ORACLE_CDB"
	DiscoveryJobSummaryResourceTypeOciOraclePdb     DiscoveryJobSummaryResourceTypeEnum = "OCI_ORACLE_PDB"
	DiscoveryJobSummaryResourceTypeHost             DiscoveryJobSummaryResourceTypeEnum = "HOST"
	DiscoveryJobSummaryResourceTypeOraclePsft       DiscoveryJobSummaryResourceTypeEnum = "ORACLE_PSFT"
	DiscoveryJobSummaryResourceTypeOracleMft        DiscoveryJobSummaryResourceTypeEnum = "ORACLE_MFT"
	DiscoveryJobSummaryResourceTypeApacheHttpServer DiscoveryJobSummaryResourceTypeEnum = "APACHE_HTTP_SERVER"
)

var mappingDiscoveryJobSummaryResourceTypeEnum = map[string]DiscoveryJobSummaryResourceTypeEnum{
	"WEBLOGIC_DOMAIN":    DiscoveryJobSummaryResourceTypeWeblogicDomain,
	"EBS_INSTANCE":       DiscoveryJobSummaryResourceTypeEbsInstance,
	"SQL_SERVER":         DiscoveryJobSummaryResourceTypeSqlServer,
	"APACHE_TOMCAT":      DiscoveryJobSummaryResourceTypeApacheTomcat,
	"ORACLE_DATABASE":    DiscoveryJobSummaryResourceTypeOracleDatabase,
	"OCI_ORACLE_DB":      DiscoveryJobSummaryResourceTypeOciOracleDb,
	"OCI_ORACLE_CDB":     DiscoveryJobSummaryResourceTypeOciOracleCdb,
	"OCI_ORACLE_PDB":     DiscoveryJobSummaryResourceTypeOciOraclePdb,
	"HOST":               DiscoveryJobSummaryResourceTypeHost,
	"ORACLE_PSFT":        DiscoveryJobSummaryResourceTypeOraclePsft,
	"ORACLE_MFT":         DiscoveryJobSummaryResourceTypeOracleMft,
	"APACHE_HTTP_SERVER": DiscoveryJobSummaryResourceTypeApacheHttpServer,
}

var mappingDiscoveryJobSummaryResourceTypeEnumLowerCase = map[string]DiscoveryJobSummaryResourceTypeEnum{
	"weblogic_domain":    DiscoveryJobSummaryResourceTypeWeblogicDomain,
	"ebs_instance":       DiscoveryJobSummaryResourceTypeEbsInstance,
	"sql_server":         DiscoveryJobSummaryResourceTypeSqlServer,
	"apache_tomcat":      DiscoveryJobSummaryResourceTypeApacheTomcat,
	"oracle_database":    DiscoveryJobSummaryResourceTypeOracleDatabase,
	"oci_oracle_db":      DiscoveryJobSummaryResourceTypeOciOracleDb,
	"oci_oracle_cdb":     DiscoveryJobSummaryResourceTypeOciOracleCdb,
	"oci_oracle_pdb":     DiscoveryJobSummaryResourceTypeOciOraclePdb,
	"host":               DiscoveryJobSummaryResourceTypeHost,
	"oracle_psft":        DiscoveryJobSummaryResourceTypeOraclePsft,
	"oracle_mft":         DiscoveryJobSummaryResourceTypeOracleMft,
	"apache_http_server": DiscoveryJobSummaryResourceTypeApacheHttpServer,
}

// GetDiscoveryJobSummaryResourceTypeEnumValues Enumerates the set of values for DiscoveryJobSummaryResourceTypeEnum
func GetDiscoveryJobSummaryResourceTypeEnumValues() []DiscoveryJobSummaryResourceTypeEnum {
	values := make([]DiscoveryJobSummaryResourceTypeEnum, 0)
	for _, v := range mappingDiscoveryJobSummaryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobSummaryResourceTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobSummaryResourceTypeEnum
func GetDiscoveryJobSummaryResourceTypeEnumStringValues() []string {
	return []string{
		"WEBLOGIC_DOMAIN",
		"EBS_INSTANCE",
		"SQL_SERVER",
		"APACHE_TOMCAT",
		"ORACLE_DATABASE",
		"OCI_ORACLE_DB",
		"OCI_ORACLE_CDB",
		"OCI_ORACLE_PDB",
		"HOST",
		"ORACLE_PSFT",
		"ORACLE_MFT",
		"APACHE_HTTP_SERVER",
	}
}

// GetMappingDiscoveryJobSummaryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobSummaryResourceTypeEnum(val string) (DiscoveryJobSummaryResourceTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobSummaryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobSummaryDiscoveryTypeEnum Enum with underlying type: string
type DiscoveryJobSummaryDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobSummaryDiscoveryTypeEnum
const (
	DiscoveryJobSummaryDiscoveryTypeAdd          DiscoveryJobSummaryDiscoveryTypeEnum = "ADD"
	DiscoveryJobSummaryDiscoveryTypeAddWithRetry DiscoveryJobSummaryDiscoveryTypeEnum = "ADD_WITH_RETRY"
	DiscoveryJobSummaryDiscoveryTypeRefresh      DiscoveryJobSummaryDiscoveryTypeEnum = "REFRESH"
)

var mappingDiscoveryJobSummaryDiscoveryTypeEnum = map[string]DiscoveryJobSummaryDiscoveryTypeEnum{
	"ADD":            DiscoveryJobSummaryDiscoveryTypeAdd,
	"ADD_WITH_RETRY": DiscoveryJobSummaryDiscoveryTypeAddWithRetry,
	"REFRESH":        DiscoveryJobSummaryDiscoveryTypeRefresh,
}

var mappingDiscoveryJobSummaryDiscoveryTypeEnumLowerCase = map[string]DiscoveryJobSummaryDiscoveryTypeEnum{
	"add":            DiscoveryJobSummaryDiscoveryTypeAdd,
	"add_with_retry": DiscoveryJobSummaryDiscoveryTypeAddWithRetry,
	"refresh":        DiscoveryJobSummaryDiscoveryTypeRefresh,
}

// GetDiscoveryJobSummaryDiscoveryTypeEnumValues Enumerates the set of values for DiscoveryJobSummaryDiscoveryTypeEnum
func GetDiscoveryJobSummaryDiscoveryTypeEnumValues() []DiscoveryJobSummaryDiscoveryTypeEnum {
	values := make([]DiscoveryJobSummaryDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoveryJobSummaryDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobSummaryDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobSummaryDiscoveryTypeEnum
func GetDiscoveryJobSummaryDiscoveryTypeEnumStringValues() []string {
	return []string{
		"ADD",
		"ADD_WITH_RETRY",
		"REFRESH",
	}
}

// GetMappingDiscoveryJobSummaryDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobSummaryDiscoveryTypeEnum(val string) (DiscoveryJobSummaryDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobSummaryDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DiscoveryJobSummaryStatusEnum Enum with underlying type: string
type DiscoveryJobSummaryStatusEnum string

// Set of constants representing the allowable values for DiscoveryJobSummaryStatusEnum
const (
	DiscoveryJobSummaryStatusSuccess    DiscoveryJobSummaryStatusEnum = "SUCCESS"
	DiscoveryJobSummaryStatusFailure    DiscoveryJobSummaryStatusEnum = "FAILURE"
	DiscoveryJobSummaryStatusInprogress DiscoveryJobSummaryStatusEnum = "INPROGRESS"
	DiscoveryJobSummaryStatusInactive   DiscoveryJobSummaryStatusEnum = "INACTIVE"
	DiscoveryJobSummaryStatusCreated    DiscoveryJobSummaryStatusEnum = "CREATED"
	DiscoveryJobSummaryStatusDeleted    DiscoveryJobSummaryStatusEnum = "DELETED"
)

var mappingDiscoveryJobSummaryStatusEnum = map[string]DiscoveryJobSummaryStatusEnum{
	"SUCCESS":    DiscoveryJobSummaryStatusSuccess,
	"FAILURE":    DiscoveryJobSummaryStatusFailure,
	"INPROGRESS": DiscoveryJobSummaryStatusInprogress,
	"INACTIVE":   DiscoveryJobSummaryStatusInactive,
	"CREATED":    DiscoveryJobSummaryStatusCreated,
	"DELETED":    DiscoveryJobSummaryStatusDeleted,
}

var mappingDiscoveryJobSummaryStatusEnumLowerCase = map[string]DiscoveryJobSummaryStatusEnum{
	"success":    DiscoveryJobSummaryStatusSuccess,
	"failure":    DiscoveryJobSummaryStatusFailure,
	"inprogress": DiscoveryJobSummaryStatusInprogress,
	"inactive":   DiscoveryJobSummaryStatusInactive,
	"created":    DiscoveryJobSummaryStatusCreated,
	"deleted":    DiscoveryJobSummaryStatusDeleted,
}

// GetDiscoveryJobSummaryStatusEnumValues Enumerates the set of values for DiscoveryJobSummaryStatusEnum
func GetDiscoveryJobSummaryStatusEnumValues() []DiscoveryJobSummaryStatusEnum {
	values := make([]DiscoveryJobSummaryStatusEnum, 0)
	for _, v := range mappingDiscoveryJobSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobSummaryStatusEnumStringValues Enumerates the set of values in String for DiscoveryJobSummaryStatusEnum
func GetDiscoveryJobSummaryStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
		"INPROGRESS",
		"INACTIVE",
		"CREATED",
		"DELETED",
	}
}

// GetMappingDiscoveryJobSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobSummaryStatusEnum(val string) (DiscoveryJobSummaryStatusEnum, bool) {
	enum, ok := mappingDiscoveryJobSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
