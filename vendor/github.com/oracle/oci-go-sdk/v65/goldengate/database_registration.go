// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseRegistration Represents the metadata description of a database used by deployments in the same compartment.
type DatabaseRegistration struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the databaseRegistration being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"true" json:"fqdn"`

	// The private IP address in the customer's VCN of the customer's endpoint, typically a
	// database.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"true" json:"username"`

	// Credential store alias.
	AliasName *string `mandatory:"true" json:"aliasName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// A Private Endpoint IP address created in the customer's subnet.
	// A customer database can expect network traffic initiated by GoldenGate Service from this IP address.
	// It can also send network traffic to this IP address, typically in response to requests from GoldenGate Service.
	// The customer may use this IP address in Security Lists or Network Security Groups (NSG) as needed.
	RcePrivateIp *string `mandatory:"false" json:"rcePrivateIp"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Connect descriptor or Easy Connect Naming method used to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The mode of the database connection session to be established by the data client.
	// 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database.
	// Connection to a RAC database involves a redirection received from the SCAN listeners
	// to the database node to connect to. By default the mode would be DIRECT.
	SessionMode DatabaseRegistrationSessionModeEnum `mandatory:"false" json:"sessionMode,omitempty"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// The OCID of the compartment where the GoldenGate Secret will be created.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	SecretCompartmentId *string `mandatory:"false" json:"secretCompartmentId"`

	// The OCID of the customer's GoldenGate Service Secret.
	// If provided, it references a key that customers will be required to ensure the policies are established
	// to permit GoldenGate to use this Secret.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m DatabaseRegistration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseRegistration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseRegistrationSessionModeEnum(string(m.SessionMode)); !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetDatabaseRegistrationSessionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseRegistrationSessionModeEnum Enum with underlying type: string
type DatabaseRegistrationSessionModeEnum string

// Set of constants representing the allowable values for DatabaseRegistrationSessionModeEnum
const (
	DatabaseRegistrationSessionModeDirect   DatabaseRegistrationSessionModeEnum = "DIRECT"
	DatabaseRegistrationSessionModeRedirect DatabaseRegistrationSessionModeEnum = "REDIRECT"
)

var mappingDatabaseRegistrationSessionModeEnum = map[string]DatabaseRegistrationSessionModeEnum{
	"DIRECT":   DatabaseRegistrationSessionModeDirect,
	"REDIRECT": DatabaseRegistrationSessionModeRedirect,
}

var mappingDatabaseRegistrationSessionModeEnumLowerCase = map[string]DatabaseRegistrationSessionModeEnum{
	"direct":   DatabaseRegistrationSessionModeDirect,
	"redirect": DatabaseRegistrationSessionModeRedirect,
}

// GetDatabaseRegistrationSessionModeEnumValues Enumerates the set of values for DatabaseRegistrationSessionModeEnum
func GetDatabaseRegistrationSessionModeEnumValues() []DatabaseRegistrationSessionModeEnum {
	values := make([]DatabaseRegistrationSessionModeEnum, 0)
	for _, v := range mappingDatabaseRegistrationSessionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseRegistrationSessionModeEnumStringValues Enumerates the set of values in String for DatabaseRegistrationSessionModeEnum
func GetDatabaseRegistrationSessionModeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"REDIRECT",
	}
}

// GetMappingDatabaseRegistrationSessionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseRegistrationSessionModeEnum(val string) (DatabaseRegistrationSessionModeEnum, bool) {
	enum, ok := mappingDatabaseRegistrationSessionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
