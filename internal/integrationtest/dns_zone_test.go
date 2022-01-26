// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	zoneDataSourceRepresentationRequiredOnly = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	zoneDataSourceRepresentationRequiredOnlyWithFilter = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: zoneDataSourceFilterRepresentation},
	})
	zoneDataSourceRepresentationWithNameOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
	})
	zoneDataSourceRepresentationWithNameContainsOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"name_contains": acctest.Representation{RepType: acctest.Optional, Create: `oci-zone-test`},
	})
	zoneDataSourceRepresentationWithStateOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	})
	zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
	})
	zoneDataSourceRepresentationWithTimeCreatedLessThanOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"time_created_less_than": acctest.Representation{RepType: acctest.Optional, Create: `2022-04-10T19:01:09.000-00:00`},
	})
	zoneDataSourceRepresentationWithZoneTypeOptional = acctest.RepresentationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"zone_type": acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY`},
		"view_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	})

	zoneDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_zone.test_zone.id}`}},
	}

	zoneRepresentationPrimary = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
		"zone_type":        acctest.Representation{RepType: acctest.Required, Create: `PRIMARY`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_masters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: zoneExternalMastersRepresentation},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":            acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	zoneRepresentation = acctest.GetUpdatedRepresentationCopy("zone_type", acctest.Representation{RepType: acctest.Required, Create: `SECONDARY`}, zoneRepresentationPrimary)

	zoneExternalMastersRepresentation = map[string]interface{}{
		"address":     acctest.Representation{RepType: acctest.Required, Create: `77.64.12.1`, Update: `address2`},
		"port":        acctest.Representation{RepType: acctest.Optional, Create: `53`, Update: `11`},
		"tsig_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_tsig_key.test_tsig_key.id}`},
	}

	ZoneResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dns_tsig_key", "test_tsig_key", acctest.Required, acctest.Create, tsigKeyRepresentation) +
		DefinedTagsDependencies + `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}
` + acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, viewRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsZoneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsZoneResource_basic")
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
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"})), nil), "dns", "zone", t)

	acctest.ResourceTest(t, testAccCheckDnsZoneDestroy, []resource.TestStep{
		// test PRIMARY zone creation
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
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

		// delete before next Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies, nil),
		},
		// verify Create with optionals
		{
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "nameservers.#"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "serial"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttrSet(resourceName, "view_id"),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Resource discovery is not supported for Zone resources created using scope field
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: tokenFn(config+compartmentIdVariableStr+compartmentIdUVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"}), map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "serial"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttrSet(resourceName, "view_id"),
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
			Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"})), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "serial"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttrSet(resourceName, "view_id"),
				resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

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
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationRequiredOnlyWithFilter)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "view_id"),
				resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.is_protected"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.serial"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.view_id"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.0.nameservers.#"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithNameOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
				resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithNameContainsOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithStateOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithZoneTypeOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		{
			Config: tokenFn(config+acctest.GenerateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", acctest.Optional, acctest.Create, zoneDataSourceRepresentationWithTimeCreatedLessThanOptional)+
				compartmentIdVariableStr+ZoneResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, zoneRepresentationPrimary), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2022-04-10T19:01:09.000-00:00"),
				resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
			),
		},
		// verify resource import
		{
			Config:            tokenFn(config, nil),
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getZoneImportId(resourceName),
			ResourceName:      resourceName,
		},
	})
}

func getZoneImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("zoneNameOrId/" + rs.Primary.Attributes["id"] + "/scope/" + rs.Primary.Attributes["scope"] + "/viewId/" + rs.Primary.Attributes["view_id"]), nil
	}
}

func testAccCheckDnsZoneDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_zone" {
			noResourceFound = false
			request := oci_dns.GetZoneRequest{}

			tmp := rs.Primary.ID
			request.ZoneNameOrId = &tmp

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["scope"]; ok {
				request.Scope = oci_dns.GetZoneScopeEnum(value)
			}

			if value, ok := rs.Primary.Attributes["view_id"]; ok {
				request.ViewId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")

			_, err := client.GetZone(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for 404.
			// after destruction
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DnsZone") {
		resource.AddTestSweepers("DnsZone", &resource.Sweeper{
			Name:         "DnsZone",
			Dependencies: acctest.DependencyGraph["zone"],
			F:            sweepDnsZoneResource,
		})
	}
}

func sweepDnsZoneResource(compartment string) error {
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()
	zoneIds, err := getZoneIds(compartment)
	if err != nil {
		return err
	}
	for _, zoneId := range zoneIds {
		if ok := acctest.SweeperDefaultResourceId[zoneId]; !ok {
			deleteZoneRequest := oci_dns.DeleteZoneRequest{}

			deleteZoneRequest.ZoneNameOrId = &zoneId

			deleteZoneRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteZone(context.Background(), deleteZoneRequest)
			if error != nil {
				fmt.Printf("Error deleting Zone %s %s, It is possible that the resource is already deleted. Please verify manually \n", zoneId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &zoneId, zoneSweepWaitCondition, time.Duration(3*time.Minute),
				zoneSweepResponseFetchOperation, "dns", true)
		}
	}
	return nil
}

func getZoneIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ZoneId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()

	listZonesRequest := oci_dns.ListZonesRequest{}
	listZonesRequest.CompartmentId = &compartmentId
	listZonesRequest.LifecycleState = oci_dns.ListZonesLifecycleStateActive
	listZonesResponse, err := dnsClient.ListZones(context.Background(), listZonesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Zone list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, zone := range listZonesResponse.Items {
		id := *zone.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ZoneId", id)
	}
	return resourceIds, nil
}

func zoneSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if zoneResponse, ok := response.Response.(oci_dns.GetZoneResponse); ok {
		return zoneResponse.LifecycleState != oci_dns.ZoneLifecycleStateDeleted
	}
	return false
}

func zoneSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DnsClient().GetZone(context.Background(), oci_dns.GetZoneRequest{
		ZoneNameOrId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
