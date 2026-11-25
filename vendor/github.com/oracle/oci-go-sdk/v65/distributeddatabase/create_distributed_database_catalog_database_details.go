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

// CreateDistributedDatabaseCatalogDatabaseDetails Details for creating a distributed database catalog.
type CreateDistributedDatabaseCatalogDatabaseDetails interface {

	// The password to open the TDE wallet.
	// The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters.
	// The allowed special characters are _, \#, and -.
	GetTdeWalletPassword() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI vault secret. This cannot be used in conjunction with tdeWalletPassword.
	GetTdeWalletPasswordSecretId() *string

	// The version of the vault secret. If no version is specified, the latest version will be used.
	GetTdeWalletPasswordSecretVersionNumber() *int
}

type createdistributeddatabasecatalogdatabasedetails struct {
	JsonData                             []byte
	TdeWalletPassword                    *string `mandatory:"false" json:"tdeWalletPassword" sensitive:"true"`
	TdeWalletPasswordSecretId            *string `mandatory:"false" json:"tdeWalletPasswordSecretId"`
	TdeWalletPasswordSecretVersionNumber *int    `mandatory:"false" json:"tdeWalletPasswordSecretVersionNumber"`
	Source                               string  `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributeddatabasecatalogdatabasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributeddatabasecatalogdatabasedetails createdistributeddatabasecatalogdatabasedetails
	s := struct {
		Model Unmarshalercreatedistributeddatabasecatalogdatabasedetails
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
func (m *createdistributeddatabasecatalogdatabasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "XS_NEW_VAULT_AND_CLUSTER":
		mm := CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXISTING_DB_HOME":
		mm := CreateDistributedDatabaseCatalogDatabaseWithDbHomeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_EXISTING_CLUSTER":
		mm := CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XD_EXISTING_CLUSTER":
		mm := CreateDistributedDatabaseCatalogDatabaseWithExadbXdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_NEW_CLUSTER":
		mm := CreateDistributedDatabaseCatalogWithExadbXsNewClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedDatabaseCatalogDatabaseDetails: %s.", m.Source)
		return *m, nil
	}
}

// GetTdeWalletPassword returns TdeWalletPassword
func (m createdistributeddatabasecatalogdatabasedetails) GetTdeWalletPassword() *string {
	return m.TdeWalletPassword
}

// GetTdeWalletPasswordSecretId returns TdeWalletPasswordSecretId
func (m createdistributeddatabasecatalogdatabasedetails) GetTdeWalletPasswordSecretId() *string {
	return m.TdeWalletPasswordSecretId
}

// GetTdeWalletPasswordSecretVersionNumber returns TdeWalletPasswordSecretVersionNumber
func (m createdistributeddatabasecatalogdatabasedetails) GetTdeWalletPasswordSecretVersionNumber() *int {
	return m.TdeWalletPasswordSecretVersionNumber
}

func (m createdistributeddatabasecatalogdatabasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributeddatabasecatalogdatabasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum Enum with underlying type: string
type CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum
const (
	CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsExistingCluster    CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = "XS_EXISTING_CLUSTER"
	CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewCluster         CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = "XS_NEW_CLUSTER"
	CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewVaultAndCluster CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = "XS_NEW_VAULT_AND_CLUSTER"
	CreateDistributedDatabaseCatalogDatabaseDetailsSourceXdExistingCluster    CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = "XD_EXISTING_CLUSTER"
	CreateDistributedDatabaseCatalogDatabaseDetailsSourceExistingDbHome       CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = "EXISTING_DB_HOME"
)

var mappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum = map[string]CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum{
	"XS_EXISTING_CLUSTER":      CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsExistingCluster,
	"XS_NEW_CLUSTER":           CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewCluster,
	"XS_NEW_VAULT_AND_CLUSTER": CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewVaultAndCluster,
	"XD_EXISTING_CLUSTER":      CreateDistributedDatabaseCatalogDatabaseDetailsSourceXdExistingCluster,
	"EXISTING_DB_HOME":         CreateDistributedDatabaseCatalogDatabaseDetailsSourceExistingDbHome,
}

var mappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumLowerCase = map[string]CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum{
	"xs_existing_cluster":      CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsExistingCluster,
	"xs_new_cluster":           CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewCluster,
	"xs_new_vault_and_cluster": CreateDistributedDatabaseCatalogDatabaseDetailsSourceXsNewVaultAndCluster,
	"xd_existing_cluster":      CreateDistributedDatabaseCatalogDatabaseDetailsSourceXdExistingCluster,
	"existing_db_home":         CreateDistributedDatabaseCatalogDatabaseDetailsSourceExistingDbHome,
}

// GetCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumValues Enumerates the set of values for CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum
func GetCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumValues() []CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum {
	values := make([]CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum
func GetCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumStringValues() []string {
	return []string{
		"XS_EXISTING_CLUSTER",
		"XS_NEW_CLUSTER",
		"XS_NEW_VAULT_AND_CLUSTER",
		"XD_EXISTING_CLUSTER",
		"EXISTING_DB_HOME",
	}
}

// GetMappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum(val string) (CreateDistributedDatabaseCatalogDatabaseDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseCatalogDatabaseDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
