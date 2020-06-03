// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

//Region type for regions
type Region string

const (
	instanceMetadataRegionInfoURLV2 = "http://169.254.169.254/opc/v2/instance/regionInfo"

	//RegionSEA region SEA
	RegionSEA Region = "sea"
	//RegionCAToronto1 region for toronto
	RegionCAToronto1 Region = "ca-toronto-1"
	//RegionCAMontreal1 region for Montreal
	RegionCAMontreal1 Region = "ca-montreal-1"
	//RegionPHX region PHX
	RegionPHX Region = "us-phoenix-1"
	//RegionIAD region IAD
	RegionIAD Region = "us-ashburn-1"
	//RegionFRA region FRA
	RegionFRA Region = "eu-frankfurt-1"
	//RegionLHR region LHR
	RegionLHR Region = "uk-london-1"
	//RegionAPTokyo1 region for tokyo
	RegionAPTokyo1 Region = "ap-tokyo-1"
	//RegionAPOsaka1 region for Osaka
	RegionAPOsaka1 Region = "ap-osaka-1"
	//RegionAPSeoul1 region for seoul
	RegionAPSeoul1 Region = "ap-seoul-1"
	//RegionAPChuncheon1 region for Chuncheon
	RegionAPChuncheon1 Region = "ap-chuncheon-1"
	//RegionAPMumbai1 region for mumbai
	RegionAPMumbai1 Region = "ap-mumbai-1"
	//RegionAPHyderabad1 region for Hyderabad
	RegionAPHyderabad1 Region = "ap-hyderabad-1"
	//RegionAPMelbourne1 region for Melbourne
	RegionAPMelbourne1 Region = "ap-melbourne-1"
	//RegionAPSydney1 region for Sydney
	RegionAPSydney1 Region = "ap-sydney-1"
	//RegionMEJeddah1 region for Jeddah
	RegionMEJeddah1 Region = "me-jeddah-1"
	//RegionEUZurich1 region for Zurich
	RegionEUZurich1 Region = "eu-zurich-1"
	//RegionEUAmsterdam1 region for Amsterdam
	RegionEUAmsterdam1 Region = "eu-amsterdam-1"
	//RegionSASaopaulo1 region for Sao Paulo
	RegionSASaopaulo1 Region = "sa-saopaulo-1"
	//RegionUSLangley1 region for langley
	RegionUSLangley1 Region = "us-langley-1"
	//RegionUSLuke1 region for luke
	RegionUSLuke1 Region = "us-luke-1"

	//RegionUSGovAshburn1 region for langley
	RegionUSGovAshburn1 Region = "us-gov-ashburn-1"
	//RegionUSGovChicago1 region for luke
	RegionUSGovChicago1 Region = "us-gov-chicago-1"
	//RegionUSGovPhoenix1 region for luke
	RegionUSGovPhoenix1 Region = "us-gov-phoenix-1"
	//RegionUKGovLondon1 gov region London
	RegionUKGovLondon1 Region = "uk-gov-london-1"

	//RegionUKGovCardiff1 gov region Cardiff
	RegionUKGovCardiff1 Region = "uk-gov-cardiff-1"

	//RegionUSTacoma1 region for us-tacoma-1
	RegionUSTacoma1 Region = "us-tacoma-1"

	// Region Metadata Configuration File
	regionMetadataCfgDirName  = ".oci"
	regionMetadataCfgFileName = "regions-config.json"

	// Region Metadata Environment Variable
	regionMetadataEnvVarName = "OCI_REGION_METADATA"

	// Region Metadata
	regionIdentifierPropertyName     = "regionIdentifier"     // e.g. "ap-sydney-1"
	realmKeyPropertyName             = "realmKey"             // e.g. "oc1"
	realmDomainComponentPropertyName = "realmDomainComponent" // e.g. "oraclecloud.com"
	regionKeyPropertyName            = "regionKey"            // e.g. "SYD"
)

var realm = map[string]string{
	"oc1": "oraclecloud.com",
	"oc2": "oraclegovcloud.com",
	"oc3": "oraclegovcloud.com",
	"oc4": "oraclegovcloud.uk",
}

var regionRealm = map[Region]string{
	RegionPHX:          "oc1",
	RegionIAD:          "oc1",
	RegionFRA:          "oc1",
	RegionLHR:          "oc1",
	RegionCAToronto1:   "oc1",
	RegionCAMontreal1:  "oc1",
	RegionAPTokyo1:     "oc1",
	RegionAPOsaka1:     "oc1",
	RegionAPSeoul1:     "oc1",
	RegionAPChuncheon1: "oc1",
	RegionAPSydney1:    "oc1",
	RegionAPMumbai1:    "oc1",
	RegionAPHyderabad1: "oc1",
	RegionAPMelbourne1: "oc1",
	RegionMEJeddah1:    "oc1",
	RegionEUZurich1:    "oc1",
	RegionEUAmsterdam1: "oc1",
	RegionSASaopaulo1:  "oc1",

	RegionUSLangley1:    "oc2",
	RegionUSLuke1:       "oc2",
	RegionUSGovAshburn1: "oc3",
	RegionUSGovChicago1: "oc3",
	RegionUSGovPhoenix1: "oc3",
	RegionUKGovLondon1:  "oc4",
}

