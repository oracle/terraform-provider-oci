// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	RecordRequiredOnlyResource = RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	#Required
	zone_name_or_id = "${oci_dns_zone.test_zone.name}"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.${var.record_items_domain}"
	rdata = "${var.record_items_rdata}"
	rtype = "${var.record_items_rtype}"
	ttl = "${var.record_items_ttl}"
}
`

	RecordResourceConfig = RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	#Required
	zone_name_or_id = "${oci_dns_zone.test_zone.name}"

	#Optional
	compartment_id = "${var.compartment_id}"

	domain = "${data.oci_identity_tenancy.test_tenancy.name}.${var.record_items_domain}"
	rdata = "${var.record_items_rdata}"
	rtype = "${var.record_items_rtype}"
	ttl = "${var.record_items_ttl}"
}
`
	RecordPropertyVariables = `
variable "record_items_domain" { default = "oci-test" }
variable "record_items_rdata" { default = "192.168.0.1" }
variable "record_items_rtype" { default = "A" }
variable "record_items_ttl" { default = 3600 }
`
	RecordResourceDependencies = `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	zone_type = "PRIMARY"
}

resource "oci_dns_zone" "test_zone2" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}2.oci-test"
	zone_type = "PRIMARY"
}`
)

func TestDnsRecordsResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_record.test_record"
	datasourceName := "data.oci_dns_records.test_records"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + RecordPropertyVariables + compartmentIdVariableStr + RecordRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RecordResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-test")),
					resource.TestCheckResourceAttr(resourceName, "is_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
					resource.TestCheckResourceAttrSet(resourceName, "record_hash"),
					resource.TestCheckResourceAttrSet(resourceName, "rrset_version"),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					TestCheckResourceAttributesEqual(resourceName, "zone_name_or_id", "oci_dns_zone.test_zone", "name"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("resource was not recreated after delete")
						}
						resId = resId2
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "record_items_domain" { default = "oci-test" }
variable "record_items_rdata" { default = "77.77.77.77" }
variable "record_items_rtype" { default = "A" }
variable "record_items_ttl" { default = 1000 }
                ` + compartmentIdVariableStr + RecordResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-test")),
					resource.TestCheckResourceAttr(resourceName, "is_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "rdata", "77.77.77.77"),
					resource.TestCheckResourceAttrSet(resourceName, "record_hash"),
					resource.TestCheckResourceAttrSet(resourceName, "rrset_version"),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1000"),
					TestCheckResourceAttributesEqual(resourceName, "zone_name_or_id", "oci_dns_zone.test_zone", "name"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("record hash was the same after an update, it should be different")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_zone.name}"

  # optional
  domain = "${oci_dns_zone.test_zone.name}"
  rtype = "NS"
  sort_by = "ttl"
  sort_order = "DESC"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "records.0.rtype", "NS"),
					resource.TestCheckResourceAttr(datasourceName, "records.0.ttl", "86400"),
				),
			},
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_zone.name}"
  domain = "${oci_dns_zone.test_zone.name}"
	filter {
	  name = "rtype"
	  values = ["SOA"]
	}
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "records.#", "1"),
				),
			},
		},
	})
}

func TestDnsRecordsResource_diffSuppression(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_record.test_record"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify AAAA ipv6 shortening does not cause diffs
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "AAAA"
	rdata = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
					resource.TestCheckResourceAttr(resourceName, "rdata", "2001:db8:85a3::8a2e:370:7334"),
				),
			},
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "AAAA"
	rdata = "0000:0000:0000:0000:0000:8a2e:0370:0001"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "::8a2e:370:1"),
				),
			},
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "AAAA"
	rdata = "8a2e:0000:0000:0000:0000:0370:0000:0000"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "8a2e::370:0:0"),
				),
			},
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "TXT"
	rdata = "arbitrary text"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "\"arbitrary\" \"text\""),
				),
			},
			// this tests represents several record types where the service appends a `.` to the rdata
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "ALIAS"
	rdata = "other.tf-provider.oci-test"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "other.tf-provider.oci-test."),
				),
			},
		},
	})
}

func TestDnsRecordsResource_badUpdate(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_record.test_record"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "3600"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
				),
			},
			{
				Config: config + RecordPropertyVariables + compartmentIdVariableStr + RecordResourceDependencies + `
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.oci-test"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "-1"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
				//resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
				// todo: this test was attempting to verify the resource is not changed if the update operation fails
				// but this terraform testing library does not run "Checks" if you add an error expectation ;_;
				),
				ExpectError: regexp.MustCompile("-1 is not a valid value for TTL"),
			},
		},
	})
}
