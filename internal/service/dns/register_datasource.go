// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dns_records", DnsRecordsDataSource())
	tfresource.RegisterDatasource("oci_dns_resolver", DnsResolverDataSource())
	tfresource.RegisterDatasource("oci_dns_resolver_endpoint", DnsResolverEndpointDataSource())
	tfresource.RegisterDatasource("oci_dns_resolver_endpoints", DnsResolverEndpointsDataSource())
	tfresource.RegisterDatasource("oci_dns_resolvers", DnsResolversDataSource())
	tfresource.RegisterDatasource("oci_dns_rrset", DnsRrsetDataSource())
	tfresource.RegisterDatasource("oci_dns_rrsets", DnsRrsetsDataSource())
	tfresource.RegisterDatasource("oci_dns_steering_policies", DnsSteeringPoliciesDataSource())
	tfresource.RegisterDatasource("oci_dns_steering_policy", DnsSteeringPolicyDataSource())
	tfresource.RegisterDatasource("oci_dns_steering_policy_attachment", DnsSteeringPolicyAttachmentDataSource())
	tfresource.RegisterDatasource("oci_dns_steering_policy_attachments", DnsSteeringPolicyAttachmentsDataSource())
	tfresource.RegisterDatasource("oci_dns_tsig_key", DnsTsigKeyDataSource())
	tfresource.RegisterDatasource("oci_dns_tsig_keys", DnsTsigKeysDataSource())
	tfresource.RegisterDatasource("oci_dns_view", DnsViewDataSource())
	tfresource.RegisterDatasource("oci_dns_views", DnsViewsDataSource())
	tfresource.RegisterDatasource("oci_dns_zones", DnsZonesDataSource())
}
