// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	oci_object_storage "github.com/oracle/oci-go-sdk/v47/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	NamespaceResourceConfig = ""
)

// issue-routing-tag: object_storage/default
func TestObjectStorageNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_objectstorage_namespace.test_namespace"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Optional, Create, namespaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NamespaceResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
			),
		},
	})
}

func getNamespaces(compartment string) ([]string, error) {
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()

	getNamespacesRequest := oci_object_storage.GetNamespaceRequest{}
	getNamespacesResponse, err := objectStorageClient.GetNamespace(context.Background(), getNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Bucket NameSpace list for compartment id : %s , %s \n", compartmentId, err)
	}

	resourceIds = append(resourceIds, *getNamespacesResponse.Value)
	return resourceIds, nil
}
