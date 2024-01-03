// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IamWorkRequest (For tenancies that support identity domains) An IAM work request object that allows users to track the status of asynchronous API requests.
type IamWorkRequest struct {

	// The OCID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The asynchronous operation tracked by this IAM work request.
	OperationType IamWorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the work request.
	Status IamWorkRequestStatusEnum `mandatory:"true" json:"status"`

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

func (m IamWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IamWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIamWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetIamWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIamWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetIamWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IamWorkRequestOperationTypeEnum Enum with underlying type: string
type IamWorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for IamWorkRequestOperationTypeEnum
const (
	IamWorkRequestOperationTypeCreateDomain               IamWorkRequestOperationTypeEnum = "CREATE_DOMAIN"
	IamWorkRequestOperationTypeReplicateDomainToRegion    IamWorkRequestOperationTypeEnum = "REPLICATE_DOMAIN_TO_REGION"
	IamWorkRequestOperationTypeUpdateDomain               IamWorkRequestOperationTypeEnum = "UPDATE_DOMAIN"
	IamWorkRequestOperationTypeActivateDomain             IamWorkRequestOperationTypeEnum = "ACTIVATE_DOMAIN"
	IamWorkRequestOperationTypeDeactivateDomain           IamWorkRequestOperationTypeEnum = "DEACTIVATE_DOMAIN"
	IamWorkRequestOperationTypeDeleteDomain               IamWorkRequestOperationTypeEnum = "DELETE_DOMAIN"
	IamWorkRequestOperationTypeChangeCompartmentForDomain IamWorkRequestOperationTypeEnum = "CHANGE_COMPARTMENT_FOR_DOMAIN"
	IamWorkRequestOperationTypeChangeLicenseTypeForDomain IamWorkRequestOperationTypeEnum = "CHANGE_LICENSE_TYPE_FOR_DOMAIN"
)

var mappingIamWorkRequestOperationTypeEnum = map[string]IamWorkRequestOperationTypeEnum{
	"CREATE_DOMAIN":                  IamWorkRequestOperationTypeCreateDomain,
	"REPLICATE_DOMAIN_TO_REGION":     IamWorkRequestOperationTypeReplicateDomainToRegion,
	"UPDATE_DOMAIN":                  IamWorkRequestOperationTypeUpdateDomain,
	"ACTIVATE_DOMAIN":                IamWorkRequestOperationTypeActivateDomain,
	"DEACTIVATE_DOMAIN":              IamWorkRequestOperationTypeDeactivateDomain,
	"DELETE_DOMAIN":                  IamWorkRequestOperationTypeDeleteDomain,
	"CHANGE_COMPARTMENT_FOR_DOMAIN":  IamWorkRequestOperationTypeChangeCompartmentForDomain,
	"CHANGE_LICENSE_TYPE_FOR_DOMAIN": IamWorkRequestOperationTypeChangeLicenseTypeForDomain,
}

var mappingIamWorkRequestOperationTypeEnumLowerCase = map[string]IamWorkRequestOperationTypeEnum{
	"create_domain":                  IamWorkRequestOperationTypeCreateDomain,
	"replicate_domain_to_region":     IamWorkRequestOperationTypeReplicateDomainToRegion,
	"update_domain":                  IamWorkRequestOperationTypeUpdateDomain,
	"activate_domain":                IamWorkRequestOperationTypeActivateDomain,
	"deactivate_domain":              IamWorkRequestOperationTypeDeactivateDomain,
	"delete_domain":                  IamWorkRequestOperationTypeDeleteDomain,
	"change_compartment_for_domain":  IamWorkRequestOperationTypeChangeCompartmentForDomain,
	"change_license_type_for_domain": IamWorkRequestOperationTypeChangeLicenseTypeForDomain,
}

// GetIamWorkRequestOperationTypeEnumValues Enumerates the set of values for IamWorkRequestOperationTypeEnum
func GetIamWorkRequestOperationTypeEnumValues() []IamWorkRequestOperationTypeEnum {
	values := make([]IamWorkRequestOperationTypeEnum, 0)
	for _, v := range mappingIamWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for IamWorkRequestOperationTypeEnum
func GetIamWorkRequestOperationTypeEnumStringValues() []string {
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

// GetMappingIamWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIamWorkRequestOperationTypeEnum(val string) (IamWorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingIamWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IamWorkRequestStatusEnum Enum with underlying type: string
type IamWorkRequestStatusEnum string

// Set of constants representing the allowable values for IamWorkRequestStatusEnum
const (
	IamWorkRequestStatusAccepted   IamWorkRequestStatusEnum = "ACCEPTED"
	IamWorkRequestStatusInProgress IamWorkRequestStatusEnum = "IN_PROGRESS"
	IamWorkRequestStatusFailed     IamWorkRequestStatusEnum = "FAILED"
	IamWorkRequestStatusSucceeded  IamWorkRequestStatusEnum = "SUCCEEDED"
	IamWorkRequestStatusCanceling  IamWorkRequestStatusEnum = "CANCELING"
	IamWorkRequestStatusCanceled   IamWorkRequestStatusEnum = "CANCELED"
)

var mappingIamWorkRequestStatusEnum = map[string]IamWorkRequestStatusEnum{
	"ACCEPTED":    IamWorkRequestStatusAccepted,
	"IN_PROGRESS": IamWorkRequestStatusInProgress,
	"FAILED":      IamWorkRequestStatusFailed,
	"SUCCEEDED":   IamWorkRequestStatusSucceeded,
	"CANCELING":   IamWorkRequestStatusCanceling,
	"CANCELED":    IamWorkRequestStatusCanceled,
}

var mappingIamWorkRequestStatusEnumLowerCase = map[string]IamWorkRequestStatusEnum{
	"accepted":    IamWorkRequestStatusAccepted,
	"in_progress": IamWorkRequestStatusInProgress,
	"failed":      IamWorkRequestStatusFailed,
	"succeeded":   IamWorkRequestStatusSucceeded,
	"canceling":   IamWorkRequestStatusCanceling,
	"canceled":    IamWorkRequestStatusCanceled,
}

// GetIamWorkRequestStatusEnumValues Enumerates the set of values for IamWorkRequestStatusEnum
func GetIamWorkRequestStatusEnumValues() []IamWorkRequestStatusEnum {
	values := make([]IamWorkRequestStatusEnum, 0)
	for _, v := range mappingIamWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetIamWorkRequestStatusEnumStringValues Enumerates the set of values in String for IamWorkRequestStatusEnum
func GetIamWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingIamWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIamWorkRequestStatusEnum(val string) (IamWorkRequestStatusEnum, bool) {
	enum, ok := mappingIamWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
