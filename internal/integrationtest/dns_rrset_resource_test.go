package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	rrsetSingularDataSourceRepresentationDefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	rrsetRepresentationDefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: rrsetItemsRepresentation},
	}

	rrsetItemsRepresentation2 = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `192.168.0.2`, Update: `77.77.77.78`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `A`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	rrsetRepresentationAAAA = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `AAAA`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           []acctest.RepresentationGroup{{RepType: acctest.Required, Group: rrsetItemsRepresentationAAAA}, {RepType: acctest.Required, Group: rrsetItemsRepresentationAAAA2}},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationAAAADefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `AAAA`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           []acctest.RepresentationGroup{{RepType: acctest.Required, Group: rrsetItemsRepresentationAAAA}, {RepType: acctest.Required, Group: rrsetItemsRepresentationAAAA2}},
	}

	rrsetItemsRepresentationAAAA = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`, Update: `0000:0000:0000:0000:0000:8a2e:0370:0001`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `AAAA`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	rrsetItemsRepresentationAAAA2 = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `8a2e:0000:0000:0000:0000:0370:0000:0000`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `AAAA`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	rrsetRepresentationCname = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: "el." + dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `CNAME`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationCname},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationCnameDefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: "el." + dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `CNAME`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationCname},
	}

	rrsetItemsRepresentationCname = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: "el." + dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `CNAME`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	rrsetRepresentationTxt = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `TXT`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationTxt},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationTxtDefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `TXT`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationTxt},
	}

	rrsetItemsRepresentationTxt = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: "arbitrary text"},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `TXT`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	rrsetRepresentationAlias = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `ALIAS`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationAlias},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	rrsetRepresentationAliasDefault = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `ALIAS`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: rrsetItemsRepresentationAlias},
	}

	rrsetItemsRepresentationAlias = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: "other.tf-provider.oci-record-test"},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `ALIAS`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	RrsetRequiredOnlyResourceDefault = RrsetResourceDependenciesDefault +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, rrsetRepresentationDefault)
	RrsetResourceConfigDefault = RrsetResourceDependenciesDefault +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, rrsetRepresentationDefault)

	RrsetResourceDependenciesDefault = `
	data "oci_identity_tenancy" "test_tenancy" {
		tenancy_id = "${var.tenancy_ocid}"
	}
	` + acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: dnsDomainName}, zoneRepresentationPrimary), []string{"scope", "view_id"}))
)

// issue-routing-tag: dns/default
func TestResourceDnsRrsetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("items", []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: rrsetItemsRepresentation}, {RepType: acctest.Optional, Group: rrsetItemsRepresentation2}}, rrsetRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Resource discovery is not supported for Rrset resources created using scope field
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify Update
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
		},

		// verify Create AAAA
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationAAAA),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify Update AAAA
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, rrsetRepresentationAAAA),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
		},
		// verify Create CNAME
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationCname),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", "el."+dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "CNAME"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
		},
		// verify Create TXT
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationTxt),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "TXT"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		/* TODO PN: Alias records are not yet supported
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
		},
		// verify Create ALIAS
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationAlias),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "ALIAS"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		}, */
	})
}

// issue-routing-tag: dns/default
func TestResourceDnsRrsetResource_default(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_default")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("items", []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: rrsetItemsRepresentation}, {RepType: acctest.Optional, Group: rrsetItemsRepresentation2}}, rrsetRepresentationDefault)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		// verify Update
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
		},

		// verify Create AAAA
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationAAAADefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "2001:db8:85a3::8a2e:370:7334",
					"rtype":  "AAAA",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		// verify Update AAAA
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, rrsetRepresentationAAAADefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "::8a2e:370:1",
					"rtype":  "AAAA",
					"ttl":    "1000",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
		},
		// verify Create CNAME
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationCnameDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", "el."+dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": "el." + dnsDomainName,
					"rdata":  dnsDomainName + ".",
					"rtype":  "CNAME",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "CNAME"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
		},
		// verify Create TXT
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationTxtDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "\"arbitrary\" \"text\"",
					"rtype":  "TXT",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "TXT"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
		},
		// verify Create ALIAS
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationAliasDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "other.tf-provider.oci-record-test.",
					"rtype":  "ALIAS",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "ALIAS"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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
	})
}

// issue-routing-tag: dns/default
func TestResourceDnsRrsetResource_iterative_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_iterative_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
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
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update
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
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.2",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: dns/default
func TestResourceDnsRrsetResource_iterative_default(t *testing.T) {
	httpreplay.SetScenario("TestResourceDnsRrsetResource_iterative_default")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
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
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update
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
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.2",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: dns/default
func TestDnsRrsetResource_default(t *testing.T) {
	httpreplay.SetScenario("TestDnsRrsetResource_default")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	singularDatasourceName := "data.oci_dns_rrset.test_rrset"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, rrsetRepresentationDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, rrsetRepresentationDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "192.168.0.1",
					"rtype":  "A",
					"ttl":    "3600",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RrsetResourceDependenciesDefault +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, rrsetRepresentationDefault),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"domain": dnsDomainName,
					"rdata":  "77.77.77.77",
					"rtype":  "A",
					"ttl":    "1000",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, rrsetSingularDataSourceRepresentationDefault) +
				compartmentIdVariableStr + RrsetResourceConfigDefault,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(singularDatasourceName, "rtype", "A"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_name_or_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
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
	})
}
