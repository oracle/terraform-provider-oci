// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	zoneDataSourceTimeCreatedLessThanTime           = time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	zoneDataSourceRepresentationRequiredOnlyDefault = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
	zoneDataSourceRepresentationRequiredOnlyWithFilterDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DnsDnsZoneDataSourceFilterRepresentation},
	})
	zoneDataSourceRepresentationWithNameOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
	})
	zoneDataSourceRepresentationWithNameContainsOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"name_contains": acctest.Representation{RepType: acctest.Optional, Create: `oci-zone-test`},
	})
	zoneDataSourceRepresentationWithStateOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	})
	zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
	})
	zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"time_created_less_than": acctest.Representation{RepType: acctest.Optional, Create: zoneDataSourceTimeCreatedLessThanTime},
	})
	zoneDataSourceRepresentationWithZoneTypeOptionalDefault = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnlyDefault, map[string]interface{}{
		"zone_type": acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY`},
	})

	zoneRepresentationPrimaryDefault = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
		"zone_type":            acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_downstreams": acctest.RepresentationGroup{RepType: acctest.Optional, Group: zoneExternalDownstreamsRepresentation},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	zoneRepresentationDefault = acctest.GetUpdatedRepresentationCopy("zone_type", acctest.Representation{RepType: acctest.Required, Create: `SECONDARY`}, zoneRepresentationPrimaryDefault)

	zoneExternalDownstreamsRepresentation = map[string]interface{}{
		"address":     acctest.Representation{RepType: acctest.Required, Create: `1.2.3.4`, Update: `2.3.4.5`},
		"port":        acctest.Representation{RepType: acctest.Optional, Create: `53`},
		"tsig_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_tsig_key.test_tsig_key.id}`},
	}

	ZoneResourceDependenciesDefault = acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Required, acctest.Create, DnsTsigKeyRepresentation) +
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_zone.test_zone"
	datasourceName := "data.oci_dns_zones.test_zones"

	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_zone")
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDnsZoneDestroy, []resource.TestStep{
		// test PRIMARY zone creation
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.address", "1.2.3.4"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_downstreams.0.tsig_key_id"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.address", "1.2.3.4"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_downstreams.0.tsig_key_id"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(zoneRepresentationPrimaryDefault, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.address", "1.2.3.4"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_downstreams.0.tsig_key_id"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Update, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.address", "2.3.4.5"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_downstreams.0.tsig_key_id"),
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationRequiredOnlyWithFilterDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.address", "2.3.4.5"),
				resource.TestCheckResourceAttr(resourceName, "external_downstreams.0.port", "53"),
				resource.TestCheckResourceAttrSet(resourceName, "external_downstreams.0.tsig_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
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
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithNameOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithNameContainsOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithStateOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithZoneTypeOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithTimeCreatedLessThanOptionalDefault)+
				compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", zoneDataSourceTimeCreatedLessThanTime),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		// verify resource import
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependenciesDefault+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimaryDefault), nil),
			ImportState:       true,
			ImportStateVerify: true,
			ResourceName:      resourceName,
		},
	})
}
