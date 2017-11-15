#!/usr/local/bin/python3.6

import os
import configparser
import argparse

env_list = {
    "tenancy": "tenancy_ocid",
    "fingerprint": "fingerprint",
    "key_file": "private_key_path",
    "user": "user_ocid",
    "pass_phrase": "private_key_password",
    "public_key": "ssh_public_key",
    "region": "region",
    "profile": "profile"
}

argparser = argparse.ArgumentParser()
argparser.add_argument('--config_path', required=False,
                       default='~/.oci/config', dest="config_path")
argparser.add_argument('--profile', required=False,
                       default='DEFAULT', dest="profile")
argparser.add_argument('--dest_file', required=False,
                       default='./tfenv.sh', dest='dest_file')
argparser.add_argument('--public_key', required=False,
                       default="~/.ssh/id_rsa.pub", dest='public_key')
args = vars(argparser.parse_args())

output_env = {}
output_file = os.open(args["dest_file"], os.O_RDWR |
                      os.O_TRUNC | os.O_CREAT, 0x0755)

if len(args["config_path"]) <= 0:
    oci_config_file = "~/.oci/config"
else:
    oci_config_file = args["config_path"]

oci_config_file = os.path.expanduser(oci_config_file)

parser = configparser.ConfigParser()
parser.read(oci_config_file)

if (args["profile"] in parser.sections()):
    for options in parser.options(args["profile"]):
        output_env[env_list[options]] = parser.get(
            args["profile"], options)
elif args["profile"] == "DEFAULT":
    for options in parser.defaults().keys():
        output_env[env_list[options]] = parser.defaults()[options]

if "~" in output_env["private_key_path"]:
    output_env["private_key_path"] = \
        os.path.expanduser(output_env["private_key_path"])

public_key_file = os.path.expanduser(args['public_key'])
if (os.path.exists(public_key_file)):
    pk_file = os.open(public_key_file, os.O_RDONLY)
    output_env["ssh_public_key"] = os.read(pk_file, 512).rstrip()
    os.close(pk_file)

output_env["profile"] = args["profile"]

os.write(output_file, b"#!/bin/bash\n")
for key in env_list.keys():
    if env_list[key] not in output_env:
        output_env[env_list[key]] = ""
    os.write(output_file, "export TF_VAR_{0}=\"{1}\"\n".format(
        env_list[key], output_env[env_list[key]]).encode())
os.close(output_file)
