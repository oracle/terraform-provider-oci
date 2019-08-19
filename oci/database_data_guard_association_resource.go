// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDataGuardAssociationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: &TwoHours,
			Delete: &TwoHours,
			Update: &TwoHours,
		},
		Create: createDatabaseDataGuardAssociation,
		Read:   readDatabaseDataGuardAssociation,
		Update: updateDatabaseDataGuardAssociation,
		Delete: deleteDatabaseDataGuardAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"creation_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ExistingDbSystem",
					"NewDbSystem",
				}, true),
			},
			"database_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"delete_standby_db_home_on_delete": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transport_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"backup_network_nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"peer_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"apply_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_data_guard_association_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
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

func createDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return DeleteResource(d, sync)
}

type DatabaseDataGuardAssociationResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DataGuardAssociation
	DisableNotFoundRetries bool
}

//after deleting the dataguard we should wait a little bit of time because the dbSystem sometime goes from terminating to available if deleted right after the standby dbHome is deleted and the dataguard is deleted.
func (s *DatabaseDataGuardAssociationResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return 10 * time.Millisecond
	}

	return time.Second * 30
}

func (s *DatabaseDataGuardAssociationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDataGuardAssociationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateProvisioning),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateAvailable),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateAvailable),
		string(oci_database.DataGuardAssociationLifecycleStateTerminating),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateTerminated),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) Create() error {
	request := oci_database.CreateDataGuardAssociationRequest{}
	err := s.populateTopLevelPolymorphicCreateDataGuardAssociationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataGuardAssociation
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) Get() error {
	request := oci_database.GetDataGuardAssociationRequest{}

	tmp := s.D.Id()
	request.DataGuardAssociationId = &tmp

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataGuardAssociation
	return nil
}

//We need to have an empty delete because otherwise the added delete_standby_db_home_on_delete would have to be marked as ForceNew, which is undesireable. This way the update will pass the new value for this property from the config to the statefile.
func (s *DatabaseDataGuardAssociationResourceCrud) Update() error {
	return s.Get()
}

func (s *DatabaseDataGuardAssociationResourceCrud) Delete() error {
	if deleteStandbyDbHomeOnDelete, ok := s.D.GetOkExists("delete_standby_db_home_on_delete"); ok {
		tmp := deleteStandbyDbHomeOnDelete.(string)
		if tmp != "true" {
			return fmt.Errorf("we do not currently support deleting the dataguard association without deleting the standby dbHome. Please set delete_standby_db_home_on_delete to \"true\" if you want to continue with the destroy. Once you change the value of delete_standby_db_home_on_delete you must do a `terraform apply` before running a `terraform destroy` so that destroy operation would get the new value")
		}
	}

	err := s.Get()
	if err != nil {
		return err
	}

	creationType, ok := s.D.GetOkExists("creation_type")
	if !ok {
		return fmt.Errorf("creation_type could not be established during the delete")
	}
	if strings.ToLower(creationType.(string)) == strings.ToLower("ExistingDbSystem") {
		var standbyDbHomeId *string
		if s.Res.PeerRole == oci_database.DataGuardAssociationPeerRoleStandby {
			standbyDbHomeId = s.Res.PeerDbHomeId
		} else if s.Res.Role == oci_database.DataGuardAssociationRoleStandby {
			standbyDbHomeId, err = s.GetDbHomeIdFromDatabaseId(s.Res.DatabaseId)
			if err != nil {
				return fmt.Errorf("could not delete the dataguard association as the standby DB Home Id could not be obtained: %v", err)
			}
		} else {
			return fmt.Errorf("could not delete the dataguard association as it is not possible to determine the standby DB home")
		}

		if standbyDbHomeId == nil {
			return fmt.Errorf("could not delete the dataguard association as the standby DB Home Id could not be obtained")
		}

		request := oci_database.DeleteDbHomeRequest{}
		request.DbHomeId = standbyDbHomeId
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
		_, err = s.Client.DeleteDbHome(context.Background(), request)
		if err != nil {
			return fmt.Errorf("could not delete standby DB Home to delete the data guard association: %v", err)
		}

		getDbHomeRequest := oci_database.GetDbHomeRequest{}
		getDbHomeRequest.DbHomeId = standbyDbHomeId
		getDbHomeRequest.RequestMetadata.RetryPolicy = waitForDbHomeToTerminateRetryPolicy(2 * time.Hour)
		getDbHomeResponse, err := s.Client.GetDbHome(context.Background(), getDbHomeRequest)
		//Sometimed the DB Home goes from a TERMINATING Lifecycle state into AVAILABLE, this is to catch that case and report it ot the user as a failure in deleting the database
		if getDbHomeResponse.LifecycleState == oci_database.DbHomeLifecycleStateAvailable {
			return fmt.Errorf("could not delete the dataguard association as the dbHome could not be deleted")
		}

		return err
	} else if strings.ToLower(creationType.(string)) == strings.ToLower("NewDbSystem") {
		var standbyDbSystemId *string
		if s.Res.PeerRole == oci_database.DataGuardAssociationPeerRoleStandby {
			standbyDbSystemId = s.Res.PeerDbSystemId
		} else if s.Res.Role == oci_database.DataGuardAssociationRoleStandby {
			standbyDbSystemId, err = s.GetDbSystemIdFromDatabaseId(s.Res.DatabaseId)
			if err != nil {
				return fmt.Errorf("could not delete the dataguard association as the standby DB System Id could not be obtained: %v", err)
			}
		} else {
			return fmt.Errorf("could not delete the dataguard association as it is not possible to determine the standby DB System")
		}

		if standbyDbSystemId == nil {
			return fmt.Errorf("could not delete the dataguard association as the standby DB System Id could not be obtained")
		}

		request := oci_database.TerminateDbSystemRequest{}
		request.DbSystemId = standbyDbSystemId
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
		_, err := s.Client.TerminateDbSystem(context.Background(), request)
		if err != nil {
			return fmt.Errorf("could not delete standby DB System to delete the data guard association: %v", err)
		}

		getDbSystemRequest := oci_database.GetDbSystemRequest{}
		getDbSystemRequest.DbSystemId = standbyDbSystemId
		getDbSystemRequest.RequestMetadata.RetryPolicy = waitForDbSystemToTerminateRetryPolicy(2 * time.Hour)
		getDbSystemResponse, err := s.Client.GetDbSystem(context.Background(), getDbSystemRequest)
		//Sometimes the DB System goes from a TERMINATING Lifecycle state into AVAILABLE, this is to catch that case and report it ot the user as a failure in deleting the db System
		if getDbSystemResponse.LifecycleState == oci_database.DbSystemLifecycleStateAvailable {
			return fmt.Errorf("could not delete the dataguard association as the dbSystem could not be deleted")
		}

		return err
	}
	return fmt.Errorf("unrecognized creation_type during delete")
}

