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

// UpdateSensitiveDataModelDetails Details to update a sensitive data model. Note that updating any attribute of a sensitive data model does not perform data discovery.
type UpdateSensitiveDataModelDetails struct {

	// The display name of the sensitive data model. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the reference target database to be associated with the sensitive data model. All operations such as
	// performing data discovery and adding columns manually are done in the context of the associated target database.
	// Note that updating the targetId attribute does not perform data discovery automatically.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The application suite name identifying a collection of applications. It's useful only if maintaining a sensitive data model for a suite of applications.
	AppSuiteName *string `mandatory:"false" json:"appSuiteName"`

	// The description of the sensitive data model.
	Description *string `mandatory:"false" json:"description"`

	// The schemas to be used for future data discovery jobs.
	SchemasForDiscovery []string `mandatory:"false" json:"schemasForDiscovery"`

	// The OCIDs of the sensitive types to be used for future data discovery jobs. If OCID of a sensitive category is
	// provided, all its child sensitive types are used for data discovery.
	SensitiveTypeIdsForDiscovery []string `mandatory:"false" json:"sensitiveTypeIdsForDiscovery"`

	// Indicates if data discovery jobs should collect and store sample data values for the discovered columns.
	// Sample data helps review the discovered columns and ensure that they actually contain sensitive data.
	// As it collects original data from the target database, it's disabled by default and should be used only
	// if it's acceptable to store sample data in Data Safe's repository in Oracle Cloud. Note that sample data
	// values are not collected for columns with the following data types: LONG, LOB, RAW, XMLTYPE and BFILE.
	IsSampleDataCollectionEnabled *bool `mandatory:"false" json:"isSampleDataCollectionEnabled"`

	// Indicates if data discovery jobs should identify potential application-level (non-dictionary) referential
	// relationships between columns. Note that data discovery automatically identifies and adds database-level
	// (dictionary-defined) relationships. This option helps identify application-level relationships that are not
	// defined in the database dictionary, which in turn, helps identify additional sensitive columns and preserve
	// referential integrity during data masking. It's disabled by default and should be used only if there is a
	// need to identify application-level relationships.
	IsAppDefinedRelationDiscoveryEnabled *bool `mandatory:"false" json:"isAppDefinedRelationDiscoveryEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSensitiveDataModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSensitiveDataModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
