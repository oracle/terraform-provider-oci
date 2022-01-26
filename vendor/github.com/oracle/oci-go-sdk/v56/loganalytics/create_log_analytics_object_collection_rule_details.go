// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateLogAnalyticsObjectCollectionRuleDetails The configuration details of collection rule to enable automatic log collection from an object storage bucket.
type CreateLogAnalyticsObjectCollectionRuleDetails struct {

	// A unique name given to the rule. The name must be unique within the tenancy, and cannot be modified.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Object Storage namespace.
	OsNamespace *string `mandatory:"true" json:"osNamespace"`

	// Name of the Object Storage bucket.
	OsBucketName *string `mandatory:"true" json:"osBucketName"`

	// Logging Analytics Log group OCID to associate the processed logs with.
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// Name of the Logging Analytics Source to use for the processing.
	LogSourceName *string `mandatory:"true" json:"logSourceName"`

	// A string that describes the details of the rule. It does not have to be unique, and can be changed.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The type of collection.
	CollectionType ObjectCollectionRuleCollectionTypesEnum `mandatory:"false" json:"collectionType,omitempty"`

	// The oldest time of the file in the bucket to consider for collection.
	// Accepted values are: BEGINNING or CURRENT_TIME or RFC3339 formatted datetime string.
	// Use this for HISTORIC or HISTORIC_LIVE collection types. When collectionType is LIVE, specifying pollSince value other than CURRENT_TIME will result in error.
	PollSince *string `mandatory:"false" json:"pollSince"`

	// The newest time of the file in the bucket to consider for collection.
	// Accepted values are: CURRENT_TIME or RFC3339 formatted datetime string.
	// Use this for HISTORIC collection type. When collectionType is LIVE or HISTORIC_LIVE, specifying pollTill will result in error.
	PollTill *string `mandatory:"false" json:"pollTill"`

	// Logging Analytics entity OCID. Associates the processed logs with the given entity (optional).
	EntityId *string `mandatory:"false" json:"entityId"`

	// An optional character encoding to aid in detecting the character encoding of the contents of the objects while processing.
	// It is recommended to set this value as ISO_8859_1 when configuring content of the objects having more numeric characters,
	// and very few alphabets.
	// For e.g. this applies when configuring VCN Flow Logs.
	CharEncoding *string `mandatory:"false" json:"charEncoding"`

	// Whether or not this rule is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The override is used to modify some important configuration properties for objects matching a specific pattern inside the bucket.
	// Supported propeties for override are: logSourceName, charEncoding, entityId.
	// Supported matchType for override are "contains".
	Overrides map[string][]PropertyOverride `mandatory:"false" json:"overrides"`

	// When the filters are provided, only the objects matching the filters are picked up for processing.
	// The matchType supported is exact match and accommodates wildcard "*".
	// For more information on filters, see Event Filters (https://docs.oracle.com/en-us/iaas/Content/Events/Concepts/filterevents.htm).
	ObjectNameFilters []string `mandatory:"false" json:"objectNameFilters"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateLogAnalyticsObjectCollectionRuleDetails) String() string {
	return common.PointerString(m)
}
