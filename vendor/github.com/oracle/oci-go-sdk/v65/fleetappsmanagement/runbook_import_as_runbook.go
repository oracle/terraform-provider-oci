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

// RunbookImportAsRunbook Import as Runbook.
type RunbookImportAsRunbook struct {

	// Only for RUNBOOK importType.
	RunbookName *string `mandatory:"true" json:"runbookName"`

	// Version number.
	Version *string `mandatory:"false" json:"version"`

	// ImportOptions for Runbook.
	ImportOption ImportRunbookDetailsImportOptionEnum `mandatory:"true" json:"importOption"`
}

// GetVersion returns Version
func (m RunbookImportAsRunbook) GetVersion() *string {
	return m.Version
}

// GetImportOption returns ImportOption
func (m RunbookImportAsRunbook) GetImportOption() ImportRunbookDetailsImportOptionEnum {
	return m.ImportOption
}

func (m RunbookImportAsRunbook) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunbookImportAsRunbook) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImportRunbookDetailsImportOptionEnum(string(m.ImportOption)); !ok && m.ImportOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImportOption: %s. Supported values are: %s.", m.ImportOption, strings.Join(GetImportRunbookDetailsImportOptionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RunbookImportAsRunbook) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRunbookImportAsRunbook RunbookImportAsRunbook
	s := struct {
		DiscriminatorParam string `json:"importType"`
		MarshalTypeRunbookImportAsRunbook
	}{
		"RUNBOOK",
		(MarshalTypeRunbookImportAsRunbook)(m),
	}

	return json.Marshal(&s)
}
