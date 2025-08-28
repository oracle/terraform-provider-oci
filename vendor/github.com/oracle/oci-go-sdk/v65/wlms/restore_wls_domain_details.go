// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreWlsDomainDetails The configuration details for the restore patches to WebLogic domain operation.
type RestoreWlsDomainDetails struct {

	// If the restore operation required to stop the servers, this option indicates to force shutdown of the servers if they have not shutdown after a period of time. The timeout can be configured in the WebLogic domain configuration.
	IsForceServersShutdown *bool `mandatory:"false" json:"isForceServersShutdown"`

	// If the WebLogic Domain being restored shares middleware with other domains, and the restore operation will restore servers in the other domains, this option allows the operation to proceed on all domains that share the same middleware. If not set to true, the restore operation will fail if there is any other domain using the same middleware.
	MustIncludeDomainsSharingMiddleware *bool `mandatory:"false" json:"mustIncludeDomainsSharingMiddleware"`

	// List of servers to restore and the backups
	ServersToRestore []ServerToRestore `mandatory:"false" json:"serversToRestore"`
}

func (m RestoreWlsDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreWlsDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
