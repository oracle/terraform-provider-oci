// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
)

func DnsSteeringPolicyAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsSteeringPolicyAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"steering_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"steering_policy_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DnsSteeringPolicyAttachmentResource()),
			},
		},
	}
}

func readDnsSteeringPolicyAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsSteeringPolicyAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListSteeringPolicyAttachmentsResponse
}

func (s *DnsSteeringPolicyAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsSteeringPolicyAttachmentsDataSourceCrud) Get() error {
	request := oci_dns.ListSteeringPolicyAttachmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if domainContains, ok := s.D.GetOkExists("domain_contains"); ok {
		tmp := domainContains.(string)
		request.DomainContains = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.SteeringPolicyAttachmentSummaryLifecycleStateEnum(state.(string))
	}

	if steeringPolicyId, ok := s.D.GetOkExists("steering_policy_id"); ok {
		tmp := steeringPolicyId.(string)
		request.SteeringPolicyId = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if zoneId, ok := s.D.GetOkExists("zone_id"); ok {
		tmp := zoneId.(string)
		request.ZoneId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.ListSteeringPolicyAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSteeringPolicyAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsSteeringPolicyAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsSteeringPolicyAttachmentsDataSource-", DnsSteeringPolicyAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		steeringPolicyAttachment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			steeringPolicyAttachment["display_name"] = *r.DisplayName
		}

		if r.DomainName != nil {
			steeringPolicyAttachment["domain_name"] = *r.DomainName
		}

		if r.Id != nil {
			steeringPolicyAttachment["id"] = *r.Id
		}

		steeringPolicyAttachment["rtypes"] = r.Rtypes

		if r.Self != nil {
			steeringPolicyAttachment["self"] = *r.Self
		}

		steeringPolicyAttachment["state"] = r.LifecycleState

		if r.SteeringPolicyId != nil {
			steeringPolicyAttachment["steering_policy_id"] = *r.SteeringPolicyId
		}

		if r.TimeCreated != nil {
			steeringPolicyAttachment["time_created"] = r.TimeCreated.String()
		}

		if r.ZoneId != nil {
			steeringPolicyAttachment["zone_id"] = *r.ZoneId
		}

		resources = append(resources, steeringPolicyAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsSteeringPolicyAttachmentsDataSource().Schema["steering_policy_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("steering_policy_attachments", resources); err != nil {
		return err
	}

	return nil
}
