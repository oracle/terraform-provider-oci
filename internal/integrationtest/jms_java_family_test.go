// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsJavaFamilySingularDataSourceRepresentation = map[string]interface{}{
		"family_version": acctest.Representation{RepType: acctest.Required, Create: `8`},
	}

	JmsJavaFamilyDataSourceRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `JDK 11`},
		"family_version": acctest.Representation{RepType: acctest.Optional, Create: `11`},
	}
)

// issue-routing-tag: jms/default
func TestJmsJavaFamilyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaFamilyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_java_families.test_java_families"
	singularDatasourceName := "data.oci_jms_java_family.test_java_family"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_families",
					"test_java_families",
					acctest.Optional,
					acctest.Create,
					JmsJavaFamilyDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "JDK 11"),
				resource.TestCheckResourceAttr(datasourceName, "family_version", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_family_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_family",
					"test_java_family",
					acctest.Required,
					acctest.Create,
					JmsJavaFamilySingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "family_version", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "JDK 8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "doc_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "end_of_support_life_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_supported_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "latest_release_artifacts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "latest_release_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "support_type"),
			),
		},
	})
}
