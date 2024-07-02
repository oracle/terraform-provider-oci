// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabasePluggableDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabasePluggableDatabase,
		Read:     readDatabasePluggableDatabase,
		Update:   updateDatabasePluggableDatabase,
		Delete:   deleteDatabasePluggableDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"container_database_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
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
			"pdb_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"pdb_creation_type_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"creation_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"LOCAL_CLONE_PDB",
								"RELOCATE_PDB",
								"REMOTE_CLONE_PDB",
							}, true),
						},
						"source_pluggable_database_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"dblink_user_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"dblink_username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_thin_clone": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"refreshable_clone_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"is_refreshable_clone": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"source_container_database_admin_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},

						// Computed
					},
				},
			},
			"should_create_pdb_backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"should_pdb_admin_account_be_locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tde_wallet_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"convert_to_regular_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"refresh_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rotate_key_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"pdb_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pdb_ip_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_restricted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"open_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_node_level_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"node_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"open_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"pluggable_database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"refreshable_clone_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_refreshable_clone": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
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

func createDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("convert_to_regular_trigger"); ok {
		err := sync.ConvertToRegularPluggableDatabase()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("refresh_trigger"); ok {
		err := sync.RefreshPluggableDatabase()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok {
		err := sync.RotatePluggableDatabaseEncryptionKey()
		if err != nil {
			return err
		}
	}
	return nil

}

func readDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	if _, ok := sync.D.GetOkExists("convert_to_regular_trigger"); ok && sync.D.HasChange("convert_to_regular_trigger") {
		oldRaw, newRaw := sync.D.GetChange("convert_to_regular_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConvertToRegularPluggableDatabase()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("convert_to_regular_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("refresh_trigger"); ok && sync.D.HasChange("refresh_trigger") {
		oldRaw, newRaw := sync.D.GetChange("refresh_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RefreshPluggableDatabase()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("refresh_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("rotate_key_trigger"); ok && sync.D.HasChange("rotate_key_trigger") {
		oldRaw, newRaw := sync.D.GetChange("rotate_key_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RotatePluggableDatabaseEncryptionKey()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("rotate_key_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabasePluggableDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.PluggableDatabase
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	DisableNotFoundRetries bool
}

func (s *DatabasePluggableDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabasePluggableDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminating),
		string(oci_database.PluggableDatabaseLifecycleStateDisabled),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabasePluggableDatabaseResourceCrud) Create() error {
	request := oci_database.CreatePluggableDatabaseRequest{}

	if containerDatabaseAdminPassword, ok := s.D.GetOkExists("container_database_admin_password"); ok {
		tmp := containerDatabaseAdminPassword.(string)
		request.ContainerDatabaseAdminPassword = &tmp
	}

	if containerDatabaseId, ok := s.D.GetOkExists("container_database_id"); ok {
		tmp := containerDatabaseId.(string)
		request.ContainerDatabaseId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if pdbAdminPassword, ok := s.D.GetOkExists("pdb_admin_password"); ok {
		tmp := pdbAdminPassword.(string)
		request.PdbAdminPassword = &tmp
	}

	if pdbCreationTypeDetails, ok := s.D.GetOkExists("pdb_creation_type_details"); ok {
		if tmpList := pdbCreationTypeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "pdb_creation_type_details", 0)
			tmp, err := s.mapToCreatePluggableDatabaseCreationTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PdbCreationTypeDetails = tmp
		}
	}

	if pdbName, ok := s.D.GetOkExists("pdb_name"); ok {
		tmp := pdbName.(string)
		request.PdbName = &tmp
	}

	if shouldCreatePdbBackup, ok := s.D.GetOkExists("should_create_pdb_backup"); ok {
		tmp := shouldCreatePdbBackup.(bool)
		request.ShouldCreatePdbBackup = &tmp
	}

	if shouldPdbAdminAccountBeLocked, ok := s.D.GetOkExists("should_pdb_admin_account_be_locked"); ok {
		tmp := shouldPdbAdminAccountBeLocked.(bool)
		request.ShouldPdbAdminAccountBeLocked = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("tde_wallet_password"); ok {
		tmp := tdeWalletPassword.(string)
		request.TdeWalletPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreatePluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabasePluggableDatabaseResourceCrud) Get() error {
	request := oci_database.GetPluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) Update() error {
	request := oci_database.UpdatePluggableDatabaseRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdatePluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) Delete() error {
	request := oci_database.DeletePluggableDatabaseRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeletePluggableDatabase(context.Background(), request)
	return err
}

func (s *DatabasePluggableDatabaseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{PluggableDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.ContainerDatabaseId != nil {
		s.D.Set("container_database_id", *s.Res.ContainerDatabaseId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRestricted != nil {
		s.D.Set("is_restricted", *s.Res.IsRestricted)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
	}

	pdbNodeLevelDetails := []interface{}{}
	for _, item := range s.Res.PdbNodeLevelDetails {
		pdbNodeLevelDetails = append(pdbNodeLevelDetails, PluggableDatabaseNodeLevelDetailsToMap(item))
	}
	s.D.Set("pdb_node_level_details", pdbNodeLevelDetails)

	if s.Res.PluggableDatabaseManagementConfig != nil {
		s.D.Set("pluggable_database_management_config", []interface{}{PluggableDatabaseManagementConfigToMap(s.Res.PluggableDatabaseManagementConfig)})
	} else {
		s.D.Set("pluggable_database_management_config", nil)
	}

	if s.Res.RefreshableCloneConfig != nil {
		s.D.Set("refreshable_clone_config", []interface{}{PluggableDatabaseRefreshableCloneConfigToMap(s.Res.RefreshableCloneConfig)})
	} else {
		s.D.Set("refreshable_clone_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) ConvertToRegularPluggableDatabase() error {
	request := oci_database.ConvertToRegularPluggableDatabaseRequest{}

	if containerDatabaseAdminPassword, ok := s.D.GetOkExists("container_database_admin_password"); ok {
		tmp := containerDatabaseAdminPassword.(string)
		request.ContainerDatabaseAdminPassword = &tmp
	}

	idTmp := s.D.Id()
	request.PluggableDatabaseId = &idTmp

	if shouldCreatePdbBackup, ok := s.D.GetOkExists("should_create_pdb_backup"); ok {
		tmp := shouldCreatePdbBackup.(bool)
		request.ShouldCreatePdbBackup = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("tde_wallet_password"); ok {
		tmp := tdeWalletPassword.(string)
		request.TdeWalletPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConvertToRegularPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("convert_to_regular_trigger")
	s.D.Set("convert_to_regular_trigger", val)

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) RefreshPluggableDatabase() error {
	request := oci_database.RefreshPluggableDatabaseRequest{}

	idTmp := s.D.Id()
	request.PluggableDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RefreshPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("refresh_trigger")
	s.D.Set("refresh_trigger", val)

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) RotatePluggableDatabaseEncryptionKey() error {
	request := oci_database.RotatePluggableDatabaseEncryptionKeyRequest{}

	idTmp := s.D.Id()
	request.PluggableDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.RotatePluggableDatabaseEncryptionKey(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabasePluggableDatabaseResourceCrud) mapToCreatePluggableDatabaseCreationTypeDetails(fieldKeyFormat string) (oci_database.CreatePluggableDatabaseCreationTypeDetails, error) {
	var baseObject oci_database.CreatePluggableDatabaseCreationTypeDetails
	//discriminator
	creationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "creation_type"))
	var creationType string
	if ok {
		creationType = creationTypeRaw.(string)
	} else {
		creationType = "" // default value
	}
	switch strings.ToLower(creationType) {
	case strings.ToLower("LOCAL_CLONE_PDB"):
		details := oci_database.CreatePluggableDatabaseFromLocalCloneDetails{}
		if isThinClone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_thin_clone")); ok {
			tmp := isThinClone.(bool)
			details.IsThinClone = &tmp
		}
		if sourcePluggableDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_pluggable_database_id")); ok {
			tmp := sourcePluggableDatabaseId.(string)
			details.SourcePluggableDatabaseId = &tmp
		}
		baseObject = details
	case strings.ToLower("RELOCATE_PDB"):
		details := oci_database.CreatePluggableDatabaseFromRelocateDetails{}
		if dblinkUserPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dblink_user_password")); ok {
			tmp := dblinkUserPassword.(string)
			details.DblinkUserPassword = &tmp
		}
		if dblinkUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dblink_username")); ok {
			tmp := dblinkUsername.(string)
			details.DblinkUsername = &tmp
		}
		if sourceContainerDatabaseAdminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_container_database_admin_password")); ok {
			tmp := sourceContainerDatabaseAdminPassword.(string)
			details.SourceContainerDatabaseAdminPassword = &tmp
		}
		if sourcePluggableDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_pluggable_database_id")); ok {
			tmp := sourcePluggableDatabaseId.(string)
			details.SourcePluggableDatabaseId = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOTE_CLONE_PDB"):
		details := oci_database.CreatePluggableDatabaseFromRemoteCloneDetails{}
		if dblinkUserPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dblink_user_password")); ok {
			tmp := dblinkUserPassword.(string)
			details.DblinkUserPassword = &tmp
		}
		if dblinkUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dblink_username")); ok {
			tmp := dblinkUsername.(string)
			details.DblinkUsername = &tmp
		}
		if isThinClone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_thin_clone")); ok {
			tmp := isThinClone.(bool)
			details.IsThinClone = &tmp
		}
		if refreshableCloneDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "refreshable_clone_details")); ok {
			if tmpList := refreshableCloneDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "refreshable_clone_details"), 0)
				tmp, err := s.mapToCreatePluggableDatabaseRefreshableCloneDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert refreshable_clone_details, encountered error: %v", err)
				}
				details.RefreshableCloneDetails = &tmp
			}
		}
		if sourceContainerDatabaseAdminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_container_database_admin_password")); ok {
			tmp := sourceContainerDatabaseAdminPassword.(string)
			details.SourceContainerDatabaseAdminPassword = &tmp
		}
		if sourcePluggableDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_pluggable_database_id")); ok {
			tmp := sourcePluggableDatabaseId.(string)
			details.SourcePluggableDatabaseId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown creation_type '%v' was specified", creationType)
	}
	return baseObject, nil
}

