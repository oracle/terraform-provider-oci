// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DataSourceDetailsAtp Data Source details for ATP
type DataSourceDetailsAtp struct {

	// wallet password Secret ID in String format
	WalletPasswordSecretId *string `mandatory:"false" json:"walletPasswordSecretId"`

	// atp db user name
	AtpUserName *string `mandatory:"false" json:"atpUserName"`

	// atp db password Secret Id
	AtpPasswordSecretId *string `mandatory:"false" json:"atpPasswordSecretId"`

	// OCID of the secret containing the containers certificates of ATP wallet
	CwalletFileSecretId *string `mandatory:"false" json:"cwalletFileSecretId"`

	// OCID of the secret containing the PDB'S certificates of ATP wallet
	EwalletFileSecretId *string `mandatory:"false" json:"ewalletFileSecretId"`

	// OCID of the secret containing Keystore.jks file of the ATP wallet
	KeyStoreFileSecretId *string `mandatory:"false" json:"keyStoreFileSecretId"`

	// OCID of the secret that contains jdbc properties file of ATP wallet
	OjdbcFileSecretId *string `mandatory:"false" json:"ojdbcFileSecretId"`

	// OCID of the secret that contains the tnsnames file of ATP wallet
	TnsnamesFileSecretId *string `mandatory:"false" json:"tnsnamesFileSecretId"`

	// OCID of the secret containing truststore.jks file of the ATP wallet
	TruststoreFileSecretId *string `mandatory:"false" json:"truststoreFileSecretId"`

	// atp database name
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// atp database table name
	TableName *string `mandatory:"false" json:"tableName"`
}

func (m DataSourceDetailsAtp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSourceDetailsAtp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataSourceDetailsAtp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataSourceDetailsAtp DataSourceDetailsAtp
	s := struct {
		DiscriminatorParam string `json:"dataSourceType"`
		MarshalTypeDataSourceDetailsAtp
	}{
		"ORACLE_ATP",
		(MarshalTypeDataSourceDetailsAtp)(m),
	}

	return json.Marshal(&s)
}
