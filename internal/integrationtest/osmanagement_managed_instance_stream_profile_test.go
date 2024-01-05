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
	OsmanagementOsmanagementManagedInstanceStreamProfileDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: managedInstanceOCID},
	}
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceStreamProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceStreamProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_osmanagement_managed_instance_stream_profiles.test_managed_instance_stream_profiles"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_stream_profiles", "test_managed_instance_stream_profiles", acctest.Required, acctest.Create, OsmanagementOsmanagementManagedInstanceStreamProfileDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_on_managed_instances.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_on_managed_instances.0.module_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_on_managed_instances.0.profile_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_on_managed_instances.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_on_managed_instances.0.stream_name"),
			),
		},
	})
}
