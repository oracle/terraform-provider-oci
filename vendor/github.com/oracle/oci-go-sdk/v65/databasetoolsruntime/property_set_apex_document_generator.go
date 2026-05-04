// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PropertySetApexDocumentGenerator Contains the details of an APEX Document Generator property set
type PropertySetApexDocumentGenerator struct {

	// Indicates whether the property set is mutable or not
	IsMutable *bool `mandatory:"true" json:"isMutable"`

	// The name of the credential used by APEX to manage Object Storage Buckets and Objects as well as invoke the Document Generator function
	CredentialKey *string `mandatory:"false" json:"credentialKey"`

	// The Object Storage Namespace containing the Object Storage Buckets managed by APEX
	ObjectStorageNamespace *string `mandatory:"false" json:"objectStorageNamespace"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Object Storage Buckets managed by APEX
	ObjectStorageBucketCompartmentId *string `mandatory:"false" json:"objectStorageBucketCompartmentId"`

	// Object Storage Endpoint
	ObjectStorageEndpoint *string `mandatory:"false" json:"objectStorageEndpoint"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Document Generator function
	FunctionId *string `mandatory:"false" json:"functionId"`

	// The base endpoint URL to use to invoke the Document Generator function
	InvokeEndpoint *string `mandatory:"false" json:"invokeEndpoint"`

	// The print server type
	PrintServerType PropertySetApexDocumentGeneratorPrintServerTypeEnum `mandatory:"false" json:"printServerType,omitempty"`

	// The status of the Autonomous Database Serverless Resource Principal (OCI$RESOURCE_PRINCIPAL)
	AutonomousDatabaseResourcePrincipalStatus PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum `mandatory:"false" json:"autonomousDatabaseResourcePrincipalStatus,omitempty"`
}

// GetIsMutable returns IsMutable
func (m PropertySetApexDocumentGenerator) GetIsMutable() *bool {
	return m.IsMutable
}

