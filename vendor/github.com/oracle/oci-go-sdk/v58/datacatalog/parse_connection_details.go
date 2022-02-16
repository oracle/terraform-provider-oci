// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ParseConnectionDetails Parse connections from the connection metadata and Oracle wallet file.
// An error will be returned if more than one of connectionPayload, walletSecretId or walletSecretName are present in the request.
type ParseConnectionDetails struct {
	ConnectionDetail *Connection `mandatory:"false" json:"connectionDetail"`

	// The information used to parse the connection from the wallet file payload.
	ConnectionPayload []byte `mandatory:"false" json:"connectionPayload"`

	// OCID of the OCI Vault secret holding the Oracle wallet to parse.
	WalletSecretId *string `mandatory:"false" json:"walletSecretId"`

	// Name of the OCI Vault secret holding the Oracle wallet to parse.
	WalletSecretName *string `mandatory:"false" json:"walletSecretName"`
}

func (m ParseConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ParseConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
