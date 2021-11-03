// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package testing

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v49/analytics"
	"github.com/oracle/oci-go-sdk/v49/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	tf_common "github.com/terraform-providers/terraform-provider-oci/oci"
	tf_client "github.com/terraform-providers/terraform-provider-oci/oci/client"
)

var (
	AnalyticsInstancePrivateAccessChannelResourceConfig = AnalyticsInstancePrivateAccessChannelResourceDependencies +
		tf_common.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", tf_common.Optional, tf_common.Update, analyticsInstancePrivateAccessChannelRepresentation)

	analyticsInstancePrivateAccessChannelRepresentation = map[string]interface{}{
		"analytics_instance_id":    tf_common.Representation{RepType: tf_common.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"display_name":             tf_common.Representation{RepType: tf_common.Required, Create: `example_private_access_channel`, Update: `example_private_access_channel2`},
		"private_source_dns_zones": tf_common.RepresentationGroup{tf_common.Required, analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation},
		"subnet_id":                tf_common.Representation{RepType: tf_common.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                   tf_common.Representation{RepType: tf_common.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}
	analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation = map[string]interface{}{
		"dns_zone":    tf_common.Representation{RepType: tf_common.Required, Create: `terraformtest.oraclevcn.com`, Update: `terraformtest2.oraclevcn.com`},
		"description": tf_common.Representation{RepType: tf_common.Optional, Create: `Tenant VCN DNS Zone`, Update: `Tenant VCN DNS Zone 2`},
	}

	AnalyticsInstancePrivateAccessChannelResourceDependencies = tf_common.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", tf_common.Required, tf_common.Create, analyticsInstanceRepresentation) +
		tf_common.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", tf_common.Required, tf_common.Create, tf_common.SubnetRepresentation) +
		tf_common.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", tf_common.Required, tf_common.Create, tf_common.VcnRepresentation)
)

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic")
	defer httpreplay.SaveScenario()

	config := tf_common.ProviderTestConfig()

	compartmentId := tf_common.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := tf_common.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_analytics_analytics_instance_private_access_channel.test_analytics_instance_private_access_channel"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	tf_common.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsInstancePrivateAccessChannelResourceDependencies+
		tf_common.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", tf_common.Required, tf_common.Create, analyticsInstancePrivateAccessChannelRepresentation), "analytics", "analyticsInstancePrivateAccessChannel", t)

	tf_common.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstancePrivateAccessChannelResourceDependencies +
				tf_common.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", tf_common.Required, tf_common.Create, analyticsInstancePrivateAccessChannelRepresentation),
			Check: tf_common.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = tf_common.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(tf_common.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := tf_common.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstancePrivateAccessChannelResourceDependencies +
				tf_common.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", tf_common.Optional, tf_common.Update, analyticsInstancePrivateAccessChannelRepresentation),
			Check: tf_common.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel2"),
				resource.TestCheckResourceAttr(resourceName, "egress_source_ip_addresses.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.description", "Tenant VCN DNS Zone 2"),
				resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest2.oraclevcn.com"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = tf_common.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := tf_common.TestAccProvider.Meta().(*tf_client.OracleAnalyticsClients).AnalyticsClient()
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

			request.RequestMetadata.RetryPolicy = tf_common.GetRetryPolicy(true, "analytics")

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
