#!/bin/bash
# ipxe_gen.sh - Generate ipxe.sh script for use in building 
# IPXE boot server specific to RHEL 7.4 image creation
# 
# Steven B. Nelson - 27 Oct 2017

# Script takes input from terraform and generates all scripts needed
# for building ipxe server.  Resulting script is then base64encoded via
# terraform and run at instance provisioning time.

# Function build_ipxe - build the uuencode file functions
# Arguments - $1: function name, $2:source directory, $3: file name, $4:desired filename
# This just contains the function template and the rest is filled out by the variables.
# Since shar (in certain linux distros) created too big of function, we had to simplify
# to the bare minimum.
 
function build_ipxe {
cat >> ./ipxe.sh <<-EOF
function $1 {
cat > ./temp.uue <<"EoF"
EOF

source_file=$2"/"$3
uuencode ${source_file} $4 >> ./ipxe.sh

cat >> ./ipxe.sh <<-"EOF"
EoF
uudecode ./temp.uue
rm ./temp.uue
}

EOF
}

# Set input for stdin to capture JSON from terraform
INPUT_JSON=`cat -`

# Set location of the template to use, plus the location of the 
# source build files to put into the ipxe boot environment
IPXE_BUILD_TEMPLATE="./ipxe.sh.template"
IPXE_KS="ks.cfg"
IPXE_FWCFG="direct.xml"
IPXE_CLOUDINIT="cloud.cfg"
IPXE_SOURCE_DIR="./sources"

# Capture all the JSON info passed by the terraform into local variables
SSH_PUBLIC_KEY=`echo ${INPUT_JSON} | jq -r '.ssh_public_key'`
OCI_API_PRIVATE_KEY=`echo ${INPUT_JSON} | jq -r '.private_key_path'`
OCI_API_FINGERPRINT=`echo ${INPUT_JSON} | jq -r '.fingerprint'`
OCI_API_TENANCY=`echo ${INPUT_JSON} | jq -r '.tenancy_ocid'`
OCI_API_USER=`echo ${INPUT_JSON} | jq -r '.user_ocid'`
OCI_API_REGION=`echo ${INPUT_JSON} | jq -r '.region'`
OCI_PRKEY_PW=`echo ${INPUT_JSON} | jq -r '.private_key_password'`
OCI_OS_SHORT_NAME=`echo ${INPUT_JSON} | jq -r '.os_short_name'`
RHEL_UNAME=`echo ${INPUT_JSON} | jq -r '.rhel_user'`
RHEL_PW=`echo ${INPUT_JSON} | jq -r '.rhel_pw'`
ZEROS_OCID=`echo ${INPUT_JSON} | jq -r '.zeros_ocid'`
ISO_URL=`echo ${INPUT_JSON} | jq -r '.iso_url'`

# Create the head of the script and initialize our first function.  All functions are 
# simply encapsulations of uuencoded files.  This design pattern repeats
# for each of the files needed during build - cloud.cfg, direct.xml (firewalld), 
# ks.cfg (kickstart), and private key (OCI CLI)

echo "#!/bin/bash" > ./ipxe.sh
build_ipxe cloud ${IPXE_SOURCE_DIR} ${IPXE_CLOUDINIT} ${IPXE_CLOUDINIT}
build_ipxe firewallcfg ${IPXE_SOURCE_DIR} ${IPXE_FWCFG} ${IPXE_FWCFG}
build_ipxe ks ${IPXE_SOURCE_DIR} ${IPXE_KS} ${IPXE_KS}
build_ipxe privkey `dirname ${OCI_API_PRIVATE_KEY}` `basename ${OCI_API_PRIVATE_KEY}` oci_api_key.pem

# Add the template file to the shell script being built
cat ${IPXE_BUILD_TEMPLATE} >> ./ipxe.sh

# Replace all the tags in the shell script with actual values
sed -i.bak 's|<PUBLIC_KEY>|\"'"${SSH_PUBLIC_KEY}"'\"|g
s|<TENANCY>|\"'"${OCI_API_TENANCY}"'\"|g
s|<USER>|\"'"${OCI_API_USER}"'\"|g
s|<FINGERPRINT>|\"'"${OCI_API_FINGERPRINT}"'\"|g
s|<REGION>|\"'"${OCI_API_REGION}"'\"|g
s|<PASSPHRASE>|\"'"${OCI_PRKEY_PW}"'\"|g
s|<OS_NAME>|'"${OCI_OS_SHORT_NAME}"'|g
s|<RHEL_UNAME>|'"${RHEL_UNAME}"'|g
s|<RHEL_PASS>|'"${RHEL_PW}"'|g 
s|<ISO_URL>|'"${ISO_URL}"'|g
s|<ZEROS_OCID>|\"'"${ZEROS_OCID}"'\"|g' ./ipxe.sh

chmod 700 ./ipxe.sh
rm -rf ./ipxe.sh.bak
rm ./temp.uue

jq -n --arg shell "./ipxe.sh" '{ "shell":$shell }'
