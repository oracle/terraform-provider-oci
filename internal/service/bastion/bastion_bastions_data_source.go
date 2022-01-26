// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bastion "github.com/oracle/oci-go-sdk/v56/bastion"
)

func BastionBastionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBastionBastions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bastion_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bastion_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bastions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BastionBastionResource()),
			},
		},
	}
}

func readBastionBastions(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

type BastionBastionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bastion.BastionClient
	Res    *oci_bastion.ListBastionsResponse
}

func (s *BastionBastionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BastionBastionsDataSourceCrud) Get() error {
	request := oci_bastion.ListBastionsRequest{}

	if bastionId, ok := s.D.GetOkExists("id"); ok {
		tmp := bastionId.(string)
		request.BastionId = &tmp
	}

	if bastionLifecycleState, ok := s.D.GetOkExists("bastion_lifecycle_state"); ok {
		request.BastionLifecycleState = oci_bastion.ListBastionsBastionLifecycleStateEnum(bastionLifecycleState.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bastion")

	response, err := s.Client.ListBastions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBastions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BastionBastionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BastionBastionsDataSource-", BastionBastionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bastion := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BastionType != nil {
			bastion["bastion_type"] = *r.BastionType
		}

		if r.DefinedTags != nil {
			bastion["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		bastion["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			bastion["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			bastion["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Name != nil {
			bastion["name"] = *r.Name
		}

		bastion["state"] = r.LifecycleState

		if r.SystemTags != nil {
			bastion["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TargetSubnetId != nil {
			bastion["target_subnet_id"] = *r.TargetSubnetId
		}

		if r.TargetVcnId != nil {
			bastion["target_vcn_id"] = *r.TargetVcnId
		}

		if r.TimeCreated != nil {
			bastion["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			bastion["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, bastion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BastionBastionsDataSource().Schema["bastions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bastions", resources); err != nil {
		return err
	}

	return nil
}
