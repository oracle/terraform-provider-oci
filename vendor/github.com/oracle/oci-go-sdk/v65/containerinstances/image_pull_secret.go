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

// ImagePullSecret The image pull secrets for accessing private registry to pull images for containers
type ImagePullSecret interface {

	// The registry endpoint of the container image.
	GetRegistryEndpoint() *string
}

type imagepullsecret struct {
	JsonData         []byte
	RegistryEndpoint *string `mandatory:"true" json:"registryEndpoint"`
	SecretType       string  `json:"secretType"`
}

// UnmarshalJSON unmarshals json
func (m *imagepullsecret) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimagepullsecret imagepullsecret
	s := struct {
		Model Unmarshalerimagepullsecret
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
func (m *imagepullsecret) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SecretType {
	case "VAULT":
		mm := VaultImagePullSecret{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BASIC":
		mm := BasicImagePullSecret{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ImagePullSecret: %s.", m.SecretType)
		return *m, nil
	}
}

// GetRegistryEndpoint returns RegistryEndpoint
func (m imagepullsecret) GetRegistryEndpoint() *string {
	return m.RegistryEndpoint
}

func (m imagepullsecret) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m imagepullsecret) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImagePullSecretSecretTypeEnum Enum with underlying type: string
type ImagePullSecretSecretTypeEnum string

// Set of constants representing the allowable values for ImagePullSecretSecretTypeEnum
const (
	ImagePullSecretSecretTypeBasic ImagePullSecretSecretTypeEnum = "BASIC"
	ImagePullSecretSecretTypeVault ImagePullSecretSecretTypeEnum = "VAULT"
)

var mappingImagePullSecretSecretTypeEnum = map[string]ImagePullSecretSecretTypeEnum{
	"BASIC": ImagePullSecretSecretTypeBasic,
	"VAULT": ImagePullSecretSecretTypeVault,
}

var mappingImagePullSecretSecretTypeEnumLowerCase = map[string]ImagePullSecretSecretTypeEnum{
	"basic": ImagePullSecretSecretTypeBasic,
	"vault": ImagePullSecretSecretTypeVault,
}

// GetImagePullSecretSecretTypeEnumValues Enumerates the set of values for ImagePullSecretSecretTypeEnum
func GetImagePullSecretSecretTypeEnumValues() []ImagePullSecretSecretTypeEnum {
	values := make([]ImagePullSecretSecretTypeEnum, 0)
	for _, v := range mappingImagePullSecretSecretTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetImagePullSecretSecretTypeEnumStringValues Enumerates the set of values in String for ImagePullSecretSecretTypeEnum
func GetImagePullSecretSecretTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"VAULT",
	}
}

// GetMappingImagePullSecretSecretTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImagePullSecretSecretTypeEnum(val string) (ImagePullSecretSecretTypeEnum, bool) {
	enum, ok := mappingImagePullSecretSecretTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
