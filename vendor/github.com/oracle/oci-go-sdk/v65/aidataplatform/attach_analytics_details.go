// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachAnalyticsDetails The data to associate Analytics during AiDataPlatform create.
type AttachAnalyticsDetails interface {
}

type attachanalyticsdetails struct {
	JsonData        []byte
	AssociationType string `json:"associationType"`
}

// UnmarshalJSON unmarshals json
func (m *attachanalyticsdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerattachanalyticsdetails attachanalyticsdetails
	s := struct {
		Model Unmarshalerattachanalyticsdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AssociationType = s.Model.AssociationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *attachanalyticsdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AssociationType {
	case "EXISTING_ANALYTICS":
		mm := AnalyticsAssociationByAnalyticsIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOMAIN_ADMIN":
		mm := AnalyticsAssociationByIamDomainDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IDCS_ACCESS_TOKEN_SECRET":
		mm := AnalyticsAssociationByIdcsAccessTokenSecretDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IDCS_ACCESS_TOKEN":
		mm := LegacyAnalyticsAssociationByIdcsAccessTokenDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AttachAnalyticsDetails: %s.", m.AssociationType)
		return *m, nil
	}
}

func (m attachanalyticsdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m attachanalyticsdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttachAnalyticsDetailsAssociationTypeEnum Enum with underlying type: string
type AttachAnalyticsDetailsAssociationTypeEnum string

// Set of constants representing the allowable values for AttachAnalyticsDetailsAssociationTypeEnum
const (
	AttachAnalyticsDetailsAssociationTypeExistingAnalytics     AttachAnalyticsDetailsAssociationTypeEnum = "EXISTING_ANALYTICS"
	AttachAnalyticsDetailsAssociationTypeDomainAdmin           AttachAnalyticsDetailsAssociationTypeEnum = "DOMAIN_ADMIN"
	AttachAnalyticsDetailsAssociationTypeIdcsAccessToken       AttachAnalyticsDetailsAssociationTypeEnum = "IDCS_ACCESS_TOKEN"
	AttachAnalyticsDetailsAssociationTypeIdcsAccessTokenSecret AttachAnalyticsDetailsAssociationTypeEnum = "IDCS_ACCESS_TOKEN_SECRET"
)

var mappingAttachAnalyticsDetailsAssociationTypeEnum = map[string]AttachAnalyticsDetailsAssociationTypeEnum{
	"EXISTING_ANALYTICS":       AttachAnalyticsDetailsAssociationTypeExistingAnalytics,
	"DOMAIN_ADMIN":             AttachAnalyticsDetailsAssociationTypeDomainAdmin,
	"IDCS_ACCESS_TOKEN":        AttachAnalyticsDetailsAssociationTypeIdcsAccessToken,
	"IDCS_ACCESS_TOKEN_SECRET": AttachAnalyticsDetailsAssociationTypeIdcsAccessTokenSecret,
}

var mappingAttachAnalyticsDetailsAssociationTypeEnumLowerCase = map[string]AttachAnalyticsDetailsAssociationTypeEnum{
	"existing_analytics":       AttachAnalyticsDetailsAssociationTypeExistingAnalytics,
	"domain_admin":             AttachAnalyticsDetailsAssociationTypeDomainAdmin,
	"idcs_access_token":        AttachAnalyticsDetailsAssociationTypeIdcsAccessToken,
	"idcs_access_token_secret": AttachAnalyticsDetailsAssociationTypeIdcsAccessTokenSecret,
}

// GetAttachAnalyticsDetailsAssociationTypeEnumValues Enumerates the set of values for AttachAnalyticsDetailsAssociationTypeEnum
func GetAttachAnalyticsDetailsAssociationTypeEnumValues() []AttachAnalyticsDetailsAssociationTypeEnum {
	values := make([]AttachAnalyticsDetailsAssociationTypeEnum, 0)
	for _, v := range mappingAttachAnalyticsDetailsAssociationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttachAnalyticsDetailsAssociationTypeEnumStringValues Enumerates the set of values in String for AttachAnalyticsDetailsAssociationTypeEnum
func GetAttachAnalyticsDetailsAssociationTypeEnumStringValues() []string {
	return []string{
		"EXISTING_ANALYTICS",
		"DOMAIN_ADMIN",
		"IDCS_ACCESS_TOKEN",
		"IDCS_ACCESS_TOKEN_SECRET",
	}
}

// GetMappingAttachAnalyticsDetailsAssociationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttachAnalyticsDetailsAssociationTypeEnum(val string) (AttachAnalyticsDetailsAssociationTypeEnum, bool) {
	enum, ok := mappingAttachAnalyticsDetailsAssociationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
