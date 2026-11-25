// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedDatabaseShardDatabaseDetails Details for creating a distributed database shard.
type CreateDistributedDatabaseShardDatabaseDetails interface {

	// The password to open the TDE wallet.
	// The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters.
	// The allowed special characters are _, \#, and -.
	GetTdeWalletPassword() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI vault secret. This cannot be used in conjunction with tdeWalletPassword.
	GetTdeWalletPasswordSecretId() *string

	// The version of the vault secret. If no version is specified, the latest version will be used.
	GetTdeWalletPasswordSecretVersionNumber() *int
}

type createdistributeddatabasesharddatabasedetails struct {
	JsonData                             []byte
	TdeWalletPassword                    *string `mandatory:"false" json:"tdeWalletPassword" sensitive:"true"`
	TdeWalletPasswordSecretId            *string `mandatory:"false" json:"tdeWalletPasswordSecretId"`
	TdeWalletPasswordSecretVersionNumber *int    `mandatory:"false" json:"tdeWalletPasswordSecretVersionNumber"`
	Source                               string  `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributeddatabasesharddatabasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributeddatabasesharddatabasedetails createdistributeddatabasesharddatabasedetails
	s := struct {
		Model Unmarshalercreatedistributeddatabasesharddatabasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TdeWalletPassword = s.Model.TdeWalletPassword
	m.TdeWalletPasswordSecretId = s.Model.TdeWalletPasswordSecretId
	m.TdeWalletPasswordSecretVersionNumber = s.Model.TdeWalletPasswordSecretVersionNumber
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributeddatabasesharddatabasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "XS_NEW_CLUSTER":
		mm := CreateDistributedDatabaseShardWithExadbXsNewClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_NEW_VAULT_AND_CLUSTER":
		mm := CreateDistributedDatabaseShardWithExadbXsNewVaultAndClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_EXISTING_CLUSTER":
		mm := CreateDistributedDatabaseShardDatabaseWithExadbXsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXISTING_DB_HOME":
		mm := CreateDistributedDatabaseShardDatabaseWithDbHomeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XD_EXISTING_CLUSTER":
		mm := CreateDistributedDatabaseShardDatabaseWithExadbXdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedDatabaseShardDatabaseDetails: %s.", m.Source)
		return *m, nil
	}
}

// GetTdeWalletPassword returns TdeWalletPassword
func (m createdistributeddatabasesharddatabasedetails) GetTdeWalletPassword() *string {
	return m.TdeWalletPassword
}

// GetTdeWalletPasswordSecretId returns TdeWalletPasswordSecretId
func (m createdistributeddatabasesharddatabasedetails) GetTdeWalletPasswordSecretId() *string {
	return m.TdeWalletPasswordSecretId
}

// GetTdeWalletPasswordSecretVersionNumber returns TdeWalletPasswordSecretVersionNumber
func (m createdistributeddatabasesharddatabasedetails) GetTdeWalletPasswordSecretVersionNumber() *int {
	return m.TdeWalletPasswordSecretVersionNumber
}

func (m createdistributeddatabasesharddatabasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributeddatabasesharddatabasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedDatabaseShardDatabaseDetailsSourceEnum Enum with underlying type: string
type CreateDistributedDatabaseShardDatabaseDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseShardDatabaseDetailsSourceEnum
const (
	CreateDistributedDatabaseShardDatabaseDetailsSourceXsExistingCluster    CreateDistributedDatabaseShardDatabaseDetailsSourceEnum = "XS_EXISTING_CLUSTER"
	CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewCluster         CreateDistributedDatabaseShardDatabaseDetailsSourceEnum = "XS_NEW_CLUSTER"
	CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewVaultAndCluster CreateDistributedDatabaseShardDatabaseDetailsSourceEnum = "XS_NEW_VAULT_AND_CLUSTER"
	CreateDistributedDatabaseShardDatabaseDetailsSourceXdExistingCluster    CreateDistributedDatabaseShardDatabaseDetailsSourceEnum = "XD_EXISTING_CLUSTER"
	CreateDistributedDatabaseShardDatabaseDetailsSourceExistingDbHome       CreateDistributedDatabaseShardDatabaseDetailsSourceEnum = "EXISTING_DB_HOME"
)

var mappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnum = map[string]CreateDistributedDatabaseShardDatabaseDetailsSourceEnum{
	"XS_EXISTING_CLUSTER":      CreateDistributedDatabaseShardDatabaseDetailsSourceXsExistingCluster,
	"XS_NEW_CLUSTER":           CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewCluster,
	"XS_NEW_VAULT_AND_CLUSTER": CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewVaultAndCluster,
	"XD_EXISTING_CLUSTER":      CreateDistributedDatabaseShardDatabaseDetailsSourceXdExistingCluster,
	"EXISTING_DB_HOME":         CreateDistributedDatabaseShardDatabaseDetailsSourceExistingDbHome,
}

var mappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnumLowerCase = map[string]CreateDistributedDatabaseShardDatabaseDetailsSourceEnum{
	"xs_existing_cluster":      CreateDistributedDatabaseShardDatabaseDetailsSourceXsExistingCluster,
	"xs_new_cluster":           CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewCluster,
	"xs_new_vault_and_cluster": CreateDistributedDatabaseShardDatabaseDetailsSourceXsNewVaultAndCluster,
	"xd_existing_cluster":      CreateDistributedDatabaseShardDatabaseDetailsSourceXdExistingCluster,
	"existing_db_home":         CreateDistributedDatabaseShardDatabaseDetailsSourceExistingDbHome,
}

// GetCreateDistributedDatabaseShardDatabaseDetailsSourceEnumValues Enumerates the set of values for CreateDistributedDatabaseShardDatabaseDetailsSourceEnum
func GetCreateDistributedDatabaseShardDatabaseDetailsSourceEnumValues() []CreateDistributedDatabaseShardDatabaseDetailsSourceEnum {
	values := make([]CreateDistributedDatabaseShardDatabaseDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseShardDatabaseDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseShardDatabaseDetailsSourceEnum
func GetCreateDistributedDatabaseShardDatabaseDetailsSourceEnumStringValues() []string {
	return []string{
		"XS_EXISTING_CLUSTER",
		"XS_NEW_CLUSTER",
		"XS_NEW_VAULT_AND_CLUSTER",
		"XD_EXISTING_CLUSTER",
		"EXISTING_DB_HOME",
	}
}

// GetMappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnum(val string) (CreateDistributedDatabaseShardDatabaseDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseShardDatabaseDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
