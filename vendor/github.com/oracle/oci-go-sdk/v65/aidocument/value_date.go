// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValueDate The date field value.
type ValueDate struct {

	// The confidence score between 0 and 1.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`

	// The indexes of the words in the field value.
	WordIndexes []int `mandatory:"true" json:"wordIndexes"`

	// The date field value as yyyy-mm-dd.
	Value *common.SDKTime `mandatory:"true" json:"value"`

	// The detected text of a field.
	Text *string `mandatory:"false" json:"text"`
}

// GetText returns Text
func (m ValueDate) GetText() *string {
	return m.Text
}

// GetConfidence returns Confidence
func (m ValueDate) GetConfidence() *float32 {
	return m.Confidence
}

// GetBoundingPolygon returns BoundingPolygon
func (m ValueDate) GetBoundingPolygon() *BoundingPolygon {
	return m.BoundingPolygon
}

// GetWordIndexes returns WordIndexes
func (m ValueDate) GetWordIndexes() []int {
	return m.WordIndexes
}

func (m ValueDate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValueDate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ValueDate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValueDate ValueDate
	s := struct {
		DiscriminatorParam string `json:"valueType"`
		MarshalTypeValueDate
	}{
		"DATE",
		(MarshalTypeValueDate)(m),
	}

	return json.Marshal(&s)
}
