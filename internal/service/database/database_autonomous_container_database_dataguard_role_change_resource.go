package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func DatabaseAutonomousContainerDatabaseDataguardRoleChangeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext,
		ReadContext:   readDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext,
		DeleteContext: deleteDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext,
		Schema: map[string]*schema.Schema{
			// Required

			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"role": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"connection_strings_type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

type DatabaseAutonomousContainerDatabaseDataguardRoleChange struct {
	// The OCID of the Data Safe private endpoint.
	Id *string `mandatory:"true" json:"id"`
}

type DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	DisableNotFoundRetries bool
	Res                    *oci_database.AutonomousContainerDatabase
}

func (s *DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud) ID() string {
	return fmt.Sprint(utils.GetStringHashcode(s.D.Get("autonomous_container_database_id").(string)))
}

func createDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDatabaseAutonomousContainerDatabaseDataguardRoleChangeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud) GetWithContext(ctx context.Context) error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardRoleChangeResourceCrud) CreateWithContext(ctx context.Context) error {

	request := oci_database.ChangeDataguardRoleRequest{}
	details := oci_database.ChangeDataguardRoleDetails{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		details.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if role, ok := s.D.GetOkExists("role"); ok {
		details.Role = oci_database.ChangeDataguardRoleDetailsRoleEnum(role.(string))
	}

	if connectionStringsType, ok := s.D.GetOkExists("connection_strings_type"); ok {
		details.ConnectionStringsType = oci_database.ChangeDataguardRoleDetailsConnectionStringsTypeEnum(connectionStringsType.(string))
	}

	request.ChangeDataguardRoleDetails = details

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeDataguardRole(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "autonomousContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
		}
	}

	s.Res = &response.AutonomousContainerDatabase

	return nil
}
