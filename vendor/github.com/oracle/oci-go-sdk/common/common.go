// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package common

import (
	"fmt"
	"strings"
)

//Region type for regions
type Region string

const (
	//RegionSEA region SEA
	RegionSEA Region = "sea"
	//RegionPHX region PHX
	RegionPHX Region = "us-phoenix-1"
	//RegionIAD region IAD
	RegionIAD Region = "us-ashburn-1"
	//RegionFRA region FRA
	RegionFRA Region = "eu-frankfurt-1"
	//RegionLHR region LHR
	RegionLHR Region = "uk-london-1"
)

//StringToRegion convert a string to Region type
func StringToRegion(stringRegion string) (r Region) {
	switch strings.ToLower(stringRegion) {
	case "sea":
		r = RegionSEA
	case "phx", "us-phoenix-1":
		r = RegionPHX
	case "iad", "us-ashburn-1":
		r = RegionIAD
	case "fra", "eu-frankfurt-1":
		r = RegionFRA
	case "lhr", "uk-london-1":
		r = RegionLHR
	default:
		r = Region(stringRegion)
		Debugf("region named: %s, is not recognized", stringRegion)
	}
	return
}

// canStringBeRegion test if the string can be a region, if it can, returns the string as is, otherwise it
// returns an error
func canStringBeRegion(stringRegion string) (region string, err error) {
	if strings.Contains(stringRegion, " ") || stringRegion == "" {
		return "", fmt.Errorf("region can not be empty or have spaces")
	}
	return stringRegion, nil
}
