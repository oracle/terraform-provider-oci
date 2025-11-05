// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeSnapshotManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDbNodeSnapshotManagement,
		Read:     readDatabaseDbNodeSnapshotManagement,
		Delete:   deleteDatabaseDbNodeSnapshotManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"exadb_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_dbnode_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mount_points": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_node_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_dbnode_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volumes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func createDatabaseDbNodeSnapshotManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeSnapshotManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if err := tfresource.CreateResource(d, sync); err != nil {
		return err
	}

	// lifecycleState becomes AVAILABLE after work request succeeds
	if snapshots, exists := sync.D.GetOkExists("snapshots"); exists {
		if snapshotList := snapshots.([]interface{}); len(snapshotList) > 0 {
			for _, snapshot := range snapshotList {
				tmp := snapshot.(map[string]interface{})
				tmp["state"] = string(oci_database.DbnodeSnapshotLifecycleStateAvailable)
			}
			sync.D.Set("snapshots", snapshotList)
		}
	}
	return nil
}

func readDatabaseDbNodeSnapshotManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseDbNodeSnapshotManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseDbNodeSnapshotManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AddDbnodeSnapshotsForExadbVmClusterResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDbNodeSnapshotManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseDbNodeSnapshotManagementResource-", DatabaseDbNodeSnapshotManagementResource(), s.D)
}

func (s *DatabaseDbNodeSnapshotManagementResourceCrud) Create() error {
	request := oci_database.AddDbnodeSnapshotsForExadbVmClusterRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if exadbVmClusterId, ok := s.D.GetOkExists("exadb_vm_cluster_id"); ok {
		tmp := exadbVmClusterId.(string)
		request.ExadbVmClusterId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if sourceDbnodeIds, ok := s.D.GetOkExists("source_dbnode_ids"); ok {
		set := sourceDbnodeIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("source_dbnode_ids") {
			request.SourceDbnodeIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.AddDbnodeSnapshotsForExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response

	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbnodesnapshot", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseDbNodeSnapshotManagementResourceCrud) SetData() error {
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	snapshots := []interface{}{}
	for _, item := range s.Res.Snapshots {
		snapshots = append(snapshots, DbnodeSnapshotToMap(item))
	}
	s.D.Set("snapshots", snapshots)
	return nil
}

func DbnodeSnapshotToMap(obj oci_database.DbnodeSnapshot) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	mountPoints := []interface{}{}
	for _, item := range obj.MountPoints {
		mountPoints = append(mountPoints, MountPointDetailsToMap(item))
	}
	result["mount_points"] = mountPoints

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SourceDbnodeId != nil {
		result["source_dbnode_id"] = string(*obj.SourceDbnodeId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	volumes := []interface{}{}
	for _, item := range obj.Volumes {
		volumes = append(volumes, VolumeDetailsToMap(item))
	}
	result["volumes"] = volumes

	return result
}
