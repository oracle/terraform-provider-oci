// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	PsqlDbSystemConnectionDetailSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_ocid}`},
	}

	PsqlDbSystemConnectionDetailResourceConfig = ""
)

// issue-routing-tag: psql/default
func TestPsqlDbSystemConnectionDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDbSystemConnectionDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("db_ocid")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"db_ocid\" { default = \"%s\" }\n", dbSystemId)

	singularDatasourceName := "data.oci_psql_db_system_connection_detail.test_db_system_connection_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + dbSystemIdVariableStr + subnetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_system_connection_detail", "test_db_system_connection_detail", acctest.Required, acctest.Create, PsqlDbSystemConnectionDetailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlDbSystemConnectionDetailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "ca_certificate"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "primary_db_endpoint.#", "1"),
			),
		},
	})
}
