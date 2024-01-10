// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScriptedRestMonitorConfiguration Configuration details for the SCRIPTED_REST monitor type.
type ScriptedRestMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// Expected HTTP response codes. For status code range, set values such as 2xx, 3xx.
	VerifyResponseCodes []string `mandatory:"false" json:"verifyResponseCodes"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// Request HTTP authentication scheme.
	ReqAuthenticationScheme RequestAuthenticationSchemesForScriptedRestEnum `mandatory:"false" json:"reqAuthenticationScheme,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m ScriptedRestMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m ScriptedRestMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m ScriptedRestMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScriptedRestMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequestAuthenticationSchemesForScriptedRestEnum(string(m.ReqAuthenticationScheme)); !ok && m.ReqAuthenticationScheme != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReqAuthenticationScheme: %s. Supported values are: %s.", m.ReqAuthenticationScheme, strings.Join(GetRequestAuthenticationSchemesForScriptedRestEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScriptedRestMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptedRestMonitorConfiguration ScriptedRestMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScriptedRestMonitorConfiguration
	}{
		"SCRIPTED_REST_CONFIG",
		(MarshalTypeScriptedRestMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
