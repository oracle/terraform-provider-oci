#! /bin/bash

LOG_FILE=/opt/omc/installer/omc-agent-container-`date +%H-%M-%S`.log
INSTALLER_DIR="/opt/omc/installer"
INSTALL_DIR="/opt/omc/omc-agent"
INSTALLER_ZIP="$INSTALLER_DIR/agent.zip"
OMCLI_LOC="$INSTALL_DIR/agent_inst/bin/omcli"
RSP_FILE="$INSTALLER_DIR/agent.rsp"
PROXY=""

echo "`date`: ******* Starting OMC Agent container ********" > $LOG_FILE

_log() {
 ts=`date`
 echo "$ts: $1" |& tee -a $LOG_FILE
}

_check_process(){
 NUM_PROCESSES=$(ps -ef | grep "$1" | grep -v "grep" | wc -l)
 if [ $NUM_PROCESSES -gt 0 ];
 then
  return $NUM_PROCESSES
 else
  return 0
 fi
}

_check_envvar(){
 return $(env | grep -iw ^"$1"= | wc -l)
}

_get_proxy(){
 PROXY="`env | grep https_ | awk -F= '{print $2}'`"
}


_check_file() {
  _log "Checking if file $1 exists and is non-zero size"
 if [ ! -s "$1" ]  || [ ! -f "$1" ];
 then
  _log "$1 does NOT exist"
  return 0
 else
  return 1
 fi
}

_init() {

 _check_file "$OMCLI_LOC"

 if [ $? -gt 0 ];
 then
  _log "Agent is already installed in $INSTALL_DIR, attempting to start agent"
  `$OMCLI_LOC start agent`
  _log "Agent started successfully"
  _log "Launching sshd"
  `/usr/sbin/sshd -D`
  _log "Container ready"
 else
  _log "Agent is NOT installed in this container - will start install routine now"
 fi

 _log "Starting env var checks"

 _check_envvar "$OMC_URL"
 if [ $? -gt 0 ];
 then
  _log "Variable OMC_URL required by OMC Agent installer can not install agent - exiting install process"
  exit -1
 else
  OMC_URL=`echo $OMC_URL | sed 's:/*$::'`
  _log "Variable OMC_URL=$OMC_URL is set"
 fi

_check_envvar "$AGENT_REPO_URL"
 if [ $? -gt 0 ];
 then
  _log "Variable AGENT_REPO_URL required by OMC Agent installer can not install agent - exiting install process"
  exit -1
 else
  AGENT_REPO_URL=`echo $AGENT_REPO_URL | sed 's:/*$::'`
  _log "Variable AGENT_REPO_URL=$AGENT_REPO_URL is set"
 fi

 _check_envvar "$REGISTRATION_KEY"
 if [ $? -gt 0 ];
  then
   _log "Variable REGISTRATION_KEY required by OMC Agent installer can not install agent - exiting install process"
   exit -1
 else
  _log "Variable REGISTRATION_KEY=$REGISTRATION_KEY is set"
 fi

 _check_envvar "$TENANT_NAME"
 if [ $? -gt 0 ];
 then
  _log "Variable TENANT_NAME required by OMC Agent installer can not install agent - exiting install process"
  exit -1
 else
  _log "Variable TENANT_NAME=$TENANT_NAME is set"
 fi

#  _check_envvar "$USER_NAME"
#  if [ $? -gt 0 ];
#  then
#   _log "Variable USER_NAME required by OMC Agent installer can not install agent - exiting install process"
#   exit -1
#  else
#   _log "Variable USER_NAME=$USER_NAME is set"
#  fi

#  _check_envvar "$PASSWORD"
#  if [ $? -gt 0 ];
#  then
#   _log "Variable $PASSWORD required by OMC Agent installer can not install agent - exiting install process"
#   exit -1
#  else
#   _log "Variable password is set (value not shown)"
#  fi

#  _check_envvar "$HTTPS_PROXY"
#  if [ $? -gt 0 ];
#  then
#    #sets the value of global variable PROXY, $PROXY will be used for RSP substituion later
#    _get_proxy
#   _log "Variable https_proxy=$PROXY is set"  
#  else
#   _log "Variable HTTPS_PROXY is not set - this is NOT fatal, will attempt install"
#  fi

}

_init

#If tenant id contains '-' in it, then maintain another tenant id without '-'
TENANT_NAME_WITHOUT_INS=$TENANT_NAME
count=`echo $TENANT_NAME | sed -e 's/-/ /g' | wc -w`
if [ $count -gt 1 ];
then
  TENANT_NAME_WITHOUT_INS=`echo $TENANT_NAME | cut -d'-' -f2`
fi

DOWNLOAD_CMD_FORPRINT="wget --no-check-certificate $AGENT_REPO_URL -O $INSTALLER_ZIP"

DOWNLOAD_CMD="wget --no-check-certificate $AGENT_REPO_URL -O $INSTALLER_ZIP"

_log "DOWNLOAD_CMD is $DOWNLOAD_CMD_FORPRINT"
_log "Begin downloading installer"
`$DOWNLOAD_CMD | tee -a $LOG_FILE`
_log "Done downloading installer"
_log "Checking downloaded file"
_check_file "$INSTALLER_ZIP"
if [ $? -gt 0 ];
 then
  _log "$INSTALLER_ZIP downloaded successfully"
 else
  _log "Error downloading installer, please make sure required environment variables are set. You may need to set HTTPS_PROXY env var. TERMINATING script - container NOT launched"
  exit -1
fi

_log "Begin unzipping installer"
cd $INSTALLER_DIR; unzip agent.zip 2>&1 |  tee -a $LOG_FILE
_log "Done unzipping installer"

_log "Replacing variables in $RSP_FILE"
sed -i "/^TENANT_NAME=/c\TENANT_NAME=$TENANT_NAME" $RSP_FILE
sed -i "/^OMC_URL=/c\OMC_URL=$OMC_URL" $RSP_FILE
sed -i "/^AGENT_REGISTRATION_KEY=/c\AGENT_REGISTRATION_KEY=$REGISTRATION_KEY" $RSP_FILE
sed -i "/^AGENT_BASE_DIRECTORY=/c\AGENT_BASE_DIRECTORY=/opt/omc/omc-agent" $RSP_FILE
sed -i "/^IGNORE_ULIMIT_CHECK=/c\IGNORE_ULIMIT_CHECK=true" $RSP_FILE
sed -i "/^AGENT_PORT=/c\AGENT_PORT=9000" $RSP_FILE

 if [ ! -z "$PROXY" ];
 then
  _log "Variable PROXY=$PROXY is set"
  PROXY_SERVER=`echo $PROXY | awk -F/ '{print $3}' | awk -F: '{print $1}'`
  PROXY_PORT=`echo $PROXY | awk -F/ '{print $3}' | awk -F: '{print $2}'`

  _log "$PROXY_SERVER will be used for OMC_PROXYSERVER"
  _log "$PROXY_PORT will be used for OMC_PROXYPORT"

 else
  _log "Variable PROXY is NOT set, no updates to $RSP_FILE will be performed"
 fi


_log "Done replacing variables in $RSP_FILE"

`rm -rf /opt/omc/omc-agent`

_log "Launching installer"
/opt/omc/installer/AgentInstall.sh 2>&1 |  tee -a $LOG_FILE
_log "Installer complete"

# _log "Launching sshd"
# `/usr/sbin/sshd -D`
_log "Container ready"
