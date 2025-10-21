// Copyright (c) 2025 Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// When the DNS zone is ForceNew (name change), an rrset referencing the zone ID is also re-created,
// and its records are re-applied to the newly created zone in the same apply.
func TestDnsZoneForceNew_RecreatesRrset(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneForceNew_RecreatesRrset")
	defer httpreplay.SaveScenario()

	// Provider HCL config
	config := acctest.ProviderTestConfig()

	// Compartment var from environment for acceptance tests
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Tokenizer so we can embed a stable token in resource names
	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_zone_rrset_force_new")

	// Minimal dependency used by many DNS tests to help construct unique names
	deps := `
data "oci_identity_tenancy" "test_tenancy" {
  tenancy_id = "${var.tenancy_ocid}"
}
`

	// Step 1: Create zone (name suffix "-a") + rrset (2 A records) that references zone ID
	zoneRepStep1 := map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-force-new-a`},
		"zone_type":      acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
	}

	// Step 2: Same rrset, but zone name changed (suffix "-b") to trigger ForceNew on zone
	zoneRepStep2 := map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-force-new-b`},
		"zone_type":      acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
	}

	// rrset references the zone by ID (ensures rrset is ForceNew when zone is recreated)
	rrsetItems1 := map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `192.168.0.1`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `A`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`},
	}
	rrsetItems2 := map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `192.168.0.2`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `A`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`},
	}

	rrsetRep := map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.name}`},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           []acctest.RepresentationGroup{{RepType: acctest.Required, Group: rrsetItems1}, {RepType: acctest.Required, Group: rrsetItems2}},
	}

	resourceNameZone := "oci_dns_zone.test_zone"
	resourceNameRrset := "oci_dns_rrset.test_rrset"

	var rrsetId1, rrsetId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 1: Create zone (name "-a") and rrset
		{
			Config: tokenFn(config+compartmentIdVariableStr+deps+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepStep1)+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRep), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameZone, "id"),
				resource.TestCheckResourceAttrSet(resourceNameRrset, "id"),
				resource.TestCheckResourceAttr(resourceNameRrset, "items.#", "2"),
				func(s *terraform.State) (err error) {
					rrsetId1, err = acctest.FromInstanceState(s, resourceNameRrset, "id")
					if err == nil {
						log.Printf("[INFO] rrsetId1=%s", rrsetId1)
					}
					return err
				},
			),
		},
		// Step 2: Change zone name to trigger ForceNew on zone; rrset references zone.id so it should be replaced too
		{
			Config: tokenFn(config+compartmentIdVariableStr+deps+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepStep2)+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRep), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameZone, "id"),
				resource.TestCheckResourceAttrSet(resourceNameRrset, "id"),
				// Verify records are present after replacement
				resource.TestCheckResourceAttr(resourceNameRrset, "items.#", "2"),
				func(s *terraform.State) (err error) {
					rrsetId2, err = acctest.FromInstanceState(s, resourceNameRrset, "id")
					if err != nil {
						return err
					}
					log.Printf("[INFO] rrsetId1=%s rrsetId2=%s", rrsetId1, rrsetId2)
					if rrsetId1 == rrsetId2 {
						return fmt.Errorf("RRSet ID should change when zone is ForceNew-replaced (expected rrset replacement)")
					}
					return nil
				},
			),
		},
	})
}
