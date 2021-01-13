// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	zoneDataSourceRepresentationRequiredOnlyDefault = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}
	zoneDataSourceRepresentationRequiredOnlyWithFilterDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"filter": RepresentationGroup{Required, zoneDataSourceFilterRepresentation},
	})
	zoneDataSourceRepresentationWithNameOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name": Representation{repType: Optional, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
	})
	zoneDataSourceRepresentationWithNameContainsOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name_contains": Representation{repType: Optional, create: `oci-zone-test`},
	})
	zoneDataSourceRepresentationWithStateOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"state": Representation{repType: Optional, create: `ACTIVE`},
	})
	zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
	})
	zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_less_than": Representation{repType: Optional, create: `2022-04-10T19:01:09.000-00:00`},
	})
	zoneDataSourceRepresentationWithZoneTypeOptionalDefault = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"zone_type": Representation{repType: Optional, create: `PRIMARY`},
	})

	zoneRepresentationPrimaryDefault = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"name":             Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
		"zone_type":        Representation{repType: Required, create: `PRIMARY`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_masters": RepresentationGroup{Optional, zoneExternalMastersRepresentation},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	zoneRepresentationDefault = getUpdatedRepresentationCopy("zone_type", Representation{repType: Required, create: `SECONDARY`}, zoneRepresentationPrimaryDefault)

	ZoneResourceDependenciesDefault = generateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Required, Create, tsigKeyRepresentation) +
		DefinedTagsDependencies + `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}
`
)

func TestDnsZoneResource_default(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneResource_default")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_zone.test_zone"
	datasourceName := "data.oci_dns_zones.test_zones"

	_, tokenFn := tokenizeWithHttpReplay("dns_zone")
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsZoneDestroy,
		Steps: []resource.TestStep{
			// test PRIMARY zone creation
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create, zoneRepresentationDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "external_masters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_masters.0.address", "77.64.12.1"),
					resource.TestCheckResourceAttr(resourceName, "external_masters.0.port", "53"),
					resource.TestCheckResourceAttrSet(resourceName, "external_masters.0.tsig_key_id"),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "SECONDARY"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("resource id should be different")
						}
						resId = resId2
						return err
					},
				),
			},

			// delete before next create
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault, nil),
			},
			// verify create with optionals
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create,
						representationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"})), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "nameservers.#"),
					resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
					resource.TestCheckResourceAttrSet(resourceName, "self"),
					resource.TestCheckResourceAttrSet(resourceName, "serial"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: tokenFn(config+compartmentIdVariableStr+compartmentIdUVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"}), map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "4"),
					resource.TestCheckResourceAttrSet(resourceName, "self"),
					resource.TestCheckResourceAttrSet(resourceName, "serial"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Update,
						representationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"})), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "4"),
					resource.TestCheckResourceAttrSet(resourceName, "self"),
					resource.TestCheckResourceAttrSet(resourceName, "serial"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationRequiredOnlyWithFilterDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.is_protected"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.scope", "GLOBAL"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.self"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.serial"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.version"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.zone_type", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.nameservers.#"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameContainsOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithStateOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithZoneTypeOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault)+
					compartmentIdVariableStr+ZoneResourceDependenciesDefault+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2022-04-10T19:01:09.000-00:00"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			// verify resource import
			{
				Config:            tokenFn(config, nil),
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}
