// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BucketDetailsDefinition Information about a particular bucket.
type BucketDetailsDefinition struct {

	// Type of bucket. SENDER pipelines can be SOURCE, TRANSFER, REJECT, or FAILED. RECEIVER pipelines have a DESTINATION bucket.
	BucketType BucketDetailsDefinitionBucketTypeEnum `mandatory:"true" json:"bucketType"`

	// Namespace of the bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Name of the bucket.
	Name *string `mandatory:"true" json:"name"`

	// OCID of the bucket.
	Id *string `mandatory:"true" json:"id"`
}

func (m BucketDetailsDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BucketDetailsDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBucketDetailsDefinitionBucketTypeEnum(string(m.BucketType)); !ok && m.BucketType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BucketType: %s. Supported values are: %s.", m.BucketType, strings.Join(GetBucketDetailsDefinitionBucketTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BucketDetailsDefinitionBucketTypeEnum Enum with underlying type: string
type BucketDetailsDefinitionBucketTypeEnum string

// Set of constants representing the allowable values for BucketDetailsDefinitionBucketTypeEnum
const (
	BucketDetailsDefinitionBucketTypeSource      BucketDetailsDefinitionBucketTypeEnum = "SOURCE"
	BucketDetailsDefinitionBucketTypeTransfer    BucketDetailsDefinitionBucketTypeEnum = "TRANSFER"
	BucketDetailsDefinitionBucketTypeReject      BucketDetailsDefinitionBucketTypeEnum = "REJECT"
	BucketDetailsDefinitionBucketTypeDestination BucketDetailsDefinitionBucketTypeEnum = "DESTINATION"
)

var mappingBucketDetailsDefinitionBucketTypeEnum = map[string]BucketDetailsDefinitionBucketTypeEnum{
	"SOURCE":      BucketDetailsDefinitionBucketTypeSource,
	"TRANSFER":    BucketDetailsDefinitionBucketTypeTransfer,
	"REJECT":      BucketDetailsDefinitionBucketTypeReject,
	"DESTINATION": BucketDetailsDefinitionBucketTypeDestination,
}

var mappingBucketDetailsDefinitionBucketTypeEnumLowerCase = map[string]BucketDetailsDefinitionBucketTypeEnum{
	"source":      BucketDetailsDefinitionBucketTypeSource,
	"transfer":    BucketDetailsDefinitionBucketTypeTransfer,
	"reject":      BucketDetailsDefinitionBucketTypeReject,
	"destination": BucketDetailsDefinitionBucketTypeDestination,
}

// GetBucketDetailsDefinitionBucketTypeEnumValues Enumerates the set of values for BucketDetailsDefinitionBucketTypeEnum
func GetBucketDetailsDefinitionBucketTypeEnumValues() []BucketDetailsDefinitionBucketTypeEnum {
	values := make([]BucketDetailsDefinitionBucketTypeEnum, 0)
	for _, v := range mappingBucketDetailsDefinitionBucketTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBucketDetailsDefinitionBucketTypeEnumStringValues Enumerates the set of values in String for BucketDetailsDefinitionBucketTypeEnum
func GetBucketDetailsDefinitionBucketTypeEnumStringValues() []string {
	return []string{
		"SOURCE",
		"TRANSFER",
		"REJECT",
		"DESTINATION",
	}
}

// GetMappingBucketDetailsDefinitionBucketTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBucketDetailsDefinitionBucketTypeEnum(val string) (BucketDetailsDefinitionBucketTypeEnum, bool) {
	enum, ok := mappingBucketDetailsDefinitionBucketTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
