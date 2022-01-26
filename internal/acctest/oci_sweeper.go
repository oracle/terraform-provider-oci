// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	utils "github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

/* This map holds the list of ocids for a given resourceType by compartment
For Example :
		Submap1[vcn] : [vcn-1_ocid, vcn-2_ocid, ...]				// In Compartment 1
		Submap1[instance] : [instance-1_ocid, instance-2_ocid, ...] // In Compartment 1
		SweeperResourceCompartmentIdMap[compartment-1_ocid] = Submap1

		Submap2[vcn] : [vcn-1_ocid, vcn-2_ocid, ...]				// In Compartment 2
		Submap2[instance] : [instance-1_ocid, instance-2_ocid, ...] // In Compartment 2
		SweeperResourceCompartmentIdMap[compartment-2_ocid] = Submap2
*/
var SweeperResourceCompartmentIdMap map[string]map[string][]string

/*
This Map hold the ocids of the default resources.
For example: vcn can have default dhcpOptions, routeTables and securityLists which should not be deleted individually.
			 These default resources are deleted as part of deleting the vcn itself.
			 In such cases we identify and add the default resource ocid into this map and the respective sweeper
			 checks if the ocid of the resource be deleted is present in this map then it will skip that resource.
*/
var SweeperDefaultResourceId = make(map[string]bool)

// issue-routing-tag: terraform/default
func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func AddResourceIdToSweeperResourceIdMap(compartmentId string, resourceType string, resourceId string) {
	if _, ok := SweeperResourceCompartmentIdMap[compartmentId]; ok {
		resourceCompartmentIdMap := SweeperResourceCompartmentIdMap[compartmentId]
		if _, ok := resourceCompartmentIdMap[resourceType]; ok {
			resourceCompartmentIdMap[resourceType] = append(resourceCompartmentIdMap[resourceType], resourceId)
		} else {
			idList := []string{resourceId}
			resourceCompartmentIdMap := SweeperResourceCompartmentIdMap[compartmentId]
			resourceCompartmentIdMap[resourceType] = idList
		}
	} else {
		resourceCompartmentIdMap := map[string]map[string][]string{}
		resourceIdMap := make(map[string][]string)
		resourceIdList := []string{resourceId}
		resourceIdMap[resourceType] = resourceIdList
		resourceCompartmentIdMap[compartmentId] = resourceIdMap
		SweeperResourceCompartmentIdMap = resourceCompartmentIdMap
	}
}

func GetResourceIdsToSweep(compartmentId string, resourceName string) []string {
	if _, ok := SweeperResourceCompartmentIdMap[compartmentId]; ok {
		resourceIdMap := SweeperResourceCompartmentIdMap[compartmentId]
		if _, ok := resourceIdMap[resourceName]; ok {
			return resourceIdMap[resourceName]
		}
	}
	return nil
}

func GetAvalabilityDomains(compartmentId string) (map[string]string, error) {
	availabilityDomains := make(map[string]string)
	identityClient := GetTestClients(&schema.ResourceData{}).IdentityClient()
	adRequest := oci_identity.ListAvailabilityDomainsRequest{}
	adRequest.CompartmentId = &compartmentId
	ads, err := identityClient.ListAvailabilityDomains(context.Background(), adRequest)
	if err != nil {
		return availabilityDomains, fmt.Errorf("Error getting availability domains for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ad := range ads.Items {
		availabilityDomains[*ad.Id] = *ad.Name
	}
	return availabilityDomains, nil
}

func InSweeperExcludeList(sweeperName string) bool {
	excludeListSweeper := strings.Split(utils.GetEnvSettingWithBlankDefault("sweep_exclude_list"), ",")

	for _, sweeper := range excludeListSweeper {
		if strings.EqualFold(strings.Trim(sweeper, " "), sweeperName) {
			log.Printf("[DEBUG] Skip sweeper for %s", sweeperName)
			return true
		}
	}

	return false
}
