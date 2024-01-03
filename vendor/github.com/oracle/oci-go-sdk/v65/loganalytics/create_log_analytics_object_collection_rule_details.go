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

	// Timezone to be used when processing log entries whose timestamps do not include an explicit timezone.
	// When this property is not specified, the timezone of the entity specified is used.
	// If the entity is also not specified or do not have a valid timezone then UTC is used.
	Timezone *string `mandatory:"false" json:"timezone"`

	// The logSet to be associated with the processed logs. The logSet feature can be used by customers with high volume of data
	// and this feature has to be enabled for a given tenancy prior to its usage.
	// When logSetExtRegex value is provided, it will take precedence over this logSet value and logSet will be computed dynamically
	// using logSetKey and logSetExtRegex.
	LogSet *string `mandatory:"false" json:"logSet"`

	// An optional parameter to indicate from where the logSet to be extracted using logSetExtRegex. Default value is OBJECT_PATH (e.g. /n/<namespace>/b/<bucketname>/o/<objectname>).
	LogSetKey LogSetKeyTypesEnum `mandatory:"false" json:"logSetKey,omitempty"`

	// The regex to be applied against given logSetKey. Regex has to be in string escaped format.
	LogSetExtRegex *string `mandatory:"false" json:"logSetExtRegex"`

	// The override is used to modify some important configuration properties for objects matching a specific pattern inside the bucket.
	// Supported propeties for override are: logSourceName, charEncoding, entityId.
	// Supported matchType for override are "contains".
	Overrides map[string][]PropertyOverride `mandatory:"false" json:"overrides"`

	// When the filters are provided, only the objects matching the filters are picked up for processing.
	// The matchType supported is exact match and accommodates wildcard "*".
	// For more information on filters, see Event Filters (https://docs.oracle.com/en-us/iaas/Content/Events/Concepts/filterevents.htm).
	ObjectNameFilters []string `mandatory:"false" json:"objectNameFilters"`

	// Type of files/objects in this object collection rule.
	LogType LogTypesEnum `mandatory:"false" json:"logType,omitempty"`

	// Flag to allow historic collection if poll period overlaps with existing ACTIVE collection rule
	IsForceHistoricCollection *bool `mandatory:"false" json:"isForceHistoricCollection"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLogAnalyticsObjectCollectionRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectCollectionRuleCollectionTypesEnum(string(m.CollectionType)); !ok && m.CollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionType: %s. Supported values are: %s.", m.CollectionType, strings.Join(GetObjectCollectionRuleCollectionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogSetKeyTypesEnum(string(m.LogSetKey)); !ok && m.LogSetKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogSetKey: %s. Supported values are: %s.", m.LogSetKey, strings.Join(GetLogSetKeyTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogTypesEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetLogTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
