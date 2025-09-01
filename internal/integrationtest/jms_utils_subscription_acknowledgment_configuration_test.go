// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsUtilsSubscriptionAcknowledgmentConfigurationCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	JmsUtilsSubscriptionAcknowledgmentConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: JmsUtilsSubscriptionAcknowledgmentConfigurationCompartmentId},
	}
)

// issue-routing-tag: jms_utils/default
func TestJmsUtilsSubscriptionAcknowledgmentConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsSubscriptionAcknowledgmentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_jms_utils_subscription_acknowledgment_configuration.test_subscription_acknowledgment_configuration"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		// note: we cannot write test for this case because
		// we don't have create API.

		// verify update
		// note: we cannot write test for this case because
		// we don't have update API.

		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_subscription_acknowledgment_configuration",
					"test_subscription_acknowledgment_configuration",
					acctest.Optional,
					acctest.Create,
					JmsUtilsSubscriptionAcknowledgmentConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				// check actual data matches data used for the GET API
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", JmsUtilsSubscriptionAcknowledgmentConfigurationCompartmentId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_acknowledged"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_acknowledged"),
			),
		},
	})
}
