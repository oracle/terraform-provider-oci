// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vcnDnsResolverAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"vcn_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	VcnDnsResolverAssociationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)
)

// issue-routing-tag: core/default
func TestCoreVcnDnsResolverAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVcnDnsResolverAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", acctest.Required, acctest.Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VcnDnsResolverAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "dns_resolver_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
