// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DesktopConnection Provides information about a connection to a desktop, including connect and disconnect time, and client properties.
type DesktopConnection struct {

	// The time when the last connection to a desktop started.
	TimeConnected *common.SDKTime `mandatory:"false" json:"timeConnected"`

	// The time when the last connection to a desktop ended.
	TimeDisconnected *common.SDKTime `mandatory:"false" json:"timeDisconnected"`

	NextAction *DesktopAction `mandatory:"false" json:"nextAction"`

	LastAction *DesktopAction `mandatory:"false" json:"lastAction"`

	// The type of Secure Desktops client connected to a desktop.
	ClientType *string `mandatory:"false" json:"clientType"`

	// The version of the Secure Desktops client connected to a desktop, applicable only to the installed client type.
	ClientVersion *string `mandatory:"false" json:"clientVersion"`

	// The platform on which the Secure Desktops client runs.
	ClientPlatform *string `mandatory:"false" json:"clientPlatform"`
}

func (m DesktopConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
