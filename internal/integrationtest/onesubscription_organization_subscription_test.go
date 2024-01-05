// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OnesubscriptionOnesubscriptionOrganizationSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionOrganizationSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionOrganizationSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_onesubscription_organization_subscriptions.test_organization_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_organization_subscriptions", "test_organization_subscriptions", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionOrganizationSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.#"),
				resource.TestCheckResourceAttr(datasourceName, "organization_subscriptions.0.currency.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.service_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.time_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.total_value"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_subscriptions.0.type"),
			),
		},
	})
}
