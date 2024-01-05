// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AbstractCommandDescriptor Generic command descriptor defining all attributes common to all querylanguage commands for parse output.
type AbstractCommandDescriptor interface {

	// Command fragment display string from user specified query string formatted by query builder.
	GetDisplayQueryString() *string

	// Command fragment internal string from user specified query string formatted by query builder.
	GetInternalQueryString() *string

	// querylanguage command designation for example; reporting vs filtering
	GetCategory() *string

	// Fields referenced in command fragment from user specified query string.
	GetReferencedFields() []AbstractField

	// Fields declared in command fragment from user specified query string.
	GetDeclaredFields() []AbstractField

	// Field denoting if this is a hidden command that is not shown in the query string.
	GetIsHidden() *bool
}

type abstractcommanddescriptor struct {
	JsonData            []byte
	Category            *string         `mandatory:"false" json:"category"`
	ReferencedFields    json.RawMessage `mandatory:"false" json:"referencedFields"`
	DeclaredFields      json.RawMessage `mandatory:"false" json:"declaredFields"`
	IsHidden            *bool           `mandatory:"false" json:"isHidden"`
	DisplayQueryString  *string         `mandatory:"true" json:"displayQueryString"`
	InternalQueryString *string         `mandatory:"true" json:"internalQueryString"`
	Name                string          `json:"name"`
}

// UnmarshalJSON unmarshals json
func (m *abstractcommanddescriptor) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractcommanddescriptor abstractcommanddescriptor
	s := struct {
		Model Unmarshalerabstractcommanddescriptor
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayQueryString = s.Model.DisplayQueryString
	m.InternalQueryString = s.Model.InternalQueryString
	m.Category = s.Model.Category
	m.ReferencedFields = s.Model.ReferencedFields
	m.DeclaredFields = s.Model.DeclaredFields
	m.IsHidden = s.Model.IsHidden
	m.Name = s.Model.Name

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractcommanddescriptor) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Name {
	case "TOP":
		mm := TopCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HIGHLIGHT":
		mm := HighlightCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATS":
		mm := StatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAIL":
		mm := TailCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OUTLIER":
		mm := OutlierCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEMO_MODE":
		mm := DemoModeCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD_SUMMARY":
		mm := FieldSummaryCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GEO_STATS":
		mm := GeoStatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MAP":
		mm := MapCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HIGHLIGHT_GROUPS":
		mm := HighlightGroupsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEDUP":
		mm := DedupCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_STATS":
		mm := TimeStatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER":
		mm := ClusterCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELETE":
		mm := DeleteCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH":
		mm := SearchCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUCKET":
		mm := BucketCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RARE":
		mm := RareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_INSIGHTS":
		mm := AddInsightsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LINK":
		mm := LinkCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SORT":
		mm := SortCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HIGHLIGHT_ROWS":
		mm := HighlightRowsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACRO":
		mm := MacroCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EVAL":
		mm := EvalCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RENAME":
		mm := RenameCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XML_EXTRACT":
		mm := XmlExtractCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MULTI_SEARCH":
		mm := MultiSearchCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARE":
		mm := CompareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_COMPARE":
		mm := TimeCompareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MODULE":
		mm := ModuleCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REGEX":
		mm := RegexCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELTA":
		mm := DeltaCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOOKUP":
		mm := LookupCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON_EXTRACT":
		mm := JsonExtractCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EVENT_STATS":
		mm := EventStatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WHERE":
		mm := WhereCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_SPLIT":
		mm := ClusterSplitCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FREQUENT":
		mm := FrequentCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_DETAILS":
		mm := ClusterDetailsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_COMPARE":
		mm := ClusterCompareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMMAND":
		mm := CommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISTINCT":
		mm := DistinctCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTRACT":
		mm := ExtractCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NLP":
		mm := NlpCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BOTTOM":
		mm := BottomCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELDS":
		mm := FieldsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANOMALY":
		mm := AnomalyCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLASSIFY":
		mm := ClassifyCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LINK_DETAILS":
		mm := LinkDetailsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH_LOOKUP":
		mm := SearchLookupCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HEAD":
		mm := HeadCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREATE_VIEW":
		mm := CreateViewCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_CLUSTER":
		mm := TimeClusterCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_FIELDS":
		mm := AddFieldsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AbstractCommandDescriptor: %s.", m.Name)
		return *m, nil
	}
}

