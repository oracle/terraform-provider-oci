package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	rrsetSingularDataSourceRepresentationDefault = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `A`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	rrsetRepresentationDefault = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `A`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           RepresentationGroup{Optional, rrsetItemsRepresentation},
	}

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
		"scope":           Representation{repType: Required, create: `PRIVATE`},
		"view_id":         Representation{repType: Required, create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationAAAADefault = map[string]interface{}{
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
		"scope":           Representation{repType: Required, create: `PRIVATE`},
		"view_id":         Representation{repType: Required, create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationCnameDefault = map[string]interface{}{
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
		"scope":           Representation{repType: Required, create: `PRIVATE`},
		"view_id":         Representation{repType: Required, create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationTxtDefault = map[string]interface{}{
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
		"scope":           Representation{repType: Required, create: `PRIVATE`},
		"view_id":         Representation{repType: Required, create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationAliasDefault = map[string]interface{}{
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

	RrsetRequiredOnlyResourceDefault = RrsetResourceDependenciesDefault +
		generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetRepresentationDefault)
	RrsetResourceConfigDefault = RrsetResourceDependenciesDefault +
		generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentationDefault)

	RrsetResourceDependenciesDefault = `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "` + dnsDomainName + `"
	zone_type = "PRIMARY"
}
`
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
						// Resource discovery is not supported for Rrset resources created using scope field
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
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
					resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
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
					resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
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
					resource.TestCheckResourceAttr(resourceName, "rtype", "CNAME"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
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
					resource.TestCheckResourceAttr(resourceName, "rtype", "TXT"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
						return err
					},
				),
			},
			/* TODO PN: Alias records are not yet supported
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
			}, */
		},
	})
}

func TestResourceDnsRrsetResource_default(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_default")
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create,
						getUpdatedRepresentationCopy("items", []RepresentationGroup{{Optional, rrsetItemsRepresentation}, {Optional, rrsetItemsRepresentation2}}, rrsetRepresentationDefault)),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationDefault),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
			},

			// verify create AAAA
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationAAAADefault),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentationAAAADefault),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
			},
			// verify create CNAME
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationCnameDefault),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
			},
			// verify create TXT
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationTxtDefault),
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
			},
			// verify create ALIAS
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationAliasDefault),
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

func TestResourceDnsRrsetResource_iterative_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_iterative_basic")
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
					`
					locals {
						test_ips = ["192.168.0.1", "192.168.0.2"]
					}
					resource "oci_dns_rrset" "test_rrset" {
						zone_name_or_id = "${oci_dns_zone.test_zone.name}"
						domain          = "${oci_dns_zone.test_zone.name}"
						rtype           = "A"
						compartment_id  = "${var.compartment_id}"
						scope           = "PRIVATE"
						view_id         = "${oci_dns_view.test_view.id}"

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
						scope           = "PRIVATE"
						view_id         = "${oci_dns_view.test_view.id}"

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

func TestResourceDnsRrsetResource_iterative_default(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_iterative_default")
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
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

func TestDnsRrsetResource_default(t *testing.T) {
	httpreplay.SetScenario("TestDnsRrsetResource_default")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	singularDatasourceName := "data.oci_dns_rrset.test_rrset"

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
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetRepresentationDefault),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentationDefault),
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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentationDefault),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "77.77.77.77",
						"rtype":  "A",
						"ttl":    "1000",
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
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetSingularDataSourceRepresentationDefault) +
					compartmentIdVariableStr + RrsetResourceConfigDefault,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(singularDatasourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_name_or_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "77.77.77.77",
						"rtype":  "A",
						"ttl":    "1000",
					},
						[]string{
							"is_protected",
							"record_hash",
							"rrset_version",
						}),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + RrsetResourceConfigDefault,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"compartment_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
