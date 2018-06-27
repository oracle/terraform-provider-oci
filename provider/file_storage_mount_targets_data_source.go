// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func MountTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMountTargets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
				Elem:     MountTargetResource(),
			},
		},
	}
}

func readMountTargets(d *schema.ResourceData, m interface{}) error {
	sync := &MountTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

type MountTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListMountTargetsResponse
}

func (s *MountTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MountTargetsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

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

func (s *MountTargetsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mountTarget := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DisplayName != nil {
			mountTarget["display_name"] = *r.DisplayName
		}

		if r.ExportSetId != nil {
			mountTarget["export_set_id"] = *r.ExportSetId
		}

		if r.Id != nil {
			mountTarget["id"] = *r.Id
		}

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
		resources = ApplyFilters(f.(*schema.Set), resources, MountTargetsDataSource().Schema["mount_targets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("mount_targets", resources); err != nil {
		panic(err)
	}

	return
}
