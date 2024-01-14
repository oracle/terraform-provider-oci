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

// DatabaseViewAccessEntry A DatabaseViewAccessEntry object is a resource corresponding to a row in view authorization report.
// It's a subresource of Security Policy Report resource and is always associated with a SecurityPolicyReport.
type DatabaseViewAccessEntry struct {

	// The unique key that identifies the table access report. It is numeric and unique within a security policy report.
	Key *string `mandatory:"true" json:"key"`

	// Grantee is the user who can access the table or view
	Grantee *string `mandatory:"true" json:"grantee"`

	// The type of the access the user has on the table, there can be one or more from SELECT, UPDATE, INSERT or DELETE.
	AccessType AccessTypeEnum `mandatory:"false" json:"accessType,omitempty"`

	// The name of the schema the table belongs to.
	TableSchema *string `mandatory:"false" json:"tableSchema"`

	// The name of the database table the user has access to.
	TableName *string `mandatory:"false" json:"tableName"`

	// Type of the privilege user has, this includes System Privilege, Schema Privilege, Object Privilege, Column Privilege,
	// Owner or Schema Privilege on a schema.
	PrivilegeType *string `mandatory:"false" json:"privilegeType"`

	// The OCID of the of the  target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The name of the privilege.
	Privilege PrivilegeNameEnum `mandatory:"false" json:"privilege,omitempty"`

	// Indicates whether the grantee can grant this privilege to other users. Privileges can be granted to a user or role with
	// GRANT_OPTION or ADMIN_OPTION
	PrivilegeGrantable PrivilegeGrantableOptionEnum `mandatory:"false" json:"privilegeGrantable,omitempty"`

	// This can be empty in case of direct grant, in case of indirect grant, this attribute displays the name of the
	// role which is granted to the user though which the user has access to the table.
	GrantFromRole *string `mandatory:"false" json:"grantFromRole"`

	// The name of the schema.
	ViewSchema *string `mandatory:"false" json:"viewSchema"`

	// The name of the view.
	ViewName *string `mandatory:"false" json:"viewName"`

	// Definition of the view.
	ViewText *string `mandatory:"false" json:"viewText"`

	// If there are column level privileges on a table or view.
	ColumnName *string `mandatory:"false" json:"columnName"`

	// The user who granted the privilege.
	Grantor *string `mandatory:"false" json:"grantor"`

	// Indicates whether the table access is constrained via Oracle Database Vault.
	IsAccessConstrainedByDatabaseVault *bool `mandatory:"false" json:"isAccessConstrainedByDatabaseVault"`

	// Indicates whether the view access is constrained via Virtual Private Database.
	IsAccessConstrainedByVirtualPrivateDatabase *bool `mandatory:"false" json:"isAccessConstrainedByVirtualPrivateDatabase"`

	// Indicates whether the view access is constrained via Oracle Data Redaction.
	IsAccessConstrainedByRedaction *bool `mandatory:"false" json:"isAccessConstrainedByRedaction"`

	// Indicates whether the view access is constrained via Real Application Security.
	IsAccessConstrainedByRealApplicationSecurity *bool `mandatory:"false" json:"isAccessConstrainedByRealApplicationSecurity"`

	// Indicates whether the view access is constrained via Oracle Database SQL Firewall.
	IsAccessConstrainedBySqlFirewall *bool `mandatory:"false" json:"isAccessConstrainedBySqlFirewall"`
}

func (m DatabaseViewAccessEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseViewAccessEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAccessTypeEnum(string(m.AccessType)); !ok && m.AccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessType: %s. Supported values are: %s.", m.AccessType, strings.Join(GetAccessTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivilegeNameEnum(string(m.Privilege)); !ok && m.Privilege != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Privilege: %s. Supported values are: %s.", m.Privilege, strings.Join(GetPrivilegeNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivilegeGrantableOptionEnum(string(m.PrivilegeGrantable)); !ok && m.PrivilegeGrantable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivilegeGrantable: %s. Supported values are: %s.", m.PrivilegeGrantable, strings.Join(GetPrivilegeGrantableOptionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
