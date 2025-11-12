// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportRunbookPrecheckDetails Request for precheck of Runbook import.
type ImportRunbookPrecheckDetails struct {

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ContentSource TransferRunbookContentDetails `mandatory:"true" json:"contentSource"`

	ContentDestination TransferRunbookContentDetails `mandatory:"true" json:"contentDestination"`

	ImportAs RunbookImportAs `mandatory:"false" json:"importAs"`

	// List of required values.
	RequiredValues []ImportRunbookPrecheckRequiredValue `mandatory:"false" json:"requiredValues"`
}

func (m ImportRunbookPrecheckDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportRunbookPrecheckDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImportRunbookPrecheckDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ImportAs           runbookimportas                      `json:"importAs"`
		RequiredValues     []ImportRunbookPrecheckRequiredValue `json:"requiredValues"`
		CompartmentId      *string                              `json:"compartmentId"`
		ContentSource      transferrunbookcontentdetails        `json:"contentSource"`
		ContentDestination transferrunbookcontentdetails        `json:"contentDestination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ImportAs.UnmarshalPolymorphicJSON(model.ImportAs.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ImportAs = nn.(RunbookImportAs)
	} else {
		m.ImportAs = nil
	}

	m.RequiredValues = make([]ImportRunbookPrecheckRequiredValue, len(model.RequiredValues))
	copy(m.RequiredValues, model.RequiredValues)
	m.CompartmentId = model.CompartmentId

	nn, e = model.ContentSource.UnmarshalPolymorphicJSON(model.ContentSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContentSource = nn.(TransferRunbookContentDetails)
	} else {
		m.ContentSource = nil
	}

	nn, e = model.ContentDestination.UnmarshalPolymorphicJSON(model.ContentDestination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContentDestination = nn.(TransferRunbookContentDetails)
	} else {
		m.ContentDestination = nil
	}

	return
}
