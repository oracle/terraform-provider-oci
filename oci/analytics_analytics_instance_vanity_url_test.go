// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"

	oci_analytics "github.com/oracle/oci-go-sdk/v46/analytics"
	"github.com/oracle/oci-go-sdk/v46/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ca_cert1              = getEnvSettingWithBlankDefault("ca_cert1")
	private_key_1_no_pass = getEnvSettingWithBlankDefault("private_key_1_no_pass")
	public_cert_1_no_pass = getEnvSettingWithBlankDefault("public_cert_1_no_pass")

	ca_cert2           = getEnvSettingWithBlankDefault("ca_cert2")
	private_key_2_pass = getEnvSettingWithBlankDefault("private_key_2_pass")
	public_cert_2_pass = getEnvSettingWithBlankDefault("public_cert_2_pass")
	passphrase_2       = getEnvSettingWithBlankDefault("passphrase_2")

	ca_cert1_val              = getEnvWithNewlineExpansion("ca_cert1")
	private_key_1_no_pass_val = getEnvWithNewlineExpansion("private_key_1_no_pass")
	public_cert_1_no_pass_val = getEnvWithNewlineExpansion("public_cert_1_no_pass")

	ca_cert2_val           = getEnvWithNewlineExpansion("ca_cert2")
	private_key_2_pass_val = getEnvWithNewlineExpansion("private_key_2_pass")
	public_cert_2_pass_val = getEnvWithNewlineExpansion("public_cert_2_pass")
	passphrase_2_val       = getEnvWithNewlineExpansion("passphrase_2")

	AnalyticsInstanceVanityUrlRequiredOnlyResource = AnalyticsInstanceVanityUrlResourceDependencies +
		generateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", Required, Create, analyticsInstanceVanityUrlRepresentation1)

	analyticsInstanceVanityUrlRepresentation1 = map[string]interface{}{
		"analytics_instance_id": Representation{repType: Required, create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},

		"ca_certificate":     Representation{repType: Required, create: ca_cert1, update: ca_cert2},
		"hosts":              Representation{repType: Required, create: []string{`test1.robcorobotics.com`}},
		"private_key":        Representation{repType: Required, create: private_key_1_no_pass, update: private_key_2_pass},
		"public_certificate": Representation{repType: Required, create: public_cert_1_no_pass, update: public_cert_2_pass},
		"description":        Representation{repType: Optional, create: `description`},
		"passphrase":         Representation{repType: Optional, create: ``, update: passphrase_2},
	}

	analyticsInstanceVanityUrlRepresentation2 = map[string]interface{}{
		"analytics_instance_id": Representation{repType: Required, create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"ca_certificate":        Representation{repType: Required, create: ca_cert2, update: ca_cert1},
		"hosts":                 Representation{repType: Required, create: []string{`test1.robcorobotics.com`}},
		"private_key":           Representation{repType: Required, create: private_key_2_pass, update: private_key_1_no_pass},
		"public_certificate":    Representation{repType: Required, create: public_cert_2_pass, update: public_cert_1_no_pass},
		"description":           Representation{repType: Optional, create: `description`},
		"passphrase":            Representation{repType: Optional, create: passphrase_2, update: ``},
	}

	AnalyticsInstanceVanityUrlResourceDependencies = generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Required, Create, analyticsInstanceRepresentation)
)

func getEnvWithNewlineExpansion(env_variable string) string {
	return strings.ReplaceAll(getEnvSettingWithBlankDefault(env_variable), "\\n", "\n")
}

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstanceVanityUrlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstanceVanityUrlResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	// Read certificates, etc. from environment variables

	resourceName := "oci_analytics_analytics_instance_vanity_url.test_analytics_instance_vanity_url"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+AnalyticsInstanceVanityUrlResourceDependencies+
		generateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", Optional, Create, analyticsInstanceVanityUrlRepresentation2), "analytics", "analyticsInstanceVanityUrl", t)

	ResourceTest(t, testAccCheckAnalyticsAnalyticsInstanceVanityUrlDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				generateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", Required, Create, analyticsInstanceVanityUrlRepresentation1),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert1_val),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_1_no_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_1_no_pass_val),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				generateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", Optional, Create, analyticsInstanceVanityUrlRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert2_val),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", passphrase_2_val),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_2_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_2_pass_val),

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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				generateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", Optional, Update, analyticsInstanceVanityUrlRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert1_val),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", ``),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_1_no_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_1_no_pass_val),

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
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ca_certificate",
				"private_key",
				"passphrase",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckAnalyticsAnalyticsInstanceVanityUrlDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).analyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_analytics_analytics_instance_vanity_url" {
			noResourceFound = false
			request := oci_analytics.GetAnalyticsInstanceRequest{}

			compositeId := rs.Primary.ID
			parts := strings.Split(compositeId, "/")
			match, _ := regexp.MatchString("analyticsInstances/.*/vanityUrls/.*", compositeId)
			if !match || len(parts) != 4 {
				return fmt.Errorf("illegal compositeId %s encountered", compositeId)
			}
			analyticsInstanceId, _ := url.PathUnescape(parts[1])
			vanityUrlKey, _ := url.PathUnescape(parts[3])

			request.AnalyticsInstanceId = &analyticsInstanceId

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "analytics")

			response, err := client.GetAnalyticsInstance(context.Background(), request)

			if err == nil {
				// Check that the vanityUrlKey is not present in the instance vanity url details
				vanityUrlDetails := response.VanityUrlDetails
				_, ok := vanityUrlDetails[vanityUrlKey]
				if ok {
					return fmt.Errorf("vanity url %v was found in analytics instance %v when it was expected to be deleted", vanityUrlKey, analyticsInstanceId)
				}
				continue
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
