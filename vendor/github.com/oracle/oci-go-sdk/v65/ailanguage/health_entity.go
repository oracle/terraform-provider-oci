// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HealthEntity Health entity object
type HealthEntity struct {

	// The number of Unicode code points preceding this entity in the submitted text.
	Offset *int `mandatory:"true" json:"offset"`

	// Length of entity text
	Length *int `mandatory:"true" json:"length"`

	// Entity text like name of person, location, and so on.
	Text *string `mandatory:"true" json:"text"`

	// Type of entity text like PER, LOC.
	Type *string `mandatory:"true" json:"type"`

	// Score or confidence for detected entity.
	Score *float64 `mandatory:"true" json:"score"`

	// Unique id of the entity
	Id *string `mandatory:"true" json:"id"`

	// Sub-type of entity text like GPE for LOCATION type
	SubType *string `mandatory:"false" json:"subType"`

	// Entity category e.g, MEDICAL_CONDITION, MEDICATION, GENERAL, ANATOMY
	Category *string `mandatory:"false" json:"category"`

	// list of all assertions associated with this entity.
	Assertions []AssertionDetails `mandatory:"false" json:"assertions"`

	// This contains the list of matched concepts which are ranked by the relevant score with the input text
	MatchedConcepts []MelConcept `mandatory:"false" json:"matchedConcepts"`
}

func (m HealthEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
