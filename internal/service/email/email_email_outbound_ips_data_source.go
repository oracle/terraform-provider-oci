// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmailEmailOutboundIpsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readEmailEmailOutboundIpsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assignment_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"outbound_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_outbound_ip_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assignment_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"outbound_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readEmailEmailOutboundIpsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailOutboundIpsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type EmailEmailOutboundIpsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListEmailOutboundIpsResponse
}

func (s *EmailEmailOutboundIpsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailEmailOutboundIpsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_email.ListEmailOutboundIpsRequest{}

	if assignmentState, ok := s.D.GetOkExists("assignment_state"); ok {
		request.AssignmentState = oci_email.EmailOutboundIpSummaryAssignmentStateEnum(assignmentState.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if outboundIp, ok := s.D.GetOkExists("outbound_ip"); ok {
		tmp := outboundIp.(string)
		request.OutboundIp = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_email.EmailOutboundIpSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.ListEmailOutboundIps(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEmailOutboundIps(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmailEmailOutboundIpsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailEmailOutboundIpsDataSource-", EmailEmailOutboundIpsDataSource(), s.D))
	resources := []map[string]interface{}{}
	emailOutboundIp := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EmailOutboundIpSummaryToMap(item))
	}
	emailOutboundIp["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, EmailEmailOutboundIpsDataSource().Schema["email_outbound_ip_collection"].Elem.(*schema.Resource).Schema)
		emailOutboundIp["items"] = items
	}

	resources = append(resources, emailOutboundIp)
	if err := s.D.Set("email_outbound_ip_collection", resources); err != nil {
		return err
	}

	return nil
}

func EmailOutboundIpSummaryToMap(obj oci_email.EmailOutboundIpSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["assignment_state"] = string(obj.AssignmentState)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OutboundIp != nil {
		result["outbound_ip"] = string(*obj.OutboundIp)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
