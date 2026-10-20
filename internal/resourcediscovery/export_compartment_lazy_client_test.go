// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package resourcediscovery

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
	tfclient "github.com/oracle/terraform-provider-oci/internal/client"
	tfexport "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func TestIdentityOperationsReturnLazyClientInitializationErrors(t *testing.T) {
	clients := &tfclient.OracleClients{SdkClientMap: make(map[string]interface{})}

	if _, err := listCompartments(clients, oci_identity.ListCompartmentsRequest{}); err == nil {
		t.Fatal("listCompartments did not return the lazy client initialization error")
	}
	if _, err := getCompartment(clients, oci_identity.GetCompartmentRequest{}); err == nil {
		t.Fatal("getCompartment did not return the lazy client initialization error")
	}
	if _, err := getTenancy(clients, oci_identity.GetTenancyRequest{}); err == nil {
		t.Fatal("getTenancy did not return the lazy client initialization error")
	}
}

func TestRefreshProviderSchemaMapsRebuildsProviderSchemasAndPreservesDiscoveryEntries(t *testing.T) {
	resources := tfexport.ResourcesMap
	datasources := tfexport.DatasourcesMap
	t.Cleanup(func() {
		tfexport.ResourcesMap = resources
		tfexport.DatasourcesMap = datasources
	})

	discoveryResource := &schema.Resource{}
	discoveryDatasource := &schema.Resource{}
	tfexport.ResourcesMap = map[string]*schema.Resource{"oci_test_discovery_resource": discoveryResource}
	tfexport.DatasourcesMap = map[string]*schema.Resource{"oci_test_discovery_datasource": discoveryDatasource}

	refreshProviderSchemaMaps()

	firstProviderResource := tfexport.ResourcesMap["oci_identity_tag_namespace"]
	if firstProviderResource == nil {
		t.Fatal("first refresh did not include a registered provider resource schema")
	}
	firstProviderDatasource := tfexport.DatasourcesMap["oci_identity_regions"]
	if firstProviderDatasource == nil {
		t.Fatal("first refresh did not include a registered provider data-source schema")
	}

	refreshProviderSchemaMaps()

	secondProviderResource := tfexport.ResourcesMap["oci_identity_tag_namespace"]
	if secondProviderResource == nil {
		t.Fatal("second refresh did not include a registered provider resource schema")
	}
	if secondProviderResource == firstProviderResource {
		t.Fatal("second refresh retained the previous provider resource schema")
	}
	secondProviderDatasource := tfexport.DatasourcesMap["oci_identity_regions"]
	if secondProviderDatasource == nil {
		t.Fatal("second refresh did not include a registered provider data-source schema")
	}
	if secondProviderDatasource == firstProviderDatasource {
		t.Fatal("second refresh retained the previous provider data-source schema")
	}
	if tfexport.ResourcesMap["oci_test_discovery_resource"] != discoveryResource {
		t.Fatal("refresh replaced a Resource Discovery-only resource schema")
	}
	if tfexport.DatasourcesMap["oci_test_discovery_datasource"] != discoveryDatasource {
		t.Fatal("refresh replaced a Resource Discovery-only data-source schema")
	}
}
