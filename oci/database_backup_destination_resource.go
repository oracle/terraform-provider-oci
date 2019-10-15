// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseBackupDestinationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"local_mount_point_path": {
				Type:     schema.TypeString,
				Optional: true,
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
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseBackupDestinationResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.BackupDestination
	DisableNotFoundRetries bool
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LocalMountPointPath != nil {
		s.D.Set("local_mount_point_path", *s.Res.LocalMountPointPath)
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
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeBackupDestinationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
