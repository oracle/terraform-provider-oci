// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"
)

func DatabaseBackupDestinationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseBackupDestination,
		Read:     readDatabaseBackupDestination,
		Update:   updateDatabaseBackupDestination,
		Delete:   deleteDatabaseBackupDestination,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"NFS",
					"RECOVERY_APPLIANCE",
				}, true),
			},

			// Optional
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"local_mount_point_path": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("local_mount_point_path", "local_mount_point_path under mount_type_details"),
			},
			"mount_type_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"mount_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AUTOMATED_MOUNT",
								"SELF_MOUNT",
							}, true),
						},

						// Optional
						"local_mount_point_path": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"local_mount_point_path"},
						},
						"nfs_server": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"nfs_server_export": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"vpc_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"associated_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nfs_mount_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nfs_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nfs_server_export": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseBackupDestinationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.BackupDestination
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseBackupDestinationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseBackupDestinationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatabaseBackupDestinationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.BackupDestinationLifecycleStateActive),
	}
}

func (s *DatabaseBackupDestinationResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseBackupDestinationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.BackupDestinationLifecycleStateDeleted),
	}
}

func (s *DatabaseBackupDestinationResourceCrud) Create() error {
	request := oci_database.CreateBackupDestinationRequest{}
	err := s.populateTopLevelPolymorphicCreateBackupDestinationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateBackupDestination(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackupDestination
	return nil
}

func (s *DatabaseBackupDestinationResourceCrud) Get() error {
	request := oci_database.GetBackupDestinationRequest{}

	tmp := s.D.Id()
	request.BackupDestinationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetBackupDestination(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackupDestination
	return nil
}

func (s *DatabaseBackupDestinationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateBackupDestinationRequest{}

	tmp := s.D.Id()
	request.BackupDestinationId = &tmp

	if connectionString, ok := s.D.GetOkExists("connection_string"); ok && s.D.HasChange("connection_string") {
		tmp := connectionString.(string)
		request.ConnectionString = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if localMountPointPath, ok := s.D.GetOkExists("local_mount_point_path"); ok && s.D.HasChange("local_mount_point_path") {
		tmp := localMountPointPath.(string)
		request.LocalMountPointPath = &tmp
	}

	if vpcUsers, ok := s.D.GetOkExists("vpc_users"); ok && s.D.HasChange("vpc_users") {
		request.VpcUsers = []string{}
		interfaces := vpcUsers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.VpcUsers = tmp
	}

	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mount_type_details", 0)
	err := s.mapMountTypeDetailsForUpdate(fieldKeyFormat, &request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateBackupDestination(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackupDestination
	return nil
}

func (s *DatabaseBackupDestinationResourceCrud) Delete() error {
	request := oci_database.DeleteBackupDestinationRequest{}

	tmp := s.D.Id()
	request.BackupDestinationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteBackupDestination(context.Background(), request)
	return err
}

func (s *DatabaseBackupDestinationResourceCrud) SetData() error {
	associatedDatabases := []interface{}{}
	for _, item := range s.Res.AssociatedDatabases {
		associatedDatabases = append(associatedDatabases, AssociatedDatabaseDetailsToMap(item))
	}
	s.D.Set("associated_databases", associatedDatabases)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	mountTypeDetails := []interface{}{}
	mountTypeDetails = append(mountTypeDetails, MountTypeDetailsToMap(s.Res))

	s.D.Set("mount_type_details", mountTypeDetails)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LocalMountPointPath != nil {
		s.D.Set("local_mount_point_path", *s.Res.LocalMountPointPath)
	}

	s.D.Set("nfs_mount_type", s.Res.NfsMountType)

	s.D.Set("nfs_server", s.Res.NfsServer)

	if s.Res.NfsServerExport != nil {
		s.D.Set("nfs_server_export", *s.Res.NfsServerExport)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VpcUsers != nil {
		s.D.Set("vpc_users", s.Res.VpcUsers)
	}

	return nil
}

func AssociatedDatabaseDetailsToMap(obj oci_database.AssociatedDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *DatabaseBackupDestinationResourceCrud) mapToMountTypeDetails(fieldKeyFormat string) (oci_database.MountTypeDetails, error) {
	var baseObject oci_database.MountTypeDetails
	//discriminator
	mountTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_type"))
	var mountType string
	if ok {
		mountType = mountTypeRaw.(string)
	} else {
		mountType = "SELF_MOUNT" // default value
	}
	switch strings.ToLower(mountType) {
	case strings.ToLower("AUTOMATED_MOUNT"):
		details := oci_database.AutomatedMountDetails{}
		if nfsServer, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nfs_server")); ok {
			interfaces := nfsServer.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nfs_server")) {
				details.NfsServer = tmp
			}
		}
		if nfsServerExport, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nfs_server_export")); ok {
			tmp := nfsServerExport.(string)
			details.NfsServerExport = &tmp
		}
		baseObject = details
	case strings.ToLower("SELF_MOUNT"):
		details := oci_database.SelfMountDetails{}
		if localMountPointPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "local_mount_point_path")); ok {
			tmp := localMountPointPath.(string)
			details.LocalMountPointPath = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown mount_type '%v' was specified", mountType)
	}
	return baseObject, nil
}

func MountTypeDetailsToMap(obj *oci_database.BackupDestination) map[string]interface{} {
	result := map[string]interface{}{}

	switch obj.NfsMountType {
	case oci_database.BackupDestinationNfsMountTypeAutomatedMount:
		result["mount_type"] = "AUTOMATED_MOUNT"

		if obj.NfsServer != nil {
			result["nfs_server"] = obj.NfsServer
		}

		if obj.NfsServerExport != nil {
			result["nfs_server_export"] = *obj.NfsServerExport
		}
	case oci_database.BackupDestinationNfsMountTypeSelfMount:
		result["mount_type"] = "SELF_MOUNT"

		if obj.LocalMountPointPath != nil {
			result["local_mount_point_path"] = *obj.LocalMountPointPath
		}
	default:
		log.Printf("[WARN] Received 'mount_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseBackupDestinationResourceCrud) mapMountTypeDetailsForUpdate(fieldKeyFormat string, request *oci_database.UpdateBackupDestinationRequest) error {
	mountTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_type"))
	var mountType string
	if ok {
		mountType = mountTypeRaw.(string)
	} else {
		mountType = "SELF_MOUNT" // default value
	}

	request.NfsMountType = oci_database.UpdateBackupDestinationDetailsNfsMountTypeEnum(mountType)

	switch strings.ToLower(mountType) {
	case strings.ToLower("AUTOMATED_MOUNT"):
		if nfsServer, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nfs_server")); ok {
			interfaces := nfsServer.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nfs_server")) {
				request.NfsServer = tmp
			}
		}

		if nfsServerExport, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nfs_server_export")); ok {
			tmp := nfsServerExport.(string)
			request.NfsServerExport = &tmp
		}
	case strings.ToLower("SELF_MOUNT"):
		if localMountPointPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "local_mount_point_path")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "local_mount_point_path")) {
			tmp := localMountPointPath.(string)
			request.LocalMountPointPath = &tmp
		}
	default:
		return fmt.Errorf("unknown mount_type '%v' was specified", mountType)
	}
	return nil
}

func (s *DatabaseBackupDestinationResourceCrud) populateTopLevelPolymorphicCreateBackupDestinationRequest(request *oci_database.CreateBackupDestinationRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("NFS"):
		details := oci_database.CreateNfsBackupDestinationDetails{}
		if localMountPointPath, ok := s.D.GetOkExists("local_mount_point_path"); ok {
			tmp := localMountPointPath.(string)
			details.LocalMountPointPath = &tmp
		}
		if mountTypeDetails, ok := s.D.GetOkExists("mount_type_details"); ok {
			if tmpList := mountTypeDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mount_type_details", 0)
				tmp, err := s.mapToMountTypeDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.MountTypeDetails = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBackupDestinationDetails = details
	case strings.ToLower("RECOVERY_APPLIANCE"):
		details := oci_database.CreateRecoveryApplianceBackupDestinationDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if vpcUsers, ok := s.D.GetOkExists("vpc_users"); ok {
			interfaces := vpcUsers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("vpc_users") {
				details.VpcUsers = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBackupDestinationDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseBackupDestinationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeBackupDestinationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BackupDestinationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeBackupDestinationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "backupDestination", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
