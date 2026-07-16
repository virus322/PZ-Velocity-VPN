#!/bin/bash

set -e


echo "================================="
echo " PZ Velocity VPN Installer"
echo " Version 0.1.0"
echo "================================="


echo "[+] Updating system"

apt update -y


echo "[+] Installing dependencies"

apt install -y \
curl \
wget \
git \
ca-certificates


echo "[+] Creating directories"


mkdir -p /opt/pz-velocity


echo "[+] Installation structure ready"


echo "PZ Velocity VPN installed"