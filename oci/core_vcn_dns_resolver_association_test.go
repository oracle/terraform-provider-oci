// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vcnDnsResolverAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"vcn_id": Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
	}

	VcnDnsResolverAssociationResourceConfig = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)
)

func TestCoreVcnDnsResolverAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVcnDnsResolverAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create dependencies
			{
				Config: config + compartmentIdVariableStr + VcnDnsResolverAssociationResourceConfig,
				Check: func(s *terraform.State) (err error) {
					log.Printf("Wait for 2 minutes for oci_dns_resolver resource to get created first")
					time.Sleep(2 * time.Minute)
					return nil
				},
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VcnDnsResolverAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "dns_resolver_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				),
			},
		},
	})
}
