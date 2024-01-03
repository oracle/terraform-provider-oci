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

// ScriptedBrowserMonitorConfiguration Configuration details for the SCRIPTED_BROWSER monitor type.
type ScriptedBrowserMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// If certificate validation is enabled, then the call will fail in case of certification errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	// If disabled, auto snapshots are not collected.
	IsDefaultSnapshotEnabled *bool `mandatory:"false" json:"isDefaultSnapshotEnabled"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m ScriptedBrowserMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m ScriptedBrowserMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m ScriptedBrowserMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScriptedBrowserMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScriptedBrowserMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptedBrowserMonitorConfiguration ScriptedBrowserMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeScriptedBrowserMonitorConfiguration
	}{
		"SCRIPTED_BROWSER_CONFIG",
		(MarshalTypeScriptedBrowserMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
