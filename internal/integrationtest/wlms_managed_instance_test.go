// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	wlmsCompartmentOcid     = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	wlmsInstanceDisplayName = utils.GetEnvSettingWithBlankDefault("wlms_instance_display_name")

	WlmsManagedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: wlmsManagedInstanceOcid},
	}

	WlmsManagedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: wlmsCompartmentOcid},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: wlmsInstanceDisplayName},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: wlmsManagedInstanceOcid},
		"plugin_status":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_wlms_managed_instance.test_managed_instance"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instances", "test_managed_instances", acctest.Optional, acctest.Create, WlmsManagedInstanceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, WlmsManagedInstanceSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_arch"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plugin_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
