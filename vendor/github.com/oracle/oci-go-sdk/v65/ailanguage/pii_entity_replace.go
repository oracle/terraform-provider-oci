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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PiiEntityReplace Replace PII entities with a given sequence of characters.
type PiiEntityReplace struct {

	// List of offsets/entities to be removed from anonymization.
	Exclude []string `mandatory:"false" json:"exclude"`

	// To include excluded entities from masking in detected entities or not.
	ShouldDetect *bool `mandatory:"false" json:"shouldDetect"`

	// Replace entities with given sequence of characters. By default PII entity will be replaced with <ENTITY_TYPE>.
	ReplaceWith *string `mandatory:"false" json:"replaceWith"`
}

// GetExclude returns Exclude
func (m PiiEntityReplace) GetExclude() []string {
	return m.Exclude
}

// GetShouldDetect returns ShouldDetect
func (m PiiEntityReplace) GetShouldDetect() *bool {
	return m.ShouldDetect
}

func (m PiiEntityReplace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PiiEntityReplace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PiiEntityReplace) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePiiEntityReplace PiiEntityReplace
	s := struct {
		DiscriminatorParam string `json:"mode"`
		MarshalTypePiiEntityReplace
	}{
		"REPLACE",
		(MarshalTypePiiEntityReplace)(m),
	}

	return json.Marshal(&s)
}
