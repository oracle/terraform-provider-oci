// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceProviderConfigCategoryDetails Defines Source Provider details.
type SourceProviderConfigCategoryDetails struct {

	// URL of the remote repository.
	Url *string `mandatory:"true" json:"url"`

	Secret SourceProviderCredentialDetails `mandatory:"true" json:"secret"`

	// Authentication type to authenticate with Remote repository.
	AuthenticationType SourceProviderConfigCategoryDetailsAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`
}

func (m SourceProviderConfigCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceProviderConfigCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetSourceProviderConfigCategoryDetailsAuthenticationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SourceProviderConfigCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceProviderConfigCategoryDetails SourceProviderConfigCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"configCategory"`
		MarshalTypeSourceProviderConfigCategoryDetails
	}{
		"SOURCE_PROVIDER",
		(MarshalTypeSourceProviderConfigCategoryDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SourceProviderConfigCategoryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AuthenticationType SourceProviderConfigCategoryDetailsAuthenticationTypeEnum `json:"authenticationType"`
		Url                *string                                                   `json:"url"`
		Secret             sourceprovidercredentialdetails                           `json:"secret"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AuthenticationType = model.AuthenticationType

	m.Url = model.Url

	nn, e = model.Secret.UnmarshalPolymorphicJSON(model.Secret.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Secret = nn.(SourceProviderCredentialDetails)
	} else {
		m.Secret = nil
	}

	return
}

// SourceProviderConfigCategoryDetailsAuthenticationTypeEnum Enum with underlying type: string
type SourceProviderConfigCategoryDetailsAuthenticationTypeEnum string

// Set of constants representing the allowable values for SourceProviderConfigCategoryDetailsAuthenticationTypeEnum
const (
	SourceProviderConfigCategoryDetailsAuthenticationTypeKeyBased   SourceProviderConfigCategoryDetailsAuthenticationTypeEnum = "KEY_BASED"
	SourceProviderConfigCategoryDetailsAuthenticationTypeTokenBased SourceProviderConfigCategoryDetailsAuthenticationTypeEnum = "TOKEN_BASED"
)

var mappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnum = map[string]SourceProviderConfigCategoryDetailsAuthenticationTypeEnum{
	"KEY_BASED":   SourceProviderConfigCategoryDetailsAuthenticationTypeKeyBased,
	"TOKEN_BASED": SourceProviderConfigCategoryDetailsAuthenticationTypeTokenBased,
}

var mappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnumLowerCase = map[string]SourceProviderConfigCategoryDetailsAuthenticationTypeEnum{
	"key_based":   SourceProviderConfigCategoryDetailsAuthenticationTypeKeyBased,
	"token_based": SourceProviderConfigCategoryDetailsAuthenticationTypeTokenBased,
}

// GetSourceProviderConfigCategoryDetailsAuthenticationTypeEnumValues Enumerates the set of values for SourceProviderConfigCategoryDetailsAuthenticationTypeEnum
func GetSourceProviderConfigCategoryDetailsAuthenticationTypeEnumValues() []SourceProviderConfigCategoryDetailsAuthenticationTypeEnum {
	values := make([]SourceProviderConfigCategoryDetailsAuthenticationTypeEnum, 0)
	for _, v := range mappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceProviderConfigCategoryDetailsAuthenticationTypeEnumStringValues Enumerates the set of values in String for SourceProviderConfigCategoryDetailsAuthenticationTypeEnum
func GetSourceProviderConfigCategoryDetailsAuthenticationTypeEnumStringValues() []string {
	return []string{
		"KEY_BASED",
		"TOKEN_BASED",
	}
}

// GetMappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnum(val string) (SourceProviderConfigCategoryDetailsAuthenticationTypeEnum, bool) {
	enum, ok := mappingSourceProviderConfigCategoryDetailsAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
