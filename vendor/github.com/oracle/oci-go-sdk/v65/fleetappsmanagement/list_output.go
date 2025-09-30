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

// ListOutput List output.
type ListOutput struct {

	// Output label shown to the user.
	Title *string `mandatory:"false" json:"title"`

	// Extended help or summary for understanding output.
	Description *string `mandatory:"false" json:"description"`

	// If true, marks this output as sensitive.
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`

	// Hint about formatting or rendering the output value.
	Format *string `mandatory:"false" json:"format"`

	// Expression to show/hide this output.
	Visible *string `mandatory:"false" json:"visible"`

	// Array of objects representing output data.
	Value []interface{} `mandatory:"false" json:"value"`
}

// GetTitle returns Title
func (m ListOutput) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m ListOutput) GetDescription() *string {
	return m.Description
}

// GetIsSensitive returns IsSensitive
func (m ListOutput) GetIsSensitive() *bool {
	return m.IsSensitive
}

// GetFormat returns Format
func (m ListOutput) GetFormat() *string {
	return m.Format
}

// GetVisible returns Visible
func (m ListOutput) GetVisible() *string {
	return m.Visible
}

func (m ListOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ListOutput) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeListOutput ListOutput
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeListOutput
	}{
		"LIST",
		(MarshalTypeListOutput)(m),
	}

	return json.Marshal(&s)
}
