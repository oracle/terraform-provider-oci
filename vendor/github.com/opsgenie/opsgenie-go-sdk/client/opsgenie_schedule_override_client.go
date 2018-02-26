package client

import (
	"errors"
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/logging"
	"github.com/opsgenie/opsgenie-go-sdk/scheduleoverride"
)

const (
	scheduleOverrideURL = "/v1/json/schedule/override"
)

// OpsGenieScheduleOverrideClient is the data type to make Schedule API requests.
type OpsGenieScheduleOverrideClient struct {
	OpsGenieClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieScheduleOverrideClient.
func (cli *OpsGenieScheduleOverrideClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// Add method adds a schedule override at OpsGenie.
func (cli *OpsGenieScheduleOverrideClient) Add(req override.AddScheduleOverrideRequest) (*override.AddScheduleOverrideResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(scheduleOverrideURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var addScheduleOverrideResp override.AddScheduleOverrideResponse

	if err = resp.Body.FromJsonTo(&addScheduleOverrideResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &addScheduleOverrideResp, nil
}

// Update method updates a schedule override at OpsGenie.
func (cli *OpsGenieScheduleOverrideClient) Update(req override.UpdateScheduleOverrideRequest) (*override.UpdateScheduleOverrideResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildPostRequest(scheduleOverrideURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updateScheduleOverrideResp override.UpdateScheduleOverrideResponse

	if err = resp.Body.FromJsonTo(&updateScheduleOverrideResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &updateScheduleOverrideResp, nil
}

// Delete method deletes a schedule override at OpsGenie.
func (cli *OpsGenieScheduleOverrideClient) Delete(req override.DeleteScheduleOverrideRequest) (*override.DeleteScheduleOverrideResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildDeleteRequest(scheduleOverrideURL, req))

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deleteScheduleOverrideResp override.DeleteScheduleOverrideResponse

	if err = resp.Body.FromJsonTo(&deleteScheduleOverrideResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	return &deleteScheduleOverrideResp, nil
}

// Get method retrieves specified schedule override details from OpsGenie.
func (cli *OpsGenieScheduleOverrideClient) Get(req override.GetScheduleOverrideRequest) (*override.GetScheduleOverrideResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(scheduleOverrideURL, req))
	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()
	var getScheduleOverrideResp override.GetScheduleOverrideResponse

	if err = resp.Body.FromJsonTo(&getScheduleOverrideResp); err != nil {
		fmt.Println("Error parsing json")
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}
	fmt.Printf("%+v", getScheduleOverrideResp)
	return &getScheduleOverrideResp, nil
}

// List method retrieves schedule overrides from OpsGenie.
func (cli *OpsGenieScheduleOverrideClient) List(req override.ListScheduleOverridesRequest) (*override.ListScheduleOverridesResponse, error) {
	req.APIKey = cli.apiKey
	resp, err := cli.sendRequest(cli.buildGetRequest(scheduleOverrideURL, req))

	if resp == nil {
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()

	var listScheduleOverridesResp override.ListScheduleOverridesResponse

	if err = resp.Body.FromJsonTo(&listScheduleOverridesResp); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return nil, errors.New(message)
	}

	return &listScheduleOverridesResp, nil
}
