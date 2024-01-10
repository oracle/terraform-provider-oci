// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateProtectedDatabaseDetails Describes the parameters required to update a protected database.
type UpdateProtectedDatabaseDetails struct {

	// The protected database name. You can change the displayName. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The size of the database is allowed to be decreased. XS - Less than 5GB, S - 5GB to 50GB, M - 50GB to 500GB, L - 500GB to 1TB, XL - 1TB to 5TB, XXL - Greater than 5TB.
	DatabaseSize DatabaseSizesEnum `mandatory:"false" json:"databaseSize,omitempty"`

	// The size of the database, in gigabytes.
	DatabaseSizeInGBs *int `mandatory:"false" json:"databaseSizeInGBs"`

	// Password credential which can be used to connect to Protected Database.
	// It must contain at least 2 uppercase, 2 lowercase, 2 numeric and 2 special characters.
	// The special characters must be underscore (_), number sign (#) or hyphen (-). The password must not contain the username "admin", regardless of casing.
	// Password must not be same as current passsword.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the protection policy associated with the protected database.
	ProtectionPolicyId *string `mandatory:"false" json:"protectionPolicyId"`

	// List of recovery service subnet resources associated with the protected database.
	RecoveryServiceSubnets []RecoveryServiceSubnetInput `mandatory:"false" json:"recoveryServiceSubnets"`

	// The value TRUE indicates that the protected database is configured to use Real-time data protection, and redo-data is sent from the protected database to Recovery Service.
	// Real-time data protection substantially reduces the window of potential data loss that exists between successive archived redo log backups. For this to be effective, additional
	// configuration is needed on client side.
	IsRedoLogsShipped *bool `mandatory:"false" json:"isRedoLogsShipped"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateProtectedDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateProtectedDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseSizesEnum(string(m.DatabaseSize)); !ok && m.DatabaseSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSize: %s. Supported values are: %s.", m.DatabaseSize, strings.Join(GetDatabaseSizesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
