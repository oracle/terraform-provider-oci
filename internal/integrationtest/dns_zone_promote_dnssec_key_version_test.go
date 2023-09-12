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
	ZonePromoteDnssecKeyVersionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dns_zone",
		"test_dnssec_zone", acctest.Required, acctest.Create, zoneRepresentationDnssec) +
		DefinedTagsDependencies + `
			data "oci_identity_tenancy" "test_tenancy" {
				tenancy_id = "${var.tenancy_ocid}"
			}
		`
)

// issue-routing-tag: dns/default
func TestDnsZonePromoteDnssecKeyVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsZonePromoteDnssecKeyVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_zone.test_dnssec_zone"

	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_resource")

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// Create a dnssec enabled zone
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZonePromoteDnssecKeyVersionResourceDependencies, nil),
		},

		// Promote the staged KSK version
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZonePromoteDnssecKeyVersionResourceDependencies+`
				resource "oci_dns_zone_promote_dnssec_key_version" "test_zone_promote_dnssec_key_version" {
					zone_id                 = oci_dns_zone.test_dnssec_zone.id
					dnssec_key_version_uuid = oci_dns_zone.test_dnssec_zone.dnssec_config[0].ksk_dnssec_key_versions[0].uuid
					scope                   = "GLOBAL"
				}
				`, nil),
		},

		// Validate that the KSK key version's time_promoted was updated.
		// This requires a separate step because it requires a refresh of the/zone resource.
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZonePromoteDnssecKeyVersionResourceDependencies, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "dnssec_config.0.ksk_dnssec_key_versions.0.time_promoted"),
			),
		},
	})
}
