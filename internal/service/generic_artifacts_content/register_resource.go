// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generic_artifacts_content

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_generic_artifacts_content_artifact_by_path", GenericArtifactsContentArtifactByPathResource())
}
