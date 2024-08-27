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
	DelegateAccessControlDelegationControlResourceDataSourceRepresentation = map[string]interface{}{
		"delegation_control_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_delegate_access_control_delegation_control.test_delegation_control.id}`},
	}

	DelegateAccessControlDelegationControlResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlRepresentation)
	//acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegationControlResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegationControlResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_delegation_control_resources.test_delegation_control_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegation_control_resources", "test_delegation_control_resources", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlResourceDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegationControlResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "delegation_control_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "delegation_control_resource_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "delegation_control_resource_collection.0.items.#", "1"),
			),
		},
	})
}
