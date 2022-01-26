// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_analytics "github.com/oracle/oci-go-sdk/v56/analytics"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ca_cert1              = utils.GetEnvSettingWithBlankDefault("ca_cert1")
	private_key_1_no_pass = utils.GetEnvSettingWithBlankDefault("private_key_1_no_pass")
	public_cert_1_no_pass = utils.GetEnvSettingWithBlankDefault("public_cert_1_no_pass")

	ca_cert2           = utils.GetEnvSettingWithBlankDefault("ca_cert2")
	private_key_2_pass = utils.GetEnvSettingWithBlankDefault("private_key_2_pass")
	public_cert_2_pass = utils.GetEnvSettingWithBlankDefault("public_cert_2_pass")
	passphrase_2       = utils.GetEnvSettingWithBlankDefault("passphrase_2")

	ca_cert1_val              = getEnvWithNewlineExpansion("ca_cert1")
	private_key_1_no_pass_val = getEnvWithNewlineExpansion("private_key_1_no_pass")
	public_cert_1_no_pass_val = getEnvWithNewlineExpansion("public_cert_1_no_pass")

	ca_cert2_val           = getEnvWithNewlineExpansion("ca_cert2")
	private_key_2_pass_val = getEnvWithNewlineExpansion("private_key_2_pass")
	public_cert_2_pass_val = getEnvWithNewlineExpansion("public_cert_2_pass")
	passphrase_2_val       = getEnvWithNewlineExpansion("passphrase_2")

	AnalyticsInstanceVanityUrlRequiredOnlyResource = AnalyticsInstanceVanityUrlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", acctest.Required, acctest.Create, analyticsInstanceVanityUrlRepresentation1)

	analyticsInstanceVanityUrlRepresentation1 = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},

		"ca_certificate":     acctest.Representation{RepType: acctest.Required, Create: ca_cert1, Update: ca_cert2},
		"hosts":              acctest.Representation{RepType: acctest.Required, Create: []string{`test1.robcorobotics.com`}},
		"private_key":        acctest.Representation{RepType: acctest.Required, Create: private_key_1_no_pass, Update: private_key_2_pass},
		"public_certificate": acctest.Representation{RepType: acctest.Required, Create: public_cert_1_no_pass, Update: public_cert_2_pass},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"passphrase":         acctest.Representation{RepType: acctest.Optional, Create: ``, Update: passphrase_2},
	}

	analyticsInstanceVanityUrlRepresentation2 = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"ca_certificate":        acctest.Representation{RepType: acctest.Required, Create: ca_cert2, Update: ca_cert1},
		"hosts":                 acctest.Representation{RepType: acctest.Required, Create: []string{`test1.robcorobotics.com`}},
		"private_key":           acctest.Representation{RepType: acctest.Required, Create: private_key_2_pass, Update: private_key_1_no_pass},
		"public_certificate":    acctest.Representation{RepType: acctest.Required, Create: public_cert_2_pass, Update: public_cert_1_no_pass},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"passphrase":            acctest.Representation{RepType: acctest.Optional, Create: passphrase_2, Update: ``},
	}

	AnalyticsInstanceVanityUrlResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", acctest.Required, acctest.Create, analyticsInstanceRepresentation)
)

func getEnvWithNewlineExpansion(env_variable string) string {
	return strings.ReplaceAll(utils.GetEnvSettingWithBlankDefault(env_variable), "\\n", "\n")
}

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstanceVanityUrlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstanceVanityUrlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	// Read certificates, etc. from environment variables

	resourceName := "oci_analytics_analytics_instance_vanity_url.test_analytics_instance_vanity_url"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsInstanceVanityUrlResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", acctest.Optional, acctest.Create, analyticsInstanceVanityUrlRepresentation2), "analytics", "analyticsInstanceVanityUrl", t)

	acctest.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstanceVanityUrlDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", acctest.Required, acctest.Create, analyticsInstanceVanityUrlRepresentation1),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert1_val),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_1_no_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_1_no_pass_val),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", acctest.Optional, acctest.Create, analyticsInstanceVanityUrlRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert2_val),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", passphrase_2_val),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_2_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_2_pass_val),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceVanityUrlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_vanity_url", "test_analytics_instance_vanity_url", acctest.Optional, acctest.Update, analyticsInstanceVanityUrlRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "ca_certificate", ca_cert1_val),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", ``),
				resource.TestCheckResourceAttr(resourceName, "private_key", private_key_1_no_pass_val),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", public_cert_1_no_pass_val),

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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnalyticsClient()
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

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")

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
