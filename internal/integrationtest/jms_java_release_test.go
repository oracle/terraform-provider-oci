// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsJavaReleaseSingularDataSourceRepresentation = map[string]interface{}{
		"release_version": acctest.Representation{RepType: acctest.Required, Create: `17.0.2`},
	}

	JmsJavaReleaseDataSourceRepresentation = map[string]interface{}{
		"family_version":      acctest.Representation{RepType: acctest.Optional, Create: `11`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `UPDATE_REQUIRED`},
		"license_type":        acctest.Representation{RepType: acctest.Optional, Create: `OTN`},
		"release_type":        acctest.Representation{RepType: acctest.Optional, Create: `CPU`},
		"release_version":     acctest.Representation{RepType: acctest.Optional, Create: `11.0.11`},
	}
)

// issue-routing-tag: jms/default
func TestJmsJavaReleaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaReleaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_java_releases.test_java_releases"
	singularDatasourceName := "data.oci_jms_java_release.test_java_release"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_releases",
					"test_java_releases",
					acctest.Optional,
					acctest.Create,
					JmsJavaReleaseDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "family_version", "11"),
				resource.TestCheckResourceAttr(datasourceName, "jre_security_status", "UPDATE_REQUIRED"),
				resource.TestCheckResourceAttr(datasourceName, "license_type", "OTN"),
				resource.TestCheckResourceAttr(datasourceName, "release_type", "CPU"),
				resource.TestCheckResourceAttr(datasourceName, "release_version", "11.0.11"),

				resource.TestCheckResourceAttrSet(datasourceName, "java_release_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_release",
					"test_java_release",
					acctest.Required,
					acctest.Create,
					JmsJavaReleaseSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_version"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_content_types.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifacts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "days_under_security_baseline"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_version", "17"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mos_patches.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_notes_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "release_version", "17.0.2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_status"),
			),
		},
	})
}
