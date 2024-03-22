// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HelmCommandSpecArtifactSource Specifies Helm command spec details
type HelmCommandSpecArtifactSource struct {

	// The Helm commands to be executed, base 64 encoded
	Base64EncodedContent *string `mandatory:"true" json:"base64EncodedContent"`

	// Specifies types of artifact sources.
	HelmArtifactSourceType HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum `mandatory:"true" json:"helmArtifactSourceType"`
}

func (m HelmCommandSpecArtifactSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HelmCommandSpecArtifactSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum(string(m.HelmArtifactSourceType)); !ok && m.HelmArtifactSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HelmArtifactSourceType: %s. Supported values are: %s.", m.HelmArtifactSourceType, strings.Join(GetHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HelmCommandSpecArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHelmCommandSpecArtifactSource HelmCommandSpecArtifactSource
	s := struct {
		DiscriminatorParam string `json:"deployArtifactSourceType"`
		MarshalTypeHelmCommandSpecArtifactSource
	}{
		"HELM_COMMAND_SPEC",
		(MarshalTypeHelmCommandSpecArtifactSource)(m),
	}

	return json.Marshal(&s)
}

// HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum Enum with underlying type: string
type HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum string

// Set of constants representing the allowable values for HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum
const (
	HelmCommandSpecArtifactSourceHelmArtifactSourceTypeInline HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum = "INLINE"
)

var mappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum = map[string]HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum{
	"INLINE": HelmCommandSpecArtifactSourceHelmArtifactSourceTypeInline,
}

var mappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumLowerCase = map[string]HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum{
	"inline": HelmCommandSpecArtifactSourceHelmArtifactSourceTypeInline,
}

// GetHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumValues Enumerates the set of values for HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum
func GetHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumValues() []HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum {
	values := make([]HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum, 0)
	for _, v := range mappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumStringValues Enumerates the set of values in String for HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum
func GetHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumStringValues() []string {
	return []string{
		"INLINE",
	}
}

// GetMappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum(val string) (HelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnum, bool) {
	enum, ok := mappingHelmCommandSpecArtifactSourceHelmArtifactSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
