// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateSensitiveDataModelDetails Details to create a new sensitive data model. If schemas and sensitive types are provided, it automatically runs
// data discovery and adds the discovered columns to the sensitive data model. Otherwise, it creates an empty sensitive
// data model that can be updated later.
// To specify some schemas and sensitive types for data discovery, use schemasForDiscovery and sensitiveTypeIdsForDiscovery
// attributes. But if you want to include all schemas and sensitive types, you can set isIncludeAllSchemas and
// isIncludeAllSensitiveTypes attributes to true. In the latter case, you do not need to list all schemas and
// sensitive types.
type CreateSensitiveDataModelDetails struct {

	// The OCID of the compartment where the sensitive data model should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the reference target database to be associated with the sensitive data model. All operations such
	// as performing data discovery and adding columns manually are done in the context of the associated target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The display name of the sensitive data model. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The application suite name identifying a collection of applications. It's useful only if maintaining a sensitive data model for a suite of applications.
	AppSuiteName *string `mandatory:"false" json:"appSuiteName"`

	// The description of the sensitive data model.
	Description *string `mandatory:"false" json:"description"`

	// The schemas to be scanned by data discovery jobs.
	SchemasForDiscovery []string `mandatory:"false" json:"schemasForDiscovery"`

	// The data discovery jobs will scan the tables specified here, including both schemas and tables.
	// For instance, the input could be in the format: [{schemaName: "HR", tableName: ["T1", "T2"]}, {schemaName:
	// "OE", tableName : ["T3", "T4"]}].
	TablesForDiscovery []TablesForDiscovery `mandatory:"false" json:"tablesForDiscovery"`

	// The OCIDs of the sensitive types to be used by data discovery jobs. If OCID of a sensitive category is provided,
	// all its child sensitive types are used for data discovery.
	SensitiveTypeIdsForDiscovery []string `mandatory:"false" json:"sensitiveTypeIdsForDiscovery"`

	// Indicates if data discovery jobs should collect and store sample data values for the discovered columns.
	// Sample data helps review the discovered columns and ensure that they actually contain sensitive data.
	// As it collects original data from the target database, it's disabled by default and should be used only
	// if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data values
	// are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE.
	IsSampleDataCollectionEnabled *bool `mandatory:"false" json:"isSampleDataCollectionEnabled"`

	// Indicates if data discovery jobs should identify potential application-level (non-dictionary) referential relationships
	// between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined) relationships.
	// This option helps identify application-level relationships that are not defined in the database dictionary, which in turn,
	// helps identify additional sensitive columns and preserve referential integrity during data masking. It's disabled by default
	// and should be used only if there is a need to identify application-level relationships.
	IsAppDefinedRelationDiscoveryEnabled *bool `mandatory:"false" json:"isAppDefinedRelationDiscoveryEnabled"`

	// Indicates if all the schemas in the associated target database should be scanned by data discovery jobs.
	// If it is set to true, sensitive data is discovered in all schemas (except for schemas maintained by Oracle).
	IsIncludeAllSchemas *bool `mandatory:"false" json:"isIncludeAllSchemas"`

	// Indicates if all the existing sensitive types should be used by data discovery jobs. If it's set to true,
	// the sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery.
	IsIncludeAllSensitiveTypes *bool `mandatory:"false" json:"isIncludeAllSensitiveTypes"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSensitiveDataModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSensitiveDataModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