func (m PropertySetApexDocumentGenerator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PropertySetApexDocumentGenerator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPropertySetApexDocumentGeneratorPrintServerTypeEnum(string(m.PrintServerType)); !ok && m.PrintServerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrintServerType: %s. Supported values are: %s.", m.PrintServerType, strings.Join(GetPropertySetApexDocumentGeneratorPrintServerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum(string(m.AutonomousDatabaseResourcePrincipalStatus)); !ok && m.AutonomousDatabaseResourcePrincipalStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousDatabaseResourcePrincipalStatus: %s. Supported values are: %s.", m.AutonomousDatabaseResourcePrincipalStatus, strings.Join(GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PropertySetApexDocumentGenerator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePropertySetApexDocumentGenerator PropertySetApexDocumentGenerator
	s := struct {
		DiscriminatorParam string `json:"key"`
		MarshalTypePropertySetApexDocumentGenerator
	}{
		"APEX_DOCUMENT_GENERATOR",
		(MarshalTypePropertySetApexDocumentGenerator)(m),
	}

	return json.Marshal(&s)
}

// PropertySetApexDocumentGeneratorPrintServerTypeEnum Enum with underlying type: string
type PropertySetApexDocumentGeneratorPrintServerTypeEnum string

// Set of constants representing the allowable values for PropertySetApexDocumentGeneratorPrintServerTypeEnum
const (
	PropertySetApexDocumentGeneratorPrintServerTypeDocumentGenerator PropertySetApexDocumentGeneratorPrintServerTypeEnum = "DOCUMENT_GENERATOR"
	PropertySetApexDocumentGeneratorPrintServerTypeNone              PropertySetApexDocumentGeneratorPrintServerTypeEnum = "NONE"
	PropertySetApexDocumentGeneratorPrintServerTypeStandard          PropertySetApexDocumentGeneratorPrintServerTypeEnum = "STANDARD"
	PropertySetApexDocumentGeneratorPrintServerTypeAdvanced          PropertySetApexDocumentGeneratorPrintServerTypeEnum = "ADVANCED"
	PropertySetApexDocumentGeneratorPrintServerTypeAop               PropertySetApexDocumentGeneratorPrintServerTypeEnum = "AOP"
	PropertySetApexDocumentGeneratorPrintServerTypeOther             PropertySetApexDocumentGeneratorPrintServerTypeEnum = "OTHER"
)

var mappingPropertySetApexDocumentGeneratorPrintServerTypeEnum = map[string]PropertySetApexDocumentGeneratorPrintServerTypeEnum{
	"DOCUMENT_GENERATOR": PropertySetApexDocumentGeneratorPrintServerTypeDocumentGenerator,
	"NONE":               PropertySetApexDocumentGeneratorPrintServerTypeNone,
	"STANDARD":           PropertySetApexDocumentGeneratorPrintServerTypeStandard,
	"ADVANCED":           PropertySetApexDocumentGeneratorPrintServerTypeAdvanced,
	"AOP":                PropertySetApexDocumentGeneratorPrintServerTypeAop,
	"OTHER":              PropertySetApexDocumentGeneratorPrintServerTypeOther,
}

var mappingPropertySetApexDocumentGeneratorPrintServerTypeEnumLowerCase = map[string]PropertySetApexDocumentGeneratorPrintServerTypeEnum{
	"document_generator": PropertySetApexDocumentGeneratorPrintServerTypeDocumentGenerator,
	"none":               PropertySetApexDocumentGeneratorPrintServerTypeNone,
	"standard":           PropertySetApexDocumentGeneratorPrintServerTypeStandard,
	"advanced":           PropertySetApexDocumentGeneratorPrintServerTypeAdvanced,
	"aop":                PropertySetApexDocumentGeneratorPrintServerTypeAop,
	"other":              PropertySetApexDocumentGeneratorPrintServerTypeOther,
}

// GetPropertySetApexDocumentGeneratorPrintServerTypeEnumValues Enumerates the set of values for PropertySetApexDocumentGeneratorPrintServerTypeEnum
func GetPropertySetApexDocumentGeneratorPrintServerTypeEnumValues() []PropertySetApexDocumentGeneratorPrintServerTypeEnum {
	values := make([]PropertySetApexDocumentGeneratorPrintServerTypeEnum, 0)
	for _, v := range mappingPropertySetApexDocumentGeneratorPrintServerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertySetApexDocumentGeneratorPrintServerTypeEnumStringValues Enumerates the set of values in String for PropertySetApexDocumentGeneratorPrintServerTypeEnum
func GetPropertySetApexDocumentGeneratorPrintServerTypeEnumStringValues() []string {
	return []string{
		"DOCUMENT_GENERATOR",
		"NONE",
		"STANDARD",
		"ADVANCED",
		"AOP",
		"OTHER",
	}
}

// GetMappingPropertySetApexDocumentGeneratorPrintServerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertySetApexDocumentGeneratorPrintServerTypeEnum(val string) (PropertySetApexDocumentGeneratorPrintServerTypeEnum, bool) {
	enum, ok := mappingPropertySetApexDocumentGeneratorPrintServerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum Enum with underlying type: string
type PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum string

// Set of constants representing the allowable values for PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum
const (
	PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnabled  PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum = "ENABLED"
	PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusDisabled PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum = "DISABLED"
)

var mappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum = map[string]PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum{
	"ENABLED":  PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnabled,
	"DISABLED": PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusDisabled,
}

var mappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumLowerCase = map[string]PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum{
	"enabled":  PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnabled,
	"disabled": PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusDisabled,
}

// GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumValues Enumerates the set of values for PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum
func GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumValues() []PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum {
	values := make([]PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum, 0)
	for _, v := range mappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumStringValues Enumerates the set of values in String for PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum
func GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum(val string) (PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum, bool) {
	enum, ok := mappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
