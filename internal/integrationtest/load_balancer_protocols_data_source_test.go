// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// issue-routing-tag: load_balancer/default
func TestAccDatasourceLoadBalancerProtocols_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerProtocols_basic")
	defer httpreplay.SaveScenario()
	providers := acctest.TestAccProviders
	config := acctest.LegacyTestProviderConfig() + `
	data "oci_load_balancer_protocols" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_protocols.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { acctest.PreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "protocols.#"),
					resource.TestCheckResourceAttrSet(resourceName, "protocols.0.name"),
				),
			},
		},
	})
}
