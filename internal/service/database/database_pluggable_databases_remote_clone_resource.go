// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabasePluggableDatabasesRemoteCloneResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabasePluggableDatabasesRemoteClone,
		Read:     readDatabasePluggableDatabasesRemoteClone,
		Delete:   deleteDatabasePluggableDatabasesRemoteClone,
		Schema: map[string]*schema.Schema{
			// Required
			"cloned_pdb_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_container_db_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"target_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"pdb_admin_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"should_pdb_admin_account_be_locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"target_tde_wallet_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
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
			"container_database_id": {
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
			"pdb_name": {
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

func createDatabasePluggableDatabasesRemoteClone(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabasesRemoteCloneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabasePluggableDatabasesRemoteClone(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabasePluggableDatabasesRemoteClone(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabasePluggableDatabasesRemoteCloneResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.PluggableDatabase
	DisableNotFoundRetries bool
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminating),
		string(oci_database.PluggableDatabaseLifecycleStateDisabled),
	}
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) Create() error {
	request := oci_database.RemoteClonePluggableDatabaseRequest{}

	if clonedPdbName, ok := s.D.GetOkExists("cloned_pdb_name"); ok {
		tmp := clonedPdbName.(string)
		request.ClonedPdbName = &tmp
	}

	if pdbAdminPassword, ok := s.D.GetOkExists("pdb_admin_password"); ok {
		tmp := pdbAdminPassword.(string)
		request.PdbAdminPassword = &tmp
	}

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	if shouldPdbAdminAccountBeLocked, ok := s.D.GetOkExists("should_pdb_admin_account_be_locked"); ok {
		tmp := shouldPdbAdminAccountBeLocked.(bool)
		request.ShouldPdbAdminAccountBeLocked = &tmp
	}

	if sourceContainerDbAdminPassword, ok := s.D.GetOkExists("source_container_db_admin_password"); ok {
		tmp := sourceContainerDbAdminPassword.(string)
		request.SourceContainerDbAdminPassword = &tmp
	}

	if targetContainerDatabaseId, ok := s.D.GetOkExists("target_container_database_id"); ok {
		tmp := targetContainerDatabaseId.(string)
		request.TargetContainerDatabaseId = &tmp
	}

	if targetTdeWalletPassword, ok := s.D.GetOkExists("target_tde_wallet_password"); ok {
		tmp := targetTdeWalletPassword.(string)
		request.TargetTdeWalletPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RemoteClonePluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabase
	return nil
}

func (s *DatabasePluggableDatabasesRemoteCloneResourceCrud) SetData() error {
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
