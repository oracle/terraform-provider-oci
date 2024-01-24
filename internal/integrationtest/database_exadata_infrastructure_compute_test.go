// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	exadataInfrastructureComputeSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	exadataInfrastructureComputeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tstExaInfra`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `REQUIRES_ACTIVATION`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: exadataInfrastructureComputeDataSourceFilterRepresentation}}
	exadataInfrastructureComputeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`}},
	}

	ExadataInfrastructureRepresentationCompute = map[string]interface{}{
		"exadata_infrastructure_id":                                acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`, Update: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"additional_compute_count_compute_managed_resource":        acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`},
		"additional_compute_system_model_compute_managed_resource": acctest.Representation{RepType: acctest.Optional, Create: `X8M`, Update: `X9M`},
	}

	ExadataInfrastructureComputeManagedResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureComputeManagedResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureComputeManagedResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//verify Create
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureComputeManagedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation),
		},

		//verify Update/scale up with optionals
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureComputeManagedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExadataInfrastructureRepresentation, map[string]interface{}{
					"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_compute", "test_database_exadata_infrastructure_compute", acctest.Optional, acctest.Create,
					ExadataInfrastructureRepresentationCompute),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "state", "REQUIRES_ACTIVATION"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X7"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_system_model", ""),

				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "storage_count", "3"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "compute_count", "4"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_count", "0"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_system_model", ""),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureComputeManagedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExadataInfrastructureRepresentation, map[string]interface{}{
					"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_compute", "test_database_exadata_infrastructure_compute", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ExadataInfrastructureRepresentationCompute, map[string]interface{}{
						"activation_file": acctest.Representation{RepType: acctest.Optional, Create: activationFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X7"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_system_model", ""),

				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "state", "ACTIVE"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "storage_count", "3"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "compute_count", "4"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_count", "0"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_system_model", ""),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify update/scale up after activation
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureComputeManagedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExadataInfrastructureRepresentation, map[string]interface{}{
					"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_compute", "test_database_exadata_infrastructure_compute", acctest.Optional, acctest.Update, ExadataInfrastructureRepresentationCompute),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X7"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_system_model", "X9M"),

				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "state", "ACTIVE"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "storage_count", "3"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "compute_count", "4"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_count", "3"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_system_model", "X9M"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify compute count after activate
		{
			Config: config + compartmentIdVariableStr + ExadataInfrastructureComputeManagedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExadataInfrastructureRepresentation, map[string]interface{}{
					"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_compute", "test_database_exadata_infrastructure_compute", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(ExadataInfrastructureRepresentationCompute, map[string]interface{}{
						"activation_file": acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "storage_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "7"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "shape", "ExadataCC.X7"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "additional_compute_system_model", ""),

				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "state", "ACTIVE"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "storage_count", "3"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "compute_count", "7"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_count", "0"),
				resource.TestCheckResourceAttr("oci_database_exadata_infrastructure_compute.test_database_exadata_infrastructure_compute", "additional_compute_system_model", ""),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseExadataInfrastructureCompute") {
		resource.AddTestSweepers("DatabaseExadataInfrastructureCompute", &resource.Sweeper{
			Name:         "DatabaseExadataInfrastructureCompute",
			Dependencies: acctest.DependencyGraph["exadataInfrastructureCompute"],
			F:            sweepDatabaseExadataInfrastructureResource,
		})
	}
}
