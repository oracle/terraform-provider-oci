// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	zoneRepresentationGlobal = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name": acctest.Representation{RepType: acctest.Required,
			Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.stage-dnssec-key-version-test`},
		"zone_type": acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"scope":     acctest.Representation{RepType: acctest.Required, Create: `GLOBAL`},
	}

	ZoneResourceDnssecDependencies = `
			data "oci_identity_tenancy" "test_tenancy" {
				tenancy_id = "${var.tenancy_ocid}"
			}
		`
)

func TestDnsZoneResourceDnssec(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneResourceDnssec")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_zone.test_zone"

	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_zone")

	acctest.ResourceTest(t, testAccCheckDnsZoneDestroy, []resource.TestStep{
		// create a zone with DNSSEC disabled
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDnssecDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required,
					acctest.Create, zoneRepresentationGlobal), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "dnssec_state", "DISABLED"),
			),
		},

		// verify enabling DNSSEC
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDnssecDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(zoneRepresentationGlobal,
						map[string]interface{}{
							"dnssec_state": acctest.Representation{RepType: acctest.Optional,
								Create: `ENABLED`},
						})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "dnssec_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "dnssec_config.0.zsk_dnssec_key_versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dnssec_config.0.ksk_dnssec_key_versions.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: tokenFn(config+compartmentIdVariableStr, nil),
		},

		// verify zone creation with DNSSEC enabled
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDnssecDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(zoneRepresentationGlobal, map[string]interface{}{
						"dnssec_state": acctest.Representation{RepType: acctest.Required,
							Create: `ENABLED`},
					})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "dnssec_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "dnssec_config.0.zsk_dnssec_key_versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dnssec_config.0.ksk_dnssec_key_versions.#", "1"),
			),
		},

		// verify disabling DNSSEC
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDnssecDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional,
					acctest.Update, acctest.RepresentationCopyWithNewProperties(zoneRepresentationGlobal,
						map[string]interface{}{
							"dnssec_state": acctest.Representation{RepType: acctest.Optional,
								Create: `DISABLED`},
						})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "dnssec_state", "DISABLED"),
			),
		},
	})
}
