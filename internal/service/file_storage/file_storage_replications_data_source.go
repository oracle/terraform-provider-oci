// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageReplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageReplications,
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
			"file_system_id": {
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
			"replications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageReplicationResource()),
			},
		},
	}
}

func readFileStorageReplications(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageReplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListReplicationsResponse
}

func (s *FileStorageReplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageReplicationsDataSourceCrud) Get() error {
	request := oci_file_storage.ListReplicationsRequest{}

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

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListReplicationsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListReplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageReplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageReplicationsDataSource-", FileStorageReplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		replication := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			replication["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			replication["display_name"] = *r.DisplayName
		}

		replication["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			replication["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			replication["lifecycle_details"] = *r.LifecycleDetails
		}

		locks := []interface{}{}
		for _, item := range r.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		replication["locks"] = locks

		if r.RecoveryPointTime != nil {
			replication["recovery_point_time"] = r.RecoveryPointTime.String()
		}

		if r.ReplicationInterval != nil {
			replication["replication_interval"] = strconv.FormatInt(*r.ReplicationInterval, 10)
		}

		replication["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			replication["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, replication)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageReplicationsDataSource().Schema["replications"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("replications", resources); err != nil {
		return err
	}

	return nil
}
