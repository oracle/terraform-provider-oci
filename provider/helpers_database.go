package provider

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/oracle/oci-go-sdk/database"
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
