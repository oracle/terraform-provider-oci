// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RelationEntity The relation entities details for health ner.
type RelationEntity struct {

	// id of the relation
	Id *string `mandatory:"true" json:"id"`

	// Subject entity of the relation. eg,MEDICINE_NAME, EXAMINATION
	SubjectId *string `mandatory:"true" json:"subjectId"`

	// Object entity of the relation. eg, MEDICINE_DOSAGE, MEDICINE_DURATION, MEASUREMENT
	ObjectId *string `mandatory:"true" json:"objectId"`

	// Type of relation between subject entity and object entity. eg,EXAMINATION_MEASUREMENT, FREQUENCY, DOSAGE
	Type *string `mandatory:"true" json:"type"`

	// Score or confidence for detected entity.
	Score *float64 `mandatory:"true" json:"score"`
}

func (m RelationEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RelationEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
