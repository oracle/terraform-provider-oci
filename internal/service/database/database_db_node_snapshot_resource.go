package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDbnodeSnapshot,
		Read:     readDatabaseDbnodeSnapshot,
		Update:   updateDatabaseDbnodeSnapshot,
		Delete:   deleteDatabaseDbnodeSnapshot,
		Schema: map[string]*schema.Schema{
			// Required
			"dbnode_snapshot_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mount_dbnode_id": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc:     stringMustBeDbNodeIdOrNull(),
			},

			// Optional

			// Computed
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_dbnode_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
	}
}

func createDatabaseDbnodeSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbnodeSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	// read "mount_dbnode_id" field from config prior to the first Create call and no Terraform state exists yet for the dbnode snapshot
	targetMountDbNodeId := "null"
	if tmp, ok := sync.D.GetOkExists("mount_dbnode_id"); ok {
		targetMountDbNodeId = strings.ToLower(tmp.(string))
	}

	// Get and populate state. Do not call ReadResource which will clear the state if dbnodeSnapshot is in terminated state
	if err := tfresource.CreateResource(d, sync); err != nil {
		return err
	}

	// read "mount_dbnode_id" field from the Terraform state of the dbnode snapshot after the first Create call
	currentMountDbNodeId := "null"
	if tmp, ok := sync.D.GetOkExists("mount_dbnode_id"); ok {
		currentMountDbNodeId = strings.ToLower(tmp.(string))
	}

	if currentMountDbNodeId != targetMountDbNodeId {
		if currentMountDbNodeId != "null" {
			if err := sync.UnmountDbNodeSnapshot(currentMountDbNodeId); err != nil {
				return err
			}
		}
		if targetMountDbNodeId != "null" {
			if err := sync.MountDbNodeSnapshot(targetMountDbNodeId); err != nil {
				return err
			}
		}
		// Get and update state. Do not call ReadResource which will clear the state if dbnodeSnapshot is in terminated state
		if err := tfresource.CreateResource(d, sync); err != nil {
			return err
		}
	}

	return nil

}

func readDatabaseDbnodeSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbnodeSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseDbnodeSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbnodeSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)

}

func deleteDatabaseDbnodeSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbnodeSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if tmp, ok := sync.D.GetOkExists("mount_dbnode_id"); ok {
		mountDbNodeId := strings.ToLower(tmp.(string))
		if mountDbNodeId != "null" {
			if err := sync.UnmountDbNodeSnapshot(mountDbNodeId); err != nil {
				return err
			}
		}
	}

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDbnodeSnapshotResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DbnodeSnapshot
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDbnodeSnapshotResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDbnodeSnapshotResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbnodeSnapshotLifecycleStateTerminating),
	}
}

func (s *DatabaseDbnodeSnapshotResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbnodeSnapshotLifecycleStateTerminated),
	}
}

func (s *DatabaseDbnodeSnapshotResourceCrud) Create() error {
	if dbNodeSnapshotId, ok := s.D.GetOkExists("dbnode_snapshot_id"); ok {
		s.D.SetId(dbNodeSnapshotId.(string))
		return s.Get()
	}
	// un-reachable code since dbnode_snapshot_id is a required parameter
	return fmt.Errorf("dbnode_snapshot resource does not have a dbnode_snapshot_id set")
}

func (s *DatabaseDbnodeSnapshotResourceCrud) Get() error {
	request := oci_database.GetDbnodeSnapshotRequest{}
	tmp := s.D.Id()
	request.DbnodeSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbnodeSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbnodeSnapshot
	return nil
}

func (s *DatabaseDbnodeSnapshotResourceCrud) Update() error {
	if _, ok := s.D.GetOkExists("mount_dbnode_id"); ok && s.D.HasChange("mount_dbnode_id") {
		oldRaw, newRaw := s.D.GetChange("mount_dbnode_id")
		targetMountDbNodeId := strings.ToLower(newRaw.(string))
		currentMountDbNodeId := strings.ToLower(oldRaw.(string))
		if currentMountDbNodeId != targetMountDbNodeId {
			if currentMountDbNodeId != "null" {
				if err := s.UnmountDbNodeSnapshot(currentMountDbNodeId); err != nil {
					return err
				}
			}
			if targetMountDbNodeId != "null" {
				if err := s.MountDbNodeSnapshot(targetMountDbNodeId); err != nil {
					return err
				}
			}
		}
	}
	return s.Get()
}

func (s *DatabaseDbnodeSnapshotResourceCrud) Delete() error {
	request := oci_database.DeleteDbnodeSnapshotRequest{}

	tmp := s.D.Id()
	request.DbnodeSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteDbnodeSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbnodesnapshot", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseDbnodeSnapshotResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("dbnode_snapshot_id", *s.Res.Id)
	}

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	mountPoints := []interface{}{}
	for _, item := range s.Res.MountPoints {
		mountPoints = append(mountPoints, MountPointDetailsToMap(item))
	}
	s.D.Set("mount_points", mountPoints)

	if len(mountPoints) > 0 {
		mountPoint0 := mountPoints[0].(map[string]interface{})
		if tmp, exist := mountPoint0["db_node_id"]; exist {
			dbNodeId := tmp.(string)
			s.D.Set("mount_dbnode_id", dbNodeId)
		}
	} else {
		s.D.Set("mount_dbnode_id", "null")
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SourceDbnodeId != nil {
		s.D.Set("source_dbnode_id", *s.Res.SourceDbnodeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	volumes := []interface{}{}
	for _, item := range s.Res.Volumes {
		volumes = append(volumes, VolumeDetailsToMap(item))
	}
	s.D.Set("volumes", volumes)

	return nil
}

func (s *DatabaseDbnodeSnapshotResourceCrud) MountDbNodeSnapshot(mountDbNodeId string) error {
	request := oci_database.MountDbnodeSnapshotRequest{}

	dbNodeSnapshotId := s.D.Id()
	request.DbnodeSnapshotId = &dbNodeSnapshotId

	details := oci_database.MountDbnodeSnapshotDetails{}
	details.DbNodeId = &mountDbNodeId
	request.MountDbnodeSnapshotDetails = details

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.MountDbnodeSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbnodesnapshot", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseDbnodeSnapshotResourceCrud) UnmountDbNodeSnapshot(unmountDbNodeId string) error {
	request := oci_database.UnmountDbnodeSnapshotRequest{}

	dbNodeSnapshotId := s.D.Id()
	request.DbnodeSnapshotId = &dbNodeSnapshotId

	details := oci_database.UnmountDbnodeSnapshotDetails{}
	details.DbNodeId = &unmountDbNodeId
	request.UnmountDbnodeSnapshotDetails = details

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UnmountDbnodeSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbnodesnapshot", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func stringMustBeDbNodeIdOrNull() func(i interface{}, k string) (warnings []string, errors []error) {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		lv := strings.ToLower(v)
		if !strings.HasPrefix(lv, "ocid1.dbnode.") && lv != "null" {
			errors = append(errors, fmt.Errorf("expected value of %s to be either a dbnode id or 'null', got '%v'", k, i))
			return warnings, errors
		}

		if strings.ContainsAny(v, " ") {
			errors = append(errors, fmt.Errorf("expected value of %s to not contain any space, got '%v'", k, i))
			return warnings, errors
		}

		return warnings, errors
	}
}
