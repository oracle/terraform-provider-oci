// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// BatchDetectHealthEntityDetails The documents details for health entities detect call.
type BatchDetectHealthEntityDetails struct {

	// List of Documents for detect health entities.
	Documents []TextDocument `mandatory:"true" json:"documents"`

	// List of NLP health ontologies to be linked
	LinkOntologies []string `mandatory:"false" json:"linkOntologies"`

	// is assertion on input text required. default value true.
	IsDetectAssertions *bool `mandatory:"false" json:"isDetectAssertions"`

	// is relationship on input text required. default value true.
	IsDetectRelationships *bool `mandatory:"false" json:"isDetectRelationships"`

	Profile *Profile `mandatory:"false" json:"profile"`

	// Unique identifier model OCID of a model that should be used for inference
	ModelId *string `mandatory:"false" json:"modelId"`
}

func (m BatchDetectHealthEntityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchDetectHealthEntityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
