// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageObjectStorageLinksDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readLustreFileStorageObjectStorageLinksWithContext,
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
			"lustre_file_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_storage_link_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LustreFileStorageObjectStorageLinkResource()),
						},
					},
				},
			},
		},
	}
}

func readLustreFileStorageObjectStorageLinksWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type LustreFileStorageObjectStorageLinksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.ListObjectStorageLinksResponse
}

func (s *LustreFileStorageObjectStorageLinksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageObjectStorageLinksDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.ListObjectStorageLinksRequest{}

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

	if lustreFileSystemId, ok := s.D.GetOkExists("lustre_file_system_id"); ok {
		tmp := lustreFileSystemId.(string)
		request.LustreFileSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_lustre_file_storage.ObjectStorageLinkLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.ListObjectStorageLinks(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListObjectStorageLinks(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LustreFileStorageObjectStorageLinksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LustreFileStorageObjectStorageLinksDataSource-", LustreFileStorageObjectStorageLinksDataSource(), s.D))
	resources := []map[string]interface{}{}
	objectStorageLink := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ObjectStorageLinkSummaryToMap(item))
	}
	objectStorageLink["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LustreFileStorageObjectStorageLinksDataSource().Schema["object_storage_link_collection"].Elem.(*schema.Resource).Schema)
		objectStorageLink["items"] = items
	}

	resources = append(resources, objectStorageLink)
	if err := s.D.Set("object_storage_link_collection", resources); err != nil {
		return err
	}

	return nil
}
