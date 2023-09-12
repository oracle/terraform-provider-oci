// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	zoneRepresentationDnssec = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name": acctest.Representation{RepType: acctest.Required,
			Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.stage-dnssec-key-version-test`},
		"zone_type":    acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"scope":        acctest.Representation{RepType: acctest.Required, Create: `GLOBAL`},
		"dnssec_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
	}

	ZoneStageDnssecKeyVersionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dns_zone",
		"test_dnssec_zone", acctest.Required, acctest.Create, zoneRepresentationDnssec) +
		DefinedTagsDependencies + `
			data "oci_identity_tenancy" "test_tenancy" {
				tenancy_id = "${var.tenancy_ocid}"
			}
		`
)

// issue-routing-tag: dns/default
func TestDnsZoneStageDnssecKeyVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneStageDnssecKeyVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_zone.test_dnssec_zone"

	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_resource")

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// Create a dnssec enabled zone
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneStageDnssecKeyVersionResourceDependencies, nil),
		},

		// Stage a replacement ZSK version
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneStageDnssecKeyVersionResourceDependencies+`
                                locals {
                                  predecessor_uuid = length(oci_dns_zone.test_dnssec_zone.dnssec_config[0].zsk_dnssec_key_versions) == 1 ? oci_dns_zone.test_dnssec_zone.dnssec_config[0].zsk_dnssec_key_versions[0].uuid : [ for zsk in oci_dns_zone.test_dnssec_zone.dnssec_config[0].zsk_dnssec_key_versions : zsk if zsk.successor_dnssec_key_version_uuid != ""][0].uuid
                                }
				resource "oci_dns_zone_stage_dnssec_key_version" "test_zone_stage_dnssec_key_version" {
					zone_id                             = oci_dns_zone.test_dnssec_zone.id
					predecessor_dnssec_key_version_uuid = local.predecessor_uuid
					scope                               = "GLOBAL"
				}
				`, nil),
		},

		// Validate that a second ZSK key version was added to the dnssec configuration.
		// This requires a separate step because it requires a refresh of the zone resource.
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneStageDnssecKeyVersionResourceDependencies, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "dnssec_config.0.zsk_dnssec_key_versions.#", "2"),
			),
		},
	})
}
