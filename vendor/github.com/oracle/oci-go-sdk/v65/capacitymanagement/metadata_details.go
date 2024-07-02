// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetadataDetails Used for representing the metadata of the catalog. This denotes the version and format of the CSV file for parsing.
type MetadataDetails struct {

	// The version for the format of the catalog file being uploaded.
	FormatVersion MetadataDetailsFormatVersionEnum `mandatory:"true" json:"formatVersion"`
}

func (m MetadataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetadataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetadataDetailsFormatVersionEnum(string(m.FormatVersion)); !ok && m.FormatVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FormatVersion: %s. Supported values are: %s.", m.FormatVersion, strings.Join(GetMetadataDetailsFormatVersionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetadataDetailsFormatVersionEnum Enum with underlying type: string
type MetadataDetailsFormatVersionEnum string

// Set of constants representing the allowable values for MetadataDetailsFormatVersionEnum
const (
	MetadataDetailsFormatVersionV1 MetadataDetailsFormatVersionEnum = "V1"
	MetadataDetailsFormatVersionV2 MetadataDetailsFormatVersionEnum = "V2"
)

var mappingMetadataDetailsFormatVersionEnum = map[string]MetadataDetailsFormatVersionEnum{
	"V1": MetadataDetailsFormatVersionV1,
	"V2": MetadataDetailsFormatVersionV2,
}

var mappingMetadataDetailsFormatVersionEnumLowerCase = map[string]MetadataDetailsFormatVersionEnum{
	"v1": MetadataDetailsFormatVersionV1,
	"v2": MetadataDetailsFormatVersionV2,
}

// GetMetadataDetailsFormatVersionEnumValues Enumerates the set of values for MetadataDetailsFormatVersionEnum
func GetMetadataDetailsFormatVersionEnumValues() []MetadataDetailsFormatVersionEnum {
	values := make([]MetadataDetailsFormatVersionEnum, 0)
	for _, v := range mappingMetadataDetailsFormatVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetMetadataDetailsFormatVersionEnumStringValues Enumerates the set of values in String for MetadataDetailsFormatVersionEnum
func GetMetadataDetailsFormatVersionEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
	}
}

// GetMappingMetadataDetailsFormatVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetadataDetailsFormatVersionEnum(val string) (MetadataDetailsFormatVersionEnum, bool) {
	enum, ok := mappingMetadataDetailsFormatVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
