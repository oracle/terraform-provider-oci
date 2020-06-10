package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	rrsetItemsRepresentation2 = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: `192.168.0.2`, update: `77.77.77.78`},
		"rtype":  Representation{repType: Required, create: `A`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	rrsetRepresentationAAAA = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `AAAA`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           []RepresentationGroup{{Required, rrsetItemsRepresentationAAAA}, {Required, rrsetItemsRepresentationAAAA2}},
	}

	rrsetItemsRepresentationAAAA = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`, update: `0000:0000:0000:0000:0000:8a2e:0370:0001`},
		"rtype":  Representation{repType: Required, create: `AAAA`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	rrsetItemsRepresentationAAAA2 = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: `8a2e:0000:0000:0000:0000:0370:0000:0000`},
		"rtype":  Representation{repType: Required, create: `AAAA`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	rrsetRepresentationCname = map[string]interface{}{
		"domain":          Representation{repType: Required, create: "el." + dnsDomainName},
		"rtype":           Representation{repType: Required, create: `CNAME`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           RepresentationGroup{Required, rrsetItemsRepresentationCname},
	}

	rrsetItemsRepresentationCname = map[string]interface{}{
		"domain": Representation{repType: Required, create: "el." + dnsDomainName},
		"rdata":  Representation{repType: Required, create: dnsDomainName},
		"rtype":  Representation{repType: Required, create: `CNAME`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	rrsetRepresentationTxt = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `TXT`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           RepresentationGroup{Required, rrsetItemsRepresentationTxt},
	}

	rrsetItemsRepresentationTxt = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: "arbitrary text"},
		"rtype":  Representation{repType: Required, create: `TXT`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	rrsetRepresentationAlias = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `ALIAS`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           RepresentationGroup{Required, rrsetItemsRepresentationAlias},
	}

	rrsetItemsRepresentationAlias = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: "other.tf-provider.oci-record-test"},
		"rtype":  Representation{repType: Required, create: `ALIAS`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}
)

func TestResourceDnsRrsetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsRrsetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create,
						getUpdatedRepresentationCopy("items", []RepresentationGroup{{Optional, rrsetItemsRepresentation}, {Optional, rrsetItemsRepresentation2}}, rrsetRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "192.168.0.1",
						"rtype":  "A",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
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
			// verify update
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "192.168.0.1",
						"rtype":  "A",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
			},

			// verify create AAAA
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationAAAA),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "2001:db8:85a3::8a2e:370:7334",
						"rtype":  "AAAA",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
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
			// verify update AAAA
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentationAAAA),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "::8a2e:370:1",
						"rtype":  "AAAA",
						"ttl":    "1000",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
			},
			// verify create CNAME
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationCname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", "el."+dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": "el." + dnsDomainName,
						"rdata":  dnsDomainName + ".",
						"rtype":  "CNAME",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "CNAME"),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
			},
			// verify create TXT
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationTxt),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "\"arbitrary\" \"text\"",
						"rtype":  "TXT",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "TXT"),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
			},
			// verify create ALIAS
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationAlias),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "other.tf-provider.oci-record-test.",
						"rtype":  "ALIAS",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "ALIAS"),
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
		},
	})
}

func TestResourceDnsRrsetResource_iterative(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsRrsetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					`
					locals {
  						test_ips = ["192.168.0.1", "192.168.0.2"]
					}
					resource "oci_dns_rrset" "test_rrset" {
  						zone_name_or_id = "${oci_dns_zone.test_zone.name}"
  						domain          = "${oci_dns_zone.test_zone.name}"
  						rtype           = "A"
						compartment_id  = "${var.compartment_id}"

  						dynamic items {
    						for_each = "${local.test_ips}"
    						content {
      							domain = "${oci_dns_zone.test_zone.name}"
      							rtype = "A"
								rdata = "${items.value}"
      							ttl = 3600
    						}
  						}
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "192.168.0.1",
						"rtype":  "A",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					`
					locals {
  						test_ips = ["192.168.0.2"]
					}
					resource "oci_dns_rrset" "test_rrset" {
  						zone_name_or_id = "${oci_dns_zone.test_zone.name}"
  						domain          = "${oci_dns_zone.test_zone.name}"
  						rtype           = "A"
						compartment_id  = "${var.compartment_id}"

  						dynamic items {
    						for_each = "${local.test_ips}"
    						content {
      							domain = "${oci_dns_zone.test_zone.name}"
      							rtype = "A"
								rdata = "${items.value}"
      							ttl = 3600
    						}
  						}
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "192.168.0.2",
						"rtype":  "A",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}
