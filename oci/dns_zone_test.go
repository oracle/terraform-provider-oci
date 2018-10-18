// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
)

var (
	ZoneRequiredOnlyResource = ZoneResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary)

	zoneDataSourceRepresentationRequiredOnly = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}
	zoneDataSourceRepresentationRequiredOnlyWithFilter = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"filter": RepresentationGroup{Required, zoneDataSourceFilterRepresentation},
	})
	zoneDataSourceRepresentationWithNameOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"name": Representation{repType: Optional, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
	})
	zoneDataSourceRepresentationWithNameContainsOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"name_contains": Representation{repType: Optional, create: `oci-zone-test`},
	})
	zoneDataSourceRepresentationWithStateOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"state": Representation{repType: Optional, create: `ACTIVE`},
	})
	zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
	})
	zoneDataSourceRepresentationWithTimeCreatedLessThanOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"time_created_less_than": Representation{repType: Optional, create: `2022-04-10T19:01:09.000-00:00`},
	})
	zoneDataSourceRepresentationWithZoneTypeOptional = representationCopyWithNewProperties(zoneDataSourceRepresentationRequiredOnly, map[string]interface{}{
		"zone_type": Representation{repType: Optional, create: `PRIMARY`},
	})

	zoneDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_zone.test_zone.id}`}},
	}

	zoneRepresentationPrimary = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"name":             Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test`},
		"zone_type":        Representation{repType: Required, create: `PRIMARY`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_masters": RepresentationGroup{Optional, zoneExternalMastersRepresentation},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}
	zoneRepresentation = getUpdatedRepresentationCopy("zone_type", "SECONDARY", zoneRepresentationPrimary)

	zoneExternalMastersRepresentation = map[string]interface{}{
		"address": Representation{repType: Required, create: `77.64.12.1`, update: `address2`},
		"port":    Representation{repType: Optional, create: `53`, update: `11`},
		"tsig":    RepresentationGroup{Optional, zoneExternalMastersTsigRepresentation},
	}
	zoneExternalMastersTsigRepresentation = map[string]interface{}{
		"algorithm": Representation{repType: Required, create: `hmac-sha1`, update: `algorithm2`},
		"name":      Representation{repType: Required, create: `name`, update: `name2`},
		"secret":    Representation{repType: Required, create: `c2VjcmV0`, update: `secret2`},
	}

	ZoneResourceDependencies = DefinedTagsDependencies + `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}
`
)

func TestDnsZoneResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_zone.test_zone"
	datasourceName := "data.oci_dns_zones.test_zones"

	_, tokenFn := tokenize()
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
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
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
			// test SECONDARY zone creation, force new at the same time
			// Disable SECONDARY zone creation test for now, since it's using a bogus external_master server.
			// This will put the zone in a bad state and cause any records in this zone to fail during PATCH.
			/*
				{
					Config: tokenFn(config + compartmentIdVariableStr + ZoneResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create, zoneRepresentation), nil),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
						resource.TestCheckResourceAttr(resourceName, "external_masters.#", "1"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.address", "77.64.12.1"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.port", "53"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.tsig.#", "1"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.tsig.0.algorithm", "hmac-sha1"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.tsig.0.name", "name"),
						resource.TestCheckResourceAttr(resourceName, "external_masters.0.tsig.0.secret", "c2VjcmV0"),
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
			*/
			// delete before next create
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies, nil),
			},
			// verify create with optionals
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Create,
						representationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"})), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "nameservers.#"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: tokenFn(config+compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Optional, Update,
						representationCopyWithRemovedProperties(zoneRepresentationPrimary, []string{"external_masters"})), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

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
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationRequiredOnlyWithFilter)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.0.nameservers.#"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithNameContainsOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithStateOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithZoneTypeOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedGreaterThanOrEqualToOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "zones.0.freeform_tags.%", "1"),
				),
			},
			{
				Config: tokenFn(config+generateDataSourceFromRepresentationMap("oci_dns_zones", "test_zones", Optional, Create, zoneDataSourceRepresentationWithTimeCreatedLessThanOptional)+
					compartmentIdVariableStr+ZoneResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_zone", "test_zone", Required, Create, zoneRepresentationPrimary), nil),
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

func testAccCheckDnsZoneDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_zone" {
			noResourceFound = false
			request := oci_dns.GetZoneRequest{}

			tmp := rs.Primary.ID
			request.ZoneNameOrId = &tmp

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			_, err := client.GetZone(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for 400.
			// Normally expect 404, but DNS service returns a "InvalidParameter. Bad Request - Invalid domain name. http status code: 400"
			// after destruction
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 400 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func initDnsZoneSweeper() {
	resource.AddTestSweepers("DnsZone", &resource.Sweeper{
		Name:         "DnsZone",
		Dependencies: DependencyGraph["zone"],
		F:            sweepDnsZoneResource,
	})
}

func sweepDnsZoneResource(compartment string) error {
	compartmentId := compartment
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient

	listZonesRequest := oci_dns.ListZonesRequest{}
	listZonesRequest.CompartmentId = &compartmentId
	listZonesRequest.LifecycleState = oci_dns.ListZonesLifecycleStateActive
	listZonesResponse, err := dnsClient.ListZones(context.Background(), listZonesRequest)

	if err != nil {
		return fmt.Errorf("Error getting Zone list for compartment id : %s , %s \n", compartmentId, err)
	}

	for _, zone := range listZonesResponse.Items {
		log.Printf("deleting zone %s ", *zone.Id)

		deleteZoneRequest := oci_dns.DeleteZoneRequest{}

		deleteZoneRequest.ZoneNameOrId = zone.Id

		deleteZoneRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")
		_, error := dnsClient.DeleteZone(context.Background(), deleteZoneRequest)
		if error != nil {
			fmt.Printf("Error deleting Zone %s %s, It is possible that the resource is already deleted. Please verify manually \n", *zone.Id, error)
			continue
		}

		getZoneRequest := oci_dns.GetZoneRequest{}

		getZoneRequest.ZoneNameOrId = zone.Id

		_, error = dnsClient.GetZone(context.Background(), getZoneRequest)
		if error != nil {
			fmt.Printf("Error retrieving Zone state %s \n", error)
			continue
		}

		waitTillCondition(testAccProvider, zone.Id, zoneSweepWaitCondition, time.Duration(3*time.Minute),
			zoneSweepResponseFetchOperation, "dns", true)
	}
	return nil
}

func zoneSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if zoneResponse, ok := response.Response.(oci_dns.GetZoneResponse); ok {
		return zoneResponse.LifecycleState == oci_dns.ZoneLifecycleStateDeleted
	}
	return false
}

func zoneSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dnsClient.GetZone(context.Background(), oci_dns.GetZoneRequest{
		ZoneNameOrId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
