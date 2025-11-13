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

func DatascienceMlApplicationImplementationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceMlApplicationImplementations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ml_application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ml_application_implementation_id": {
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
			"ml_application_implementation_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatascienceMlApplicationImplementationResource()),
						},
					},
				},
			},
		},
	}
}

func readDatascienceMlApplicationImplementations(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceMlApplicationImplementationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListMlApplicationImplementationsResponse
}

func (s *DatascienceMlApplicationImplementationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceMlApplicationImplementationsDataSourceCrud) Get() error {
	request := oci_datascience.ListMlApplicationImplementationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if mlApplicationId, ok := s.D.GetOkExists("ml_application_id"); ok {
		tmp := mlApplicationId.(string)
		request.MlApplicationId = &tmp
	}

	if mlApplicationImplementationId, ok := s.D.GetOkExists("id"); ok {
		tmp := mlApplicationImplementationId.(string)
		request.MlApplicationImplementationId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.MlApplicationImplementationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListMlApplicationImplementations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMlApplicationImplementations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceMlApplicationImplementationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceMlApplicationImplementationsDataSource-", DatascienceMlApplicationImplementationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	mlApplicationImplementation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MlApplicationImplementationSummaryToMap(item))
	}
	mlApplicationImplementation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatascienceMlApplicationImplementationsDataSource().Schema["ml_application_implementation_collection"].Elem.(*schema.Resource).Schema)
		mlApplicationImplementation["items"] = items
	}

	resources = append(resources, mlApplicationImplementation)
	if err := s.D.Set("ml_application_implementation_collection", resources); err != nil {
		return err
	}

	return nil
}
