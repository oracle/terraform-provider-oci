// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

var (
	DnsRrsetRequiredOnlyResource = DnsRrsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, DnsRrsetRepresentation)

	DnsRrsetResourceConfig = DnsRrsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, DnsRrsetRepresentation)

	DnsDnsRrsetSingularDataSourceRepresentation = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	DnsDnsRrsetDataSourceRepresentation = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}

	dnsDomainName = utils.RandomString(5, utils.CharsetWithoutDigits) + ".token.oci-record-test"

	DnsRrsetRepresentation = map[string]interface{}{
		"domain":          acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rtype":           acctest.Representation{RepType: acctest.Required, Create: `A`},
		"zone_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DnsRrsetItemsRepresentation},
		"scope":           acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"view_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
	}
	DnsRrsetItemsRepresentation = map[string]interface{}{
		"domain": acctest.Representation{RepType: acctest.Required, Create: dnsDomainName},
		"rdata":  acctest.Representation{RepType: acctest.Required, Create: `192.168.0.1`, Update: `77.77.77.77`},
		"rtype":  acctest.Representation{RepType: acctest.Required, Create: `A`},
		"ttl":    acctest.Representation{RepType: acctest.Required, Create: `3600`, Update: `1000`},
	}

	DnsRrsetResourceDependencies = `
	data "oci_identity_tenancy" "test_tenancy" {
		tenancy_id = "${var.tenancy_ocid}"
	}
	` + acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: dnsDomainName}, DnsDnsZoneRepresentationPrimary)) +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, DnsViewRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsRrsetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsRrsetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"
	datasourceName := "data.oci_dns_rrsets.test_rrsets"
	singularDatasourceName := "data.oci_dns_rrset.test_rrset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DnsRrsetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, DnsRrsetRepresentation), "dns", "rrset", t)

	acctest.ResourceTest(t, testAccCheckDnsRrsetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DnsRrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, DnsRrsetRepresentation),
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
			Config: config + compartmentIdVariableStr + DnsRrsetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DnsRrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Create, DnsRrsetRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "view_id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DnsRrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, DnsRrsetRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "view_id"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_rrsets", "test_rrsets", acctest.Optional, acctest.Update, DnsDnsRrsetDataSourceRepresentation) +
				compartmentIdVariableStr + DnsRrsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Optional, acctest.Update, DnsRrsetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "rrsets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rrsets.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rrsets.0.domain", dnsDomainName),
				resource.TestCheckResourceAttr(datasourceName, "rrsets.0.rtype", "A"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "rrsets.0.items", map[string]string{
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
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_rrset", "test_rrset", acctest.Required, acctest.Create, DnsDnsRrsetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DnsRrsetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", dnsDomainName),
				resource.TestCheckResourceAttr(singularDatasourceName, "rtype", "A"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "view_id"),
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
		// verify resource import
		{
			Config:            config + DnsRrsetRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getRrSetImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"compartment_id",
				"scope",
				"view_id",
			},
			ResourceName: resourceName,
		},
	})
}

func getRrSetImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("zoneNameOrId/" + rs.Primary.Attributes["zone_name_or_id"] + "/domain/" + rs.Primary.Attributes["domain"] + "/rtype/" + rs.Primary.Attributes["rtype"] + "/scope/" +
			rs.Primary.Attributes["scope"] + "/viewId/" + rs.Primary.Attributes["view_id"]), nil
	}
}

func testAccCheckDnsRrsetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_rrset" {
			noResourceFound = false
			request := oci_dns.GetRRSetRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["domain"]; ok {
				request.Domain = &value
			}

			if value, ok := rs.Primary.Attributes["rtype"]; ok {
				request.Rtype = &value
			}

			if value, ok := rs.Primary.Attributes["scope"]; ok {
				request.Scope = oci_dns.GetRRSetScopeEnum(value)
			}

			if value, ok := rs.Primary.Attributes["view_id"]; ok {
				request.ViewId = &value
			}

			if value, ok := rs.Primary.Attributes["zone_name_or_id"]; ok {
				request.ZoneNameOrId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")

			_, err := client.GetRRSet(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

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
