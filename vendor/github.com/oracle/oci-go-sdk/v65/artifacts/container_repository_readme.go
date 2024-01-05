// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// API covering the Artifacts and Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as generic artifacts and container images.
//

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerRepositoryReadme Container repository readme.
type ContainerRepositoryReadme struct {

	// Readme content. Avoid entering confidential information.
	Content *string `mandatory:"true" json:"content"`

	// Readme format. Supported formats are text/plain and text/markdown.
	Format ContainerRepositoryReadmeFormatEnum `mandatory:"true" json:"format"`
}

func (m ContainerRepositoryReadme) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerRepositoryReadme) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerRepositoryReadmeFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetContainerRepositoryReadmeFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerRepositoryReadmeFormatEnum Enum with underlying type: string
type ContainerRepositoryReadmeFormatEnum string

// Set of constants representing the allowable values for ContainerRepositoryReadmeFormatEnum
const (
	ContainerRepositoryReadmeFormatMarkdown ContainerRepositoryReadmeFormatEnum = "TEXT_MARKDOWN"
	ContainerRepositoryReadmeFormatPlain    ContainerRepositoryReadmeFormatEnum = "TEXT_PLAIN"
)

var mappingContainerRepositoryReadmeFormatEnum = map[string]ContainerRepositoryReadmeFormatEnum{
	"TEXT_MARKDOWN": ContainerRepositoryReadmeFormatMarkdown,
	"TEXT_PLAIN":    ContainerRepositoryReadmeFormatPlain,
}

var mappingContainerRepositoryReadmeFormatEnumLowerCase = map[string]ContainerRepositoryReadmeFormatEnum{
	"text_markdown": ContainerRepositoryReadmeFormatMarkdown,
	"text_plain":    ContainerRepositoryReadmeFormatPlain,
}

// GetContainerRepositoryReadmeFormatEnumValues Enumerates the set of values for ContainerRepositoryReadmeFormatEnum
func GetContainerRepositoryReadmeFormatEnumValues() []ContainerRepositoryReadmeFormatEnum {
	values := make([]ContainerRepositoryReadmeFormatEnum, 0)
	for _, v := range mappingContainerRepositoryReadmeFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerRepositoryReadmeFormatEnumStringValues Enumerates the set of values in String for ContainerRepositoryReadmeFormatEnum
func GetContainerRepositoryReadmeFormatEnumStringValues() []string {
	return []string{
		"TEXT_MARKDOWN",
		"TEXT_PLAIN",
	}
}

// GetMappingContainerRepositoryReadmeFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerRepositoryReadmeFormatEnum(val string) (ContainerRepositoryReadmeFormatEnum, bool) {
	enum, ok := mappingContainerRepositoryReadmeFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
