package adm

import (
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("adm", admResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAdmVulnerabilityAuditHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_adm_vulnerability_audit",
	DatasourceClass:        "oci_adm_vulnerability_audits",
	DatasourceItemsAttr:    "vulnerability_audit_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vulnerability_audit",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_adm.VulnerabilityAuditLifecycleStateActive),
	},
}

var exportAdmKnowledgeBaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_adm_knowledge_base",
	DatasourceClass:        "oci_adm_knowledge_bases",
	DatasourceItemsAttr:    "knowledge_base_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "knowledge_base",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_adm.KnowledgeBaseLifecycleStateActive),
	},
}

var admResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAdmVulnerabilityAuditHints},
		{TerraformResourceHints: exportAdmKnowledgeBaseHints},
	},
}
