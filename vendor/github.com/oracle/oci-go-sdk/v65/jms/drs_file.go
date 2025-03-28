// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrsFile A Deployment Rule Set(DRS) is a JAR (Java ARchive) file used in Java applications to enforce security and manage compatibility
// between different versions of Java applets and web start applications
// (https://docs.oracle.com/javase/8/docs/technotes/guides/deploy/deployment_rules.html).
type DrsFile struct {

	// The Object Storage bucket name where the DRS file is located.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The namespace for Object Storage.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the DRS file in Object Store.
	DrsFileName *string `mandatory:"true" json:"drsFileName"`

	// The unique identifier of the DRS file in Object Storage.
	DrsFileKey *string `mandatory:"true" json:"drsFileKey"`

	// The checksum type for the DRS file in Object Storage.
	ChecksumType DrsFileChecksumTypeEnum `mandatory:"true" json:"checksumType"`

	// The checksum value for the DRS file in Object Storage.
	ChecksumValue *string `mandatory:"true" json:"checksumValue"`

	// To check if the DRS file is the detfault ones.
	IsDefault *bool `mandatory:"true" json:"isDefault"`
}

func (m DrsFile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrsFile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrsFileChecksumTypeEnum(string(m.ChecksumType)); !ok && m.ChecksumType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChecksumType: %s. Supported values are: %s.", m.ChecksumType, strings.Join(GetDrsFileChecksumTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrsFileChecksumTypeEnum Enum with underlying type: string
type DrsFileChecksumTypeEnum string

// Set of constants representing the allowable values for DrsFileChecksumTypeEnum
const (
	DrsFileChecksumTypeSha256 DrsFileChecksumTypeEnum = "SHA256"
)

var mappingDrsFileChecksumTypeEnum = map[string]DrsFileChecksumTypeEnum{
	"SHA256": DrsFileChecksumTypeSha256,
}

var mappingDrsFileChecksumTypeEnumLowerCase = map[string]DrsFileChecksumTypeEnum{
	"sha256": DrsFileChecksumTypeSha256,
}

// GetDrsFileChecksumTypeEnumValues Enumerates the set of values for DrsFileChecksumTypeEnum
func GetDrsFileChecksumTypeEnumValues() []DrsFileChecksumTypeEnum {
	values := make([]DrsFileChecksumTypeEnum, 0)
	for _, v := range mappingDrsFileChecksumTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrsFileChecksumTypeEnumStringValues Enumerates the set of values in String for DrsFileChecksumTypeEnum
func GetDrsFileChecksumTypeEnumStringValues() []string {
	return []string{
		"SHA256",
	}
}

// GetMappingDrsFileChecksumTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrsFileChecksumTypeEnum(val string) (DrsFileChecksumTypeEnum, bool) {
	enum, ok := mappingDrsFileChecksumTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
