// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"strings"
)

// DocumentLanguageEnum Enum with underlying type: string
type DocumentLanguageEnum string

// Set of constants representing the allowable values for DocumentLanguageEnum
const (
	DocumentLanguageEng    DocumentLanguageEnum = "ENG"
	DocumentLanguageCes    DocumentLanguageEnum = "CES"
	DocumentLanguageDan    DocumentLanguageEnum = "DAN"
	DocumentLanguageNld    DocumentLanguageEnum = "NLD"
	DocumentLanguageFin    DocumentLanguageEnum = "FIN"
	DocumentLanguageFra    DocumentLanguageEnum = "FRA"
	DocumentLanguageDeu    DocumentLanguageEnum = "DEU"
	DocumentLanguageEll    DocumentLanguageEnum = "ELL"
	DocumentLanguageHun    DocumentLanguageEnum = "HUN"
	DocumentLanguageIta    DocumentLanguageEnum = "ITA"
	DocumentLanguageNor    DocumentLanguageEnum = "NOR"
	DocumentLanguagePol    DocumentLanguageEnum = "POL"
	DocumentLanguagePor    DocumentLanguageEnum = "POR"
	DocumentLanguageRon    DocumentLanguageEnum = "RON"
	DocumentLanguageRus    DocumentLanguageEnum = "RUS"
	DocumentLanguageSlk    DocumentLanguageEnum = "SLK"
	DocumentLanguageSpa    DocumentLanguageEnum = "SPA"
	DocumentLanguageSwe    DocumentLanguageEnum = "SWE"
	DocumentLanguageTur    DocumentLanguageEnum = "TUR"
	DocumentLanguageAra    DocumentLanguageEnum = "ARA"
	DocumentLanguageChiSim DocumentLanguageEnum = "CHI_SIM"
	DocumentLanguageHin    DocumentLanguageEnum = "HIN"
	DocumentLanguageJpn    DocumentLanguageEnum = "JPN"
	DocumentLanguageKor    DocumentLanguageEnum = "KOR"
	DocumentLanguageOthers DocumentLanguageEnum = "OTHERS"
)

var mappingDocumentLanguageEnum = map[string]DocumentLanguageEnum{
	"ENG":     DocumentLanguageEng,
	"CES":     DocumentLanguageCes,
	"DAN":     DocumentLanguageDan,
	"NLD":     DocumentLanguageNld,
	"FIN":     DocumentLanguageFin,
	"FRA":     DocumentLanguageFra,
	"DEU":     DocumentLanguageDeu,
	"ELL":     DocumentLanguageEll,
	"HUN":     DocumentLanguageHun,
	"ITA":     DocumentLanguageIta,
	"NOR":     DocumentLanguageNor,
	"POL":     DocumentLanguagePol,
	"POR":     DocumentLanguagePor,
	"RON":     DocumentLanguageRon,
	"RUS":     DocumentLanguageRus,
	"SLK":     DocumentLanguageSlk,
	"SPA":     DocumentLanguageSpa,
	"SWE":     DocumentLanguageSwe,
	"TUR":     DocumentLanguageTur,
	"ARA":     DocumentLanguageAra,
	"CHI_SIM": DocumentLanguageChiSim,
	"HIN":     DocumentLanguageHin,
	"JPN":     DocumentLanguageJpn,
	"KOR":     DocumentLanguageKor,
	"OTHERS":  DocumentLanguageOthers,
}

// GetDocumentLanguageEnumValues Enumerates the set of values for DocumentLanguageEnum
func GetDocumentLanguageEnumValues() []DocumentLanguageEnum {
	values := make([]DocumentLanguageEnum, 0)
	for _, v := range mappingDocumentLanguageEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentLanguageEnumStringValues Enumerates the set of values in String for DocumentLanguageEnum
func GetDocumentLanguageEnumStringValues() []string {
	return []string{
		"ENG",
		"CES",
		"DAN",
		"NLD",
		"FIN",
		"FRA",
		"DEU",
		"ELL",
		"HUN",
		"ITA",
		"NOR",
		"POL",
		"POR",
		"RON",
		"RUS",
		"SLK",
		"SPA",
		"SWE",
		"TUR",
		"ARA",
		"CHI_SIM",
		"HIN",
		"JPN",
		"KOR",
		"OTHERS",
	}
}

// GetMappingDocumentLanguageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentLanguageEnum(val string) (DocumentLanguageEnum, bool) {
	mappingDocumentLanguageEnumIgnoreCase := make(map[string]DocumentLanguageEnum)
	for k, v := range mappingDocumentLanguageEnum {
		mappingDocumentLanguageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDocumentLanguageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
