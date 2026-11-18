// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// The focused cluster tests own their complete temporary network so they do not
// depend on shared VCN or NAT-gateway fixtures.
const databaseAcknowledgementNetwork = `
resource "oci_core_vcn" "ack_vcn" {
  compartment_id = var.compartment_id
  cidr_block = "10.252.0.0/16"
  display_name = "tf-ack-vcn"
  dns_label = "tfackvcn"
}
resource "oci_core_nat_gateway" "ack_nat" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.ack_vcn.id
  display_name = "tf-ack-nat"
}
resource "oci_core_route_table" "ack_routes" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.ack_vcn.id
  display_name = "tf-ack-routes"
  route_rules {
    destination = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    network_entity_id = oci_core_nat_gateway.ack_nat.id
  }
}
resource "oci_core_security_list" "ack_security" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.ack_vcn.id
  display_name = "tf-ack-security"
  ingress_security_rules {
    source = oci_core_vcn.ack_vcn.cidr_block
    protocol = "all"
  }
  egress_security_rules {
    destination = oci_core_vcn.ack_vcn.cidr_block
    protocol = "all"
  }
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "6"
    tcp_options {
      min = 443
      max = 443
    }
  }
}
resource "oci_core_subnet" "ack_client" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.ack_vcn.id
  route_table_id = oci_core_route_table.ack_routes.id
  security_list_ids = [oci_core_security_list.ack_security.id]
  cidr_block = cidrsubnet(oci_core_vcn.ack_vcn.cidr_block, 8, 40)
  display_name = "tf-ack-client"
  dns_label = "tfackclient"
  prohibit_public_ip_on_vnic = true
}
resource "oci_core_subnet" "ack_backup" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.ack_vcn.id
  route_table_id = oci_core_route_table.ack_routes.id
  security_list_ids = [oci_core_security_list.ack_security.id]
  cidr_block = cidrsubnet(oci_core_vcn.ack_vcn.cidr_block, 8, 41)
  display_name = "tf-ack-backup"
  dns_label = "tfackbackup"
  prohibit_public_ip_on_vnic = true
}
`

func databaseAcknowledgementResourceTest(t *testing.T, checkDestroy resource.TestCheckFunc, steps []resource.TestStep) {
	t.Helper()
	for i, step := range steps {
		if step.Config == "" {
			continue
		}
		file, diagnostics := hclsyntax.ParseConfig([]byte(step.Config), "acknowledgement_test.tf", hcl.InitialPos)
		if diagnostics.HasErrors() {
			t.Fatalf("step %d has invalid Terraform configuration: %s", i+1, diagnostics.Error())
		}
		for _, block := range file.Body.(*hclsyntax.Body).Blocks {
			if len(block.Labels) != 2 {
				continue
			}
			resourceType := block.Labels[0]
			if block.Type == "data" && resourceType == "oci_core_vcn" {
				t.Fatalf("step %d must create its test VCN rather than read an existing VCN", i+1)
			}
			if block.Type != "resource" {
				continue
			}
			switch resourceType {
			case "oci_core_default_route_table", "oci_core_default_security_list", "oci_core_default_dhcp_options":
				t.Fatalf("step %d must not manage a VCN default resource: %s", i+1, resourceType)
			}
		}
	}
	// Validate every step before the first apply, including import/read steps.
	acctest.ResourceTest(t, checkDestroy, steps)
}

var databaseAcknowledgementDiscoveryMutex sync.Mutex

// databaseAcknowledgementScopedDiscoveryCheck preserves the conventional
// enable_export_compartment switch while avoiding unrelated database exporter
// associations during a focused acknowledgement test run.
func databaseAcknowledgementScopedDiscoveryCheck(t *testing.T, resourceName string) resource.TestCheckFunc {
	t.Helper()
	return func(state *terraform.State) error {
		isEnabled, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true"))
		if !isEnabled {
			return nil
		}
		return databaseAcknowledgementDiscoveryCheck(t, resourceName)(state)
	}
}

// This exercises production discovery for the created resource: the original
// resource association, SDK list/read, HCL generation, import, and Terraform plan.
// The database graph is scoped to that association for this synchronous test so
// unrelated resources requiring parent IDs are not queried from compartment
// scope. The original association/hints are unchanged and restored on return.
func databaseAcknowledgementDiscoveryCheck(t *testing.T, resourceName string) resource.TestCheckFunc {
	t.Helper()
	return func(state *terraform.State) error {
		instance, ok := state.RootModule().Resources[resourceName]
		if !ok || instance.Primary == nil || instance.Primary.ID == "" {
			return fmt.Errorf("resource discovery requires an existing %s", resourceName)
		}
		resourceType := instance.Type
		resourceID := instance.Primary.ID
		compartmentID := instance.Primary.Attributes["compartment_id"]
		if compartmentID == "" {
			return fmt.Errorf("%s has no compartment_id for discovery", resourceName)
		}

		databaseAcknowledgementDiscoveryMutex.Lock()
		defer databaseAcknowledgementDiscoveryMutex.Unlock()
		originalGraph, ok := tf_export.CompartmentResourceGraphs["database"]
		if !ok {
			return fmt.Errorf("database resource discovery graph is not registered")
		}
		scopedGraph, err := databaseAcknowledgementDiscoveryGraph(originalGraph, resourceType)
		if err != nil {
			return err
		}
		tf_export.CompartmentResourceGraphs["database"] = scopedGraph
		defer func() { tf_export.CompartmentResourceGraphs["database"] = originalGraph }()

		t.Logf("Running standard resource discovery, state import, and plan for %s", resourceType)
		if err := resourcediscovery.TestExportCompartmentWithResourceName(&resourceID, &compartmentID, resourceType); err != nil {
			return fmt.Errorf("resource discovery for %s failed: %w", resourceType, err)
		}
		return nil
	}
}

