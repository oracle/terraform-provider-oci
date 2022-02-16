// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerEnrollmentStatusesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerEnrollmentStatuses,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enrollment_status_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OptimizerEnrollmentStatusResource()),
						},
					},
				},
			},
		},
	}
}

func readOptimizerEnrollmentStatuses(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerEnrollmentStatusesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerEnrollmentStatusesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListEnrollmentStatusesResponse
}

func (s *OptimizerEnrollmentStatusesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerEnrollmentStatusesDataSourceCrud) Get() error {
	request := oci_optimizer.ListEnrollmentStatusesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_optimizer.ListEnrollmentStatusesLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.ListEnrollmentStatusesStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.ListEnrollmentStatuses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEnrollmentStatuses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerEnrollmentStatusesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerEnrollmentStatusesDataSource-", OptimizerEnrollmentStatusesDataSource(), s.D))
	resources := []map[string]interface{}{}
	enrollmentStatus := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EnrollmentStatusSummaryToMap(item))
	}
	enrollmentStatus["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerEnrollmentStatusesDataSource().Schema["enrollment_status_collection"].Elem.(*schema.Resource).Schema)
		enrollmentStatus["items"] = items
	}

	resources = append(resources, enrollmentStatus)
	if err := s.D.Set("enrollment_status_collection", resources); err != nil {
		return err
	}

	return nil
}
