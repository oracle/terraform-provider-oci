// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v58/filestorage"
)

func FileStorageMountTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageMountTargets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"export_set_id": {
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
			"mount_targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageMountTargetResource()),
			},
		},
	}
}

func readFileStorageMountTargets(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageMountTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageMountTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListMountTargetsResponse
}

func (s *FileStorageMountTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageMountTargetsDataSourceCrud) Get() error {
	request := oci_file_storage.ListMountTargetsRequest{}

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

	if exportSetId, ok := s.D.GetOkExists("export_set_id"); ok {
		tmp := exportSetId.(string)
		request.ExportSetId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListMountTargetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListMountTargets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMountTargets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageMountTargetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageMountTargetsDataSource-", FileStorageMountTargetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mountTarget := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			mountTarget["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			mountTarget["display_name"] = *r.DisplayName
		}

		if r.ExportSetId != nil {
			mountTarget["export_set_id"] = *r.ExportSetId
		}

		mountTarget["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mountTarget["id"] = *r.Id
		}

		mountTarget["nsg_ids"] = r.NsgIds

		mountTarget["private_ip_ids"] = r.PrivateIpIds

		mountTarget["state"] = r.LifecycleState

		if r.SubnetId != nil {
			mountTarget["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			mountTarget["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, mountTarget)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageMountTargetsDataSource().Schema["mount_targets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("mount_targets", resources); err != nil {
		return err
	}

	return nil
}
