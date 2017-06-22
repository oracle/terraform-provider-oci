#!/usr/bin/env bash

cat /omc/stage/omc_entity.json | jq '.entities[0].name="'$(hostname -f)'"' | cat > /omc/stage/omc_entity_update.json
/omc/stage/AgentInstall.sh AGENT_TYPE=cloud_agent AGENT_BASE_DIR='/omc/app/cloud_agent' -staged  AGENT_PROPERTIES=$PWD/agent.properties AGENT_REGISTRATION_KEY=${registration_key} ORACLE_HOSTNAME=$(hostname -f)
export core=$(/omc/app/cloud_agent/agent_inst/bin/omcli status agent | grep '^Binaries Location' | awk -F: '{print $2}')
sudo $core/root.sh
/omc/app/cloud_agent/agent_inst/bin/omcli update_entity agent /omc/stage/omc_entity_update.json