// databaseAcknowledgementDiscoveryGraph keeps only the discovery path needed
// for the acknowledgement-test resource. Most Database resources are listed at
// compartment scope. DB Homes and databases are different: listing DB Homes
// requires their parent and listing databases requires their DB Home.
func databaseAcknowledgementDiscoveryGraph(originalGraph tf_export.TerraformResourceGraph, resourceType string) (tf_export.TerraformResourceGraph, error) {
	if resourceType == "oci_database_db_home" {
		const parentType = "oci_database_cloud_vm_cluster"
		var parentAssociations []tf_export.TerraformResourceAssociation
		for _, association := range originalGraph["oci_identity_compartment"] {
			if association.ResourceClass == parentType {
				parentAssociations = append(parentAssociations, association)
			}
		}
		var childAssociations []tf_export.TerraformResourceAssociation
		for _, association := range originalGraph[parentType] {
			if association.ResourceClass == resourceType {
				childAssociations = append(childAssociations, association)
			}
		}
		if len(parentAssociations) != 1 || len(childAssociations) != 1 {
			return nil, fmt.Errorf("expected one Cloud VM Cluster to DB Home discovery path, found %d parent and %d child associations", len(parentAssociations), len(childAssociations))
		}
		return tf_export.TerraformResourceGraph{
			"oci_identity_compartment": parentAssociations,
			parentType:                 childAssociations,
		}, nil
	}

	if resourceType == "oci_database_database" {
		const (
			parentType = "oci_database_exadb_vm_cluster"
			dbHomeType = "oci_database_db_home"
		)
		var parentAssociations []tf_export.TerraformResourceAssociation
		for _, association := range originalGraph["oci_identity_compartment"] {
			if association.ResourceClass == parentType {
				parentAssociations = append(parentAssociations, association)
			}
		}

		// The Database graph defines DB Home discovery below a Cloud VM Cluster.
		// Exadb VM Clusters use the same ListDbHomes vm_cluster_id query parameter,
		// so reuse that association for the disposable Exadb acknowledgement parent.
		var dbHomeAssociations []tf_export.TerraformResourceAssociation
		for _, association := range originalGraph["oci_database_cloud_vm_cluster"] {
			if association.ResourceClass == dbHomeType {
				dbHomeAssociations = append(dbHomeAssociations, association)
			}
		}
		var databaseAssociations []tf_export.TerraformResourceAssociation
		for _, association := range originalGraph[dbHomeType] {
			if association.ResourceClass == resourceType {
				databaseAssociations = append(databaseAssociations, association)
			}
		}
		if len(parentAssociations) != 1 || len(dbHomeAssociations) != 1 || len(databaseAssociations) != 1 {
			return nil, fmt.Errorf("expected one Exadb VM Cluster to DB Home to Database discovery path, found %d parent, %d DB Home, and %d Database associations", len(parentAssociations), len(dbHomeAssociations), len(databaseAssociations))
		}
		return tf_export.TerraformResourceGraph{
			"oci_identity_compartment": parentAssociations,
			parentType:                 dbHomeAssociations,
			dbHomeType:                 databaseAssociations,
		}, nil
	}

	var associations []tf_export.TerraformResourceAssociation
	for _, association := range originalGraph["oci_identity_compartment"] {
		if association.ResourceClass == resourceType {
			associations = append(associations, association)
		}
	}
	if len(associations) != 1 {
		return nil, fmt.Errorf("expected one original discovery association for %s, found %d", resourceType, len(associations))
	}
	return tf_export.TerraformResourceGraph{
		"oci_identity_compartment": associations,
	}, nil
}

func databaseAcknowledgementIDCheck(resourceName string, previousID *string, expectReplacement bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := acctest.FromInstanceState(state, resourceName, "id")
		if err != nil {
			return err
		}
		if *previousID != "" && expectReplacement == (id == *previousID) {
			return fmt.Errorf("%s resource ID changed=%t; expected replacement=%t", resourceName, id != *previousID, expectReplacement)
		}
		*previousID = id
		return nil
	}
}
