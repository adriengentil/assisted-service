#!/usr/bin/env bash

set -o nounset
set -o pipefail
set -o errexit

# Stop running this script if we have not running on Oracle Cloud
chassis_asset_tag=$(dmidecode --string "chassis-asset-tag")
if [ "${chassis_asset_tag}" != "OracleCloud.com" ]
then
  systemctl disable oci-network-if-config.service
  exit
fi

# Retrieve and run the script intended to configure the extra network interfaces on the machine

network_if_config_script=$(mktemp)
trap 'rm f -- "${network_if_config_script}"' EXIT

curl --silent \
	--fail \
	--retry 10 \
	--connect-timeout 30 \
	--output "${network_if_config_script}" \
	--header "Authorization: Bearer Oracle" \
	--location "http://169.254.169.254/opc/v2/metadata/network-if-config"

chmod +x "${network_if_config_script}"
"${network_if_config_script}"
