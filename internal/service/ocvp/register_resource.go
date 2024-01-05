// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_ocvp_cluster", OcvpClusterResource())
	tfresource.RegisterResource("oci_ocvp_esxi_host", OcvpEsxiHostResource())
	tfresource.RegisterResource("oci_ocvp_sddc", OcvpSddcResource())
}
