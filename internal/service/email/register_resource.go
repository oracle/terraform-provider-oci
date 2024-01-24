// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_email_dkim", EmailDkimResource())
	tfresource.RegisterResource("oci_email_email_domain", EmailEmailDomainResource())
	tfresource.RegisterResource("oci_email_sender", EmailSenderResource())
	tfresource.RegisterResource("oci_email_suppression", EmailSuppressionResource())
}
