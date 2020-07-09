// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_dns "github.com/oracle/oci-go-sdk/v25/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	recordDataSourceRepresentation = map[string]interface{}{
		"zone_name_or_id":   Representation{repType: Required, create: `${oci_dns_zone.test_zone.name}`},
		"if_modified_since": Representation{repType: Optional, create: `ifModifiedSince`},
		"if_none_match":     Representation{repType: Optional, create: `ifNoneMatch`},
		"compartment_id":    Representation{repType: Optional, create: `${var.compartment_id}`},
		"domain":            Representation{repType: Optional, create: `domain`},
		"domain_contains":   Representation{repType: Optional, create: `domainContains`},
		"rtype":             Representation{repType: Optional, create: `rtype`},
		"zone_version":      Representation{repType: Optional, create: `zoneVersion`},
		"filter":            RepresentationGroup{Required, recordDataSourceFilterRepresentation}}
	recordDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_record.test_record.id}`}},
	}

	recordRepresentation = map[string]interface{}{
		"domain":          Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"rdata":           Representation{repType: Required, create: `192.168.0.1`, update: `77.77.77.77`},
		"rtype":           Representation{repType: Required, create: `A`},
		"ttl":             Representation{repType: Required, create: `3600`, update: `1000`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.name}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	RecordResourceDependencies = `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	zone_type = "PRIMARY"
}
`
)

func TestDnsRecordsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_dns_record.test_record"

	_, tokenFn := tokenizeWithHttpReplay("dns_resource")
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsRecordDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_record", "test_record", Required, Create, recordRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// delete before next create
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies, nil),
			},
			// verify create with optionals
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_record", "test_record", Optional, Create, recordRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-record-test")),
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
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_record", "test_record", Optional, Update, recordRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-record-test")),
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
		},
	})
}

// The datasource tests are kept separate from the previous test steps.
// This was because the datasource steps do not create a record resource (and won't need one because, because a zone has default records).
// If this was kept in the previous test case, the CheckDestroy step would run after the datasource steps ran and would fail
// because it wouldn't have a record resource to delete and to verify destruction for.
func TestDnsRecordsResource_datasources(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_datasources")
	defer httpreplay.SaveScenario()
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := tokenizeWithHttpReplay("dns_data_source")
	datasourceName := "data.oci_dns_records.test_records"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_zone.name}"

  # optional
  domain = "${oci_dns_zone.test_zone.name}"
  rtype = "NS"
  sort_by = "ttl"
  sort_order = "DESC"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "records.0.rtype", "NS"),
					resource.TestCheckResourceAttr(datasourceName, "records.0.ttl", "86400"),
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_zone.name}"
  domain = "${oci_dns_zone.test_zone.name}"
	filter {
	  name = "rtype"
	  values = ["SOA"]
	}
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "records.#", "1"),
				),
			},
		},
	})
}

func TestDnsRecordsResource_diffSuppression(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_diffSuppression")
	defer httpreplay.SaveScenario()
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := tokenizeWithHttpReplay("dns_diff")
	resourceName := "oci_dns_record.test_record"
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify AAAA ipv6 shortening does not cause diffs
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "AAAA"
	rdata = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
					resource.TestCheckResourceAttr(resourceName, "rdata", "2001:db8:85a3::8a2e:370:7334"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.OCI-record-test"
	rtype = "AAAA"
	rdata = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
					resource.TestCheckResourceAttr(resourceName, "rdata", "2001:db8:85a3::8a2e:370:7334"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "AAAA"
	rdata = "0000:0000:0000:0000:0000:8a2e:0370:0001"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "::8a2e:370:1"),
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "AAAA"
	rdata = "8a2e:0000:0000:0000:0000:0370:0000:0000"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "8a2e::370:0:0"),
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "TXT"
	rdata = "arbitrary text"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "\"arbitrary\" \"text\""),
				),
			},
			// this tests represents several record types where the service appends a `.` to the rdata
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "ALIAS"
	rdata = "other.tf-provider.oci-record-test"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "other.tf-provider.oci-record-test."),
				),
			},
		},
	})
}

func TestDnsRecordsResource_badUpdate(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_badUpdate")
	defer httpreplay.SaveScenario()
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := tokenizeWithHttpReplay("dns_bad_update")
	resourceName := "oci_dns_record.test_record"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "3600"
}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
				),
			},
			{
				Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	domain = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "-1"
}`, nil),
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

func testAccCheckDnsRecordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_record" {
			noResourceFound = false
			request := oci_dns.GetZoneRecordsRequest{}

			if value, ok := rs.Primary.Attributes["zone_name_or_id"]; ok {
				request.ZoneNameOrId = &value
			}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			response, err := client.GetZoneRecords(context.Background(), request)
			if err == nil {
				// Convert the InstanceState attributes to a ResourceData expected by the lookup function
				attributes := convertToObjectMap(rs.Primary.Attributes)
				resourceData := schema.TestResourceDataRaw(&testing.T{}, DnsRecordResource().Schema, attributes)

				//page through records
				recordCollection := response.RecordCollection
				request.Page = response.OpcNextPage

				for request.Page != nil {
					listResponse, err := client.GetZoneRecords(context.Background(), request)
					if err != nil {
						return err
					}

					recordCollection.Items = append(recordCollection.Items, listResponse.Items...)
					request.Page = listResponse.OpcNextPage
				}
				_, err = findItem(&response.Items, resourceData)
				if err == nil {
					return fmt.Errorf("resource still exists")
				}

				// no error and item not found, item is deleted
				return nil
			}

			// TODO: If we get here, then technically this isn't verifying that the record resource was destroyed.
			// But it is verifying that at least the zone was destroyed (which guarantees that the records were destroyed)
			// This is a test gap because of Terraform test framework destroying all resources.
			// Ideally, the test framework should do a targeted destroy of the record prior to calling CheckDestroy.
			//Verify that exception is for '404 not found'.
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
