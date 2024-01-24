// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// PricingCurrencyEnumEnum Enum with underlying type: string
type PricingCurrencyEnumEnum string

// Set of constants representing the allowable values for PricingCurrencyEnumEnum
const (
	PricingCurrencyEnumUsd PricingCurrencyEnumEnum = "USD"
	PricingCurrencyEnumCad PricingCurrencyEnumEnum = "CAD"
	PricingCurrencyEnumInr PricingCurrencyEnumEnum = "INR"
	PricingCurrencyEnumGbp PricingCurrencyEnumEnum = "GBP"
	PricingCurrencyEnumBrl PricingCurrencyEnumEnum = "BRL"
	PricingCurrencyEnumJpy PricingCurrencyEnumEnum = "JPY"
	PricingCurrencyEnumOmr PricingCurrencyEnumEnum = "OMR"
	PricingCurrencyEnumEur PricingCurrencyEnumEnum = "EUR"
	PricingCurrencyEnumChf PricingCurrencyEnumEnum = "CHF"
	PricingCurrencyEnumMxn PricingCurrencyEnumEnum = "MXN"
	PricingCurrencyEnumClp PricingCurrencyEnumEnum = "CLP"
	PricingCurrencyEnumAll PricingCurrencyEnumEnum = "ALL"
	PricingCurrencyEnumArs PricingCurrencyEnumEnum = "ARS"
	PricingCurrencyEnumAud PricingCurrencyEnumEnum = "AUD"
	PricingCurrencyEnumBdt PricingCurrencyEnumEnum = "BDT"
	PricingCurrencyEnumBam PricingCurrencyEnumEnum = "BAM"
	PricingCurrencyEnumBgn PricingCurrencyEnumEnum = "BGN"
	PricingCurrencyEnumCny PricingCurrencyEnumEnum = "CNY"
	PricingCurrencyEnumCop PricingCurrencyEnumEnum = "COP"
	PricingCurrencyEnumCrc PricingCurrencyEnumEnum = "CRC"
	PricingCurrencyEnumHrk PricingCurrencyEnumEnum = "HRK"
	PricingCurrencyEnumCzk PricingCurrencyEnumEnum = "CZK"
	PricingCurrencyEnumDkk PricingCurrencyEnumEnum = "DKK"
	PricingCurrencyEnumEgp PricingCurrencyEnumEnum = "EGP"
	PricingCurrencyEnumHkd PricingCurrencyEnumEnum = "HKD"
	PricingCurrencyEnumHuf PricingCurrencyEnumEnum = "HUF"
	PricingCurrencyEnumIsk PricingCurrencyEnumEnum = "ISK"
	PricingCurrencyEnumIdr PricingCurrencyEnumEnum = "IDR"
	PricingCurrencyEnumIls PricingCurrencyEnumEnum = "ILS"
	PricingCurrencyEnumJmd PricingCurrencyEnumEnum = "JMD"
	PricingCurrencyEnumKzt PricingCurrencyEnumEnum = "KZT"
	PricingCurrencyEnumKes PricingCurrencyEnumEnum = "KES"
	PricingCurrencyEnumKrw PricingCurrencyEnumEnum = "KRW"
	PricingCurrencyEnumKwd PricingCurrencyEnumEnum = "KWD"
	PricingCurrencyEnumLbp PricingCurrencyEnumEnum = "LBP"
	PricingCurrencyEnumMop PricingCurrencyEnumEnum = "MOP"
	PricingCurrencyEnumMyr PricingCurrencyEnumEnum = "MYR"
	PricingCurrencyEnumMvr PricingCurrencyEnumEnum = "MVR"
	PricingCurrencyEnumAed PricingCurrencyEnumEnum = "AED"
	PricingCurrencyEnumNzd PricingCurrencyEnumEnum = "NZD"
	PricingCurrencyEnumNok PricingCurrencyEnumEnum = "NOK"
	PricingCurrencyEnumPkr PricingCurrencyEnumEnum = "PKR"
	PricingCurrencyEnumPen PricingCurrencyEnumEnum = "PEN"
	PricingCurrencyEnumPhp PricingCurrencyEnumEnum = "PHP"
	PricingCurrencyEnumPln PricingCurrencyEnumEnum = "PLN"
	PricingCurrencyEnumQar PricingCurrencyEnumEnum = "QAR"
	PricingCurrencyEnumRon PricingCurrencyEnumEnum = "RON"
	PricingCurrencyEnumSar PricingCurrencyEnumEnum = "SAR"
	PricingCurrencyEnumRsd PricingCurrencyEnumEnum = "RSD"
	PricingCurrencyEnumSgd PricingCurrencyEnumEnum = "SGD"
	PricingCurrencyEnumZar PricingCurrencyEnumEnum = "ZAR"
	PricingCurrencyEnumSek PricingCurrencyEnumEnum = "SEK"
	PricingCurrencyEnumTwd PricingCurrencyEnumEnum = "TWD"
	PricingCurrencyEnumThb PricingCurrencyEnumEnum = "THB"
	PricingCurrencyEnumTry PricingCurrencyEnumEnum = "TRY"
	PricingCurrencyEnumVnd PricingCurrencyEnumEnum = "VND"
)

