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
	DataSafeAuditPolicyAnalyticResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Required, acctest.Create, auditPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Required, acctest.Create, ruleRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditPolicyAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditPolicyAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{})
}
