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
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}
	zoneDataSourceRepresentationRequiredOnlyWithFilterDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"filter": RepresentationGroup{Required, zoneDataSourceFilterRepresentation},
	})
	zoneDataSourceRepresentationWithNameOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name": Representation{RepType: Optional, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
	})
	zoneDataSourceRepresentationWithNameContainsOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name_contains": Representation{RepType: Optional, Create: `oci-zone-test`},
	})
	zoneDataSourceRepresentationWithStateOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"state": Representation{RepType: Optional, Create: `ACTIVE`},
	})
	zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-01-01T00:00:00.000Z`},
	})
	zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_less_than": Representation{RepType: Optional, Create: `2022-04-10T19:01:09.000-00:00`},
	})
	zoneDataSourceRepresentationWithZoneTypeOptionalDefault = RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"zone_type": Representation{RepType: Optional, Create: `PRIMARY`},
	})

	zoneRepresentationPrimaryDefault = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":             Representation{RepType: Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
		"zone_type":        Representation{RepType: Required, Create: `PRIMARY`},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_masters": RepresentationGroup{Optional, zoneExternalMastersRepresentation},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	zoneRepresentationDefault = GetUpdatedRepresentationCopy("zone_type", Representation{RepType: Required, Create: `SECONDARY`}, zoneRepresentationPrimaryDefault)

	ZoneResourceDependenciesDefault = GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", Required, Create, tsigKeyRepresentation) +
		DefinedTagsDependencies + `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}
`
)

// issue-routing-tag: dns/default
func TestDnsZoneResource_default(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneResource_default")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_zone.test_zone"
	datasourceName := "data.oci_dns_zones.test_zones"

	_, tokenFn := TokenizeWithHttpReplay("dns_zone")
	var resId, resId2 string

	ResourceTest(t, testAccCheckDnsZoneDestroy, []resource.TestStep{
		// test PRIMARY zone creation
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create, zoneRepresentationDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_masters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_masters.0.address", "77.64.12.1"),
				resource.TestCheckResourceAttr(resourceName, "external_masters.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_masters.0.tsig_key_id"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "SECONDARY"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource id should be different")
					}
					resId = resId2
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault, nil),
		},
		// verify Create with optionals
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create,
					RepresentationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"})), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: tokenFn(config+compartmentIdVariableStr+compartmentIdUVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create,
					RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"}), map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Update,
					RepresentationCopyWithRemovedProperties(zoneRepresentationPrimaryDefault, []string{"external_masters"})), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationRequiredOnlyWithFilterDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameContainsOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithStateOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithZoneTypeOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimaryDefault), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
