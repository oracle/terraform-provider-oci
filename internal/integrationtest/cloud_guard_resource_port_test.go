package integrationtest

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

/*
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

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
	CloudGuardResourcePortDataSourceRepresentation = map[string]interface{}{
		"resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"open_port":   acctest.Representation{RepType: acctest.Optional, Create: `openPort`},
	}

	CloudGuardResourcePortResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardResourcePortResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardResourcePortResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_guard_resource_ports.test_resource_ports"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resource_ports", "test_resource_ports", acctest.Required, acctest.Create, CloudGuardResourcePortDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardResourcePortResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "open_port", "openPort"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_port_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "resource_port_collection.0.items.#", "1"),
			),
		},
	})
}*/
