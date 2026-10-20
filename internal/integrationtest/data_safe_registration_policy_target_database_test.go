// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeRegistrationPolicyTargetDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"registration_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_registration_policy.test_registration_policy.id}`},
	}

	DataSafeRegistrationPolicyTargetDatabaseResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_registration_policy", "test_registration_policy", acctest.Required, acctest.Create, DataSafeRegistrationPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeRegistrationPolicyTargetDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeRegistrationPolicyTargetDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	registrationPolicyResourceID := utils.GetEnvSettingWithBlankDefault("registration_policy_resource_id")
	registrationPolicyResourceIDVariableStr := fmt.Sprintf("variable \"registration_policy_resource_id\" { default = \"%s\" }\n", registrationPolicyResourceID)
	compartmentIdVariableStr += registrationPolicyResourceIDVariableStr

	registrationPolicyProvisioningWaitSeconds, err := strconv.Atoi(utils.GetEnvSettingWithDefault("registration_policy_provisioning_wait_seconds", "120"))
	if err != nil || registrationPolicyProvisioningWaitSeconds < 0 {
		t.Fatalf("registration_policy_provisioning_wait_seconds must be a non-negative integer")
	}
	waitForRegistrationPolicyTargetDatabases := func() {
		if registrationPolicyProvisioningWaitSeconds > 0 {
			t.Logf("waiting %d seconds for registration policy target databases to provision", registrationPolicyProvisioningWaitSeconds)
			time.Sleep(time.Duration(registrationPolicyProvisioningWaitSeconds) * time.Second)
		}
	}

	datasourceName := "data.oci_data_safe_registration_policy_target_databases.test_registration_policy_target_databases"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create the registration policy. Data Safe creates its target databases asynchronously.
		{
			Config: config + compartmentIdVariableStr + DataSafeRegistrationPolicyTargetDatabaseResourceConfig,
			Check:  resource.TestCheckResourceAttrSet("oci_data_safe_registration_policy.test_registration_policy", "id"),
		},
		// Query all target databases created by the registration policy.
		{
			PreConfig: waitForRegistrationPolicyTargetDatabases,
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_registration_policy_target_databases", "test_registration_policy_target_databases", acctest.Required, acctest.Create, DataSafeRegistrationPolicyTargetDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeRegistrationPolicyTargetDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "registration_policy_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "registration_policy_target_database_summary_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "registration_policy_target_database_summary_collection.0.items.#", "5"),
				resource.TestCheckResourceAttrSet(datasourceName, "registration_policy_target_database_summary_collection.0.items.0.target_database_id"),
			),
		},
		// Allow the policy-created target databases to finish provisioning before deleting the policy.
		{
			PreConfig: waitForRegistrationPolicyTargetDatabases,
			Config:    config + compartmentIdVariableStr,
		},
	})
}
