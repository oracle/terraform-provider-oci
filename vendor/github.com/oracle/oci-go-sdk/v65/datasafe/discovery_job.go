// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryJob A data discovery job. It helps track job's metadata as well as result statistics.
type DiscoveryJob struct {

	// The OCID of the discovery job.
	Id *string `mandatory:"true" json:"id"`

	// The type of the discovery job. It defines the job's scope.
	// NEW identifies new sensitive columns in the target database that are not in the sensitive data model.
	// DELETED identifies columns that are present in the sensitive data model but have been deleted from the target database.
	// MODIFIED identifies columns that are present in the target database as well as the sensitive data model but some of their attributes have been modified.
	// ALL covers all the above three scenarios and reports new, deleted and modified columns.
	DiscoveryType DiscoveryJobDiscoveryTypeEnum `mandatory:"true" json:"discoveryType"`

	// The display name of the discovery job.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the discovery job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the discovery job started, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the discovery job finished, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)..
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// The current state of the discovery job.
	LifecycleState DiscoveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the sensitive data model associated with the discovery job.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The OCID of the target database associated with the discovery job.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Indicates if the discovery job should collect and store sample data values for the discovered columns.
	// Sample data helps review the discovered columns and ensure that they actually contain sensitive data.
	// As it collects original data from the target database, it's disabled by default and should be used only
	// if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data
	// values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE.
	IsSampleDataCollectionEnabled *bool `mandatory:"true" json:"isSampleDataCollectionEnabled"`

	// Indicates if the discovery job should identify potential application-level (non-dictionary) referential
	// relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined)
	// relationships. This option helps identify application-level relationships that are not defined in the database dictionary,
	// which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking.
	// It's disabled by default and should be used only if there is a need to identify application-level relationships.
	IsAppDefinedRelationDiscoveryEnabled *bool `mandatory:"true" json:"isAppDefinedRelationDiscoveryEnabled"`

	// Indicates if all the schemas in the associated target database are used for data discovery.
	// If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle).
	IsIncludeAllSchemas *bool `mandatory:"true" json:"isIncludeAllSchemas"`

	// Indicates if all the existing sensitive types are used for data discovery. If it's set to true, the
	// sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used.
	IsIncludeAllSensitiveTypes *bool `mandatory:"true" json:"isIncludeAllSensitiveTypes"`

	// The total number of schemas scanned by the discovery job.
	TotalSchemasScanned *int64 `mandatory:"true" json:"totalSchemasScanned"`

	// The total number of objects (tables and editioning views) scanned by the discovery job.
	TotalObjectsScanned *int64 `mandatory:"true" json:"totalObjectsScanned"`

	// The total number of columns scanned by the discovery job.
	TotalColumnsScanned *int64 `mandatory:"true" json:"totalColumnsScanned"`

	// The total number of new sensitive columns identified by the discovery job.
	TotalNewSensitiveColumns *int64 `mandatory:"true" json:"totalNewSensitiveColumns"`

	// The total number of modified sensitive columns identified by the discovery job.
	TotalModifiedSensitiveColumns *int64 `mandatory:"true" json:"totalModifiedSensitiveColumns"`

	// The total number of deleted sensitive columns identified by the discovery job.
	TotalDeletedSensitiveColumns *int64 `mandatory:"true" json:"totalDeletedSensitiveColumns"`

	// The schemas used for data discovery.
	SchemasForDiscovery []string `mandatory:"false" json:"schemasForDiscovery"`

	// The OCIDs of the sensitive types used for data discovery.
	SensitiveTypeIdsForDiscovery []string `mandatory:"false" json:"sensitiveTypeIdsForDiscovery"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DiscoveryJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryJobDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryJobDiscoveryTypeEnum Enum with underlying type: string
type DiscoveryJobDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoveryJobDiscoveryTypeEnum
const (
	DiscoveryJobDiscoveryTypeAll      DiscoveryJobDiscoveryTypeEnum = "ALL"
	DiscoveryJobDiscoveryTypeNew      DiscoveryJobDiscoveryTypeEnum = "NEW"
	DiscoveryJobDiscoveryTypeModified DiscoveryJobDiscoveryTypeEnum = "MODIFIED"
	DiscoveryJobDiscoveryTypeDeleted  DiscoveryJobDiscoveryTypeEnum = "DELETED"
)

var mappingDiscoveryJobDiscoveryTypeEnum = map[string]DiscoveryJobDiscoveryTypeEnum{
	"ALL":      DiscoveryJobDiscoveryTypeAll,
	"NEW":      DiscoveryJobDiscoveryTypeNew,
	"MODIFIED": DiscoveryJobDiscoveryTypeModified,
	"DELETED":  DiscoveryJobDiscoveryTypeDeleted,
}

var mappingDiscoveryJobDiscoveryTypeEnumLowerCase = map[string]DiscoveryJobDiscoveryTypeEnum{
	"all":      DiscoveryJobDiscoveryTypeAll,
	"new":      DiscoveryJobDiscoveryTypeNew,
	"modified": DiscoveryJobDiscoveryTypeModified,
	"deleted":  DiscoveryJobDiscoveryTypeDeleted,
}

// GetDiscoveryJobDiscoveryTypeEnumValues Enumerates the set of values for DiscoveryJobDiscoveryTypeEnum
func GetDiscoveryJobDiscoveryTypeEnumValues() []DiscoveryJobDiscoveryTypeEnum {
	values := make([]DiscoveryJobDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoveryJobDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryJobDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoveryJobDiscoveryTypeEnum
func GetDiscoveryJobDiscoveryTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"NEW",
		"MODIFIED",
		"DELETED",
	}
}

// GetMappingDiscoveryJobDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryJobDiscoveryTypeEnum(val string) (DiscoveryJobDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoveryJobDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
