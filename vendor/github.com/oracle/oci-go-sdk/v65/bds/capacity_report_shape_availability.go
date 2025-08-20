// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CapacityReportShapeAvailability Detailed information about the availabilities of a shape in each domain.
type CapacityReportShapeAvailability struct {

	// The shape that the capacity report was requested for.
	Shape *string `mandatory:"true" json:"shape"`

	// Information about the capacity in each domain.
	DomainLevelCapacityReports []DomainTypeCapacityReport `mandatory:"true" json:"domainLevelCapacityReports"`

	ShapeConfig *ShapeConfigDetails `mandatory:"false" json:"shapeConfig"`
}

func (m CapacityReportShapeAvailability) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CapacityReportShapeAvailability) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CapacityReportShapeAvailability) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ShapeConfig                *ShapeConfigDetails        `json:"shapeConfig"`
		Shape                      *string                    `json:"shape"`
		DomainLevelCapacityReports []domaintypecapacityreport `json:"domainLevelCapacityReports"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ShapeConfig = model.ShapeConfig

	m.Shape = model.Shape

	m.DomainLevelCapacityReports = make([]DomainTypeCapacityReport, len(model.DomainLevelCapacityReports))
	for i, n := range model.DomainLevelCapacityReports {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DomainLevelCapacityReports[i] = nn.(DomainTypeCapacityReport)
		} else {
			m.DomainLevelCapacityReports[i] = nil
		}
	}
	return
}