// GetRegionInfoFromInstanceMetadataService gets the region information
var GetRegionInfoFromInstanceMetadataService = GetRegionInfoFromInstanceMetadataServiceProd

// Endpoint returns a endpoint for a service
func (region Region) Endpoint(service string) string {
	return fmt.Sprintf("%s.%s.%s", service, region, region.secondLevelDomain())
}

// EndpointForTemplate returns a endpoint for a service based on template, only unknown region name can fall back to "oc1", but not short code region name.
func (region Region) EndpointForTemplate(service string, serviceEndpointTemplate string) string {
	if serviceEndpointTemplate == "" {
		return region.Endpoint(service)
	}

	// replace service prefix
	endpoint := strings.Replace(serviceEndpointTemplate, "{serviceEndpointPrefix}", service, 1)

	// replace region
	endpoint = strings.Replace(endpoint, "{region}", string(region), 1)

	// replace second level domain
	endpoint = strings.Replace(endpoint, "{secondLevelDomain}", region.secondLevelDomain(), 1)

	return endpoint
}

func (region Region) secondLevelDomain() string {
	if realmID, ok := regionRealm[region]; ok {
		if secondLevelDomain, ok := realm[realmID]; ok {
			return secondLevelDomain
		}
	}

	Debugf("cannot find realm for region : %s, return default realm value.", region)
	return realm["oc1"]
}

//StringToRegion convert a string to Region type
func StringToRegion(stringRegion string) (r Region) {
	switch strings.ToLower(stringRegion) {
	case "sea":
		r = RegionSEA
	case "yyz", "ca-toronto-1":
		r = RegionCAToronto1
	case "yul", "ca-montreal-1":
		r = RegionCAMontreal1
	case "phx", "us-phoenix-1":
		r = RegionPHX
	case "iad", "us-ashburn-1":
		r = RegionIAD
	case "fra", "eu-frankfurt-1":
		r = RegionFRA
	case "lhr", "uk-london-1":
		r = RegionLHR
	case "nrt", "ap-tokyo-1":
		r = RegionAPTokyo1
	case "kix", "ap-osaka-1":
		r = RegionAPOsaka1
	case "icn", "ap-seoul-1":
		r = RegionAPSeoul1
	case "yny", "ap-chuncheon-1":
		r = RegionAPChuncheon1
	case "bom", "ap-mumbai-1":
		r = RegionAPMumbai1
	case "hyd", "ap-hyderabad-1":
		r = RegionAPHyderabad1
	case "mel", "ap-melbourne-1":
		r = RegionAPMelbourne1
	case "syd", "ap-sydney-1":
		r = RegionAPSydney1
	case "jed", "me-jeddah-1":
		r = RegionMEJeddah1
	case "zrh", "eu-zurich-1":
		r = RegionEUZurich1
	case "ams", "eu-amsterdam-1":
		r = RegionEUAmsterdam1
	case "gru", "sa-saopaulo-1":
		r = RegionSASaopaulo1
	case "us-langley-1":
		r = RegionUSLangley1
	case "us-luke-1":
		r = RegionUSLuke1
	case "us-gov-ashburn-1":
		r = RegionUSGovAshburn1
	case "us-gov-chicago-1":
		r = RegionUSGovChicago1
	case "us-gov-phoenix-1":
		r = RegionUSGovPhoenix1
	case "ltn", "uk-gov-london-1":
		r = RegionUKGovLondon1
	default:
		Debugf("region named: %s, is not recognized", stringRegion)
		r = checkAndAddRegionMetadata(stringRegion)
	}
	return
}

// canStringBeRegion test if the string can be a region, if it can, returns the string as is, otherwise it
// returns an error
var blankRegex = regexp.MustCompile("\\s")

func canStringBeRegion(stringRegion string) (region string, err error) {
	if blankRegex.MatchString(stringRegion) || stringRegion == "" {
		return "", fmt.Errorf("region can not be empty or have spaces")
	}
	return stringRegion, nil
}

// check region info from original map
func checkAndAddRegionMetadata(region string) Region {
	switch {
	case SetRegionMetadataFromCfgFile(&region):
	case SetRegionMetadataFromEnvVar(&region):
	case SetRegionFromInstanceMetadataService(&region):
	default:
		//err := fmt.Errorf("failed to get region metadata information.")
		return Region(region)
	}
	return Region(region)
}

// SetRegionMetadataFromEnvVar checks if region metadata env variable is provided, once it's there, parse and added it to region map
func SetRegionMetadataFromEnvVar(region *string) bool {
	// check from env variable
	if jsonStr, existed := os.LookupEnv(regionMetadataEnvVarName); existed {
		Debugf("Raw content of region metadata env var:", jsonStr)
		var regionSchema map[string]string
		if err := json.Unmarshal([]byte(jsonStr), &regionSchema); err != nil {
			Debugf("Can't unmarshal env var, the error info is", err)
			return false
		}
		// check if the specified region is in the env var.
		if checkSchemaItems(regionSchema) {
			// set mapping table
			addRegionSchema(regionSchema)
			if regionSchema[regionKeyPropertyName] == *region ||
				regionSchema[regionIdentifierPropertyName] == *region {
				*region = regionSchema[regionIdentifierPropertyName]
				return true
			}
		}
		return false
	}
	Debugf("The Region Metadata Schema wasn't set in env variable - OCI_REGION_METADATA.")
	return false
}

