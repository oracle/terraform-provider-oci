// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberDataSourceRepresentation = map[string]interface{}{
		"target_alert_policy_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id}`},
	}

	DataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Create, targetAlertPolicyAssociationTargetGroupRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetDatabaseGroupCompartmentId := utils.GetEnvSettingWithDefault("target_database_group_compartment_ocid", compartmentId)
	targetDatabaseGroupCompartmentIdVariableStr := fmt.Sprintf("variable \"target_database_group_compartment_id\" { default = \"%s\" }\n", targetDatabaseGroupCompartmentId)

	targetDatabaseGroupId := utils.GetEnvSettingWithBlankDefault("data_safe_target_database_group_ocid")
	targetDatabaseGroupIdVariableStr := fmt.Sprintf("variable \"target_database_group_id\" { default = \"%s\" }\n", targetDatabaseGroupId)

	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	datasourceName := "data.oci_data_safe_target_alert_policy_association_unassociated_target_members.test_target_alert_policy_association_unassociated_target_members"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_alert_policy_association_unassociated_target_members", "test_target_alert_policy_association_unassociated_target_members", acctest.Required, acctest.Create, DataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberDataSourceRepresentation) +
				compartmentIdVariableStr + targetDatabaseGroupCompartmentIdVariableStr + targetDatabaseGroupIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationUnassociatedTargetMemberResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "target_alert_policy_association_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "target_alert_policy_unassociated_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_alert_policy_unassociated_collection.0.items.#"),
			),
		},
	})
}
