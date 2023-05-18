// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatapathMessageDetails Details to create data path message
type CreateDatapathMessageDetails struct {

	// Client-supplied unique identifier of the message
	Id *string `mandatory:"true" json:"id"`

	// Type of datapath message. Can be blockStorage or ipLearning.
	Type *string `mandatory:"true" json:"type"`

	// Datapath message in protobuf
	Content *string `mandatory:"true" json:"content"`

	// OCID of the primary VNIC for the instance. Only applies to blockstorage calls.
	VnicId *string `mandatory:"false" json:"vnicId"`

	// The substrate IP address of the smartNic where the VNIC is attached to and the message has to be delivered.
	// **Note:** This is a required field if vnic is getting live migrated.
	SubstrateIp *string `mandatory:"false" json:"substrateIp"`
}

func (m CreateDatapathMessageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatapathMessageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
