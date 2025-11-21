// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
)

func AnalyticsAnalyticsInstancePrivateAccessChannelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["analytics_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["private_access_channel_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AnalyticsAnalyticsInstancePrivateAccessChannelResource(), fieldMap, readSingularAnalyticsAnalyticsInstancePrivateAccessChannel)
}

func readSingularAnalyticsAnalyticsInstancePrivateAccessChannel(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancePrivateAccessChannelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.ReadResource(sync)
}

type AnalyticsAnalyticsInstancePrivateAccessChannelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_analytics.AnalyticsClient
	Res    *oci_analytics.GetPrivateAccessChannelResponse
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelDataSourceCrud) Get() error {
	request := oci_analytics.GetPrivateAccessChannelRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if privateAccessChannelKey, ok := s.D.GetOkExists("private_access_channel_key"); ok {
		tmp := privateAccessChannelKey.(string)
		request.PrivateAccessChannelKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "analytics")

	response, err := s.Client.GetPrivateAccessChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AnalyticsAnalyticsInstancePrivateAccessChannelDataSource-", AnalyticsAnalyticsInstancePrivateAccessChannelDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("egress_source_ip_addresses", s.Res.EgressSourceIpAddresses)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("network_security_group_ids", s.Res.NetworkSecurityGroupIds)

	privateSourceDnsZones := []interface{}{}
	for _, item := range s.Res.PrivateSourceDnsZones {
		privateSourceDnsZones = append(privateSourceDnsZones, PrivateSourceDnsZoneToMap(item))
	}
	s.D.Set("private_source_dns_zones", privateSourceDnsZones)

	privateSourceScanHosts := []interface{}{}
	for _, item := range s.Res.PrivateSourceScanHosts {
		privateSourceScanHosts = append(privateSourceScanHosts, PrivateSourceScanHostToMap(item))
	}
	s.D.Set("private_source_scan_hosts", privateSourceScanHosts)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
