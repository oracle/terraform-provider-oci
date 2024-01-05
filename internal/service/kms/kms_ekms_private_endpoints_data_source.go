// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsEkmsPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsEkmsPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ekms_private_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(KmsEkmsPrivateEndpointResource()),
			},
		},
	}
}

func readKmsEkmsPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()

	return tfresource.ReadResource(sync)
}

type KmsEkmsPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.EkmClient
	Res    *oci_kms.ListEkmsPrivateEndpointsResponse
}

func (s *KmsEkmsPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsEkmsPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_kms.ListEkmsPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.ListEkmsPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEkmsPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *KmsEkmsPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsEkmsPrivateEndpointsDataSource-", KmsEkmsPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ekmsPrivateEndpoint := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			ekmsPrivateEndpoint["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			ekmsPrivateEndpoint["display_name"] = *r.DisplayName
		}

		ekmsPrivateEndpoint["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			ekmsPrivateEndpoint["id"] = *r.Id
		}

		ekmsPrivateEndpoint["state"] = r.LifecycleState

		if r.SubnetId != nil {
			ekmsPrivateEndpoint["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			ekmsPrivateEndpoint["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			ekmsPrivateEndpoint["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, ekmsPrivateEndpoint)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, KmsEkmsPrivateEndpointsDataSource().Schema["ekms_private_endpoints"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ekms_private_endpoints", resources); err != nil {
		return err
	}

	return nil
}
