// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsEmBridgeSummary Enterprise manager bridge summary.
type LogAnalyticsEmBridgeSummary struct {

	// The enterprise manager bridge OCID.
	Id *string `mandatory:"true" json:"id"`

	// Log analytics enterprise manager bridge display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Compartment for entities created from enterprise manager.
	EmEntitiesCompartmentId *string `mandatory:"true" json:"emEntitiesCompartmentId"`

	// Object store bucket name where enterprise manager harvested entities will be uploaded.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the enterprise manager bridge.
	LifecycleState EmBridgeLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The status from last processing status of enterprise manager upload.
	LastImportProcessingStatus EmBridgeLatestImportProcessingStatusEnum `mandatory:"true" json:"lastImportProcessingStatus"`

	// A description for log analytics enterprise manager bridge.
	Description *string `mandatory:"false" json:"description"`

	// lifecycleDetails has additional information regarding substeps such as verifying connection to object store.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Processing status details of enterprise manager upload. This provides additional details
	// for failed status
	LastImportProcessingDetails *string `mandatory:"false" json:"lastImportProcessingDetails"`

	// The last time of enterprise manager upload was processed. This is in the format defined by RFC3339
	TimeImportLastProcessed *common.SDKTime `mandatory:"false" json:"timeImportLastProcessed"`

	// The timestamp of last enterprise manager upload to OCI Object Store. This is in the format defined by RFC3339
	TimeEmDataLastExtracted *common.SDKTime `mandatory:"false" json:"timeEmDataLastExtracted"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LogAnalyticsEmBridgeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsEmBridgeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmBridgeLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmBridgeLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmBridgeLatestImportProcessingStatusEnum(string(m.LastImportProcessingStatus)); !ok && m.LastImportProcessingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastImportProcessingStatus: %s. Supported values are: %s.", m.LastImportProcessingStatus, strings.Join(GetEmBridgeLatestImportProcessingStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
