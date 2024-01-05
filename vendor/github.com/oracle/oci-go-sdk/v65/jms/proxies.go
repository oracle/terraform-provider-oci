// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Proxies List of proxy properties to be configured in net.properties file.
type Proxies struct {

	// Sets "java.net.useSystemProxies=true" in net.properties when they exist.
	UseSystemProxies *bool `mandatory:"false" json:"useSystemProxies"`

	// Http host to be set in net.properties file.
	HttpProxyHost *string `mandatory:"false" json:"httpProxyHost"`

	// Http port number to be set in net.properties file.
	HttpProxyPort *int `mandatory:"false" json:"httpProxyPort"`

	// Https host to be set in net.properties file.
	HttpsProxyHost *string `mandatory:"false" json:"httpsProxyHost"`

	// Https port number to be set in net.properties file.
	HttpsProxyPort *int `mandatory:"false" json:"httpsProxyPort"`

	// Ftp host to be set in net.properties file.
	FtpProxyHost *string `mandatory:"false" json:"ftpProxyHost"`

	// Ftp port number to be set in net.properties file.
	FtpProxyPort *int `mandatory:"false" json:"ftpProxyPort"`

	// Socks host to be set in net.properties file.
	SocksProxyHost *string `mandatory:"false" json:"socksProxyHost"`

	// Socks port number to be set in net.properties file.
	SocksProxyPort *int `mandatory:"false" json:"socksProxyPort"`
}

func (m Proxies) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Proxies) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