// GetCategory returns Category
func (m abstractcommanddescriptor) GetCategory() *string {
	return m.Category
}

// GetReferencedFields returns ReferencedFields
func (m abstractcommanddescriptor) GetReferencedFields() json.RawMessage {
	return m.ReferencedFields
}

// GetDeclaredFields returns DeclaredFields
func (m abstractcommanddescriptor) GetDeclaredFields() json.RawMessage {
	return m.DeclaredFields
}

// GetIsHidden returns IsHidden
func (m abstractcommanddescriptor) GetIsHidden() *bool {
	return m.IsHidden
}

// GetDisplayQueryString returns DisplayQueryString
func (m abstractcommanddescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

// GetInternalQueryString returns InternalQueryString
func (m abstractcommanddescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

func (m abstractcommanddescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractcommanddescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractCommandDescriptorNameEnum Enum with underlying type: string
type AbstractCommandDescriptorNameEnum string

// Set of constants representing the allowable values for AbstractCommandDescriptorNameEnum
const (
	AbstractCommandDescriptorNameCommand         AbstractCommandDescriptorNameEnum = "COMMAND"
	AbstractCommandDescriptorNameSearch          AbstractCommandDescriptorNameEnum = "SEARCH"
	AbstractCommandDescriptorNameStats           AbstractCommandDescriptorNameEnum = "STATS"
	AbstractCommandDescriptorNameGeoStats        AbstractCommandDescriptorNameEnum = "GEO_STATS"
	AbstractCommandDescriptorNameTimeStats       AbstractCommandDescriptorNameEnum = "TIME_STATS"
	AbstractCommandDescriptorNameSort            AbstractCommandDescriptorNameEnum = "SORT"
	AbstractCommandDescriptorNameFields          AbstractCommandDescriptorNameEnum = "FIELDS"
	AbstractCommandDescriptorNameAddFields       AbstractCommandDescriptorNameEnum = "ADD_FIELDS"
	AbstractCommandDescriptorNameLink            AbstractCommandDescriptorNameEnum = "LINK"
	AbstractCommandDescriptorNameLinkDetails     AbstractCommandDescriptorNameEnum = "LINK_DETAILS"
	AbstractCommandDescriptorNameCluster         AbstractCommandDescriptorNameEnum = "CLUSTER"
	AbstractCommandDescriptorNameClusterDetails  AbstractCommandDescriptorNameEnum = "CLUSTER_DETAILS"
	AbstractCommandDescriptorNameClusterSplit    AbstractCommandDescriptorNameEnum = "CLUSTER_SPLIT"
	AbstractCommandDescriptorNameEval            AbstractCommandDescriptorNameEnum = "EVAL"
	AbstractCommandDescriptorNameExtract         AbstractCommandDescriptorNameEnum = "EXTRACT"
	AbstractCommandDescriptorNameJsonExtract     AbstractCommandDescriptorNameEnum = "JSON_EXTRACT"
	AbstractCommandDescriptorNameXmlExtract      AbstractCommandDescriptorNameEnum = "XML_EXTRACT"
	AbstractCommandDescriptorNameEventStats      AbstractCommandDescriptorNameEnum = "EVENT_STATS"
	AbstractCommandDescriptorNameBucket          AbstractCommandDescriptorNameEnum = "BUCKET"
	AbstractCommandDescriptorNameClassify        AbstractCommandDescriptorNameEnum = "CLASSIFY"
	AbstractCommandDescriptorNameTop             AbstractCommandDescriptorNameEnum = "TOP"
	AbstractCommandDescriptorNameBottom          AbstractCommandDescriptorNameEnum = "BOTTOM"
	AbstractCommandDescriptorNameHead            AbstractCommandDescriptorNameEnum = "HEAD"
	AbstractCommandDescriptorNameTail            AbstractCommandDescriptorNameEnum = "TAIL"
	AbstractCommandDescriptorNameFieldSummary    AbstractCommandDescriptorNameEnum = "FIELD_SUMMARY"
	AbstractCommandDescriptorNameRegex           AbstractCommandDescriptorNameEnum = "REGEX"
	AbstractCommandDescriptorNameRename          AbstractCommandDescriptorNameEnum = "RENAME"
	AbstractCommandDescriptorNameTimeCompare     AbstractCommandDescriptorNameEnum = "TIME_COMPARE"
	AbstractCommandDescriptorNameWhere           AbstractCommandDescriptorNameEnum = "WHERE"
	AbstractCommandDescriptorNameClusterCompare  AbstractCommandDescriptorNameEnum = "CLUSTER_COMPARE"
	AbstractCommandDescriptorNameDelete          AbstractCommandDescriptorNameEnum = "DELETE"
	AbstractCommandDescriptorNameDelta           AbstractCommandDescriptorNameEnum = "DELTA"
	AbstractCommandDescriptorNameDistinct        AbstractCommandDescriptorNameEnum = "DISTINCT"
	AbstractCommandDescriptorNameSearchLookup    AbstractCommandDescriptorNameEnum = "SEARCH_LOOKUP"
	AbstractCommandDescriptorNameLookup          AbstractCommandDescriptorNameEnum = "LOOKUP"
	AbstractCommandDescriptorNameDemoMode        AbstractCommandDescriptorNameEnum = "DEMO_MODE"
	AbstractCommandDescriptorNameMacro           AbstractCommandDescriptorNameEnum = "MACRO"
	AbstractCommandDescriptorNameModule          AbstractCommandDescriptorNameEnum = "MODULE"
	AbstractCommandDescriptorNameMultiSearch     AbstractCommandDescriptorNameEnum = "MULTI_SEARCH"
	AbstractCommandDescriptorNameHighlight       AbstractCommandDescriptorNameEnum = "HIGHLIGHT"
	AbstractCommandDescriptorNameHighlightRows   AbstractCommandDescriptorNameEnum = "HIGHLIGHT_ROWS"
	AbstractCommandDescriptorNameHighlightGroups AbstractCommandDescriptorNameEnum = "HIGHLIGHT_GROUPS"
	AbstractCommandDescriptorNameCreateView      AbstractCommandDescriptorNameEnum = "CREATE_VIEW"
	AbstractCommandDescriptorNameMap             AbstractCommandDescriptorNameEnum = "MAP"
	AbstractCommandDescriptorNameNlp             AbstractCommandDescriptorNameEnum = "NLP"
	AbstractCommandDescriptorNameCompare         AbstractCommandDescriptorNameEnum = "COMPARE"
	AbstractCommandDescriptorNameAddInsights     AbstractCommandDescriptorNameEnum = "ADD_INSIGHTS"
	AbstractCommandDescriptorNameAnomaly         AbstractCommandDescriptorNameEnum = "ANOMALY"
	AbstractCommandDescriptorNameDedup           AbstractCommandDescriptorNameEnum = "DEDUP"
	AbstractCommandDescriptorNameTimeCluster     AbstractCommandDescriptorNameEnum = "TIME_CLUSTER"
	AbstractCommandDescriptorNameFrequent        AbstractCommandDescriptorNameEnum = "FREQUENT"
	AbstractCommandDescriptorNameRare            AbstractCommandDescriptorNameEnum = "RARE"
	AbstractCommandDescriptorNameOutlier         AbstractCommandDescriptorNameEnum = "OUTLIER"
)

var mappingAbstractCommandDescriptorNameEnum = map[string]AbstractCommandDescriptorNameEnum{
	"COMMAND":          AbstractCommandDescriptorNameCommand,
	"SEARCH":           AbstractCommandDescriptorNameSearch,
	"STATS":            AbstractCommandDescriptorNameStats,
	"GEO_STATS":        AbstractCommandDescriptorNameGeoStats,
	"TIME_STATS":       AbstractCommandDescriptorNameTimeStats,
	"SORT":             AbstractCommandDescriptorNameSort,
	"FIELDS":           AbstractCommandDescriptorNameFields,
	"ADD_FIELDS":       AbstractCommandDescriptorNameAddFields,
	"LINK":             AbstractCommandDescriptorNameLink,
	"LINK_DETAILS":     AbstractCommandDescriptorNameLinkDetails,
	"CLUSTER":          AbstractCommandDescriptorNameCluster,
	"CLUSTER_DETAILS":  AbstractCommandDescriptorNameClusterDetails,
	"CLUSTER_SPLIT":    AbstractCommandDescriptorNameClusterSplit,
	"EVAL":             AbstractCommandDescriptorNameEval,
	"EXTRACT":          AbstractCommandDescriptorNameExtract,
	"JSON_EXTRACT":     AbstractCommandDescriptorNameJsonExtract,
	"XML_EXTRACT":      AbstractCommandDescriptorNameXmlExtract,
	"EVENT_STATS":      AbstractCommandDescriptorNameEventStats,
	"BUCKET":           AbstractCommandDescriptorNameBucket,
	"CLASSIFY":         AbstractCommandDescriptorNameClassify,
	"TOP":              AbstractCommandDescriptorNameTop,
	"BOTTOM":           AbstractCommandDescriptorNameBottom,
	"HEAD":             AbstractCommandDescriptorNameHead,
	"TAIL":             AbstractCommandDescriptorNameTail,
	"FIELD_SUMMARY":    AbstractCommandDescriptorNameFieldSummary,
	"REGEX":            AbstractCommandDescriptorNameRegex,
	"RENAME":           AbstractCommandDescriptorNameRename,
	"TIME_COMPARE":     AbstractCommandDescriptorNameTimeCompare,
	"WHERE":            AbstractCommandDescriptorNameWhere,
	"CLUSTER_COMPARE":  AbstractCommandDescriptorNameClusterCompare,
	"DELETE":           AbstractCommandDescriptorNameDelete,
	"DELTA":            AbstractCommandDescriptorNameDelta,
	"DISTINCT":         AbstractCommandDescriptorNameDistinct,
	"SEARCH_LOOKUP":    AbstractCommandDescriptorNameSearchLookup,
	"LOOKUP":           AbstractCommandDescriptorNameLookup,
	"DEMO_MODE":        AbstractCommandDescriptorNameDemoMode,
	"MACRO":            AbstractCommandDescriptorNameMacro,
	"MODULE":           AbstractCommandDescriptorNameModule,
	"MULTI_SEARCH":     AbstractCommandDescriptorNameMultiSearch,
	"HIGHLIGHT":        AbstractCommandDescriptorNameHighlight,
	"HIGHLIGHT_ROWS":   AbstractCommandDescriptorNameHighlightRows,
	"HIGHLIGHT_GROUPS": AbstractCommandDescriptorNameHighlightGroups,
	"CREATE_VIEW":      AbstractCommandDescriptorNameCreateView,
	"MAP":              AbstractCommandDescriptorNameMap,
	"NLP":              AbstractCommandDescriptorNameNlp,
	"COMPARE":          AbstractCommandDescriptorNameCompare,
	"ADD_INSIGHTS":     AbstractCommandDescriptorNameAddInsights,
	"ANOMALY":          AbstractCommandDescriptorNameAnomaly,
	"DEDUP":            AbstractCommandDescriptorNameDedup,
	"TIME_CLUSTER":     AbstractCommandDescriptorNameTimeCluster,
	"FREQUENT":         AbstractCommandDescriptorNameFrequent,
	"RARE":             AbstractCommandDescriptorNameRare,
	"OUTLIER":          AbstractCommandDescriptorNameOutlier,
}

var mappingAbstractCommandDescriptorNameEnumLowerCase = map[string]AbstractCommandDescriptorNameEnum{
	"command":          AbstractCommandDescriptorNameCommand,
	"search":           AbstractCommandDescriptorNameSearch,
	"stats":            AbstractCommandDescriptorNameStats,
	"geo_stats":        AbstractCommandDescriptorNameGeoStats,
	"time_stats":       AbstractCommandDescriptorNameTimeStats,
	"sort":             AbstractCommandDescriptorNameSort,
	"fields":           AbstractCommandDescriptorNameFields,
	"add_fields":       AbstractCommandDescriptorNameAddFields,
	"link":             AbstractCommandDescriptorNameLink,
	"link_details":     AbstractCommandDescriptorNameLinkDetails,
	"cluster":          AbstractCommandDescriptorNameCluster,
	"cluster_details":  AbstractCommandDescriptorNameClusterDetails,
	"cluster_split":    AbstractCommandDescriptorNameClusterSplit,
	"eval":             AbstractCommandDescriptorNameEval,
	"extract":          AbstractCommandDescriptorNameExtract,
	"json_extract":     AbstractCommandDescriptorNameJsonExtract,
	"xml_extract":      AbstractCommandDescriptorNameXmlExtract,
	"event_stats":      AbstractCommandDescriptorNameEventStats,
	"bucket":           AbstractCommandDescriptorNameBucket,
	"classify":         AbstractCommandDescriptorNameClassify,
	"top":              AbstractCommandDescriptorNameTop,
	"bottom":           AbstractCommandDescriptorNameBottom,
	"head":             AbstractCommandDescriptorNameHead,
	"tail":             AbstractCommandDescriptorNameTail,
	"field_summary":    AbstractCommandDescriptorNameFieldSummary,
	"regex":            AbstractCommandDescriptorNameRegex,
	"rename":           AbstractCommandDescriptorNameRename,
	"time_compare":     AbstractCommandDescriptorNameTimeCompare,
	"where":            AbstractCommandDescriptorNameWhere,
	"cluster_compare":  AbstractCommandDescriptorNameClusterCompare,
	"delete":           AbstractCommandDescriptorNameDelete,
	"delta":            AbstractCommandDescriptorNameDelta,
	"distinct":         AbstractCommandDescriptorNameDistinct,
	"search_lookup":    AbstractCommandDescriptorNameSearchLookup,
	"lookup":           AbstractCommandDescriptorNameLookup,
	"demo_mode":        AbstractCommandDescriptorNameDemoMode,
	"macro":            AbstractCommandDescriptorNameMacro,
	"module":           AbstractCommandDescriptorNameModule,
	"multi_search":     AbstractCommandDescriptorNameMultiSearch,
	"highlight":        AbstractCommandDescriptorNameHighlight,
	"highlight_rows":   AbstractCommandDescriptorNameHighlightRows,
	"highlight_groups": AbstractCommandDescriptorNameHighlightGroups,
	"create_view":      AbstractCommandDescriptorNameCreateView,
	"map":              AbstractCommandDescriptorNameMap,
	"nlp":              AbstractCommandDescriptorNameNlp,
	"compare":          AbstractCommandDescriptorNameCompare,
	"add_insights":     AbstractCommandDescriptorNameAddInsights,
	"anomaly":          AbstractCommandDescriptorNameAnomaly,
	"dedup":            AbstractCommandDescriptorNameDedup,
	"time_cluster":     AbstractCommandDescriptorNameTimeCluster,
	"frequent":         AbstractCommandDescriptorNameFrequent,
	"rare":             AbstractCommandDescriptorNameRare,
	"outlier":          AbstractCommandDescriptorNameOutlier,
}

// GetAbstractCommandDescriptorNameEnumValues Enumerates the set of values for AbstractCommandDescriptorNameEnum
func GetAbstractCommandDescriptorNameEnumValues() []AbstractCommandDescriptorNameEnum {
	values := make([]AbstractCommandDescriptorNameEnum, 0)
	for _, v := range mappingAbstractCommandDescriptorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractCommandDescriptorNameEnumStringValues Enumerates the set of values in String for AbstractCommandDescriptorNameEnum
func GetAbstractCommandDescriptorNameEnumStringValues() []string {
	return []string{
		"COMMAND",
		"SEARCH",
		"STATS",
		"GEO_STATS",
		"TIME_STATS",
		"SORT",
		"FIELDS",
		"ADD_FIELDS",
		"LINK",
		"LINK_DETAILS",
		"CLUSTER",
		"CLUSTER_DETAILS",
		"CLUSTER_SPLIT",
		"EVAL",
		"EXTRACT",
		"JSON_EXTRACT",
		"XML_EXTRACT",
		"EVENT_STATS",
		"BUCKET",
		"CLASSIFY",
		"TOP",
		"BOTTOM",
		"HEAD",
		"TAIL",
		"FIELD_SUMMARY",
		"REGEX",
		"RENAME",
		"TIME_COMPARE",
		"WHERE",
		"CLUSTER_COMPARE",
		"DELETE",
		"DELTA",
		"DISTINCT",
		"SEARCH_LOOKUP",
		"LOOKUP",
		"DEMO_MODE",
		"MACRO",
		"MODULE",
		"MULTI_SEARCH",
		"HIGHLIGHT",
		"HIGHLIGHT_ROWS",
		"HIGHLIGHT_GROUPS",
		"CREATE_VIEW",
		"MAP",
		"NLP",
		"COMPARE",
		"ADD_INSIGHTS",
		"ANOMALY",
		"DEDUP",
		"TIME_CLUSTER",
		"FREQUENT",
		"RARE",
		"OUTLIER",
	}
}

// GetMappingAbstractCommandDescriptorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractCommandDescriptorNameEnum(val string) (AbstractCommandDescriptorNameEnum, bool) {
	enum, ok := mappingAbstractCommandDescriptorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
