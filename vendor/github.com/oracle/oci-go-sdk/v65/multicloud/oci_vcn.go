// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciVcn Oracle Cloud Infrastructure VCN basic information object. It is optional and planned to used for future for network anchor
type OciVcn struct {

	// Oracle Cloud Infrastructure VCN OCID. CSP can not set this property.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// Oracle Cloud Infrastructure primary cidr block. CSP can set this property
	// It's optional only if disconnect anchor is allowed
	// IPv4 CIDR blocks for the VCN that meet the following criteria
	// Type: [string (length: 1–32), ...]
	// The CIDR blocks must be valid.
	// They must not overlap with each other or with the on-premises network CIDR block.
	CidrBlocks []string `mandatory:"false" json:"cidrBlocks"`

	// Oracle Cloud Infrastructure backup cidr block. CSP can set this property
	// It's optional only if disconnect anchor is allowed.
	// IPv4 CIDR blocks for the VCN that meet the following criteria
	// Type: [string (length: 1–32), ...]
	// The CIDR blocks must be valid.
	// They must not overlap with each other or with the on-premises network CIDR block.
	BackupCidrBlocks []string `mandatory:"false" json:"backupCidrBlocks"`

	// Oracle Cloud Infrastructure DNS label. This is optional if DNS config is provided.
	DnsLabel *string `mandatory:"false" json:"dnsLabel"`
}

func (m OciVcn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciVcn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
