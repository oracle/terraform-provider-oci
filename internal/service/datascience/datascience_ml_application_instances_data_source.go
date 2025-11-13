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

func DatascienceMlApplicationInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceMlApplicationInstances,
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
			"ml_application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ml_application_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatascienceMlApplicationInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readDatascienceMlApplicationInstances(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListMlApplicationInstancesResponse
}

func (s *DatascienceMlApplicationInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationInstancesDataSourceCrud) Get() error {
	request := oci_datascience.ListMlApplicationInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if mlApplicationId, ok := s.D.GetOkExists("ml_application_id"); ok {
		tmp := mlApplicationId.(string)
		request.MlApplicationId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.MlApplicationInstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListMlApplicationInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMlApplicationInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceMlApplicationInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceMlApplicationInstancesDataSource-", DatascienceMlApplicationInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	mlApplicationInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MlApplicationInstanceSummaryToMap(item))
	}
	mlApplicationInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatascienceMlApplicationInstancesDataSource().Schema["ml_application_instance_collection"].Elem.(*schema.Resource).Schema)
		mlApplicationInstance["items"] = items
	}

	resources = append(resources, mlApplicationInstance)
	if err := s.D.Set("ml_application_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
