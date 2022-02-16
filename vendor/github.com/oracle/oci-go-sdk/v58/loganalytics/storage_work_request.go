// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StorageWorkRequest This shows the storage work request details.
type StorageWorkRequest struct {

	// This is the OCID of the storage work Request.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is the work request status.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// This is the end of the time interval
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// Thie is the type of data to be purged
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// This is the type of the work request.
	OperationType StorageOperationTypeEnum `mandatory:"true" json:"operationType"`

	// When the work request started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// When the work request was accepted. Should match timeStarted in all cases.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// When the work request finished execution.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// When the work request will expire.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Percentage progress completion of the work request.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// This is the start of the time interval
	TimeDataStarted *common.SDKTime `mandatory:"false" json:"timeDataStarted"`

	// This is the solr query used to filter data for purge, '*' means all
	PurgeQueryString *string `mandatory:"false" json:"purgeQueryString"`

	// This provides more detailed status if applicable
	StatusDetails *string `mandatory:"false" json:"statusDetails"`

	// This provides more detailed info about the work request if applicable
	OperationDetails *string `mandatory:"false" json:"operationDetails"`

	// This is the policy name if applicable (e.g. purge policy)
	PolicyName *string `mandatory:"false" json:"policyName"`

	// This is the purge policy ID if applicable
	PolicyId *string `mandatory:"false" json:"policyId"`

	// This is the data usage in bytes if applicable
	StorageUsageInBytes *int64 `mandatory:"false" json:"storageUsageInBytes"`

	// If true, purge child compartments data, only applicable to purge request
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`
}

func (m StorageWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StorageWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStorageDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetStorageDataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStorageOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetStorageOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
