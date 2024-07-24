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

// FtpMonitorConfiguration Request configuration details for the FTP monitor type.
type FtpMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// If enabled, Active mode will be used for the FTP connection.
	IsActiveMode *bool `mandatory:"false" json:"isActiveMode"`

	FtpBasicAuthenticationDetails *BasicAuthenticationDetails `mandatory:"false" json:"ftpBasicAuthenticationDetails"`

	// Download size limit in Bytes, at which to stop the transfer. Maximum download size limit is 5 MiB.
	DownloadSizeLimitInBytes *int `mandatory:"false" json:"downloadSizeLimitInBytes"`

	// File upload size in Bytes, at which to stop the transfer. Maximum upload size is 5 MiB.
	UploadFileSizeInBytes *int `mandatory:"false" json:"uploadFileSizeInBytes"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// Expected FTP response codes. For status code range, set values such as 2xx, 3xx.
	VerifyResponseCodes []string `mandatory:"false" json:"verifyResponseCodes"`

	// Verify response content against regular expression based string.
	// If response content does not match the verifyResponseContent value, then it will be considered a failure.
	VerifyResponseContent *string `mandatory:"false" json:"verifyResponseContent"`

	// FTP protocol type.
	FtpProtocol FtpProtocolEnum `mandatory:"false" json:"ftpProtocol,omitempty"`

	// FTP monitor request type.
	FtpRequestType FtpRequestTypeEnum `mandatory:"false" json:"ftpRequestType,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m FtpMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m FtpMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m FtpMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FtpMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFtpProtocolEnum(string(m.FtpProtocol)); !ok && m.FtpProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FtpProtocol: %s. Supported values are: %s.", m.FtpProtocol, strings.Join(GetFtpProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFtpRequestTypeEnum(string(m.FtpRequestType)); !ok && m.FtpRequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FtpRequestType: %s. Supported values are: %s.", m.FtpRequestType, strings.Join(GetFtpRequestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FtpMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFtpMonitorConfiguration FtpMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeFtpMonitorConfiguration
	}{
		"FTP_CONFIG",
		(MarshalTypeFtpMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
