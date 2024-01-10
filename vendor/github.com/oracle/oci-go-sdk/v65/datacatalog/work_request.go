// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest A description of workrequest status.
type WorkRequest struct {

	// Type of the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, it is up to the service team to pick the primary
	// resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateCatalog                WorkRequestOperationTypeEnum = "CREATE_CATALOG"
	WorkRequestOperationTypeUpdateCatalog                WorkRequestOperationTypeEnum = "UPDATE_CATALOG"
	WorkRequestOperationTypeDeleteCatalog                WorkRequestOperationTypeEnum = "DELETE_CATALOG"
	WorkRequestOperationTypeMoveCatalog                  WorkRequestOperationTypeEnum = "MOVE_CATALOG"
	WorkRequestOperationTypeCreateCatalogPrivateEndpoint WorkRequestOperationTypeEnum = "CREATE_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDeleteCatalogPrivateEndpoint WorkRequestOperationTypeEnum = "DELETE_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeUpdateCatalogPrivateEndpoint WorkRequestOperationTypeEnum = "UPDATE_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeMoveCatalogPrivateEndpoint   WorkRequestOperationTypeEnum = "MOVE_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeAttachCatalogPrivateEndpoint WorkRequestOperationTypeEnum = "ATTACH_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDetachCatalogPrivateEndpoint WorkRequestOperationTypeEnum = "DETACH_CATALOG_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeCreateMetastore              WorkRequestOperationTypeEnum = "CREATE_METASTORE"
	WorkRequestOperationTypeUpdateMetastore              WorkRequestOperationTypeEnum = "UPDATE_METASTORE"
	WorkRequestOperationTypeDeleteMetastore              WorkRequestOperationTypeEnum = "DELETE_METASTORE"
	WorkRequestOperationTypeMoveMetastore                WorkRequestOperationTypeEnum = "MOVE_METASTORE"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_CATALOG":                  WorkRequestOperationTypeCreateCatalog,
	"UPDATE_CATALOG":                  WorkRequestOperationTypeUpdateCatalog,
	"DELETE_CATALOG":                  WorkRequestOperationTypeDeleteCatalog,
	"MOVE_CATALOG":                    WorkRequestOperationTypeMoveCatalog,
	"CREATE_CATALOG_PRIVATE_ENDPOINT": WorkRequestOperationTypeCreateCatalogPrivateEndpoint,
	"DELETE_CATALOG_PRIVATE_ENDPOINT": WorkRequestOperationTypeDeleteCatalogPrivateEndpoint,
	"UPDATE_CATALOG_PRIVATE_ENDPOINT": WorkRequestOperationTypeUpdateCatalogPrivateEndpoint,
	"MOVE_CATALOG_PRIVATE_ENDPOINT":   WorkRequestOperationTypeMoveCatalogPrivateEndpoint,
	"ATTACH_CATALOG_PRIVATE_ENDPOINT": WorkRequestOperationTypeAttachCatalogPrivateEndpoint,
	"DETACH_CATALOG_PRIVATE_ENDPOINT": WorkRequestOperationTypeDetachCatalogPrivateEndpoint,
	"CREATE_METASTORE":                WorkRequestOperationTypeCreateMetastore,
	"UPDATE_METASTORE":                WorkRequestOperationTypeUpdateMetastore,
	"DELETE_METASTORE":                WorkRequestOperationTypeDeleteMetastore,
	"MOVE_METASTORE":                  WorkRequestOperationTypeMoveMetastore,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_catalog":                  WorkRequestOperationTypeCreateCatalog,
	"update_catalog":                  WorkRequestOperationTypeUpdateCatalog,
	"delete_catalog":                  WorkRequestOperationTypeDeleteCatalog,
	"move_catalog":                    WorkRequestOperationTypeMoveCatalog,
	"create_catalog_private_endpoint": WorkRequestOperationTypeCreateCatalogPrivateEndpoint,
	"delete_catalog_private_endpoint": WorkRequestOperationTypeDeleteCatalogPrivateEndpoint,
	"update_catalog_private_endpoint": WorkRequestOperationTypeUpdateCatalogPrivateEndpoint,
	"move_catalog_private_endpoint":   WorkRequestOperationTypeMoveCatalogPrivateEndpoint,
	"attach_catalog_private_endpoint": WorkRequestOperationTypeAttachCatalogPrivateEndpoint,
	"detach_catalog_private_endpoint": WorkRequestOperationTypeDetachCatalogPrivateEndpoint,
	"create_metastore":                WorkRequestOperationTypeCreateMetastore,
	"update_metastore":                WorkRequestOperationTypeUpdateMetastore,
	"delete_metastore":                WorkRequestOperationTypeDeleteMetastore,
	"move_metastore":                  WorkRequestOperationTypeMoveMetastore,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_CATALOG",
		"UPDATE_CATALOG",
		"DELETE_CATALOG",
		"MOVE_CATALOG",
		"CREATE_CATALOG_PRIVATE_ENDPOINT",
		"DELETE_CATALOG_PRIVATE_ENDPOINT",
		"UPDATE_CATALOG_PRIVATE_ENDPOINT",
		"MOVE_CATALOG_PRIVATE_ENDPOINT",
		"ATTACH_CATALOG_PRIVATE_ENDPOINT",
		"DETACH_CATALOG_PRIVATE_ENDPOINT",
		"CREATE_METASTORE",
		"UPDATE_METASTORE",
		"DELETE_METASTORE",
		"MOVE_METASTORE",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"in_progress": WorkRequestStatusInProgress,
	"failed":      WorkRequestStatusFailed,
	"succeeded":   WorkRequestStatusSucceeded,
	"canceling":   WorkRequestStatusCanceling,
	"canceled":    WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusEnumStringValues Enumerates the set of values in String for WorkRequestStatusEnum
func GetWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
