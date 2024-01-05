// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourcePrincipalAuthConfig Authentication configuration that uses OCI Resource Principal Auth for Generic REST invocation.
type ResourcePrincipalAuthConfig struct {

	// Generated key that can be used in API calls to identify this object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The OCI resource type that will supply the authentication token
	ResourcePrincipalSource ResourcePrincipalAuthConfigResourcePrincipalSourceEnum `mandatory:"false" json:"resourcePrincipalSource,omitempty"`
}

// GetKey returns Key
func (m ResourcePrincipalAuthConfig) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m ResourcePrincipalAuthConfig) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ResourcePrincipalAuthConfig) GetParentRef() *ParentReference {
	return m.ParentRef
}

func (m ResourcePrincipalAuthConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcePrincipalAuthConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourcePrincipalAuthConfigResourcePrincipalSourceEnum(string(m.ResourcePrincipalSource)); !ok && m.ResourcePrincipalSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourcePrincipalSource: %s. Supported values are: %s.", m.ResourcePrincipalSource, strings.Join(GetResourcePrincipalAuthConfigResourcePrincipalSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ResourcePrincipalAuthConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeResourcePrincipalAuthConfig ResourcePrincipalAuthConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeResourcePrincipalAuthConfig
	}{
		"OCI_RESOURCE_AUTH_CONFIG",
		(MarshalTypeResourcePrincipalAuthConfig)(m),
	}

	return json.Marshal(&s)
}

// ResourcePrincipalAuthConfigResourcePrincipalSourceEnum Enum with underlying type: string
type ResourcePrincipalAuthConfigResourcePrincipalSourceEnum string

// Set of constants representing the allowable values for ResourcePrincipalAuthConfigResourcePrincipalSourceEnum
const (
	ResourcePrincipalAuthConfigResourcePrincipalSourceWorkspace   ResourcePrincipalAuthConfigResourcePrincipalSourceEnum = "WORKSPACE"
	ResourcePrincipalAuthConfigResourcePrincipalSourceApplication ResourcePrincipalAuthConfigResourcePrincipalSourceEnum = "APPLICATION"
)

var mappingResourcePrincipalAuthConfigResourcePrincipalSourceEnum = map[string]ResourcePrincipalAuthConfigResourcePrincipalSourceEnum{
	"WORKSPACE":   ResourcePrincipalAuthConfigResourcePrincipalSourceWorkspace,
	"APPLICATION": ResourcePrincipalAuthConfigResourcePrincipalSourceApplication,
}

var mappingResourcePrincipalAuthConfigResourcePrincipalSourceEnumLowerCase = map[string]ResourcePrincipalAuthConfigResourcePrincipalSourceEnum{
	"workspace":   ResourcePrincipalAuthConfigResourcePrincipalSourceWorkspace,
	"application": ResourcePrincipalAuthConfigResourcePrincipalSourceApplication,
}

// GetResourcePrincipalAuthConfigResourcePrincipalSourceEnumValues Enumerates the set of values for ResourcePrincipalAuthConfigResourcePrincipalSourceEnum
func GetResourcePrincipalAuthConfigResourcePrincipalSourceEnumValues() []ResourcePrincipalAuthConfigResourcePrincipalSourceEnum {
	values := make([]ResourcePrincipalAuthConfigResourcePrincipalSourceEnum, 0)
	for _, v := range mappingResourcePrincipalAuthConfigResourcePrincipalSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetResourcePrincipalAuthConfigResourcePrincipalSourceEnumStringValues Enumerates the set of values in String for ResourcePrincipalAuthConfigResourcePrincipalSourceEnum
func GetResourcePrincipalAuthConfigResourcePrincipalSourceEnumStringValues() []string {
	return []string{
		"WORKSPACE",
		"APPLICATION",
	}
}

// GetMappingResourcePrincipalAuthConfigResourcePrincipalSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourcePrincipalAuthConfigResourcePrincipalSourceEnum(val string) (ResourcePrincipalAuthConfigResourcePrincipalSourceEnum, bool) {
	enum, ok := mappingResourcePrincipalAuthConfigResourcePrincipalSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
