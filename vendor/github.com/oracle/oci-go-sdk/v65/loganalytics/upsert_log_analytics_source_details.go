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

// UpsertLogAnalyticsSourceDetails UpsertLogAnalyticsSourceDetails
type UpsertLogAnalyticsSourceDetails struct {

	// An array of source label conditions.
	LabelConditions []LogAnalyticsSourceLabelCondition `mandatory:"false" json:"labelConditions"`

	// An array of data filter definitions.
	DataFilterDefinitions []LogAnalyticsSourceDataFilter `mandatory:"false" json:"dataFilterDefinitions"`

	// The database credential name.
	DatabaseCredential *string `mandatory:"false" json:"databaseCredential"`

	// An array of extended field definitions.
	ExtendedFieldDefinitions []LogAnalyticsSourceExtendedFieldDefinition `mandatory:"false" json:"extendedFieldDefinitions"`

	// A flag indicating whether or not this is a cloud source.
	IsForCloud *bool `mandatory:"false" json:"isForCloud"`

	// An array of labels.
	Labels []LogAnalyticsLabelView `mandatory:"false" json:"labels"`

	// An array of metric definitions.
	MetricDefinitions []LogAnalyticsMetric `mandatory:"false" json:"metricDefinitions"`

	// An array of metrics.
	Metrics []LogAnalyticsSourceMetric `mandatory:"false" json:"metrics"`

	// An array of built in source parsers.
	OobParsers []LogAnalyticsParser `mandatory:"false" json:"oobParsers"`

	// An array of parameters.
	Parameters []LogAnalyticsParameter `mandatory:"false" json:"parameters"`

	// An array of patterns.
	Patterns []LogAnalyticsSourcePattern `mandatory:"false" json:"patterns"`

	// The source description.
	Description *string `mandatory:"false" json:"description"`

	// The source display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The source edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// An array of source functions.
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

	// An array of parser.
	Parsers []LogAnalyticsParser `mandatory:"false" json:"parsers"`

	// The rule unique identifier.
	RuleId *int64 `mandatory:"false" json:"ruleId"`

	// The source type internal name.
	TypeName *string `mandatory:"false" json:"typeName"`

	// The source warning configuration.
	WarningConfig *int64 `mandatory:"false" json:"warningConfig"`

	// An array of source metadata fields.
	MetadataFields []LogAnalyticsSourceMetadataField `mandatory:"false" json:"metadataFields"`

	// An array of labels.
	LabelDefinitions []LogAnalyticsLabelDefinition `mandatory:"false" json:"labelDefinitions"`

	// An array of entity types.
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// A flag indicating whether or not the source has a time zone override.
	IsTimezoneOverride *bool `mandatory:"false" json:"isTimezoneOverride"`

	// An array of custom parsers.
	UserParsers []LogAnalyticsParser `mandatory:"false" json:"userParsers"`

	// An array of categories to assign to the source. Specifying the name attribute for each category would suffice.
	// Oracle-defined category assignments cannot be removed.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`

	// An array of REST API endpoints for log collection.
	Endpoints []LogAnalyticsEndpoint `mandatory:"false" json:"endpoints"`

	// A list of source properties.
	SourceProperties []LogAnalyticsProperty `mandatory:"false" json:"sourceProperties"`
}

func (m UpsertLogAnalyticsSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpsertLogAnalyticsSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpsertLogAnalyticsSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LabelConditions          []LogAnalyticsSourceLabelCondition          `json:"labelConditions"`
		DataFilterDefinitions    []LogAnalyticsSourceDataFilter              `json:"dataFilterDefinitions"`
		DatabaseCredential       *string                                     `json:"databaseCredential"`
		ExtendedFieldDefinitions []LogAnalyticsSourceExtendedFieldDefinition `json:"extendedFieldDefinitions"`
		IsForCloud               *bool                                       `json:"isForCloud"`
		Labels                   []LogAnalyticsLabelView                     `json:"labels"`
		MetricDefinitions        []LogAnalyticsMetric                        `json:"metricDefinitions"`
		Metrics                  []LogAnalyticsSourceMetric                  `json:"metrics"`
		OobParsers               []LogAnalyticsParser                        `json:"oobParsers"`
		Parameters               []LogAnalyticsParameter                     `json:"parameters"`
		Patterns                 []LogAnalyticsSourcePattern                 `json:"patterns"`
		Description              *string                                     `json:"description"`
		DisplayName              *string                                     `json:"displayName"`
		EditVersion              *int64                                      `json:"editVersion"`
		Functions                []LogAnalyticsSourceFunction                `json:"functions"`
		SourceId                 *int64                                      `json:"sourceId"`
		Name                     *string                                     `json:"name"`
		IsSecureContent          *bool                                       `json:"isSecureContent"`
		IsSystem                 *bool                                       `json:"isSystem"`
		Parsers                  []LogAnalyticsParser                        `json:"parsers"`
		RuleId                   *int64                                      `json:"ruleId"`
		TypeName                 *string                                     `json:"typeName"`
		WarningConfig            *int64                                      `json:"warningConfig"`
		MetadataFields           []LogAnalyticsSourceMetadataField           `json:"metadataFields"`
		LabelDefinitions         []LogAnalyticsLabelDefinition               `json:"labelDefinitions"`
		EntityTypes              []LogAnalyticsSourceEntityType              `json:"entityTypes"`
		IsTimezoneOverride       *bool                                       `json:"isTimezoneOverride"`
		UserParsers              []LogAnalyticsParser                        `json:"userParsers"`
		Categories               []LogAnalyticsCategory                      `json:"categories"`
		Endpoints                []loganalyticsendpoint                      `json:"endpoints"`
		SourceProperties         []LogAnalyticsProperty                      `json:"sourceProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LabelConditions = make([]LogAnalyticsSourceLabelCondition, len(model.LabelConditions))
	copy(model.LabelConditions, m.LabelConditions)
	m.DataFilterDefinitions = make([]LogAnalyticsSourceDataFilter, len(model.DataFilterDefinitions))
	copy(model.DataFilterDefinitions, m.DataFilterDefinitions)
	m.DatabaseCredential = model.DatabaseCredential

	m.ExtendedFieldDefinitions = make([]LogAnalyticsSourceExtendedFieldDefinition, len(model.ExtendedFieldDefinitions))
	copy(model.ExtendedFieldDefinitions, m.ExtendedFieldDefinitions)
	m.IsForCloud = model.IsForCloud

	m.Labels = make([]LogAnalyticsLabelView, len(model.Labels))
	copy(model.Labels, m.Labels)
	m.MetricDefinitions = make([]LogAnalyticsMetric, len(model.MetricDefinitions))
	copy(model.MetricDefinitions, m.MetricDefinitions)
	m.Metrics = make([]LogAnalyticsSourceMetric, len(model.Metrics))
	copy(model.Metrics, m.Metrics)
	m.OobParsers = make([]LogAnalyticsParser, len(model.OobParsers))
	copy(model.OobParsers, m.OobParsers)
	m.Parameters = make([]LogAnalyticsParameter, len(model.Parameters))
	copy(model.Parameters, m.Parameters)
	m.Patterns = make([]LogAnalyticsSourcePattern, len(model.Patterns))
	copy(model.Patterns, m.Patterns)
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.EditVersion = model.EditVersion

	m.Functions = make([]LogAnalyticsSourceFunction, len(model.Functions))
	copy(model.Functions, m.Functions)
	m.SourceId = model.SourceId

	m.Name = model.Name

	m.IsSecureContent = model.IsSecureContent

	m.IsSystem = model.IsSystem

	m.Parsers = make([]LogAnalyticsParser, len(model.Parsers))
	copy(model.Parsers, m.Parsers)
	m.RuleId = model.RuleId

	m.TypeName = model.TypeName

	m.WarningConfig = model.WarningConfig

	m.MetadataFields = make([]LogAnalyticsSourceMetadataField, len(model.MetadataFields))
	copy(model.MetadataFields, m.MetadataFields)
	m.LabelDefinitions = make([]LogAnalyticsLabelDefinition, len(model.LabelDefinitions))
	copy(model.LabelDefinitions, m.LabelDefinitions)
	m.EntityTypes = make([]LogAnalyticsSourceEntityType, len(model.EntityTypes))
	copy(model.EntityTypes, m.EntityTypes)
	m.IsTimezoneOverride = model.IsTimezoneOverride

	m.UserParsers = make([]LogAnalyticsParser, len(model.UserParsers))
	copy(model.UserParsers, m.UserParsers)
	m.Categories = make([]LogAnalyticsCategory, len(model.Categories))
	copy(model.Categories, m.Categories)
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
	copy(model.SourceProperties, m.SourceProperties)
	return
}
