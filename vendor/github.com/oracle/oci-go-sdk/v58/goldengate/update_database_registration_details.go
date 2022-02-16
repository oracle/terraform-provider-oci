// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateDatabaseRegistrationDetails The information to update for a DatabaseRegistration.
type UpdateDatabaseRegistrationDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`

	// Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The mode of the database connection session to be established by the data client. REDIRECT - for a RAC database, DIRECT - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT.
	SessionMode UpdateDatabaseRegistrationDetailsSessionModeEnum `mandatory:"false" json:"sessionMode,omitempty"`

	// The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded.
	Wallet *string `mandatory:"false" json:"wallet"`

	// Credential store alias.
	AliasName *string `mandatory:"false" json:"aliasName"`
}

func (m UpdateDatabaseRegistrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseRegistrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDatabaseRegistrationDetailsSessionModeEnum(string(m.SessionMode)); !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetUpdateDatabaseRegistrationDetailsSessionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseRegistrationDetailsSessionModeEnum Enum with underlying type: string
type UpdateDatabaseRegistrationDetailsSessionModeEnum string

// Set of constants representing the allowable values for UpdateDatabaseRegistrationDetailsSessionModeEnum
const (
	UpdateDatabaseRegistrationDetailsSessionModeDirect   UpdateDatabaseRegistrationDetailsSessionModeEnum = "DIRECT"
	UpdateDatabaseRegistrationDetailsSessionModeRedirect UpdateDatabaseRegistrationDetailsSessionModeEnum = "REDIRECT"
)

var mappingUpdateDatabaseRegistrationDetailsSessionModeEnum = map[string]UpdateDatabaseRegistrationDetailsSessionModeEnum{
	"DIRECT":   UpdateDatabaseRegistrationDetailsSessionModeDirect,
	"REDIRECT": UpdateDatabaseRegistrationDetailsSessionModeRedirect,
}

// GetUpdateDatabaseRegistrationDetailsSessionModeEnumValues Enumerates the set of values for UpdateDatabaseRegistrationDetailsSessionModeEnum
func GetUpdateDatabaseRegistrationDetailsSessionModeEnumValues() []UpdateDatabaseRegistrationDetailsSessionModeEnum {
	values := make([]UpdateDatabaseRegistrationDetailsSessionModeEnum, 0)
	for _, v := range mappingUpdateDatabaseRegistrationDetailsSessionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseRegistrationDetailsSessionModeEnumStringValues Enumerates the set of values in String for UpdateDatabaseRegistrationDetailsSessionModeEnum
func GetUpdateDatabaseRegistrationDetailsSessionModeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"REDIRECT",
	}
}

// GetMappingUpdateDatabaseRegistrationDetailsSessionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseRegistrationDetailsSessionModeEnum(val string) (UpdateDatabaseRegistrationDetailsSessionModeEnum, bool) {
	mappingUpdateDatabaseRegistrationDetailsSessionModeEnumIgnoreCase := make(map[string]UpdateDatabaseRegistrationDetailsSessionModeEnum)
	for k, v := range mappingUpdateDatabaseRegistrationDetailsSessionModeEnum {
		mappingUpdateDatabaseRegistrationDetailsSessionModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateDatabaseRegistrationDetailsSessionModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
