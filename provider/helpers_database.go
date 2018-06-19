package provider

import (
	"net/http"
	"strconv"
	"strings"

	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/database"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func handleDbSimulationFlag(client *database.DatabaseClient) {
	// Internal, not intended for public use.
	// This flag allows faster testing but requires a whitelisted tenancy to use.
	// To use set environment variable: simulate_db=true
	simulateDb, _ := strconv.ParseBool(getEnvSetting("simulate_db", "false"))
	if simulateDb {
		client.Interceptor = func(r *http.Request) error {
			if r.Method == http.MethodPost && strings.Contains(r.URL.Path, "/dbSystems") {
				r.Header.Set("opc-host-serial", "FAKEHOSTSERIAL")
			}
			return nil
		}
	}
}

func readDatabaseFromState(data *schema.ResourceData) (ID string, databaseObj map[string]interface{}, err error) {
	_, dbHomeObj, err := readDBHomeFromState(data)
	if err != nil {
		return
	}

	if objs, ok := dbHomeObj["database"]; !ok {
		err = fmt.Errorf("database object is missing from state can not read database information")
		return
	} else {
		databaseObj = objs.([]interface{})[0].(map[string]interface{})
	}

	if id, ok := databaseObj["id"]; !ok {
		err = fmt.Errorf("database id is missing from from state can not read database information")
		return
	} else {
		ID = id.(string)
	}

	return
}

func readDBHomeFromState(data *schema.ResourceData) (ID string, dbHomeObj map[string]interface{}, err error) {
	currHome, ok := data.GetOk("db_home")
	if !ok {
		err = fmt.Errorf("state missing db_home object, can not read db homes nor databases")
		return
	}

	dbHomeObjs, ok := currHome.([]interface{})
	if !ok {
		err = fmt.Errorf("db home state is in the wrong format, can not read db homes nor database")
		return
	}

	dbHomeObj = dbHomeObjs[0].(map[string]interface{})
	dbHomeID, ok := dbHomeObj["id"]
	if !ok {
		err = fmt.Errorf("id is missing from db_home object, can not read db homes or databases")
		return
	}

	ID = dbHomeID.(string)
	return
}

func getDBHomesByDBSystem(client *oci_database.DatabaseClient, dbSystemID *string, compartmentID *string, disableNotFoundRetries bool) (res []oci_database.DbHome, err error) {
	res = make([]oci_database.DbHome, 0)
	// getDBHomes from DbSummaries
	appendToResultFn := func(summaries []oci_database.DbHomeSummary) error {
		for _, dbSummary := range summaries {
			resGet := oci_database.GetDbHomeRequest{}
			resGet.DbHomeId = dbSummary.Id
			resGet.RequestMetadata.RetryPolicy = getRetryPolicy(disableNotFoundRetries, "database")
			dbHomeResponse, errGet := client.GetDbHome(context.Background(), resGet)
			if errGet != nil {
				return errGet
			}
			res = append(res, dbHomeResponse.DbHome)
		}
		return nil
	}

	//List and paginate through all db homes
	listReq := oci_database.ListDbHomesRequest{DbSystemId: dbSystemID, CompartmentId: compartmentID}
	listReq.RequestMetadata.RetryPolicy = getRetryPolicy(disableNotFoundRetries, "database")
	for r, listErr := client.ListDbHomes(context.Background(), listReq); ; r, listErr = client.ListDbHomes(context.Background(), listReq) {
		if listErr != nil {
			err = listErr
			return
		}

		if err = appendToResultFn(r.Items); err != nil {
			return
		}

		//More pages
		if r.OpcNextPage != nil {
			listReq.Page = r.OpcNextPage
		} else {
			break
		}
	}
	return
}

func getDatabasesByDBHome(client *oci_database.DatabaseClient, dbHomeID *string, compartmentID *string, disableNotFoundRetries bool) (res *oci_database.Database, err error) {
	listReq := oci_database.ListDatabasesRequest{DbHomeId: dbHomeID, CompartmentId: compartmentID}
	listReq.RequestMetadata.RetryPolicy = getRetryPolicy(disableNotFoundRetries, "database")

	//No pagination necessary only one db per db home
	r, err := client.ListDatabases(context.Background(), listReq)
	if err != nil {
		return
	}

	if len(r.Items) != 1 {
		err = fmt.Errorf("no databases found in db home: %s", *dbHomeID)
		return
	}

	dbSummary := r.Items[0]
	getRequest := oci_database.GetDatabaseRequest{DatabaseId: dbSummary.Id}
	getRequest.RequestMetadata.RetryPolicy = getRetryPolicy(disableNotFoundRetries, "database")
	getRes, err := client.GetDatabase(context.Background(), getRequest)
	if err != nil {
		return
	}
	res = &getRes.Database
	return
}

func getDatabaseUpdateRetryPolicy(disableNotFoundRetries bool) *oci_common.RetryPolicy {
	retryPolicy := &oci_common.RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if shouldRetry(response, disableNotFoundRetries, "database") {
				return true
			}

			var updateRes database.UpdateDatabaseResponse
			var ok bool
			if updateRes, ok = response.Response.(oci_database.UpdateDatabaseResponse); !ok || response.Error != nil {
				return false
			}

			return updateRes.Database.LifecycleState == oci_database.DatabaseLifecycleStateUpdating
		},
		NextDuration: nextDuration,
	}

	return retryPolicy

}
