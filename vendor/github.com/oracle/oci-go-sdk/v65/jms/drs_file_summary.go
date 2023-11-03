// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrsFileSummary A Deployment Rule Set(DRS) is a JAR (Java ARchive) file used in Java applications to enforce security and manage compatibility
// between different versions of Java applets and web start applications
// (https://docs.oracle.com/javase/8/docs/technotes/guides/deploy/deployment_rules.html).
type DrsFileSummary struct {

	// The Object Storage bucket name where the DRS file is located.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The namespace for Object Storage.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the DRS file in Object Store.
	DrsFileName *string `mandatory:"true" json:"drsFileName"`

	// The unique identifier of the DRS file in Object Storage.
	DrsFileKey *string `mandatory:"true" json:"drsFileKey"`

	// The checksum type for the DRS file in Object Storage.
	ChecksumType DrsFileSummaryChecksumTypeEnum `mandatory:"true" json:"checksumType"`

	// The checksum value for the DRS file in Object Storage.
	ChecksumValue *string `mandatory:"true" json:"checksumValue"`

	// To check if the DRS file is the detfault ones.
	IsDefault *bool `mandatory:"true" json:"isDefault"`
}

func (m DrsFileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrsFileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrsFileSummaryChecksumTypeEnum(string(m.ChecksumType)); !ok && m.ChecksumType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChecksumType: %s. Supported values are: %s.", m.ChecksumType, strings.Join(GetDrsFileSummaryChecksumTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrsFileSummaryChecksumTypeEnum Enum with underlying type: string
type DrsFileSummaryChecksumTypeEnum string

// Set of constants representing the allowable values for DrsFileSummaryChecksumTypeEnum
const (
	DrsFileSummaryChecksumTypeSha256 DrsFileSummaryChecksumTypeEnum = "SHA256"
)

var mappingDrsFileSummaryChecksumTypeEnum = map[string]DrsFileSummaryChecksumTypeEnum{
	"SHA256": DrsFileSummaryChecksumTypeSha256,
}

var mappingDrsFileSummaryChecksumTypeEnumLowerCase = map[string]DrsFileSummaryChecksumTypeEnum{
	"sha256": DrsFileSummaryChecksumTypeSha256,
}

// GetDrsFileSummaryChecksumTypeEnumValues Enumerates the set of values for DrsFileSummaryChecksumTypeEnum
func GetDrsFileSummaryChecksumTypeEnumValues() []DrsFileSummaryChecksumTypeEnum {
	values := make([]DrsFileSummaryChecksumTypeEnum, 0)
	for _, v := range mappingDrsFileSummaryChecksumTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrsFileSummaryChecksumTypeEnumStringValues Enumerates the set of values in String for DrsFileSummaryChecksumTypeEnum
func GetDrsFileSummaryChecksumTypeEnumStringValues() []string {
	return []string{
		"SHA256",
	}
}

// GetMappingDrsFileSummaryChecksumTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrsFileSummaryChecksumTypeEnum(val string) (DrsFileSummaryChecksumTypeEnum, bool) {
	enum, ok := mappingDrsFileSummaryChecksumTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
