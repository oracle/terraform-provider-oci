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

// DifferenceColumnSummary Summary of a SDM masking policy difference column.
type DifferenceColumnSummary struct {

	// The unique key that identifies the SDM masking policy difference column.
	Key *string `mandatory:"true" json:"key"`

	// The type of the SDM masking policy difference column. It can be one of the following three types:
	// NEW: A new sensitive column in the sensitive data model that is not in the masking policy.
	// DELETED: A column that is present in the masking policy but has been deleted from the sensitive data model.
	// MODIFIED: A column that is present in the masking policy as well as the sensitive data model but some of its attributes have been modified.
	DifferenceType DifferenceColumnDifferenceTypeEnum `mandatory:"true" json:"differenceType"`

	// The database schema that contains the difference column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The database object that contains the difference column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the difference column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// Specifies how to process the difference column. It's set to SYNC by default. Use the PatchSdmMaskingPolicyDifferenceColumns operation to update this attribute. You can choose one of the following options:
	// SYNC: To sync the difference column and update the masking policy to reflect the changes.
	// NO_SYNC: To not sync the difference column so that it doesn't change the masking policy.
	// After specifying the planned action, you can use the ApplySdmMaskingPolicyDifference operation to automatically process the difference columns.
	PlannedAction DifferenceColumnPlannedActionEnum `mandatory:"true" json:"plannedAction"`

	// Indicates if the difference column has been processed.Use GetDifferenceColumn operation to track whether the difference column has
	// already been processed and applied to the masking policy.
	SyncStatus DifferenceColumnSyncStatusEnum `mandatory:"true" json:"syncStatus"`

	// The unique key that identifies the sensitive column represented by the SDM masking policy difference column.
	SensitiveColumnkey *string `mandatory:"false" json:"sensitiveColumnkey"`

	// The unique key that identifies the masking column represented by the SDM masking policy difference column.
	MaskingColumnkey *string `mandatory:"false" json:"maskingColumnkey"`

	// The OCID of the sensitive type associated with the difference column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// The date and time the SDM masking policy difference column was last synced, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastSynced *common.SDKTime `mandatory:"false" json:"timeLastSynced"`
}

func (m DifferenceColumnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DifferenceColumnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDifferenceColumnDifferenceTypeEnum(string(m.DifferenceType)); !ok && m.DifferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceType: %s. Supported values are: %s.", m.DifferenceType, strings.Join(GetDifferenceColumnDifferenceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDifferenceColumnPlannedActionEnum(string(m.PlannedAction)); !ok && m.PlannedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlannedAction: %s. Supported values are: %s.", m.PlannedAction, strings.Join(GetDifferenceColumnPlannedActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDifferenceColumnSyncStatusEnum(string(m.SyncStatus)); !ok && m.SyncStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyncStatus: %s. Supported values are: %s.", m.SyncStatus, strings.Join(GetDifferenceColumnSyncStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
