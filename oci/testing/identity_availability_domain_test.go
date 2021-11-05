// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

// issue-routing-tag: identity/default
func TestIdentityAvailabilityDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityAvailabilityDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_identity_availability_domains.test_availability_domains"
	singularDatasourceName := "data.oci_identity_availability_domain.test_availability_domain"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", Required, Create, availabilityDomainDataSourceRepresentation) +
				compartmentIdVariableStr + AvailabilityDomainResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.0.name"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_availability_domain", "test_availability_domain", Optional, Create, availabilityDomainSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AvailabilityDomainResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ad_number", "2"),
				resource.TestMatchResourceAttr(singularDatasourceName, "name", regexp.MustCompile(`\w+-AD-2`)),
				func(s *terraform.State) (err error) {
					adName, err := FromInstanceState(s, singularDatasourceName, "name")
					if err != nil {
						return err
					}

					regex := regexp.MustCompile(`(?i)AD-(\d)`)
					res := regex.FindAllStringSubmatch(adName, -1)

					// no matching AD name
					if res == nil || len(res) < 1 {
						err = fmt.Errorf("no match found for case insensitive search")
					}
					return err
				},
			),
		},
	})
}
