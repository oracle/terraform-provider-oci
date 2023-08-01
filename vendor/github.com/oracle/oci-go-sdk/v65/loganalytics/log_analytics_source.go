// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsSource LogAnalyticsSource
type LogAnalyticsSource struct {

	// The label alert conditions.
	LabelConditions []LogAnalyticsSourceLabelCondition `mandatory:"false" json:"labelConditions"`

	// The association count.
	AssociationCount *int `mandatory:"false" json:"associationCount"`

	// The association entities.
	AssociationEntity []LogAnalyticsAssociation `mandatory:"false" json:"associationEntity"`

	// The data filter definitions.
	DataFilterDefinitions []LogAnalyticsSourceDataFilter `mandatory:"false" json:"dataFilterDefinitions"`

	// The database credential.
	DatabaseCredential *string `mandatory:"false" json:"databaseCredential"`

	// The extended field definitions.
	ExtendedFieldDefinitions []LogAnalyticsSourceExtendedFieldDefinition `mandatory:"false" json:"extendedFieldDefinitions"`

	// A flag indicating whether or not this is a cloud source.
	IsForCloud *bool `mandatory:"false" json:"isForCloud"`

	// The labels associated with the source.
	Labels []LogAnalyticsLabelView `mandatory:"false" json:"labels"`

	// The metric definitions.
	MetricDefinitions []LogAnalyticsMetric `mandatory:"false" json:"metricDefinitions"`

	// The metric source map.
	Metrics []LogAnalyticsSourceMetric `mandatory:"false" json:"metrics"`

	// The built in parsers associated with source.
	OobParsers []LogAnalyticsParser `mandatory:"false" json:"oobParsers"`

	// The source parameters.
	Parameters []LogAnalyticsParameter `mandatory:"false" json:"parameters"`

	// The pattern count.
	PatternCount *int `mandatory:"false" json:"patternCount"`

	// The source patterns.
	Patterns []LogAnalyticsSourcePattern `mandatory:"false" json:"patterns"`

	// The source description.
	Description *string `mandatory:"false" json:"description"`

	// The source display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The source edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The source functions.
	Functions []LogAnalyticsSourceFunction `mandatory:"false" json:"functions"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The source internal name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating whether or not the source content is secure.
	IsSecureContent *bool `mandatory:"false" json:"isSecureContent"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The list of parsers used by the source.
	Parsers []LogAnalyticsParser `mandatory:"false" json:"parsers"`

	// A flag indicating whether or not the source is marked for auto-association.
	IsAutoAssociationEnabled *bool `mandatory:"false" json:"isAutoAssociationEnabled"`

	// A flag indicating whether or not the auto-association state should be overriden.
	IsAutoAssociationOverride *bool `mandatory:"false" json:"isAutoAssociationOverride"`

	// The rule unique identifier.
	RuleId *int64 `mandatory:"false" json:"ruleId"`

	// The source type internal name.
	TypeName *string `mandatory:"false" json:"typeName"`

	// The source type name.
	TypeDisplayName *string `mandatory:"false" json:"typeDisplayName"`

	// The source warning configuration.
	WarningConfig *int64 `mandatory:"false" json:"warningConfig"`

	// The source metadata fields.
	MetadataFields []LogAnalyticsSourceMetadataField `mandatory:"false" json:"metadataFields"`

	// The labls used by the source.
	LabelDefinitions []LogAnalyticsLabelDefinition `mandatory:"false" json:"labelDefinitions"`

	// The entity types.
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// A flag indicating whether or not the source has a time zone override.
	IsTimezoneOverride *bool `mandatory:"false" json:"isTimezoneOverride"`

	// An array of custom parsers.
	UserParsers []LogAnalyticsParser `mandatory:"false" json:"userParsers"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An array of event types.
	EventTypes []EventType `mandatory:"false" json:"eventTypes"`

	// An array of categories assigned to this source.
	// The isSystem flag denotes if each category assignment is user-created or Oracle-defined.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`

	// An array of REST API endpoints for log collection.
	Endpoints []LogAnalyticsEndpoint `mandatory:"false" json:"endpoints"`

	// A list of source properties.
	SourceProperties []LogAnalyticsProperty `mandatory:"false" json:"sourceProperties"`
}

