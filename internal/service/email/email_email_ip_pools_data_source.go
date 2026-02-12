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

func EmailEmailIpPoolsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readEmailEmailIpPoolsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_ip_pool_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(EmailEmailIpPoolResource()),
						},
					},
				},
			},
		},
	}
}

func readEmailEmailIpPoolsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &EmailEmailIpPoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type EmailEmailIpPoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListEmailIpPoolsResponse
}

func (s *EmailEmailIpPoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailEmailIpPoolsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_email.ListEmailIpPoolsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_email.EmailIpPoolLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.ListEmailIpPools(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEmailIpPools(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmailEmailIpPoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailEmailIpPoolsDataSource-", EmailEmailIpPoolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	emailIpPool := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EmailIpPoolSummaryToMap(item))
	}
	emailIpPool["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, EmailEmailIpPoolsDataSource().Schema["email_ip_pool_collection"].Elem.(*schema.Resource).Schema)
		emailIpPool["items"] = items
	}

	resources = append(resources, emailIpPool)
	if err := s.D.Set("email_ip_pool_collection", resources); err != nil {
		return err
	}

	return nil
}
