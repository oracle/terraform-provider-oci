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

// DatabaseWalletDetails Details for database wallet.
type DatabaseWalletDetails struct {

	// The database wallet configuration zip file.
	DatabaseWallet *string `mandatory:"true" json:"databaseWallet"`

	// Service name of the database.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	WalletPassword Password `mandatory:"true" json:"walletPassword"`
}

func (m DatabaseWalletDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseWalletDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseWalletDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DatabaseWallet *string  `json:"databaseWallet"`
		ServiceName    *string  `json:"serviceName"`
		WalletPassword password `json:"walletPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DatabaseWallet = model.DatabaseWallet

	m.ServiceName = model.ServiceName

	nn, e = model.WalletPassword.UnmarshalPolymorphicJSON(model.WalletPassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WalletPassword = nn.(Password)
	} else {
		m.WalletPassword = nil
	}

	return
}
