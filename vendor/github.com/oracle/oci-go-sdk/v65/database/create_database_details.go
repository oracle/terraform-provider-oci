// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatabaseDetails Details for creating a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseDetails struct {

	// The database name. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"true" json:"dbName"`

	// A strong password for SYS, SYSTEM, and PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The database software image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName *string `mandatory:"false" json:"pdbName"`

	// The optional password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	TdeWalletPassword *string `mandatory:"false" json:"tdeWalletPassword"`

	// The character set for the database.  The default is AL32UTF8. Allowed values are:
	// AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character set for the database.  The default is AL16UTF16. Allowed values are:
	// AL16UTF16 or UTF8.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service.
	// Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.
	// The database workload type.
	DbWorkload CreateDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	EncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"encryptionKeyLocationDetails"`
}

func (m CreateDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDatabaseDetailsDbWorkloadEnum(string(m.DbWorkload)); !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetCreateDatabaseDetailsDbWorkloadEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DbUniqueName                 *string                             `json:"dbUniqueName"`
		DatabaseSoftwareImageId      *string                             `json:"databaseSoftwareImageId"`
		PdbName                      *string                             `json:"pdbName"`
		TdeWalletPassword            *string                             `json:"tdeWalletPassword"`
		CharacterSet                 *string                             `json:"characterSet"`
		NcharacterSet                *string                             `json:"ncharacterSet"`
		DbWorkload                   CreateDatabaseDetailsDbWorkloadEnum `json:"dbWorkload"`
		DbBackupConfig               *DbBackupConfig                     `json:"dbBackupConfig"`
		FreeformTags                 map[string]string                   `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{}   `json:"definedTags"`
		KmsKeyId                     *string                             `json:"kmsKeyId"`
		KmsKeyVersionId              *string                             `json:"kmsKeyVersionId"`
		VaultId                      *string                             `json:"vaultId"`
		SidPrefix                    *string                             `json:"sidPrefix"`
		KeyStoreId                   *string                             `json:"keyStoreId"`
		EncryptionKeyLocationDetails encryptionkeylocationdetails        `json:"encryptionKeyLocationDetails"`
		DbName                       *string                             `json:"dbName"`
		AdminPassword                *string                             `json:"adminPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DbUniqueName = model.DbUniqueName

	m.DatabaseSoftwareImageId = model.DatabaseSoftwareImageId

	m.PdbName = model.PdbName

	m.TdeWalletPassword = model.TdeWalletPassword

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.DbWorkload = model.DbWorkload

	m.DbBackupConfig = model.DbBackupConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.KmsKeyId = model.KmsKeyId

	m.KmsKeyVersionId = model.KmsKeyVersionId

	m.VaultId = model.VaultId

	m.SidPrefix = model.SidPrefix

	m.KeyStoreId = model.KeyStoreId

	nn, e = model.EncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.EncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.EncryptionKeyLocationDetails = nil
	}

	m.DbName = model.DbName

	m.AdminPassword = model.AdminPassword

	return
}

// CreateDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type CreateDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for CreateDatabaseDetailsDbWorkloadEnum
const (
	CreateDatabaseDetailsDbWorkloadOltp CreateDatabaseDetailsDbWorkloadEnum = "OLTP"
	CreateDatabaseDetailsDbWorkloadDss  CreateDatabaseDetailsDbWorkloadEnum = "DSS"
)

var mappingCreateDatabaseDetailsDbWorkloadEnum = map[string]CreateDatabaseDetailsDbWorkloadEnum{
	"OLTP": CreateDatabaseDetailsDbWorkloadOltp,
	"DSS":  CreateDatabaseDetailsDbWorkloadDss,
}

var mappingCreateDatabaseDetailsDbWorkloadEnumLowerCase = map[string]CreateDatabaseDetailsDbWorkloadEnum{
	"oltp": CreateDatabaseDetailsDbWorkloadOltp,
	"dss":  CreateDatabaseDetailsDbWorkloadDss,
}

// GetCreateDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for CreateDatabaseDetailsDbWorkloadEnum
func GetCreateDatabaseDetailsDbWorkloadEnumValues() []CreateDatabaseDetailsDbWorkloadEnum {
	values := make([]CreateDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingCreateDatabaseDetailsDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseDetailsDbWorkloadEnumStringValues Enumerates the set of values in String for CreateDatabaseDetailsDbWorkloadEnum
func GetCreateDatabaseDetailsDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DSS",
	}
}

// GetMappingCreateDatabaseDetailsDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseDetailsDbWorkloadEnum(val string) (CreateDatabaseDetailsDbWorkloadEnum, bool) {
	enum, ok := mappingCreateDatabaseDetailsDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
