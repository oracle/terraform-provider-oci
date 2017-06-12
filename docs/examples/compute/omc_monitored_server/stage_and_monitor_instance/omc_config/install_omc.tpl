#!/usr/bin/env bash

#OMC Agent Install
# - Update the entity file with the hostname of this server instance.
cat /omc/stage/omc_entity.json | jq '.entities[0].name="'$(hostname -f)'"' | cat > /omc/stage/omc_entity_update.json

# - Complete second stage of agent install
/omc/stage/AgentInstall.sh AGENT_TYPE=cloud_agent AGENT_BASE_DIR='/omc/app/cloud_agent' -staged  AGENT_PROPERTIES=$PWD/agent.properties AGENT_REGISTRATION_KEY=${registration_key} ORACLE_HOSTNAME=$(hostname -f)

# - Setup the OMC CLI
export core=$(/omc/app/cloud_agent/agent_inst/bin/omcli status agent | grep '^Binaries Location' | awk -F: '{print $2}')
sudo $core/root.sh

# - Register this instance with OMC for monitoring
/omc/app/cloud_agent/agent_inst/bin/omcli update_entity agent /omc/stage/omc_entity_update.json
