// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVanityDomainActivityDetails Schedule activity details to vanity domain on pod
type CreateVanityDomainActivityDetails struct {

	// Vanity domain ID
	VanityDomainId *string `mandatory:"true" json:"vanityDomainId"`

	// Activity start time
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// The type of operation
	OperationType CreateVanityDomainActivityDetailsOperationTypeEnum `mandatory:"true" json:"operationType"`
}

func (m CreateVanityDomainActivityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVanityDomainActivityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateVanityDomainActivityDetailsOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetCreateVanityDomainActivityDetailsOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateVanityDomainActivityDetailsOperationTypeEnum Enum with underlying type: string
type CreateVanityDomainActivityDetailsOperationTypeEnum string

// Set of constants representing the allowable values for CreateVanityDomainActivityDetailsOperationTypeEnum
const (
	CreateVanityDomainActivityDetailsOperationTypeEnableVanityDomain CreateVanityDomainActivityDetailsOperationTypeEnum = "ENABLE_VANITY_DOMAIN"
	CreateVanityDomainActivityDetailsOperationTypeDeleteVanityDomain CreateVanityDomainActivityDetailsOperationTypeEnum = "DELETE_VANITY_DOMAIN"
)

var mappingCreateVanityDomainActivityDetailsOperationTypeEnum = map[string]CreateVanityDomainActivityDetailsOperationTypeEnum{
	"ENABLE_VANITY_DOMAIN": CreateVanityDomainActivityDetailsOperationTypeEnableVanityDomain,
	"DELETE_VANITY_DOMAIN": CreateVanityDomainActivityDetailsOperationTypeDeleteVanityDomain,
}

var mappingCreateVanityDomainActivityDetailsOperationTypeEnumLowerCase = map[string]CreateVanityDomainActivityDetailsOperationTypeEnum{
	"enable_vanity_domain": CreateVanityDomainActivityDetailsOperationTypeEnableVanityDomain,
	"delete_vanity_domain": CreateVanityDomainActivityDetailsOperationTypeDeleteVanityDomain,
}

// GetCreateVanityDomainActivityDetailsOperationTypeEnumValues Enumerates the set of values for CreateVanityDomainActivityDetailsOperationTypeEnum
func GetCreateVanityDomainActivityDetailsOperationTypeEnumValues() []CreateVanityDomainActivityDetailsOperationTypeEnum {
	values := make([]CreateVanityDomainActivityDetailsOperationTypeEnum, 0)
	for _, v := range mappingCreateVanityDomainActivityDetailsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVanityDomainActivityDetailsOperationTypeEnumStringValues Enumerates the set of values in String for CreateVanityDomainActivityDetailsOperationTypeEnum
func GetCreateVanityDomainActivityDetailsOperationTypeEnumStringValues() []string {
	return []string{
		"ENABLE_VANITY_DOMAIN",
		"DELETE_VANITY_DOMAIN",
	}
}

// GetMappingCreateVanityDomainActivityDetailsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVanityDomainActivityDetailsOperationTypeEnum(val string) (CreateVanityDomainActivityDetailsOperationTypeEnum, bool) {
	enum, ok := mappingCreateVanityDomainActivityDetailsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
