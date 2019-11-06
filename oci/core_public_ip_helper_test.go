package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func getPublicIpIdsForRegionScope(compartmentId string, listPublicIpsRequest oci_core.ListPublicIpsRequest) ([]oci_core.PublicIp, error) {
	return getPublicIpIdsByScope(compartmentId, listPublicIpsRequest)
}

func getPublicIpIdsForADScope(compartmentId string, listPublicIpsRequest oci_core.ListPublicIpsRequest) ([]oci_core.PublicIp, error) {
	var publicIps []oci_core.PublicIp
	availabilityDomains, err := getAvalabilityDomains(compartmentId)
	if err != nil {
		return nil, fmt.Errorf("Error getting availabilityDomains required for MountTarget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listPublicIpsRequest.AvailabilityDomain = &availabilityDomainName
		ips, err := getPublicIpIdsByScope(compartmentId, listPublicIpsRequest)
		if err != nil {
			return nil, err
		}
		publicIps = append(publicIps, ips...)
	}
	return publicIps, nil
}

func getPublicIpIdsByScope(compartmentId string, listPublicIpsRequest oci_core.ListPublicIpsRequest) ([]oci_core.PublicIp, error) {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	listPublicIpsResponse, err := virtualNetworkClient.ListPublicIps(context.Background(), listPublicIpsRequest)

	if err != nil {
		return nil, fmt.Errorf("Error getting PublicIp list for compartment id : %s , %s \n", compartmentId, err)
	}
	return listPublicIpsResponse.Items, nil
}
