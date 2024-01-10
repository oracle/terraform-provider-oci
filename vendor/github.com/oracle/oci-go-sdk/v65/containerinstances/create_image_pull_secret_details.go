// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateImagePullSecretDetails The image pull secrets for accessing private registry to pull images for containers
type CreateImagePullSecretDetails interface {

	// The registry endpoint of the container image.
	GetRegistryEndpoint() *string
}

type createimagepullsecretdetails struct {
	JsonData         []byte
	RegistryEndpoint *string `mandatory:"true" json:"registryEndpoint"`
	SecretType       string  `json:"secretType"`
}

// UnmarshalJSON unmarshals json
func (m *createimagepullsecretdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateimagepullsecretdetails createimagepullsecretdetails
	s := struct {
		Model Unmarshalercreateimagepullsecretdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RegistryEndpoint = s.Model.RegistryEndpoint
	m.SecretType = s.Model.SecretType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createimagepullsecretdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SecretType {
	case "VAULT":
		mm := CreateVaultImagePullSecretDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BASIC":
		mm := CreateBasicImagePullSecretDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateImagePullSecretDetails: %s.", m.SecretType)
		return *m, nil
	}
}

// GetRegistryEndpoint returns RegistryEndpoint
func (m createimagepullsecretdetails) GetRegistryEndpoint() *string {
	return m.RegistryEndpoint
}

func (m createimagepullsecretdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createimagepullsecretdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateImagePullSecretDetailsSecretTypeEnum Enum with underlying type: string
type CreateImagePullSecretDetailsSecretTypeEnum string

// Set of constants representing the allowable values for CreateImagePullSecretDetailsSecretTypeEnum
const (
	CreateImagePullSecretDetailsSecretTypeBasic CreateImagePullSecretDetailsSecretTypeEnum = "BASIC"
	CreateImagePullSecretDetailsSecretTypeVault CreateImagePullSecretDetailsSecretTypeEnum = "VAULT"
)

var mappingCreateImagePullSecretDetailsSecretTypeEnum = map[string]CreateImagePullSecretDetailsSecretTypeEnum{
	"BASIC": CreateImagePullSecretDetailsSecretTypeBasic,
	"VAULT": CreateImagePullSecretDetailsSecretTypeVault,
}

var mappingCreateImagePullSecretDetailsSecretTypeEnumLowerCase = map[string]CreateImagePullSecretDetailsSecretTypeEnum{
	"basic": CreateImagePullSecretDetailsSecretTypeBasic,
	"vault": CreateImagePullSecretDetailsSecretTypeVault,
}

// GetCreateImagePullSecretDetailsSecretTypeEnumValues Enumerates the set of values for CreateImagePullSecretDetailsSecretTypeEnum
func GetCreateImagePullSecretDetailsSecretTypeEnumValues() []CreateImagePullSecretDetailsSecretTypeEnum {
	values := make([]CreateImagePullSecretDetailsSecretTypeEnum, 0)
	for _, v := range mappingCreateImagePullSecretDetailsSecretTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateImagePullSecretDetailsSecretTypeEnumStringValues Enumerates the set of values in String for CreateImagePullSecretDetailsSecretTypeEnum
func GetCreateImagePullSecretDetailsSecretTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"VAULT",
	}
}

// GetMappingCreateImagePullSecretDetailsSecretTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateImagePullSecretDetailsSecretTypeEnum(val string) (CreateImagePullSecretDetailsSecretTypeEnum, bool) {
	enum, ok := mappingCreateImagePullSecretDetailsSecretTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
