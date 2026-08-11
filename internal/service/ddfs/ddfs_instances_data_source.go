// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ddfs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ddfs "github.com/oracle/oci-go-sdk/v65/ddfs"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DdfsInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDdfsInstancesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DdfsInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readDdfsInstancesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DdfsInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InstanceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DdfsInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ddfs.InstanceClient
	Res    *oci_ddfs.ListInstancesResponse
}

func (s *DdfsInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DdfsInstancesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ddfs.ListInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ddfs.InstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ddfs")

	response, err := s.Client.ListInstances(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstances(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DdfsInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DdfsInstancesDataSource-", DdfsInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	instance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstanceSummaryToMap(item))
	}
	instance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DdfsInstancesDataSource().Schema["instance_collection"].Elem.(*schema.Resource).Schema)
		instance["items"] = items
	}

	resources = append(resources, instance)
	if err := s.D.Set("instance_collection", resources); err != nil {
		return err
	}

	return nil
}
