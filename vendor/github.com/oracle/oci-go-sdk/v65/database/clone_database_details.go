// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CloneDatabaseDetails The representation of CloneDatabaseDetails
type CloneDatabaseDetails struct {

	// The source database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SourceDatabaseId *string `mandatory:"true" json:"sourceDatabaseId"`

	// The password to open the TDE wallet.
	SourceTdePassword *string `mandatory:"true" json:"sourceTdePassword"`

	// The password for the target Database TDE wallet.
	TdeWalletPassword *string `mandatory:"true" json:"tdeWalletPassword"`

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The display name of the database to be created from the snapshot/database. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"true" json:"dbName"`

	// Specifies whether thick or thin database cloning is to be performed.
	IsThinClone *bool `mandatory:"false" json:"isThinClone"`

	// The source database's password for SYS, SYSTEM.
	SourceAdminPassword *string `mandatory:"false" json:"sourceAdminPassword"`

	SourceEncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"sourceEncryptionKeyLocationDetails"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The `DB_UNIQUE_NAME` for the new Oracle Database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The character set for the database.  The default is AL32UTF8. Allowed values are:
	// AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	// The national character set for the database.  The default is AL16UTF16. Allowed values are:
	// AL16UTF16 or UTF8.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	StorageSizeDetails *DatabaseStorageSizeDetails `mandatory:"false" json:"storageSizeDetails"`

	ManagedSoftwareUpdateDetails *ManagedSoftwareUpdateInputDetails `mandatory:"false" json:"managedSoftwareUpdateDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Specifies whether to register the database with Data Safe.
	IsDataSafeRegistered *bool `mandatory:"false" json:"isDataSafeRegistered"`
}

func (m CloneDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloneDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CloneDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsThinClone                        *bool                              `json:"isThinClone"`
		SourceAdminPassword                *string                            `json:"sourceAdminPassword"`
		SourceEncryptionKeyLocationDetails encryptionkeylocationdetails       `json:"sourceEncryptionKeyLocationDetails"`
		DbBackupConfig                     *DbBackupConfig                    `json:"dbBackupConfig"`
		DbUniqueName                       *string                            `json:"dbUniqueName"`
		CharacterSet                       *string                            `json:"characterSet"`
		NcharacterSet                      *string                            `json:"ncharacterSet"`
		SidPrefix                          *string                            `json:"sidPrefix"`
		KeyStoreId                         *string                            `json:"keyStoreId"`
		StorageSizeDetails                 *DatabaseStorageSizeDetails        `json:"storageSizeDetails"`
		ManagedSoftwareUpdateDetails       *ManagedSoftwareUpdateInputDetails `json:"managedSoftwareUpdateDetails"`
		VmClusterId                        *string                            `json:"vmClusterId"`
		FreeformTags                       map[string]string                  `json:"freeformTags"`
		DefinedTags                        map[string]map[string]interface{}  `json:"definedTags"`
		IsDataSafeRegistered               *bool                              `json:"isDataSafeRegistered"`
		SourceDatabaseId                   *string                            `json:"sourceDatabaseId"`
		SourceTdePassword                  *string                            `json:"sourceTdePassword"`
		TdeWalletPassword                  *string                            `json:"tdeWalletPassword"`
		AdminPassword                      *string                            `json:"adminPassword"`
		DbName                             *string                            `json:"dbName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsThinClone = model.IsThinClone

	m.SourceAdminPassword = model.SourceAdminPassword

	nn, e = model.SourceEncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.SourceEncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.SourceEncryptionKeyLocationDetails = nil
	}

	m.DbBackupConfig = model.DbBackupConfig

	m.DbUniqueName = model.DbUniqueName

	m.CharacterSet = model.CharacterSet

	m.NcharacterSet = model.NcharacterSet

	m.SidPrefix = model.SidPrefix

	m.KeyStoreId = model.KeyStoreId

	m.StorageSizeDetails = model.StorageSizeDetails

	m.ManagedSoftwareUpdateDetails = model.ManagedSoftwareUpdateDetails

	m.VmClusterId = model.VmClusterId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IsDataSafeRegistered = model.IsDataSafeRegistered

	m.SourceDatabaseId = model.SourceDatabaseId

	m.SourceTdePassword = model.SourceTdePassword

	m.TdeWalletPassword = model.TdeWalletPassword

	m.AdminPassword = model.AdminPassword

	m.DbName = model.DbName

	return
}
