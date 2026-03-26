// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MatchingRule An object that defines the set of rules that identifies the target instances in a dynamic set.
type MatchingRule struct {

	// The list of the managed instance tags.
	Tags []Tag `mandatory:"false" json:"tags"`

	// The list of managed instance ids.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The list of managed instance group IDs.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The list of managed instance display names.
	DisplayNames []string `mandatory:"false" json:"displayNames"`

	// The list of managed instance OS names.
	OsNames []OsNameEnum `mandatory:"false" json:"osNames,omitempty"`

	// The list of managed instance architectures.
	Architectures []CpuArchTypeEnum `mandatory:"false" json:"architectures,omitempty"`

	// The list of managed instance OS families.
	OsFamilies []OsFamilyEnum `mandatory:"false" json:"osFamilies,omitempty"`

	// The list of managed instance statuses.
	Statuses []ManagedInstanceStatusEnum `mandatory:"false" json:"statuses,omitempty"`

	// The list of managed instance locations.
	Locations []ManagedInstanceLocationEnum `mandatory:"false" json:"locations,omitempty"`

	// Indicates if the managed instance needs to be rebooted.
	IsRebootRequired *bool `mandatory:"false" json:"isRebootRequired"`
}

func (m MatchingRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MatchingRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.OsNames {
		if _, ok := GetMappingOsNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsNames: %s. Supported values are: %s.", val, strings.Join(GetOsNameEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Architectures {
		if _, ok := GetMappingCpuArchTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architectures: %s. Supported values are: %s.", val, strings.Join(GetCpuArchTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range m.OsFamilies {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamilies: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Statuses {
		if _, ok := GetMappingManagedInstanceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statuses: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range m.Locations {
		if _, ok := GetMappingManagedInstanceLocationEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Locations: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MatchingRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Tags                    []tag                         `json:"tags"`
		ManagedInstanceIds      []string                      `json:"managedInstanceIds"`
		ManagedInstanceGroupIds []string                      `json:"managedInstanceGroupIds"`
		DisplayNames            []string                      `json:"displayNames"`
		OsNames                 []OsNameEnum                  `json:"osNames"`
		Architectures           []CpuArchTypeEnum             `json:"architectures"`
		OsFamilies              []OsFamilyEnum                `json:"osFamilies"`
		Statuses                []ManagedInstanceStatusEnum   `json:"statuses"`
		Locations               []ManagedInstanceLocationEnum `json:"locations"`
		IsRebootRequired        *bool                         `json:"isRebootRequired"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Tags = make([]Tag, len(model.Tags))
	for i, n := range model.Tags {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tags[i] = nn.(Tag)
		} else {
			m.Tags[i] = nil
		}
	}
	m.ManagedInstanceIds = make([]string, len(model.ManagedInstanceIds))
	copy(m.ManagedInstanceIds, model.ManagedInstanceIds)
	m.ManagedInstanceGroupIds = make([]string, len(model.ManagedInstanceGroupIds))
	copy(m.ManagedInstanceGroupIds, model.ManagedInstanceGroupIds)
	m.DisplayNames = make([]string, len(model.DisplayNames))
	copy(m.DisplayNames, model.DisplayNames)
	m.OsNames = make([]OsNameEnum, len(model.OsNames))
	copy(m.OsNames, model.OsNames)
	m.Architectures = make([]CpuArchTypeEnum, len(model.Architectures))
	copy(m.Architectures, model.Architectures)
	m.OsFamilies = make([]OsFamilyEnum, len(model.OsFamilies))
	copy(m.OsFamilies, model.OsFamilies)
	m.Statuses = make([]ManagedInstanceStatusEnum, len(model.Statuses))
	copy(m.Statuses, model.Statuses)
	m.Locations = make([]ManagedInstanceLocationEnum, len(model.Locations))
	copy(m.Locations, model.Locations)
	m.IsRebootRequired = model.IsRebootRequired

	return
}
