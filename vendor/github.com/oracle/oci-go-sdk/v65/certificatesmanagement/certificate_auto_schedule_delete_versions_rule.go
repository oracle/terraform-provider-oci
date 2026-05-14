// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateAutoScheduleDeleteVersionsRule A rule that automatically schedules the deletion of deprecated certificate versions.
type CertificateAutoScheduleDeleteVersionsRule struct {

	// A property that specifies when, in days, to delete the certificate version.
	// The minimum value is 1 day, and the maximum is 30 days, expressed in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format.
	ScheduleDeleteDuration *string `mandatory:"true" json:"scheduleDeleteDuration"`
}

func (m CertificateAutoScheduleDeleteVersionsRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateAutoScheduleDeleteVersionsRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CertificateAutoScheduleDeleteVersionsRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateAutoScheduleDeleteVersionsRule CertificateAutoScheduleDeleteVersionsRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeCertificateAutoScheduleDeleteVersionsRule
	}{
		"CERTIFICATE_AUTO_SCHEDULE_DELETE_VERSIONS_RULE",
		(MarshalTypeCertificateAutoScheduleDeleteVersionsRule)(m),
	}

	return json.Marshal(&s)
}
