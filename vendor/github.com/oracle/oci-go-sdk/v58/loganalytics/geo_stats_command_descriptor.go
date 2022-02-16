// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// GeoStatsCommandDescriptor Command descriptor for querylanguage GEOSTATS command.  This is similiar to STATS with some built in functions for City, State and Country by Coordinates.
type GeoStatsCommandDescriptor struct {

	// Command fragment display string from user specified query string formatted by query builder.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Command fragment internal string from user specified query string formatted by query builder.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// querylanguage command designation for example; reporting vs filtering
	Category *string `mandatory:"false" json:"category"`

	// Fields referenced in command fragment from user specified query string.
	ReferencedFields []AbstractField `mandatory:"false" json:"referencedFields"`

	// Fields declared in command fragment from user specified query string.
	DeclaredFields []AbstractField `mandatory:"false" json:"declaredFields"`

	// Group by fields if specified in the query string.
	GroupByFields []AbstractField `mandatory:"false" json:"groupByFields"`

	// Statistical functions specified in the query string. Atleast 1 is required for a GEOSTATS command.
	Functions []FunctionField `mandatory:"false" json:"functions"`

	// Indicates which coordinates to show.  Either client, server or both.  Defaults to client.
	Include GeoStatsCommandDescriptorIncludeEnum `mandatory:"false" json:"include,omitempty"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m GeoStatsCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m GeoStatsCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m GeoStatsCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m GeoStatsCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m GeoStatsCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m GeoStatsCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GeoStatsCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGeoStatsCommandDescriptorIncludeEnum(string(m.Include)); !ok && m.Include != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Include: %s. Supported values are: %s.", m.Include, strings.Join(GetGeoStatsCommandDescriptorIncludeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GeoStatsCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGeoStatsCommandDescriptor GeoStatsCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeGeoStatsCommandDescriptor
	}{
		"GEO_STATS",
		(MarshalTypeGeoStatsCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GeoStatsCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string                              `json:"category"`
		ReferencedFields    []abstractfield                      `json:"referencedFields"`
		DeclaredFields      []abstractfield                      `json:"declaredFields"`
		Include             GeoStatsCommandDescriptorIncludeEnum `json:"include"`
		GroupByFields       []abstractfield                      `json:"groupByFields"`
		Functions           []FunctionField                      `json:"functions"`
		DisplayQueryString  *string                              `json:"displayQueryString"`
		InternalQueryString *string                              `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.ReferencedFields = make([]AbstractField, len(model.ReferencedFields))
	for i, n := range model.ReferencedFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ReferencedFields[i] = nn.(AbstractField)
		} else {
			m.ReferencedFields[i] = nil
		}
	}

	m.DeclaredFields = make([]AbstractField, len(model.DeclaredFields))
	for i, n := range model.DeclaredFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeclaredFields[i] = nn.(AbstractField)
		} else {
			m.DeclaredFields[i] = nil
		}
	}

	m.Include = model.Include

	m.GroupByFields = make([]AbstractField, len(model.GroupByFields))
	for i, n := range model.GroupByFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.GroupByFields[i] = nn.(AbstractField)
		} else {
			m.GroupByFields[i] = nil
		}
	}

	m.Functions = make([]FunctionField, len(model.Functions))
	for i, n := range model.Functions {
		m.Functions[i] = n
	}

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}

// GeoStatsCommandDescriptorIncludeEnum Enum with underlying type: string
type GeoStatsCommandDescriptorIncludeEnum string

// Set of constants representing the allowable values for GeoStatsCommandDescriptorIncludeEnum
const (
	GeoStatsCommandDescriptorIncludeClient          GeoStatsCommandDescriptorIncludeEnum = "CLIENT"
	GeoStatsCommandDescriptorIncludeServer          GeoStatsCommandDescriptorIncludeEnum = "SERVER"
	GeoStatsCommandDescriptorIncludeClientAndServer GeoStatsCommandDescriptorIncludeEnum = "CLIENT_AND_SERVER"
)

var mappingGeoStatsCommandDescriptorIncludeEnum = map[string]GeoStatsCommandDescriptorIncludeEnum{
	"CLIENT":            GeoStatsCommandDescriptorIncludeClient,
	"SERVER":            GeoStatsCommandDescriptorIncludeServer,
	"CLIENT_AND_SERVER": GeoStatsCommandDescriptorIncludeClientAndServer,
}

// GetGeoStatsCommandDescriptorIncludeEnumValues Enumerates the set of values for GeoStatsCommandDescriptorIncludeEnum
func GetGeoStatsCommandDescriptorIncludeEnumValues() []GeoStatsCommandDescriptorIncludeEnum {
	values := make([]GeoStatsCommandDescriptorIncludeEnum, 0)
	for _, v := range mappingGeoStatsCommandDescriptorIncludeEnum {
		values = append(values, v)
	}
	return values
}

// GetGeoStatsCommandDescriptorIncludeEnumStringValues Enumerates the set of values in String for GeoStatsCommandDescriptorIncludeEnum
func GetGeoStatsCommandDescriptorIncludeEnumStringValues() []string {
	return []string{
		"CLIENT",
		"SERVER",
		"CLIENT_AND_SERVER",
	}
}

// GetMappingGeoStatsCommandDescriptorIncludeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGeoStatsCommandDescriptorIncludeEnum(val string) (GeoStatsCommandDescriptorIncludeEnum, bool) {
	mappingGeoStatsCommandDescriptorIncludeEnumIgnoreCase := make(map[string]GeoStatsCommandDescriptorIncludeEnum)
	for k, v := range mappingGeoStatsCommandDescriptorIncludeEnum {
		mappingGeoStatsCommandDescriptorIncludeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGeoStatsCommandDescriptorIncludeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
