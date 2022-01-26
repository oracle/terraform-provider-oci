// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v56/email"
)

func EmailDkimsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEmailDkims,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"email_domain_id": {
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
			"dkim_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(EmailDkimResource()),
						},
					},
				},
			},
		},
	}
}

func readEmailDkims(d *schema.ResourceData, m interface{}) error {
	sync := &EmailDkimsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailDkimsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListDkimsResponse
}

func (s *EmailDkimsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailDkimsDataSourceCrud) Get() error {
	request := oci_email.ListDkimsRequest{}

	if emailDomainId, ok := s.D.GetOkExists("email_domain_id"); ok {
		tmp := emailDomainId.(string)
		request.EmailDomainId = &tmp
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
		request.LifecycleState = oci_email.DkimLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.ListDkims(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDkims(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmailDkimsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailDkimsDataSource-", EmailDkimsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dkim := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DkimSummaryToMap(item))
	}
	dkim["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, EmailDkimsDataSource().Schema["dkim_collection"].Elem.(*schema.Resource).Schema)
		dkim["items"] = items
	}

	resources = append(resources, dkim)
	if err := s.D.Set("dkim_collection", resources); err != nil {
		return err
	}

	return nil
}