var mappingPricingCurrencyEnumEnum = map[string]PricingCurrencyEnumEnum{
	"USD": PricingCurrencyEnumUsd,
	"CAD": PricingCurrencyEnumCad,
	"INR": PricingCurrencyEnumInr,
	"GBP": PricingCurrencyEnumGbp,
	"BRL": PricingCurrencyEnumBrl,
	"JPY": PricingCurrencyEnumJpy,
	"OMR": PricingCurrencyEnumOmr,
	"EUR": PricingCurrencyEnumEur,
	"CHF": PricingCurrencyEnumChf,
	"MXN": PricingCurrencyEnumMxn,
	"CLP": PricingCurrencyEnumClp,
	"ALL": PricingCurrencyEnumAll,
	"ARS": PricingCurrencyEnumArs,
	"AUD": PricingCurrencyEnumAud,
	"BDT": PricingCurrencyEnumBdt,
	"BAM": PricingCurrencyEnumBam,
	"BGN": PricingCurrencyEnumBgn,
	"CNY": PricingCurrencyEnumCny,
	"COP": PricingCurrencyEnumCop,
	"CRC": PricingCurrencyEnumCrc,
	"HRK": PricingCurrencyEnumHrk,
	"CZK": PricingCurrencyEnumCzk,
	"DKK": PricingCurrencyEnumDkk,
	"EGP": PricingCurrencyEnumEgp,
	"HKD": PricingCurrencyEnumHkd,
	"HUF": PricingCurrencyEnumHuf,
	"ISK": PricingCurrencyEnumIsk,
	"IDR": PricingCurrencyEnumIdr,
	"ILS": PricingCurrencyEnumIls,
	"JMD": PricingCurrencyEnumJmd,
	"KZT": PricingCurrencyEnumKzt,
	"KES": PricingCurrencyEnumKes,
	"KRW": PricingCurrencyEnumKrw,
	"KWD": PricingCurrencyEnumKwd,
	"LBP": PricingCurrencyEnumLbp,
	"MOP": PricingCurrencyEnumMop,
	"MYR": PricingCurrencyEnumMyr,
	"MVR": PricingCurrencyEnumMvr,
	"AED": PricingCurrencyEnumAed,
	"NZD": PricingCurrencyEnumNzd,
	"NOK": PricingCurrencyEnumNok,
	"PKR": PricingCurrencyEnumPkr,
	"PEN": PricingCurrencyEnumPen,
	"PHP": PricingCurrencyEnumPhp,
	"PLN": PricingCurrencyEnumPln,
	"QAR": PricingCurrencyEnumQar,
	"RON": PricingCurrencyEnumRon,
	"SAR": PricingCurrencyEnumSar,
	"RSD": PricingCurrencyEnumRsd,
	"SGD": PricingCurrencyEnumSgd,
	"ZAR": PricingCurrencyEnumZar,
	"SEK": PricingCurrencyEnumSek,
	"TWD": PricingCurrencyEnumTwd,
	"THB": PricingCurrencyEnumThb,
	"TRY": PricingCurrencyEnumTry,
	"VND": PricingCurrencyEnumVnd,
}

