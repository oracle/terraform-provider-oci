// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NosqlTableReplicaRequiredOnlyResource = NosqlTableReplicaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table_replica", "test_table_replica", acctest.Required, acctest.Create, NosqlTableReplicaRepresentation)

	NosqlTableReplicaRepresentation = map[string]interface{}{
		"region":           acctest.Representation{RepType: acctest.Required, Create: `ca-montreal-1`},
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_mr_table.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"max_read_units":   acctest.Representation{RepType: acctest.Optional, Create: `11`},
		"max_write_units":  acctest.Representation{RepType: acctest.Optional, Create: `11`},
	}

	mrtable_ddl                 = "CREATE TABLE IF NOT EXISTS test_mr_table(id integer, info JSON, PRIMARY KEY(id)) WITH SCHEMA FROZEN"
	mrTableLimitsRepresentation = map[string]interface{}{
		"max_read_units":     acctest.Representation{RepType: acctest.Required, Create: `10`},
		"max_write_units":    acctest.Representation{RepType: acctest.Required, Create: `10`},
		"max_storage_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	mrTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ddl_statement":  acctest.Representation{RepType: acctest.Required, Create: mrtable_ddl},
		"name":           acctest.Representation{RepType: acctest.Required, Create: "test_mr_table"},
		"table_limits":   acctest.RepresentationGroup{RepType: acctest.Required, Group: mrTableLimitsRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreTableDefinedTags},
	}
	NosqlTableReplicaResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_mr_table", acctest.Required, acctest.Create, mrTableRepresentation)
)

// issue-routing-tag: nosql/default
func TestNosqlTableReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_nosql_table_replica.test_table_replica"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NosqlTableReplicaResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table_replica", "test_table_replica", acctest.Optional, acctest.Create, NosqlTableReplicaRepresentation), "nosql", "tableReplica", t)

	acctest.ResourceTest(t, testAccCheckNosqlTableReplicaDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NosqlTableReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table_replica", "test_table_replica", acctest.Required, acctest.Create, NosqlTableReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "region", "ca-montreal-1"),
				resource.TestCheckResourceAttrSet(resourceName, "table_name_or_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NosqlTableReplicaResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NosqlTableReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_nosql_table_replica", "test_table_replica", acctest.Optional, acctest.Create, NosqlTableReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "max_read_units", "11"),
				resource.TestCheckResourceAttr(resourceName, "max_write_units", "11"),
				resource.TestCheckResourceAttr(resourceName, "region", "ca-montreal-1"),
				resource.TestCheckResourceAttrSet(resourceName, "table_name_or_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}

func testAccCheckNosqlTableReplicaDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NosqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_nosql_table_replica" {

			noResourceFound = false

			request := oci_nosql.GetTableRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}
			if value, ok := rs.Primary.Attributes["table_name_or_id"]; ok {
				request.TableNameOrId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "nosql")
			response, err := client.GetTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_nosql.TableLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}

	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}
	return nil
}
