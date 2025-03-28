// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Record A DNS resource record. For more information, see
// Supported DNS Resource Record Types (https://docs.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm).
type Record struct {

	// The fully qualified domain name where the record can be located.
	Domain *string `mandatory:"false" json:"domain"`

	// A unique identifier for the record within its zone.
	RecordHash *string `mandatory:"false" json:"recordHash"`

	// A Boolean flag indicating whether or not parts of the record
	// are unable to be explicitly managed.
	IsProtected *bool `mandatory:"false" json:"isProtected"`

	// The record's data, as whitespace-delimited tokens in
	// type-specific presentation format. All RDATA is normalized and the
	// returned presentation of your RDATA may differ from its initial input.
	// For more information about RDATA, see Supported DNS Resource Record Types (https://docs.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)
	Rdata *string `mandatory:"false" json:"rdata"`

	// The latest version of the record's zone in which its RRSet differs
	// from the preceding version.
	RrsetVersion *string `mandatory:"false" json:"rrsetVersion"`

	// The type of DNS record, such as A or CNAME. For more information, see Resource Record (RR) TYPEs (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4).
	Rtype *string `mandatory:"false" json:"rtype"`

	// The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended.
	Ttl *int `mandatory:"false" json:"ttl"`
}

func (m Record) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Record) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
