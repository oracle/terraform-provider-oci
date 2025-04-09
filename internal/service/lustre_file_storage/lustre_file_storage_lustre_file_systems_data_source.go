// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageLustreFileSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLustreFileStorageLustreFileSystems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"lustre_file_system_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LustreFileStorageLustreFileSystemResource()),
						},
					},
				},
			},
		},
	}
}

func readLustreFileStorageLustreFileSystems(d *schema.ResourceData, m interface{}) error {
	sync := &LustreFileStorageLustreFileSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.ReadResource(sync)
}

type LustreFileStorageLustreFileSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.ListLustreFileSystemsResponse
}

func (s *LustreFileStorageLustreFileSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageLustreFileSystemsDataSourceCrud) Get() error {
	request := oci_lustre_file_storage.ListLustreFileSystemsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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
		request.LifecycleState = oci_lustre_file_storage.LustreFileSystemLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.ListLustreFileSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLustreFileSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LustreFileStorageLustreFileSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LustreFileStorageLustreFileSystemsDataSource-", LustreFileStorageLustreFileSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	lustreFileSystem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LustreFileSystemSummaryToMap(item, true))
	}
	lustreFileSystem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LustreFileStorageLustreFileSystemsDataSource().Schema["lustre_file_system_collection"].Elem.(*schema.Resource).Schema)
		lustreFileSystem["items"] = items
	}

	resources = append(resources, lustreFileSystem)
	if err := s.D.Set("lustre_file_system_collection", resources); err != nil {
		return err
	}

	return nil
}
