// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_adm_knowledge_base", AdmKnowledgeBaseResource())
	tfresource.RegisterResource("oci_adm_vulnerability_audit", AdmVulnerabilityAuditResource())
}
