// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoteReindexSourceDetails Connection and certificate details for a remote cluster.
type RemoteReindexSourceDetails struct {

	// HTTPS endpoint of the remote Elasticsearch or OpenSearch cluster. The hostname and port are used to allow remote reindexing. Use a DNS hostname, rather than an IP address, when endpointResolutionIp is specified.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// OCID of the certificate or CA chain used to verify the remote cluster's HTTPS certificate.
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// Optional IP address to which the endpoint hostname is resolved from within the OpenSearch cluster. The original hostname remains in the connection URL and is used for TLS verification.
	EndpointResolutionIp *string `mandatory:"false" json:"endpointResolutionIp"`
}

func (m RemoteReindexSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoteReindexSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
