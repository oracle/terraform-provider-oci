// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceDataSciencePrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_science_resource_type": {
				Type:     schema.TypeString,
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
			"data_science_private_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatasciencePrivateEndpointResource()),
			},
		},
	}
}

func readDatascienceDataSciencePrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceDataSciencePrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceDataSciencePrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListDataSciencePrivateEndpointsResponse
}

func (s *DatascienceDataSciencePrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceDataSciencePrivateEndpointsDataSourceCrud) Get() error {
	request := oci_datascience.ListDataSciencePrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if dataScienceResourceType, ok := s.D.GetOkExists("data_science_resource_type"); ok {
		request.DataScienceResourceType = oci_datascience.ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum(dataScienceResourceType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListDataSciencePrivateEndpointsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListDataSciencePrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataSciencePrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceDataSciencePrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatasciencePrivateEndpointsDataSource-", DatasciencePrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dataSciencePrivateEndpoint := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			dataSciencePrivateEndpoint["created_by"] = *r.CreatedBy
		}

		dataSciencePrivateEndpoint["data_science_resource_type"] = r.DataScienceResourceType

		if r.DefinedTags != nil {
			dataSciencePrivateEndpoint["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			dataSciencePrivateEndpoint["display_name"] = *r.DisplayName
		}

		if r.Fqdn != nil {
			dataSciencePrivateEndpoint["fqdn"] = *r.Fqdn
		}

		dataSciencePrivateEndpoint["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dataSciencePrivateEndpoint["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dataSciencePrivateEndpoint["lifecycle_details"] = *r.LifecycleDetails
		}

		dataSciencePrivateEndpoint["nsg_ids"] = r.NsgIds

		dataSciencePrivateEndpoint["state"] = r.LifecycleState

		if r.SubnetId != nil {
			dataSciencePrivateEndpoint["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			dataSciencePrivateEndpoint["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			dataSciencePrivateEndpoint["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, dataSciencePrivateEndpoint)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatasciencePrivateEndpointsDataSource().Schema["data_science_private_endpoints"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("data_science_private_endpoints", resources); err != nil {
		return err
	}

	return nil
}
