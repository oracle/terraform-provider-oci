// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BrowserMonitorConfiguration Configuration details for the BROWSER monitor type.
type BrowserMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// If certificate validation is enabled, then the call will fail in case of certification errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	// If disabled, auto snapshots are not collected.
	IsDefaultSnapshotEnabled *bool `mandatory:"false" json:"isDefaultSnapshotEnabled"`

	// Verifies all the search strings present in the response.
	// If any search string is not present in the response, then it will be considered as a failure.
	VerifyTexts []VerifyText `mandatory:"false" json:"verifyTexts"`

	// Expected HTTP response codes. For status code range, set values such as 2xx, 3xx.
	VerifyResponseCodes []string `mandatory:"false" json:"verifyResponseCodes"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m BrowserMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m BrowserMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m BrowserMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BrowserMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BrowserMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBrowserMonitorConfiguration BrowserMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeBrowserMonitorConfiguration
	}{
		"BROWSER_CONFIG",
		(MarshalTypeBrowserMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