func CreatePluggableDatabaseCreationTypeDetailsToMap(obj *oci_database.CreatePluggableDatabaseCreationTypeDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.CreatePluggableDatabaseFromLocalCloneDetails:
		result["creation_type"] = "LOCAL_CLONE_PDB"

		if v.IsThinClone != nil {
			result["is_thin_clone"] = bool(*v.IsThinClone)
		}

		if v.SourcePluggableDatabaseId != nil {
			result["source_pluggable_database_id"] = string(*v.SourcePluggableDatabaseId)
		}
	case oci_database.CreatePluggableDatabaseFromRelocateDetails:
		result["creation_type"] = "RELOCATE_PDB"

		if v.DblinkUserPassword != nil {
			result["dblink_user_password"] = string(*v.DblinkUserPassword)
		}

		if v.DblinkUsername != nil {
			result["dblink_username"] = string(*v.DblinkUsername)
		}

		if v.SourceContainerDatabaseAdminPassword != nil {
			result["source_container_database_admin_password"] = string(*v.SourceContainerDatabaseAdminPassword)
		}

		if v.SourcePluggableDatabaseId != nil {
			result["source_pluggable_database_id"] = string(*v.SourcePluggableDatabaseId)
		}
	case oci_database.CreatePluggableDatabaseFromRemoteCloneDetails:
		result["creation_type"] = "REMOTE_CLONE_PDB"

		if v.DblinkUserPassword != nil {
			result["dblink_user_password"] = string(*v.DblinkUserPassword)
		}

		if v.DblinkUsername != nil {
			result["dblink_username"] = string(*v.DblinkUsername)
		}

		if v.IsThinClone != nil {
			result["is_thin_clone"] = bool(*v.IsThinClone)
		}

		if v.RefreshableCloneDetails != nil {
			result["refreshable_clone_details"] = []interface{}{CreatePluggableDatabaseRefreshableCloneDetailsToMap(v.RefreshableCloneDetails)}
		}

		if v.SourceContainerDatabaseAdminPassword != nil {
			result["source_container_database_admin_password"] = string(*v.SourceContainerDatabaseAdminPassword)
		}

		if v.SourcePluggableDatabaseId != nil {
			result["source_pluggable_database_id"] = string(*v.SourcePluggableDatabaseId)
		}
	default:
		log.Printf("[WARN] Received 'creation_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabasePluggableDatabaseResourceCrud) mapToCreatePluggableDatabaseRefreshableCloneDetails(fieldKeyFormat string) (oci_database.CreatePluggableDatabaseRefreshableCloneDetails, error) {
	result := oci_database.CreatePluggableDatabaseRefreshableCloneDetails{}

	if isRefreshableClone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_refreshable_clone")); ok {
		tmp := isRefreshableClone.(bool)
		result.IsRefreshableClone = &tmp
	}

	return result, nil
}

