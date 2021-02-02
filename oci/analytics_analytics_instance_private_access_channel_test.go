// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v35/analytics"
	"github.com/oracle/oci-go-sdk/v35/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsInstancePrivateAccessChannelResourceConfig = AnalyticsInstancePrivateAccessChannelResourceDependencies +
		generateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", Optional, Update, analyticsInstancePrivateAccessChannelRepresentation)

	analyticsInstancePrivateAccessChannelRepresentation = map[string]interface{}{
		"analytics_instance_id":    Representation{repType: Required, create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"display_name":             Representation{repType: Required, create: `example_private_access_channel`, update: `example_private_access_channel2`},
		"private_source_dns_zones": RepresentationGroup{Required, analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation},
		"subnet_id":                Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":                   Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
	}
	analyticsInstancePrivateAccessChannelPrivateSourceDnsZonesRepresentation = map[string]interface{}{
		"dns_zone":    Representation{repType: Required, create: `terraformtest.oraclevcn.com`, update: `terraformtest2.oraclevcn.com`},
		"description": Representation{repType: Optional, create: `Tenant VCN DNS Zone`, update: `Tenant VCN DNS Zone 2`},
	}

	AnalyticsInstancePrivateAccessChannelResourceDependencies = generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Required, Create, analyticsInstanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)
)

func TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstancePrivateAccessChannelResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_analytics_analytics_instance_private_access_channel.test_analytics_instance_private_access_channel"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstancePrivateAccessChannelResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", Required, Create, analyticsInstancePrivateAccessChannelRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_private_access_channel"),
					resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_source_dns_zones.0.dns_zone", "terraformtest.oraclevcn.com"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstancePrivateAccessChannelResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance_private_access_channel", "test_analytics_instance_private_access_channel", Optional, Update, analyticsInstancePrivateAccessChannelRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
		},
	})
}

func testAccCheckAnalyticsAnalyticsInstancePrivateAccessChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).analyticsClient()
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "analytics")

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
