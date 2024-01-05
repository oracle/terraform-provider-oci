// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserPhotos URLs of photos for the User
type UserPhotos struct {

	// URL of a photo for the User
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// A label indicating the attribute's function; e.g., 'photo' or 'thumbnail'.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type UserPhotosTypeEnum `mandatory:"true" json:"type"`

	// A human readable name, primarily used for display purposes.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred photo or thumbnail. The primary attribute value 'true' MUST appear no more than once.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Primary *bool `mandatory:"false" json:"primary"`
}

func (m UserPhotos) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserPhotos) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserPhotosTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserPhotosTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserPhotosTypeEnum Enum with underlying type: string
type UserPhotosTypeEnum string

// Set of constants representing the allowable values for UserPhotosTypeEnum
const (
	UserPhotosTypePhoto     UserPhotosTypeEnum = "photo"
	UserPhotosTypeThumbnail UserPhotosTypeEnum = "thumbnail"
)

var mappingUserPhotosTypeEnum = map[string]UserPhotosTypeEnum{
	"photo":     UserPhotosTypePhoto,
	"thumbnail": UserPhotosTypeThumbnail,
}

var mappingUserPhotosTypeEnumLowerCase = map[string]UserPhotosTypeEnum{
	"photo":     UserPhotosTypePhoto,
	"thumbnail": UserPhotosTypeThumbnail,
}

// GetUserPhotosTypeEnumValues Enumerates the set of values for UserPhotosTypeEnum
func GetUserPhotosTypeEnumValues() []UserPhotosTypeEnum {
	values := make([]UserPhotosTypeEnum, 0)
	for _, v := range mappingUserPhotosTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserPhotosTypeEnumStringValues Enumerates the set of values in String for UserPhotosTypeEnum
func GetUserPhotosTypeEnumStringValues() []string {
	return []string{
		"photo",
		"thumbnail",
	}
}

// GetMappingUserPhotosTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserPhotosTypeEnum(val string) (UserPhotosTypeEnum, bool) {
	enum, ok := mappingUserPhotosTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
