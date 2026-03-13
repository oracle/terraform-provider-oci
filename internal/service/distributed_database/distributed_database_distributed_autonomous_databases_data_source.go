// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDistributedDatabaseDistributedAutonomousDatabasesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//"metadata": {
			// WORKAROUND / FIX REQUIRED FOR GENERATED CODE ISSUE:
			//
			// Terraform provider internal validation fails with:
			//   "metadata: Elem must be set for lists"
			//   "metadata: One of optional, required, or computed must be set"
			//
			// Root cause:
			// The code generator emitted an invalid schema definition for the `metadata`
			// field in multiple Distributed Database data sources:
			//
			//   - metadata is defined as TypeList but missing Elem
			//   - metadata has no Optional / Required / Computed flag set
			//
			// Terraform schema rules require:
			//   - TypeList / TypeSet MUST define Elem
			//   - Every schema field MUST specify exactly one of:
			//       Optional, Required, or Computed
			//
			// Because of this, Terraform fails during InternalValidate *before*
			// any user configuration is evaluated, making this a provider-side bug.
			//
			// Correct schema shape must be:
			//
			//   "metadata": {
			//       Type:     schema.TypeList,
			//       Computed: true,
			//       Elem: &schema.Resource{
			//           Schema: <metadata schema>
			//       },
			//   }
			//
			// Affected data sources:
			//   - oci_distributed_database_distributed_autonomous_databases
			//   - oci_distributed_database_distributed_autonomous_database
			//   - oci_distributed_database_distributed_databases
			//   - oci_distributed_database_distributed_database
			//
			// This must be fixed in the code generator to avoid recurring regressions.
			//
			// See JIRA: TOP-9438
			//Type:     schema.TypeList,
			//Type:     schema.TypeString,
			//Elem:     &schema.Schema{Type: schema.TypeString},
			/*Type:     schema.TypeString,
				Optional: true,
			},*/
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"map": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distributed_autonomous_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DistributedDatabaseDistributedAutonomousDatabaseResource()),
						},
					},
				},
			},
		},
	}
}

func readDistributedDatabaseDistributedAutonomousDatabasesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedAutonomousDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedAutonomousDbServiceClient
	Res    *oci_distributed_database.ListDistributedAutonomousDatabasesResponse
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.ListDistributedAutonomousDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbDeploymentType, ok := s.D.GetOkExists("db_deployment_type"); ok {
		request.DbDeploymentType = oci_distributed_database.DistributedAutonomousDatabaseDbDeploymentTypeEnum(dbDeploymentType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The data source treats 'metadata' as a list and calls an undefined helper (mapTostring),
	// but GetDistributedAutonomousDatabaseRequest.Metadata is a *string in the OCI Go SDK.
	// Fix: assign the schema string directly to request.Metadata; do not index into a list
	// or call a non-existent mapper.
	// See JIRA: TOP-9424

	/*if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapTostring(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}*/

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.ListDistributedAutonomousDatabases(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDistributedAutonomousDatabases(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedAutonomousDatabasesDataSource-", DistributedDatabaseDistributedAutonomousDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	distributedAutonomousDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DistributedAutonomousDatabaseSummaryToMap(item))
	}
	distributedAutonomousDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DistributedDatabaseDistributedAutonomousDatabasesDataSource().Schema["distributed_autonomous_database_collection"].Elem.(*schema.Resource).Schema)
		distributedAutonomousDatabase["items"] = items
	}

	resources = append(resources, distributedAutonomousDatabase)
	if err := s.D.Set("distributed_autonomous_database_collection", resources); err != nil {
		return err
	}

	return nil
}
