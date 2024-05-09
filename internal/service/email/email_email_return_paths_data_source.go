// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmailEmailReturnPathsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEmailEmailReturnPaths,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_return_path_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(EmailEmailReturnPathResource()),
						},
					},
				},
			},
		},
	}
}

func readEmailEmailReturnPaths(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailReturnPathsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailEmailReturnPathsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListEmailReturnPathsResponse
}

func (s *EmailEmailReturnPathsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailEmailReturnPathsDataSourceCrud) Get() error {
	request := oci_email.ListEmailReturnPathsRequest{}

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

	if parentResourceId, ok := s.D.GetOkExists("parent_resource_id"); ok {
		tmp := parentResourceId.(string)
		request.ParentResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_email.EmailReturnPathLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.ListEmailReturnPaths(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEmailReturnPaths(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmailEmailReturnPathsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailEmailReturnPathsDataSource-", EmailEmailReturnPathsDataSource(), s.D))
	resources := []map[string]interface{}{}
	emailReturnPath := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EmailReturnPathSummaryToMap(item))
	}
	emailReturnPath["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, EmailEmailReturnPathsDataSource().Schema["email_return_path_collection"].Elem.(*schema.Resource).Schema)
		emailReturnPath["items"] = items
	}

	resources = append(resources, emailReturnPath)
	if err := s.D.Set("email_return_path_collection", resources); err != nil {
		return err
	}

	return nil
}
