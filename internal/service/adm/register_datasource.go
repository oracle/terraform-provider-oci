// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_adm_knowledge_base", AdmKnowledgeBaseDataSource())
	tfresource.RegisterDatasource("oci_adm_knowledge_bases", AdmKnowledgeBasesDataSource())
	tfresource.RegisterDatasource("oci_adm_vulnerability_audit", AdmVulnerabilityAuditDataSource())
	tfresource.RegisterDatasource("oci_adm_vulnerability_audit_application_dependency_vulnerabilities", AdmVulnerabilityAuditApplicationDependencyVulnerabilitiesDataSource())
	tfresource.RegisterDatasource("oci_adm_vulnerability_audit_application_dependency_vulnerability", AdmVulnerabilityAuditApplicationDependencyVulnerabilityDataSource())
	tfresource.RegisterDatasource("oci_adm_vulnerability_audits", AdmVulnerabilityAuditsDataSource())
}