func CreatePluggableDatabaseRefreshableCloneDetailsToMap(obj *oci_database.CreatePluggableDatabaseRefreshableCloneDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsRefreshableClone != nil {
		result["is_refreshable_clone"] = bool(*obj.IsRefreshableClone)
	}

	return result
}

func PluggableDatabaseConnectionStringsToMap(obj *oci_database.PluggableDatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.PdbDefault != nil {
		result["pdb_default"] = string(*obj.PdbDefault)
	}

	if obj.PdbIpDefault != nil {
		result["pdb_ip_default"] = string(*obj.PdbIpDefault)
	}

	return result
}

func PluggableDatabaseManagementConfigToMap(obj *oci_database.PluggableDatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["management_status"] = string(obj.ManagementStatus)

	return result
}

func PluggableDatabaseNodeLevelDetailsToMap(obj oci_database.PluggableDatabaseNodeLevelDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NodeName != nil {
		result["node_name"] = string(*obj.NodeName)
	}

	result["open_mode"] = string(obj.OpenMode)

	return result
}

func PluggableDatabaseRefreshableCloneConfigToMap(obj *oci_database.PluggableDatabaseRefreshableCloneConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsRefreshableClone != nil {
		result["is_refreshable_clone"] = bool(*obj.IsRefreshableClone)
	}

	return result
}
