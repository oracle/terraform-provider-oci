// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciIamAuthenticationDetails OCI IAM authentication details.
type OciIamAuthenticationDetails struct {

	// The database schema of Oracle Autonomous AI Database.
	DatabaseSchema *string `mandatory:"true" json:"databaseSchema"`

	// External authentication domain name.
	IamDomainName *string `mandatory:"true" json:"iamDomainName"`

	// Authenticated User or Group name.
	IamName *string `mandatory:"true" json:"iamName"`

	// Privileges granted to database schema.
	UserRoles []string `mandatory:"true" json:"userRoles"`

	// Indicates identity type is user or group.
	IdentityType OciIamAuthenticationDetailsIdentityTypeEnum `mandatory:"true" json:"identityType"`
}

func (m OciIamAuthenticationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciIamAuthenticationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciIamAuthenticationDetailsIdentityTypeEnum(string(m.IdentityType)); !ok && m.IdentityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdentityType: %s. Supported values are: %s.", m.IdentityType, strings.Join(GetOciIamAuthenticationDetailsIdentityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciIamAuthenticationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciIamAuthenticationDetails OciIamAuthenticationDetails
	s := struct {
		DiscriminatorParam string `json:"method"`
		MarshalTypeOciIamAuthenticationDetails
	}{
		"OCI_IAM",
		(MarshalTypeOciIamAuthenticationDetails)(m),
	}

	return json.Marshal(&s)
}

// OciIamAuthenticationDetailsIdentityTypeEnum Enum with underlying type: string
type OciIamAuthenticationDetailsIdentityTypeEnum string

// Set of constants representing the allowable values for OciIamAuthenticationDetailsIdentityTypeEnum
const (
	OciIamAuthenticationDetailsIdentityTypeUser  OciIamAuthenticationDetailsIdentityTypeEnum = "USER"
	OciIamAuthenticationDetailsIdentityTypeGroup OciIamAuthenticationDetailsIdentityTypeEnum = "GROUP"
)

var mappingOciIamAuthenticationDetailsIdentityTypeEnum = map[string]OciIamAuthenticationDetailsIdentityTypeEnum{
	"USER":  OciIamAuthenticationDetailsIdentityTypeUser,
	"GROUP": OciIamAuthenticationDetailsIdentityTypeGroup,
}

var mappingOciIamAuthenticationDetailsIdentityTypeEnumLowerCase = map[string]OciIamAuthenticationDetailsIdentityTypeEnum{
	"user":  OciIamAuthenticationDetailsIdentityTypeUser,
	"group": OciIamAuthenticationDetailsIdentityTypeGroup,
}

// GetOciIamAuthenticationDetailsIdentityTypeEnumValues Enumerates the set of values for OciIamAuthenticationDetailsIdentityTypeEnum
func GetOciIamAuthenticationDetailsIdentityTypeEnumValues() []OciIamAuthenticationDetailsIdentityTypeEnum {
	values := make([]OciIamAuthenticationDetailsIdentityTypeEnum, 0)
	for _, v := range mappingOciIamAuthenticationDetailsIdentityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciIamAuthenticationDetailsIdentityTypeEnumStringValues Enumerates the set of values in String for OciIamAuthenticationDetailsIdentityTypeEnum
func GetOciIamAuthenticationDetailsIdentityTypeEnumStringValues() []string {
	return []string{
		"USER",
		"GROUP",
	}
}

// GetMappingOciIamAuthenticationDetailsIdentityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciIamAuthenticationDetailsIdentityTypeEnum(val string) (OciIamAuthenticationDetailsIdentityTypeEnum, bool) {
	enum, ok := mappingOciIamAuthenticationDetailsIdentityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
