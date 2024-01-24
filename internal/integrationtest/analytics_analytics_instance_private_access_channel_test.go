// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsAnalyticsInstancePrivateAccessChannelRequiredOnlyResource = AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Required, acctest.Create, analyticsInstancePrivateAccessChannelRepresentation)

	AnalyticsAnalyticsInstancePrivateAccessChannelResourceConfig = AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Optional, acctest.Update, analyticsInstancePrivateAccessChannelRepresentation)

	analyticsInstancePrivateAccessChannelRepresentation = map[string]interface{}{
		"analytics_instance_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `example_private_access_channel`, Update: `example_private_access_channel2`},
		"private_source_dns_zones":   acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`, Update: `${oci_core_vcn.test_vcn.id}`},
		"network_security_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"private_source_scan_hosts":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: analyticsInstancePrivateAccessChannelPrivateSourceScanHostsRepresentation},
	}
	analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation = map[string]interface{}{
		"dns_zone":    acctest.Representation{RepType: acctest.Required, Create: `terraformtest.oraclevcn.com`, Update: `terraformtest2.oraclevcn.com`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `Tenant VCN DNS Zone`, Update: `Tenant VCN DNS Zone 2`},
	}
	analyticsInstancePrivateAccessChannelPrivateSourceScanHostsRepresentation = map[string]interface{}{
		"scan_hostname": acctest.Representation{RepType: acctest.Required, Create: `scanHostname`, Update: `scanHostname2`},
		"scan_port":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `Example OCI SCAN label`, Update: `Example OCI SCAN label 2`},
	}

	analyticsinstanceNameForPac = utils.RandomString(15, utils.CharsetWithoutDigits)

	//Please note that for create oci_analytics_analytics_instance we have to change the instance name otherwise tests for analytics will fail with error: Instance name it is still in use even if instance got deleted successful
	AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: analyticsinstanceNameForPac}, analyticsInstanceRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_analytics_analytics_instance_private_access_channel.test_analytics_instance_private_access_channel"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Required, acctest.Create, analyticsInstancePrivateAccessChannelRepresentation), "analytics", "analyticsInstancePrivateAccessChannel", t)

	acctest.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Required, acctest.Create, analyticsInstancePrivateAccessChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Optional, acctest.Create, analyticsInstancePrivateAccessChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.description", "Tenant VCN DNS Zone"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.description", "Example OCI SCAN label"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.scan_hostname", "scanHostname"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.scan_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsAnalyticsInstancePrivateAccessChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", acctest.Optional, acctest.Update, analyticsInstancePrivateAccessChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel2"),
				resource.TestCheckResourceAttr(resourceName, "egress_source_ip_addresses.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.description", "Tenant VCN DNS Zone 2"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest2.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.description", "Example OCI SCAN label 2"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.scan_hostname", "scanHostname2"),
				resource.TestCheckResourceAttr(resourceName, "private_source_scan_hosts.0.scan_port", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		{
			Config:                  config + AnalyticsAnalyticsInstancePrivateAccessChannelRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_analytics_analytics_instance_private_access_channel" {
			noResourceFound = false
			request := oci_analytics.GetPrivateAccessChannelRequest{}

			if value, ok := rs.Primary.Attributes["analytics_instance_id"]; ok {
				request.AnalyticsInstanceId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.PrivateAccessChannelKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")

			_, err := client.GetPrivateAccessChannel(context.Background(), request)

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
