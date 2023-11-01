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

// CreateDiscoveryJobDetails Details to create a new data discovery job.
type CreateDiscoveryJobDetails struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The OCID of the compartment where the discovery job resource should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the discovery job. It defines the job's scope.
	// NEW identifies new sensitive columns in the target database that are not in the sensitive data model.
	// DELETED identifies columns that are present in the sensitive data model but have been deleted from the target database.
	// MODIFIED identifies columns that are present in the target database as well as the sensitive data model but some of their attributes have been modified.
	// ALL covers all the above three scenarios and reports new, deleted and modified columns.
	DiscoveryType DiscoveryJobDiscoveryTypeEnum `mandatory:"false" json:"discoveryType,omitempty"`

	// A user-friendly name for the discovery job. Does not have to be unique, and it is changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The schemas to be scanned by the discovery job. If not provided, the schemasForDiscovery attribute of the sensitive
	// data model is used to get the list of schemas.
	SchemasForDiscovery []string `mandatory:"false" json:"schemasForDiscovery"`

	// The OCIDs of the sensitive types to be used by the discovery job. If not provided, the sensitiveTypeIdsForDiscovery
	// attribute of the sensitive data model is used to get the list of sensitive types.
	SensitiveTypeIdsForDiscovery []string `mandatory:"false" json:"sensitiveTypeIdsForDiscovery"`

	// Indicates if the discovery job should collect and store sample data values for the discovered columns. Sample data
	// helps review the discovered columns and ensure that they actually contain sensitive data. As it collects original
	// data from the target database, it's disabled by default and should be used only if it's acceptable to store sample
	// data in Data Safe's repository in Oracle Cloud. Note that sample data values are not collected for columns with the
	// following data types: LONG, LOB, RAW, XMLTYPE and BFILE.
	IsSampleDataCollectionEnabled *bool `mandatory:"false" json:"isSampleDataCollectionEnabled"`

	// Indicates if the discovery job should identify potential application-level (non-dictionary) referential relationships
	// between columns. Note that data discovery automatically identifies and adds database-level (dictionary-defined)
	// relationships. This option helps identify application-level relationships that are not defined in the database
	// dictionary, which in turn, helps identify additional sensitive columns and preserve referential integrity during
	// data masking. It's disabled by default and should be used only if there is a need to identify application-level
	// relationships.
	IsAppDefinedRelationDiscoveryEnabled *bool `mandatory:"false" json:"isAppDefinedRelationDiscoveryEnabled"`

	// Indicates if all the schemas should be scanned by the discovery job. If it is set to true, sensitive data is discovered
	// in all schemas (except for schemas maintained by Oracle). If both attributes are not provided, the configuration
	// from the sensitive data model is used.
	IsIncludeAllSchemas *bool `mandatory:"false" json:"isIncludeAllSchemas"`

	// Indicates if all the existing sensitive types should be used by the discovery job. If it's set to true, the
	// sensitiveTypeIdsForDiscovery attribute is ignored and all sensitive types are used for data discovery. If both
	// attributes are not provided, the configuration from the sensitive data model is used.
	IsIncludeAllSensitiveTypes *bool `mandatory:"false" json:"isIncludeAllSensitiveTypes"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDiscoveryJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDiscoveryJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveryJobDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoveryJobDiscoveryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
