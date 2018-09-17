// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
)

const (
	ZoneRequiredOnlyResource = ZoneResourceDependencies + `
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test"
	zone_type = "${var.zone_zone_type}"
}
`

	ZoneResourceConfig = ZoneResourceDependencies + `
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test"
	zone_type = "SECONDARY"

	#Optional
	external_masters {
		#Required
		address = "${var.zone_external_masters_address}"

		#Optional
		port = "${var.zone_external_masters_port}"
		tsig {
			#Required
			algorithm = "${var.zone_external_masters_tsig_algorithm}"
			name = "${var.zone_external_masters_tsig_name}"
			secret = "${var.zone_external_masters_tsig_secret}"
		}
	}
}
`
	ZonePropertyVariables = `
variable "zone_external_masters_address" { default = "77.64.12.1" }
variable "zone_external_masters_port" { default = 53 }  // (the only allowed value)
variable "zone_external_masters_tsig_algorithm" { default = "hmac-sha1" }
variable "zone_external_masters_tsig_name" { default = "name" }
variable "zone_external_masters_tsig_secret" { default = "c2VjcmV0" }
variable "zone_name" { default = "oci-zone-test" }
variable "zone_zone_type" { default = "PRIMARY" }

`
	ZoneResourceDependencies = `
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
	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsZoneDestroy,
		Steps: []resource.TestStep{
			// test PRIMARY zone creation
			{
				Config: tokenFn(config+ZonePropertyVariables+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(resourceName, "zone_type", "PRIMARY"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// test SECONDARY zone creation, force new at the same time
			// Disable SECONDARY zone creation test for now, since it's using a bogus external_master server.
			// This will put the zone in a bad state and cause any records in this zone to fail during PATCH.
			/*
				{
					Config: tokenFn(config + ZonePropertyVariables + compartmentIdVariableStr + ZoneResourceConfig, nil),
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
			// verify datasource
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  filter {
    name = "id"
    values = ["${oci_dns_zone.test_zone.id}"]
  }
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  name = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-zone-test"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("\\.oci-zone-test")),
					resource.TestCheckResourceAttr(datasourceName, "zones.#", "1"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  name_contains = "oci-zone-test"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name_contains", "oci-zone-test"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  state = "ACTIVE"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  zone_type = "PRIMARY"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "zone_type", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  time_created_greater_than_or_equal_to = "2018-04-10T19:01:09.000-00:00"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-04-10T19:01:09.000-00:00"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
				),
			},
			{
				Config: tokenFn(config+ZonePropertyVariables+`
data "oci_dns_zones" "test_zones" {
  compartment_id = "${var.compartment_id}"
  time_created_less_than = "2022-04-10T19:01:09.000-00:00"
}
                `+compartmentIdVariableStr+ZoneRequiredOnlyResource, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2022-04-10T19:01:09.000-00:00"),
					resource.TestCheckResourceAttrSet(datasourceName, "zones.#"),
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
