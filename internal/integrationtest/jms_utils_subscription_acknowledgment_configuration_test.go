// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsUtilsSubscriptionAcknowledgmentConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsTenancyId},
	}

	JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: JmsTenancyId},
		"is_acknowledged": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
)

// issue-routing-tag: jms_utils/default
func TestJmsUtilsSubscriptionAcknowledgmentConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsSubscriptionAcknowledgmentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceName := "oci_jms_utils_subscription_acknowledgment_configuration.test_subscription_acknowledgment_configuration"

	singularDatasourceName := "data.oci_jms_utils_subscription_acknowledgment_configuration.test_subscription_acknowledgment_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_jms_utils_subscription_acknowledgment_configuration",
			"test_subscription_acknowledgment_configuration",
			acctest.Optional,
			acctest.Create,
			JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
		),
		"jmsutils",
		"subscriptionAcknowledgmentConfiguration",
		t,
	)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Required,
					acctest.Create,
					JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},

		// verify Create with optionals
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Optional,
					acctest.Create,
					JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", JmsTenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_acknowledged", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &JmsTenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Optional,
					acctest.Update,
					JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", JmsTenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_acknowledged", "true"),

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
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Optional,
					acctest.Update,
					JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Required,
					acctest.Create,
					JmsUtilsSubscriptionAcknowledgmentConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", JmsTenancyId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "acknowledged_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_acknowledged", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_acknowledged"),
			),
		},

		// verify resource import
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Required,
					acctest.Create,
					JmsUtilsSubscriptionAcknowledgmentConfigurationRepresentation,
				),
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"compartment_id",
			},
			ResourceName: resourceName,
		},
	})
}
