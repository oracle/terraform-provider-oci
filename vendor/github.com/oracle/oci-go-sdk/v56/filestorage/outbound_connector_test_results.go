// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"strings"
)

// OutboundConnectorTestResults Diagnostic test results between mount target and LDAP server using outbound connector.
type OutboundConnectorTestResults struct {

	// Status of overall connection test against all endpoints.
	// Success represents all endpoints tested and functioning.
	// PartialSuccess represents all endpoints tested, some function properly, some do not.
	// Fail represents no endpoints are functioning properly, MT cannot communicate with LDAP server.
	OverallResult OutboundConnectorTestResultsOverallResultEnum `mandatory:"true" json:"overallResult"`

	// Array of result of connecting to each endpoint
	ResultsDetail []TestOutboundConnectorResult `mandatory:"true" json:"resultsDetail"`
}

func (m OutboundConnectorTestResults) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OutboundConnectorTestResults) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingOutboundConnectorTestResultsOverallResultEnum[string(m.OverallResult)]; !ok && m.OverallResult != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OverallResult: %s. Supported values are: %s.", m.OverallResult, strings.Join(GetOutboundConnectorTestResultsOverallResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutboundConnectorTestResultsOverallResultEnum Enum with underlying type: string
type OutboundConnectorTestResultsOverallResultEnum string

// Set of constants representing the allowable values for OutboundConnectorTestResultsOverallResultEnum
const (
	OutboundConnectorTestResultsOverallResultSuccess        OutboundConnectorTestResultsOverallResultEnum = "SUCCESS"
	OutboundConnectorTestResultsOverallResultPartialsuccess OutboundConnectorTestResultsOverallResultEnum = "PARTIALSUCCESS"
	OutboundConnectorTestResultsOverallResultFail           OutboundConnectorTestResultsOverallResultEnum = "FAIL"
)

var mappingOutboundConnectorTestResultsOverallResultEnum = map[string]OutboundConnectorTestResultsOverallResultEnum{
	"SUCCESS":        OutboundConnectorTestResultsOverallResultSuccess,
	"PARTIALSUCCESS": OutboundConnectorTestResultsOverallResultPartialsuccess,
	"FAIL":           OutboundConnectorTestResultsOverallResultFail,
}

// GetOutboundConnectorTestResultsOverallResultEnumValues Enumerates the set of values for OutboundConnectorTestResultsOverallResultEnum
func GetOutboundConnectorTestResultsOverallResultEnumValues() []OutboundConnectorTestResultsOverallResultEnum {
	values := make([]OutboundConnectorTestResultsOverallResultEnum, 0)
	for _, v := range mappingOutboundConnectorTestResultsOverallResultEnum {
		values = append(values, v)
	}
	return values
}

// GetOutboundConnectorTestResultsOverallResultEnumStringValues Enumerates the set of values in String for OutboundConnectorTestResultsOverallResultEnum
func GetOutboundConnectorTestResultsOverallResultEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"PARTIALSUCCESS",
		"FAIL",
	}
}
