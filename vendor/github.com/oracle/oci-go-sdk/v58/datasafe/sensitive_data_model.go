// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SensitiveDataModel A sensitive data model is a collection of sensitive columns and their referential relationships. It helps
// understand the sensitive data landscape, track changes, and efficiently enable security controls such as data
// masking. It can be managed either manually or by performing sensitive data discovery on a reference target
// database. Learn more (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/sensitive-data-models1.html#GUID-849CA7D2-1809-40DD-B6D7-44E46EFF7EB5).
type SensitiveDataModel struct {

	// The OCID of the sensitive data model.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the sensitive data model.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the sensitive data model.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the reference target database associated with the sensitive data model. All operations such as
	// performing data discovery and adding columns manually are done in the context of the associated target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The date and time the sensitive data model was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the sensitive data model was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the sensitive data model.
	LifecycleState DiscoveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The application suite name identifying a collection of applications. The default value is GENERIC. It's useful
	// only if maintaining a sensitive data model for a suite of applications.
	AppSuiteName *string `mandatory:"true" json:"appSuiteName"`

	// Indicates if data discovery jobs should collect and store sample data values for the discovered columns.
	// Sample data helps review the discovered columns and ensure that they actually contain sensitive data.
	// As it collects original data from the target database, it's disabled by default and should be used only if
	// it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values
	// are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE.
	IsSampleDataCollectionEnabled *bool `mandatory:"true" json:"isSampleDataCollectionEnabled"`

	// Indicates if data discovery jobs should identify potential application-level (non-dictionary) referential
	// relationships between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined)
	// relationships. This option helps identify application-level relationships that are not defined in the database dictionary,
	// which in turn, helps identify additional sensitive columns and preserve referential integrity during data masking.
	// It's disabled by default and should be used only if there is a need to identify application-level relationships.
	IsAppDefinedRelationDiscoveryEnabled *bool `mandatory:"true" json:"isAppDefinedRelationDiscoveryEnabled"`

	// Indicates if all the schemas in the associated target database should be scanned by data discovery jobs.
	// If it's set to true, the schemasForDiscovery attribute is ignored and all schemas are used for data discovery.
	IsIncludeAllSchemas *bool `mandatory:"true" json:"isIncludeAllSchemas"`

	// Indicates if all the existing sensitive types should be used by data discovery jobs.If it's set to true, the
	// sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery.
	IsIncludeAllSensitiveTypes *bool `mandatory:"true" json:"isIncludeAllSensitiveTypes"`

	// The description of the sensitive data model.
	Description *string `mandatory:"false" json:"description"`

	// The schemas to be scanned by data discovery jobs.
	SchemasForDiscovery []string `mandatory:"false" json:"schemasForDiscovery"`

	// The OCIDs of the sensitive types to be used by data discovery jobs.
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

func (m SensitiveDataModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveDataModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
