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
	BdsBdsInstanceSoftwareUpdateSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
		"software_update_key": acctest.Representation{RepType: acctest.Required, Create: `${var.software_update_key}`},
	}

	BdsBdsInstanceSoftwareUpdateDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`}}

	BdsBdsInstanceSoftwareUpdateActionRepresentation = map[string]interface{}{
		"bds_instance_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
		"software_update_keys": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.software_update_keys}`}},
	}
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceSoftwareUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceSoftwareUpdateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	bdsInstanceId := utils.GetEnvSettingWithBlankDefault("bds_instance_ocid")
	bdsInstanceIdVariableStr := fmt.Sprintf("variable \"bds_instance_id\" { default = \"%s\" }\n", bdsInstanceId)

	softwareUpdateKeys := utils.GetEnvSettingWithBlankDefault("software_update_keys")
	softwareUpdateKeysVariableStr := fmt.Sprintf("variable \"software_update_keys\" { default = \"%s\" }\n", softwareUpdateKeys)

	softwareUpdateKey := utils.GetEnvSettingWithBlankDefault("software_update_key")
	softwareUpdateKeyVariableStr := fmt.Sprintf("variable \"software_update_key\" { default = \"%s\" }\n", softwareUpdateKey)

	datasourceName := "data.oci_bds_bds_instance_software_updates.test_bds_instance_software_updates"
	singularDatasourceName := "data.oci_bds_bds_instance_software_update.test_bds_instance_software_update"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//Verify Create
		{
			Config: config + compartmentIdVariableStr + bdsInstanceIdVariableStr + softwareUpdateKeysVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_software_update_action", "test_bds_instance_software_update_action", acctest.Required, acctest.Create, BdsBdsInstanceSoftwareUpdateActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_update_keys"),
			),
		},

		// verify datasource
		{
			Config: config + compartmentIdVariableStr + bdsInstanceIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_software_updates", "test_bds_instance_software_updates", acctest.Required, acctest.Create, BdsBdsInstanceSoftwareUpdateDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_update_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + bdsInstanceIdVariableStr + softwareUpdateKeyVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_software_update", "test_bds_instance_software_update", acctest.Required, acctest.Create, BdsBdsInstanceSoftwareUpdateSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_update_key"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_update_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_update_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_update_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_due"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
			),
		},
	})
}
