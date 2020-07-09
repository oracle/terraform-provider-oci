package oci

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v27/database"
)

func init() {
	RegisterResource("oci_database_autonomous_container_database_dataguard_association_operation", DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource())
}

func DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("12h"),
			Delete: getTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
		Read:   readDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
		Delete: deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"operation": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationOperation struct {
	// The OCID of the Data Safe private endpoint.
	Id *string `mandatory:"true" json:"id"`
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud) ID() string {
	return string(hashcode.String(s.D.Get("autonomous_container_database_id").(string)))
}

func createDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud) Get() error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud) SetData() error {
	return nil
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	DisableNotFoundRetries bool
	Res                    *DatabaseAutonomousContainerDatabaseDataguardAssociationOperation
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud) Create() error {
	return s.dataguardOperation()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationCrud) dataguardOperation() error {
	dataguardAssociationId := ""
	tmpId := ""
	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmpId = autonomousContainerDatabaseId.(string)
	}
	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		dataguardAssociationId = autonomousContainerDatabaseDataguardAssociationId.(string)
	}
	if operation, ok := s.D.GetOkExists("operation"); ok {
		if strings.ToLower(operation.(string)) == "switchover" {
			switchoverRequest := oci_database.SwitchoverAutonomousContainerDatabaseDataguardAssociationRequest{}
			switchoverRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			switchoverRequest.AutonomousContainerDatabaseId = &tmpId
			switchoverRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
			switchoverRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			_, err := s.Client.SwitchoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), switchoverRequest)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(operation.(string)) == "failover" {
			failoverRequest := oci_database.FailoverAutonomousContainerDatabaseDataguardAssociationRequest{}
			failoverRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			failoverRequest.AutonomousContainerDatabaseId = &tmpId
			failoverRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
			failoverRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			_, err := s.Client.FailoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), failoverRequest)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(operation.(string)) == "reinstate" {
			reinstateRequest := oci_database.ReinstateAutonomousContainerDatabaseDataguardAssociationRequest{}
			reinstateRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			reinstateRequest.AutonomousContainerDatabaseId = &tmpId
			reinstateRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")
			reinstateRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			_, err := s.Client.ReinstateAutonomousContainerDatabaseDataguardAssociation(context.Background(), reinstateRequest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