func (s *DatabaseDataGuardAssociationResourceCrud) SetData() error {
	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be set in state as the feature may or may not be supported
	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		s.D.Set("backup_network_nsg_ids", backupNetworkNsgIds)
	} else {
		s.D.Set("backup_network_nsg_ids", nil)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}
	//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be set in state as the feature may or may not be supported
	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		s.D.Set("nsg_ids", nsgIds)
	} else {
		s.D.Set("nsg_ids", nil)
	}

	if s.Res.PeerDataGuardAssociationId != nil {
		s.D.Set("peer_data_guard_association_id", *s.Res.PeerDataGuardAssociationId)
	}

	if s.Res.PeerDatabaseId != nil {
		s.D.Set("peer_database_id", *s.Res.PeerDatabaseId)
	}

	if s.Res.PeerDbHomeId != nil {
		s.D.Set("peer_db_home_id", *s.Res.PeerDbHomeId)
	}

	if s.Res.PeerDbSystemId != nil {
		s.D.Set("peer_db_system_id", *s.Res.PeerDbSystemId)
	}

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("transport_type", s.Res.TransportType)

	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) populateTopLevelPolymorphicCreateDataGuardAssociationRequest(request *oci_database.CreateDataGuardAssociationRequest) error {
	//discriminator
	creationTypeRaw, ok := s.D.GetOkExists("creation_type")
	var creationType string
	if ok {
		creationType = creationTypeRaw.(string)
	} else {
		creationType = "" // default value
	}
	switch strings.ToLower(creationType) {
	case strings.ToLower("ExistingDbSystem"):
		details := oci_database.CreateDataGuardAssociationToExistingDbSystemDetails{}
		if peerDbSystemId, ok := s.D.GetOkExists("peer_db_system_id"); ok {
			tmp := peerDbSystemId.(string)
			details.PeerDbSystemId = &tmp
		}
		if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
			tmp := databaseAdminPassword.(string)
			details.DatabaseAdminPassword = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
			details.ProtectionMode = oci_database.CreateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists("transport_type"); ok {
			details.TransportType = oci_database.CreateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
		}
		request.CreateDataGuardAssociationDetails = details
	case strings.ToLower("NewDbSystem"):
		details := oci_database.CreateDataGuardAssociationWithNewDbSystemDetails{}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be supplied as the feature may or may not be supported
		if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
			set := backupNetworkNsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.BackupNetworkNsgIds = tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be supplied as the feature may or may not be supported
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.NsgIds = tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
			tmp := databaseAdminPassword.(string)
			details.DatabaseAdminPassword = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
			details.ProtectionMode = oci_database.CreateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists("transport_type"); ok {
			details.TransportType = oci_database.CreateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
		}
		request.CreateDataGuardAssociationDetails = details
	default:
		return fmt.Errorf("unknown creation_type '%v' was specified", creationType)
	}
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbHomeIdFromDatabaseId(databaseId *string) (*string, error) {
	request := oci_database.GetDatabaseRequest{}

	request.DatabaseId = databaseId

	request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.DbHomeId, nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbSystemIdFromDatabaseId(databaseId *string) (*string, error) {
	dbHomeId, err := s.GetDbHomeIdFromDatabaseId(databaseId)
	if err != nil {
		return dbHomeId, err
	}
	return s.GetDbSystemIdFromDbHomeId(dbHomeId)
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbSystemIdFromDbHomeId(dbHomeId *string) (*string, error) {
	request := oci_database.GetDbHomeRequest{}

	request.DbHomeId = dbHomeId

	request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

	response, err := s.Client.GetDbHome(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.DbSystemId, nil
}

func waitForDbHomeToTerminateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()
	// wait for peering status to not be Pending
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if shouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDbHomeResponse, ok := response.Response.(oci_database.GetDbHomeResponse); ok {
				if getDbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateTerminated && getDbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateAvailable {
					timeWaited := getElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return getRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func waitForDbSystemToTerminateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()
	// wait for peering status to not be Pending
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if shouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDbSystemResponse, ok := response.Response.(oci_database.GetDbSystemResponse); ok {
				if getDbSystemResponse.LifecycleState != oci_database.DbSystemLifecycleStateTerminated && getDbSystemResponse.LifecycleState != oci_database.DbSystemLifecycleStateAvailable {
					timeWaited := getElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return getRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}
