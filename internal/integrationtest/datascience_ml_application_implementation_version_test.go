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
	DatascienceMlApplicationImplementationVersionSingularDataSourceRepresentation = map[string]interface{}{
		"ml_application_implementation_version_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datascience_ml_application_implementation_versions.test_ml_application_implementation_versions.ml_application_implementation_version_collection[0].items[0].id}`},
	}

	DatascienceMlApplicationImplementationVersionDataSourceRepresentation = map[string]interface{}{
		"ml_application_implementation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application_implementation.test_ml_application_implementation.id}`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatascienceMlAppImplementationRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ml_application_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `TestResource`},
		"opc_ml_app_package_args":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bucket_namespace": "idtlxnfdweil"}},
		"ml_application_package":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"source_type": "object_storage_download", "uri": "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/windows.zip"}, Update: map[string]string{"source_type": "object_storage_download", "path": "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/ml-app-package-1.8.zip"}},
		"allowed_migration_destinations": acctest.Representation{RepType: acctest.Optional, Update: []string{`allowedMigrationDestinations2`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationImplementationDefinedTagsChangesRepresentation},
	}

	DatascienceMlAppImplementationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create, DatascienceMlAppImplementationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationRepresentation)
)

// issue-routing-tag: datascience
func TestDatascienceMlApplicationImplementationVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceMlApplicationImplementationVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_ml_application_implementation_versions.test_ml_application_implementation_versions"
	singularDatasourceName := "data.oci_datascience_ml_application_implementation_version.test_ml_application_implementation_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_implementation_versions", "test_ml_application_implementation_versions", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlAppImplementationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_version_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_version_collection.0.items.0.ml_application_name", "ml-app-name-test"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_version_collection.0.items.0.name", "TestResource"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_version_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ml_application_implementation_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_implementation_versions", "test_ml_application_implementation_versions", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlAppImplementationResourceConfig + acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_implementation_version", "test_ml_application_implementation_version", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationVersionSingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_schema.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ml_application_package_arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TestResource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ml_application_name", "ml-app-name-test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_version"),
			),
		},
	})

}
