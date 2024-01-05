// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeDataSafePrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeDataSafePrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_safe_private_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataSafeDataSafePrivateEndpointResource()),
			},
		},
	}
}

func readDataSafeDataSafePrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafePrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDataSafePrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDataSafePrivateEndpointsResponse
}

func (s *DataSafeDataSafePrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDataSafePrivateEndpointsDataSourceCrud) Get() error {
	request := oci_data_safe.ListDataSafePrivateEndpointsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListDataSafePrivateEndpointsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListDataSafePrivateEndpointsLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDataSafePrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataSafePrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeDataSafePrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeDataSafePrivateEndpointsDataSource-", DataSafeDataSafePrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dataSafePrivateEndpoint := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			dataSafePrivateEndpoint["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			dataSafePrivateEndpoint["description"] = *r.Description
		}

		if r.DisplayName != nil {
			dataSafePrivateEndpoint["display_name"] = *r.DisplayName
		}

		dataSafePrivateEndpoint["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dataSafePrivateEndpoint["id"] = *r.Id
		}

		if r.PrivateEndpointId != nil {
			dataSafePrivateEndpoint["private_endpoint_id"] = *r.PrivateEndpointId
		}

		dataSafePrivateEndpoint["state"] = r.LifecycleState

		if r.SubnetId != nil {
			dataSafePrivateEndpoint["subnet_id"] = *r.SubnetId
		}

		if r.SystemTags != nil {
			dataSafePrivateEndpoint["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			dataSafePrivateEndpoint["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			dataSafePrivateEndpoint["vcn_id"] = *r.VcnId
		}

		resources = append(resources, dataSafePrivateEndpoint)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeDataSafePrivateEndpointsDataSource().Schema["data_safe_private_endpoints"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("data_safe_private_endpoints", resources); err != nil {
		return err
	}

	return nil
}
