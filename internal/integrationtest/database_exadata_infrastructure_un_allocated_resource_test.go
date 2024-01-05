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
	DatabaseDatabaseExadataInfrastructureUnAllocatedResourceSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"db_servers":                acctest.Representation{RepType: acctest.Optional, Create: []string{`dbServers`}},
	}

	DatabaseExadataInfrastructureUnAllocatedResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureUnAllocatedResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureUnAllocatedResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_exadata_infrastructure_un_allocated_resource.test_exadata_infrastructure_un_allocated_resource"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure_un_allocated_resource", "test_exadata_infrastructure_un_allocated_resource", acctest.Required, acctest.Create, DatabaseDatabaseExadataInfrastructureUnAllocatedResourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadataInfrastructureUnAllocatedResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_vm_clusters.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_storage_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "local_storage_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocpus"),
			),
		},
	})
}