func (m LogAnalyticsSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LogAnalyticsSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LabelConditions           []LogAnalyticsSourceLabelCondition          `json:"labelConditions"`
		AssociationCount          *int                                        `json:"associationCount"`
		AssociationEntity         []LogAnalyticsAssociation                   `json:"associationEntity"`
		DataFilterDefinitions     []LogAnalyticsSourceDataFilter              `json:"dataFilterDefinitions"`
		DatabaseCredential        *string                                     `json:"databaseCredential"`
		ExtendedFieldDefinitions  []LogAnalyticsSourceExtendedFieldDefinition `json:"extendedFieldDefinitions"`
		IsForCloud                *bool                                       `json:"isForCloud"`
		Labels                    []LogAnalyticsLabelView                     `json:"labels"`
		MetricDefinitions         []LogAnalyticsMetric                        `json:"metricDefinitions"`
		Metrics                   []LogAnalyticsSourceMetric                  `json:"metrics"`
		OobParsers                []LogAnalyticsParser                        `json:"oobParsers"`
		Parameters                []LogAnalyticsParameter                     `json:"parameters"`
		PatternCount              *int                                        `json:"patternCount"`
		Patterns                  []LogAnalyticsSourcePattern                 `json:"patterns"`
		Description               *string                                     `json:"description"`
		DisplayName               *string                                     `json:"displayName"`
		EditVersion               *int64                                      `json:"editVersion"`
		Functions                 []LogAnalyticsSourceFunction                `json:"functions"`
		SourceId                  *int64                                      `json:"sourceId"`
		Name                      *string                                     `json:"name"`
		IsSecureContent           *bool                                       `json:"isSecureContent"`
		IsSystem                  *bool                                       `json:"isSystem"`
		Parsers                   []LogAnalyticsParser                        `json:"parsers"`
		IsAutoAssociationEnabled  *bool                                       `json:"isAutoAssociationEnabled"`
		IsAutoAssociationOverride *bool                                       `json:"isAutoAssociationOverride"`
		RuleId                    *int64                                      `json:"ruleId"`
		TypeName                  *string                                     `json:"typeName"`
		TypeDisplayName           *string                                     `json:"typeDisplayName"`
		WarningConfig             *int64                                      `json:"warningConfig"`
		MetadataFields            []LogAnalyticsSourceMetadataField           `json:"metadataFields"`
		LabelDefinitions          []LogAnalyticsLabelDefinition               `json:"labelDefinitions"`
		EntityTypes               []LogAnalyticsSourceEntityType              `json:"entityTypes"`
		IsTimezoneOverride        *bool                                       `json:"isTimezoneOverride"`
		UserParsers               []LogAnalyticsParser                        `json:"userParsers"`
		TimeUpdated               *common.SDKTime                             `json:"timeUpdated"`
		EventTypes                []EventType                                 `json:"eventTypes"`
		Categories                []LogAnalyticsCategory                      `json:"categories"`
		Endpoints                 []loganalyticsendpoint                      `json:"endpoints"`
		SourceProperties          []LogAnalyticsProperty                      `json:"sourceProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LabelConditions = make([]LogAnalyticsSourceLabelCondition, len(model.LabelConditions))
	for i, n := range model.LabelConditions {
		m.LabelConditions[i] = n
	}

	m.AssociationCount = model.AssociationCount

	m.AssociationEntity = make([]LogAnalyticsAssociation, len(model.AssociationEntity))
	for i, n := range model.AssociationEntity {
		m.AssociationEntity[i] = n
	}

	m.DataFilterDefinitions = make([]LogAnalyticsSourceDataFilter, len(model.DataFilterDefinitions))
	for i, n := range model.DataFilterDefinitions {
		m.DataFilterDefinitions[i] = n
	}

	m.DatabaseCredential = model.DatabaseCredential

	m.ExtendedFieldDefinitions = make([]LogAnalyticsSourceExtendedFieldDefinition, len(model.ExtendedFieldDefinitions))
	for i, n := range model.ExtendedFieldDefinitions {
		m.ExtendedFieldDefinitions[i] = n
	}

	m.IsForCloud = model.IsForCloud

	m.Labels = make([]LogAnalyticsLabelView, len(model.Labels))
	for i, n := range model.Labels {
		m.Labels[i] = n
	}

	m.MetricDefinitions = make([]LogAnalyticsMetric, len(model.MetricDefinitions))
	for i, n := range model.MetricDefinitions {
		m.MetricDefinitions[i] = n
	}

	m.Metrics = make([]LogAnalyticsSourceMetric, len(model.Metrics))
	for i, n := range model.Metrics {
		m.Metrics[i] = n
	}

	m.OobParsers = make([]LogAnalyticsParser, len(model.OobParsers))
	for i, n := range model.OobParsers {
		m.OobParsers[i] = n
	}

	m.Parameters = make([]LogAnalyticsParameter, len(model.Parameters))
	for i, n := range model.Parameters {
		m.Parameters[i] = n
	}

	m.PatternCount = model.PatternCount

	m.Patterns = make([]LogAnalyticsSourcePattern, len(model.Patterns))
	for i, n := range model.Patterns {
		m.Patterns[i] = n
	}

	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.EditVersion = model.EditVersion

	m.Functions = make([]LogAnalyticsSourceFunction, len(model.Functions))
	for i, n := range model.Functions {
		m.Functions[i] = n
	}

	m.SourceId = model.SourceId

	m.Name = model.Name

	m.IsSecureContent = model.IsSecureContent

	m.IsSystem = model.IsSystem

	m.Parsers = make([]LogAnalyticsParser, len(model.Parsers))
	for i, n := range model.Parsers {
		m.Parsers[i] = n
	}

	m.IsAutoAssociationEnabled = model.IsAutoAssociationEnabled

	m.IsAutoAssociationOverride = model.IsAutoAssociationOverride

	m.RuleId = model.RuleId

	m.TypeName = model.TypeName

	m.TypeDisplayName = model.TypeDisplayName

	m.WarningConfig = model.WarningConfig

	m.MetadataFields = make([]LogAnalyticsSourceMetadataField, len(model.MetadataFields))
	for i, n := range model.MetadataFields {
		m.MetadataFields[i] = n
	}

	m.LabelDefinitions = make([]LogAnalyticsLabelDefinition, len(model.LabelDefinitions))
	for i, n := range model.LabelDefinitions {
		m.LabelDefinitions[i] = n
	}

	m.EntityTypes = make([]LogAnalyticsSourceEntityType, len(model.EntityTypes))
	for i, n := range model.EntityTypes {
		m.EntityTypes[i] = n
	}

	m.IsTimezoneOverride = model.IsTimezoneOverride

	m.UserParsers = make([]LogAnalyticsParser, len(model.UserParsers))
	for i, n := range model.UserParsers {
		m.UserParsers[i] = n
	}

	m.TimeUpdated = model.TimeUpdated

	m.EventTypes = make([]EventType, len(model.EventTypes))
	for i, n := range model.EventTypes {
		m.EventTypes[i] = n
	}

	m.Categories = make([]LogAnalyticsCategory, len(model.Categories))
	for i, n := range model.Categories {
		m.Categories[i] = n
	}

	m.Endpoints = make([]LogAnalyticsEndpoint, len(model.Endpoints))
	for i, n := range model.Endpoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Endpoints[i] = nn.(LogAnalyticsEndpoint)
		} else {
			m.Endpoints[i] = nil
		}
	}

	m.SourceProperties = make([]LogAnalyticsProperty, len(model.SourceProperties))
	for i, n := range model.SourceProperties {
		m.SourceProperties[i] = n
	}

	return
}
