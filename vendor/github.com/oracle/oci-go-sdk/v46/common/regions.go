// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

const (
	//RegionSEA region SEA
	RegionSEA Region = "sea"
	//RegionCAToronto1 region for Toronto
	RegionCAToronto1 Region = "ca-toronto-1"
	//RegionCAMontreal1 region for Montreal
	RegionCAMontreal1 Region = "ca-montreal-1"
	//RegionPHX region PHX
	RegionPHX Region = "us-phoenix-1"
	//RegionIAD region IAD
	RegionIAD Region = "us-ashburn-1"
	//RegionSJC1 region SJC
	RegionSJC1 Region = "us-sanjose-1"

	//RegionFRA region FRA
	RegionFRA Region = "eu-frankfurt-1"

	//RegionUKCardiff1 region for Cardiff
	RegionUKCardiff1 Region = "uk-cardiff-1"
	//RegionLHR region LHR
	RegionLHR Region = "uk-london-1"

	//RegionAPTokyo1 region for Tokyo
	RegionAPTokyo1 Region = "ap-tokyo-1"
	//RegionAPOsaka1 region for Osaka
	RegionAPOsaka1 Region = "ap-osaka-1"
	//RegionAPChiyoda1 region for Chiyoda
	RegionAPChiyoda1 Region = "ap-chiyoda-1"
	//RegionAPSeoul1 region for Seoul
	RegionAPSeoul1 Region = "ap-seoul-1"
	//RegionAPChuncheon1 region for Chuncheon
	RegionAPChuncheon1 Region = "ap-chuncheon-1"
	//RegionAPMumbai1 region for Mumbai
	RegionAPMumbai1 Region = "ap-mumbai-1"
	//RegionAPHyderabad1 region for Hyderabad
	RegionAPHyderabad1 Region = "ap-hyderabad-1"
	//RegionAPMelbourne1 region for Melbourne
	RegionAPMelbourne1 Region = "ap-melbourne-1"
	//RegionAPSydney1 region for Sydney
	RegionAPSydney1 Region = "ap-sydney-1"

	//RegionMEJeddah1 region for Jeddah
	RegionMEJeddah1 Region = "me-jeddah-1"
	//RegionMEDubai1 region for Dubai
	RegionMEDubai1 Region = "me-dubai-1"

	//RegionEUZurich1 region for Zurich
	RegionEUZurich1 Region = "eu-zurich-1"
	//RegionEUAmsterdam1 region for Amsterdam
	RegionEUAmsterdam1 Region = "eu-amsterdam-1"

	//RegionSASaopaulo1 region for Sao Paulo
	RegionSASaopaulo1 Region = "sa-saopaulo-1"
	//RegionSASantiago1 region for santiago
	RegionSASantiago1 Region = "sa-santiago-1"
	//RegionSAVinhedo1 region for vinhedo
	RegionSAVinhedo1 Region = "sa-vinhedo-1"

	//RegionUSLangley1 region for Langley
	RegionUSLangley1 Region = "us-langley-1"
	//RegionUSLuke1 region for Luke
	RegionUSLuke1 Region = "us-luke-1"

	//RegionUSGovAshburn1 gov region Ashburn
	RegionUSGovAshburn1 Region = "us-gov-ashburn-1"
	//RegionUSGovChicago1 gov region Chicago
	RegionUSGovChicago1 Region = "us-gov-chicago-1"
	//RegionUSGovPhoenix1 region for Phoenix
	RegionUSGovPhoenix1 Region = "us-gov-phoenix-1"
	//RegionUKGovLondon1 gov region London
	RegionUKGovLondon1 Region = "uk-gov-london-1"
	//RegionUKGovCardiff1 gov region Cardiff
	RegionUKGovCardiff1 Region = "uk-gov-cardiff-1"
)

var shortNameRegion = map[string]Region{
	"sea": RegionSEA,
	"phx": RegionPHX,
	"iad": RegionIAD,
	"fra": RegionFRA,
	"lhr": RegionLHR,
	"cwl": RegionUKCardiff1,
	"ams": RegionEUAmsterdam1,
	"zrh": RegionEUZurich1,
	"mel": RegionAPMelbourne1,
	"bom": RegionAPMumbai1,
	"hyd": RegionAPHyderabad1,
	"icn": RegionAPSeoul1,
	"yny": RegionAPChuncheon1,
	"nrt": RegionAPTokyo1,
	"kix": RegionAPOsaka1,
	"nja": RegionAPChiyoda1,
	"jed": RegionMEJeddah1,
	"dxb": RegionMEDubai1,
	"syd": RegionAPSydney1,
	"yul": RegionCAMontreal1,
	"yyz": RegionCAToronto1,
	"sjc": RegionSJC1,
	"gru": RegionSASaopaulo1,
	"scl": RegionSASantiago1,
	"vcp": RegionSAVinhedo1,
	"ltn": RegionUKGovLondon1,
	"brs": RegionUKGovCardiff1,
}

var realm = map[string]string{
	"oc1": "oraclecloud.com",
	"oc2": "oraclegovcloud.com",
	"oc3": "oraclegovcloud.com",
	"oc4": "oraclegovcloud.uk",
	"oc8": "oraclecloud8.com",
}

var regionRealm = map[Region]string{
	RegionPHX:  "oc1",
	RegionIAD:  "oc1",
	RegionFRA:  "oc1",
	RegionLHR:  "oc1",
	RegionSJC1: "oc1",

	RegionUKCardiff1: "oc1",

	RegionCAToronto1:  "oc1",
	RegionCAMontreal1: "oc1",

	RegionAPTokyo1:     "oc1",
	RegionAPOsaka1:     "oc1",
	RegionAPSeoul1:     "oc1",
	RegionAPChuncheon1: "oc1",
	RegionAPSydney1:    "oc1",
	RegionAPMumbai1:    "oc1",
	RegionAPHyderabad1: "oc1",
	RegionAPMelbourne1: "oc1",

	RegionMEJeddah1: "oc1",
	RegionMEDubai1:  "oc1",

	RegionEUZurich1:    "oc1",
	RegionEUAmsterdam1: "oc1",

	RegionSASaopaulo1: "oc1",
	RegionSASantiago1: "oc1",
	RegionSAVinhedo1:  "oc1",

	RegionUSLangley1: "oc2",
	RegionUSLuke1:    "oc2",

	RegionUSGovAshburn1: "oc3",
	RegionUSGovChicago1: "oc3",
	RegionUSGovPhoenix1: "oc3",
	RegionUKGovCardiff1: "oc4",
	RegionUKGovLondon1:  "oc4",

	RegionAPChiyoda1: "oc8",
}