// SetRegionMetadataFromCfgFile checks if region metadata config file is provided, once it's there, parse and added it to region map
func SetRegionMetadataFromCfgFile(region *string) bool {
	homeFolder := getHomeFolder()
	configFile := path.Join(homeFolder, regionMetadataCfgDirName, regionMetadataCfgFileName)
	if jsonArr, ok := readAndParseConfigFile(&configFile); ok {
		added := false
		for _, jsonItem := range jsonArr {
			if checkSchemaItems(jsonItem) {
				addRegionSchema(jsonItem)
				if jsonItem[regionKeyPropertyName] == *region ||
					jsonItem[regionIdentifierPropertyName] == *region {
					*region = jsonItem[regionIdentifierPropertyName]
					added = true
				}
			}
		}
		return added
	}
	return false
}

func readAndParseConfigFile(configFileName *string) (fileContent []map[string]string, ok bool) {

	if content, err := ioutil.ReadFile(*configFileName); err == nil {
		Debugf("Raw content of region metadata config file content:", string(content[:]))
		if err := json.Unmarshal(content, &fileContent); err != nil {
			Debugf("Can't unmarshal env var, the error info is", err)
			return
		}
		ok = true
		return
	}
	Debugf("No Region Metadata Config File provided.")
	return

}

// check map regionRealm's region name, if it's already there, no need to add it.
func addRegionSchema(regionSchema map[string]string) {
	r := Region(strings.ToLower(regionSchema[regionIdentifierPropertyName]))
	if _, ok := regionRealm[r]; !ok {
		// set mapping table
		realm[regionSchema[realmKeyPropertyName]] = regionSchema[realmDomainComponentPropertyName]
		regionRealm[r] = regionSchema[realmKeyPropertyName]
		return
	}
	Debugf("Region {} has already been added, no need to add again.", regionSchema[regionIdentifierPropertyName])
}

// check region schema content if all the required contents are provided
func checkSchemaItems(regionSchema map[string]string) bool {
	if checkSchemaItem(regionSchema, regionIdentifierPropertyName) &&
		checkSchemaItem(regionSchema, realmKeyPropertyName) &&
		checkSchemaItem(regionSchema, realmDomainComponentPropertyName) &&
		checkSchemaItem(regionSchema, regionKeyPropertyName) {
		return true
	}
	return false
}

// check region schema item is valid, if so, convert it to lower case.
func checkSchemaItem(regionSchema map[string]string, key string) bool {
	if val, ok := regionSchema[key]; ok {
		if val != "" {
			regionSchema[key] = strings.ToLower(val)
			return true
		}
		Debugf("Region metadata schema {} is provided,but content is empty.", key)
		return false
	}
	Debugf("Region metadata schema {} is not provided, please update the content", key)
	return false
}

// SetRegionFromInstanceMetadataService checks if region metadata can be provided from InstanceMetadataService.
func SetRegionFromInstanceMetadataService(region *string) bool {
	// example of content:
	// {
	// 	"realmKey" : "oc1",
	// 	"realmDomainComponent" : "oraclecloud.com",
	// 	"regionKey" : "YUL",
	// 	"regionIdentifier" : "ca-montreal-1"
	// }
	content, err := GetRegionInfoFromInstanceMetadataService()
	if err != nil {
		Debugf("Failed to get instance metadata. Error: %v", err)
		return false
	}

	var regionInfo map[string]string
	err = json.Unmarshal(content, &regionInfo)
	if err != nil {
		Debugf("Failed to unmarshal the response content: %v \nError: %v", string(content), err)
		return false
	}

	if checkSchemaItems(regionInfo) {
		addRegionSchema(regionInfo)
		if regionInfo[regionKeyPropertyName] == *region ||
			regionInfo[regionIdentifierPropertyName] == *region {
			*region = regionInfo[regionIdentifierPropertyName]
		}
	} else {
		Debugf("Region information is not valid.")
		return false
	}

	return true
}

// GetRegionInfoFromInstanceMetadataServiceProd calls instance metadata service and get the region information
func GetRegionInfoFromInstanceMetadataServiceProd() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, instanceMetadataRegionInfoURLV2, nil)
	request.Header.Add("Authorization", "Bearer Oracle")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to call instance metadata service. Error: %v", err)
	}

	statusCode := resp.StatusCode

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to get region information from response body. Error: %v", err)
	}

	if statusCode != http.StatusOK {
		err = fmt.Errorf("HTTP Get failed: URL: %s, Status: %s, Message: %s",
			instanceMetadataRegionInfoURLV2, resp.Status, string(content))
		return nil, err
	}

	return content, nil
}
