// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DnsZonePromoteDnssecKeyVersionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsZonePromoteDnssecKeyVersion,
		Read:     readDnsZonePromoteDnssecKeyVersion,
		Delete:   deleteDnsZonePromoteDnssecKeyVersion,
		Schema: map[string]*schema.Schema{
			// Required
			"dnssec_key_version_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDnsZonePromoteDnssecKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZonePromoteDnssecKeyVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDnsZonePromoteDnssecKeyVersion(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDnsZonePromoteDnssecKeyVersion(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DnsZonePromoteDnssecKeyVersionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DnsZonePromoteDnssecKeyVersionResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DnsZonePromoteDnssecKeyVersionResourceCrud) Create() error {
	request := oci_dns.PromoteZoneDnssecKeyVersionRequest{}

	if dnssecKeyVersionUuid, ok := s.D.GetOkExists("dnssec_key_version_uuid"); ok {
		tmp := dnssecKeyVersionUuid.(string)
		request.DnssecKeyVersionUuid = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.PromoteZoneDnssecKeyVersionScopeEnum(scope.(string))
	}

	if zoneId, ok := s.D.GetOkExists("zone_id"); ok {
		tmp := zoneId.(string)
		request.ZoneId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.PromoteZoneDnssecKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "zone",
			oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.D.SetId(fmt.Sprintf("%s-%s", *request.ZoneId, *request.DnssecKeyVersionUuid))

	return nil
}

func (s *DnsZonePromoteDnssecKeyVersionResourceCrud) SetData() error {
	return nil
}
