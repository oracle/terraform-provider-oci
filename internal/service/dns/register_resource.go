// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_dns_action_create_zone_from_zone_file", DnsActionCreateZoneFromZoneFileResource())
	tfresource.RegisterResource("oci_dns_record", DnsRecordResource())
	tfresource.RegisterResource("oci_dns_resolver", DnsResolverResource())
	tfresource.RegisterResource("oci_dns_resolver_endpoint", DnsResolverEndpointResource())
	tfresource.RegisterResource("oci_dns_rrset", DnsRrsetResource())
	tfresource.RegisterResource("oci_dns_steering_policy", DnsSteeringPolicyResource())
	tfresource.RegisterResource("oci_dns_steering_policy_attachment", DnsSteeringPolicyAttachmentResource())
	tfresource.RegisterResource("oci_dns_tsig_key", DnsTsigKeyResource())
	tfresource.RegisterResource("oci_dns_view", DnsViewResource())
	tfresource.RegisterResource("oci_dns_zone", DnsZoneResource())
}