var mappingPricingCurrencyEnumEnumLowerCase = map[string]PricingCurrencyEnumEnum{
	"usd": PricingCurrencyEnumUsd,
	"cad": PricingCurrencyEnumCad,
	"inr": PricingCurrencyEnumInr,
	"gbp": PricingCurrencyEnumGbp,
	"brl": PricingCurrencyEnumBrl,
	"jpy": PricingCurrencyEnumJpy,
	"omr": PricingCurrencyEnumOmr,
	"eur": PricingCurrencyEnumEur,
	"chf": PricingCurrencyEnumChf,
	"mxn": PricingCurrencyEnumMxn,
	"clp": PricingCurrencyEnumClp,
	"all": PricingCurrencyEnumAll,
	"ars": PricingCurrencyEnumArs,
	"aud": PricingCurrencyEnumAud,
	"bdt": PricingCurrencyEnumBdt,
	"bam": PricingCurrencyEnumBam,
	"bgn": PricingCurrencyEnumBgn,
	"cny": PricingCurrencyEnumCny,
	"cop": PricingCurrencyEnumCop,
	"crc": PricingCurrencyEnumCrc,
	"hrk": PricingCurrencyEnumHrk,
	"czk": PricingCurrencyEnumCzk,
	"dkk": PricingCurrencyEnumDkk,
	"egp": PricingCurrencyEnumEgp,
	"hkd": PricingCurrencyEnumHkd,
	"huf": PricingCurrencyEnumHuf,
	"isk": PricingCurrencyEnumIsk,
	"idr": PricingCurrencyEnumIdr,
	"ils": PricingCurrencyEnumIls,
	"jmd": PricingCurrencyEnumJmd,
	"kzt": PricingCurrencyEnumKzt,
	"kes": PricingCurrencyEnumKes,
	"krw": PricingCurrencyEnumKrw,
	"kwd": PricingCurrencyEnumKwd,
	"lbp": PricingCurrencyEnumLbp,
	"mop": PricingCurrencyEnumMop,
	"myr": PricingCurrencyEnumMyr,
	"mvr": PricingCurrencyEnumMvr,
	"aed": PricingCurrencyEnumAed,
	"nzd": PricingCurrencyEnumNzd,
	"nok": PricingCurrencyEnumNok,
	"pkr": PricingCurrencyEnumPkr,
	"pen": PricingCurrencyEnumPen,
	"php": PricingCurrencyEnumPhp,
	"pln": PricingCurrencyEnumPln,
	"qar": PricingCurrencyEnumQar,
	"ron": PricingCurrencyEnumRon,
	"sar": PricingCurrencyEnumSar,
	"rsd": PricingCurrencyEnumRsd,
	"sgd": PricingCurrencyEnumSgd,
	"zar": PricingCurrencyEnumZar,
	"sek": PricingCurrencyEnumSek,
	"twd": PricingCurrencyEnumTwd,
	"thb": PricingCurrencyEnumThb,
	"try": PricingCurrencyEnumTry,
	"vnd": PricingCurrencyEnumVnd,
}

// GetPricingCurrencyEnumEnumValues Enumerates the set of values for PricingCurrencyEnumEnum
func GetPricingCurrencyEnumEnumValues() []PricingCurrencyEnumEnum {
	values := make([]PricingCurrencyEnumEnum, 0)
	for _, v := range mappingPricingCurrencyEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingCurrencyEnumEnumStringValues Enumerates the set of values in String for PricingCurrencyEnumEnum
func GetPricingCurrencyEnumEnumStringValues() []string {
	return []string{
		"USD",
		"CAD",
		"INR",
		"GBP",
		"BRL",
		"JPY",
		"OMR",
		"EUR",
		"CHF",
		"MXN",
		"CLP",
		"ALL",
		"ARS",
		"AUD",
		"BDT",
		"BAM",
		"BGN",
		"CNY",
		"COP",
		"CRC",
		"HRK",
		"CZK",
		"DKK",
		"EGP",
		"HKD",
		"HUF",
		"ISK",
		"IDR",
		"ILS",
		"JMD",
		"KZT",
		"KES",
		"KRW",
		"KWD",
		"LBP",
		"MOP",
		"MYR",
		"MVR",
		"AED",
		"NZD",
		"NOK",
		"PKR",
		"PEN",
		"PHP",
		"PLN",
		"QAR",
		"RON",
		"SAR",
		"RSD",
		"SGD",
		"ZAR",
		"SEK",
		"TWD",
		"THB",
		"TRY",
		"VND",
	}
}

// GetMappingPricingCurrencyEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingCurrencyEnumEnum(val string) (PricingCurrencyEnumEnum, bool) {
	enum, ok := mappingPricingCurrencyEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
