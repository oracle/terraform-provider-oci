// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// IamWorkRequestSummary The IAM work request summary. Tracks the status of the asynchronous operations.
type IamWorkRequestSummary struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The asynchronous operation tracked by this IAM work request.
	OperationType IamWorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the work request
	Status IamWorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the compartment containing this IAM work request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The resources this work request affects.
	Resources []IamWorkRequestResource `mandatory:"false" json:"resources"`

	// How much progress the operation has made.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Date and time the work was accepted, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Date and time the work started, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Date and time the work completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m IamWorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IamWorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIamWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetIamWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIamWorkRequestSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetIamWorkRequestSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IamWorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type IamWorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for IamWorkRequestSummaryOperationTypeEnum
const (
	IamWorkRequestSummaryOperationTypeCreateDomain               IamWorkRequestSummaryOperationTypeEnum = "CREATE_DOMAIN"
	IamWorkRequestSummaryOperationTypeReplicateDomainToRegion    IamWorkRequestSummaryOperationTypeEnum = "REPLICATE_DOMAIN_TO_REGION"
	IamWorkRequestSummaryOperationTypeUpdateDomain               IamWorkRequestSummaryOperationTypeEnum = "UPDATE_DOMAIN"
	IamWorkRequestSummaryOperationTypeActivateDomain             IamWorkRequestSummaryOperationTypeEnum = "ACTIVATE_DOMAIN"
	IamWorkRequestSummaryOperationTypeDeactivateDomain           IamWorkRequestSummaryOperationTypeEnum = "DEACTIVATE_DOMAIN"
	IamWorkRequestSummaryOperationTypeDeleteDomain               IamWorkRequestSummaryOperationTypeEnum = "DELETE_DOMAIN"
	IamWorkRequestSummaryOperationTypeChangeCompartmentForDomain IamWorkRequestSummaryOperationTypeEnum = "CHANGE_COMPARTMENT_FOR_DOMAIN"
	IamWorkRequestSummaryOperationTypeChangeLicenseTypeForDomain IamWorkRequestSummaryOperationTypeEnum = "CHANGE_LICENSE_TYPE_FOR_DOMAIN"
)

var mappingIamWorkRequestSummaryOperationTypeEnum = map[string]IamWorkRequestSummaryOperationTypeEnum{
	"CREATE_DOMAIN":                  IamWorkRequestSummaryOperationTypeCreateDomain,
	"REPLICATE_DOMAIN_TO_REGION":     IamWorkRequestSummaryOperationTypeReplicateDomainToRegion,
	"UPDATE_DOMAIN":                  IamWorkRequestSummaryOperationTypeUpdateDomain,
	"ACTIVATE_DOMAIN":                IamWorkRequestSummaryOperationTypeActivateDomain,
	"DEACTIVATE_DOMAIN":              IamWorkRequestSummaryOperationTypeDeactivateDomain,
	"DELETE_DOMAIN":                  IamWorkRequestSummaryOperationTypeDeleteDomain,
	"CHANGE_COMPARTMENT_FOR_DOMAIN":  IamWorkRequestSummaryOperationTypeChangeCompartmentForDomain,
	"CHANGE_LICENSE_TYPE_FOR_DOMAIN": IamWorkRequestSummaryOperationTypeChangeLicenseTypeForDomain,
}

// GetIamWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for IamWorkRequestSummaryOperationTypeEnum
func GetIamWorkRequestSummaryOperationTypeEnumValues() []IamWorkRequestSummaryOperationTypeEnum {
	values := make([]IamWorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingIamWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for IamWorkRequestSummaryOperationTypeEnum
func GetIamWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DOMAIN",
		"REPLICATE_DOMAIN_TO_REGION",
		"UPDATE_DOMAIN",
		"ACTIVATE_DOMAIN",
		"DEACTIVATE_DOMAIN",
		"DELETE_DOMAIN",
		"CHANGE_COMPARTMENT_FOR_DOMAIN",
		"CHANGE_LICENSE_TYPE_FOR_DOMAIN",
	}
}

// GetMappingIamWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIamWorkRequestSummaryOperationTypeEnum(val string) (IamWorkRequestSummaryOperationTypeEnum, bool) {
	mappingIamWorkRequestSummaryOperationTypeEnumIgnoreCase := make(map[string]IamWorkRequestSummaryOperationTypeEnum)
	for k, v := range mappingIamWorkRequestSummaryOperationTypeEnum {
		mappingIamWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIamWorkRequestSummaryOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// IamWorkRequestSummaryStatusEnum Enum with underlying type: string
type IamWorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for IamWorkRequestSummaryStatusEnum
const (
	IamWorkRequestSummaryStatusAccepted   IamWorkRequestSummaryStatusEnum = "ACCEPTED"
	IamWorkRequestSummaryStatusInProgress IamWorkRequestSummaryStatusEnum = "IN_PROGRESS"
	IamWorkRequestSummaryStatusFailed     IamWorkRequestSummaryStatusEnum = "FAILED"
	IamWorkRequestSummaryStatusSucceeded  IamWorkRequestSummaryStatusEnum = "SUCCEEDED"
	IamWorkRequestSummaryStatusCanceling  IamWorkRequestSummaryStatusEnum = "CANCELING"
	IamWorkRequestSummaryStatusCanceled   IamWorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingIamWorkRequestSummaryStatusEnum = map[string]IamWorkRequestSummaryStatusEnum{
	"ACCEPTED":    IamWorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": IamWorkRequestSummaryStatusInProgress,
	"FAILED":      IamWorkRequestSummaryStatusFailed,
	"SUCCEEDED":   IamWorkRequestSummaryStatusSucceeded,
	"CANCELING":   IamWorkRequestSummaryStatusCanceling,
	"CANCELED":    IamWorkRequestSummaryStatusCanceled,
}

// GetIamWorkRequestSummaryStatusEnumValues Enumerates the set of values for IamWorkRequestSummaryStatusEnum
func GetIamWorkRequestSummaryStatusEnumValues() []IamWorkRequestSummaryStatusEnum {
	values := make([]IamWorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingIamWorkRequestSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestSummaryStatusEnumStringValues Enumerates the set of values in String for IamWorkRequestSummaryStatusEnum
func GetIamWorkRequestSummaryStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingIamWorkRequestSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIamWorkRequestSummaryStatusEnum(val string) (IamWorkRequestSummaryStatusEnum, bool) {
	mappingIamWorkRequestSummaryStatusEnumIgnoreCase := make(map[string]IamWorkRequestSummaryStatusEnum)
	for k, v := range mappingIamWorkRequestSummaryStatusEnum {
		mappingIamWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIamWorkRequestSummaryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
