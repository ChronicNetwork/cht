# CHT MAINNET MOROCCO-1

## Hardware Requirements
* **Minimal**
    * 4 GB RAM
    * 100 GB SSD
    * 3.2 x4 GHz CPU
* **Recommended**
    * 8 GB RAM
    * 1 TB NVME SSD
    * 3.2 GHz x4 GHz CPU

## Operating System

* **Recommended**
    * Linux(x86_64)


## Installation Steps
#### 1. Basic Packages
```bash:
# update the local package list and install any available upgrades 
sudo apt-get update && sudo apt upgrade -y 
# install toolchain and ensure accurate time synchronization 
sudo apt-get install make build-essential gcc git jq chrony -y
```
```bash:
# install gcc & make
sudo apt install gcc && sudo apt install make
```

#### 2. Install Go
Follow the instructions [here](https://golang.org/doc/install) to install Go.

Alternatively, for Ubuntu LTS, you can do:
```bash:
wget https://golang.org/dl/go1.18.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
```

Unless you want to configure in a non standard way, then set these in the `.profile` in the user's home (i.e. `~/`) folder.

```bash:
cat <<EOF >> ~/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
EOF
source ~/.profile
go version
```

Output should be: `go version go1.18.2 linux/amd64`

<a id="install-chtd"></a>
### Install chtd from source

#### 1. Clone repository

* Clone git repository
```shell
git clone https://github.com/ChronicNetwork/cht.git
```
* Checkout latest tag
```shell
cd cht
git fetch --tags
git checkout v1.1.0
```
#### 2. Install CLI
```shell
make build && make install
```

To confirm that the installation was successful, you can run:

```bash:
chtd version
```
Output should be: `v1.1.0`

## Instruction for new validators

### Init
```bash:
chtd init "$MONIKER_NAME" --chain-id $CHAIN_ID
```

### Generate keys

```bash:
# To create new keypair - make sure you save the mnemonics!
chtd keys add <key-name> 
```

or
```
# Restore existing chronic wallet with mnemonic seed phrase. 
# You will be prompted to enter mnemonic seed. 
chtd keys add <key-name> --recover
```
or
```
# Add keys using ledger
chtd keys show <key-name> --ledger
```

Check your key:
```
# Query the keystore for your public address 
chtd keys show <key-name> -a
```

## Validator Setup Instructions

### Set minimum gas fees
```bash:
perl -i -pe 's/^minimum-gas-prices = .+?$/minimum-gas-prices = "0.0125ucgas"/' ~/.cht/config/app.toml
```

### Add persistent peers


### Download new genesis file
```bash:
curl https://raw.githubusercontent.com/ChronicNetwork/net/main/mainnet/v1.1/genesis.json > ~/.cht/config/genesis.json
```

### Setup Unit/Daemon file

```bash:
# 1. create daemon file
touch /etc/systemd/system/chtd.service
# 2. run:
cat <<EOF >> /etc/systemd/system/chtd.service
[Unit]
Description=Chronic daemon
After=network-online.target
[Service]
User=<USER>
ExecStart=/home/<USER>/go/bin/chtd start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
EOF
# 3. reload the daemon
systemctl daemon-reload
# 4. enable service - this means the service will start up 
# automatically after a system reboot
systemctl enable chtd.service
# 5. start daemon
systemctl start chtd.service
```

In order to watch the service run, you can do the following:
```
journalctl -u chtd.service -f
```

Congratulations! You now have a full node. Once the node is synced with the network, 
you can then make your node a validator.

### Create validator
1. Transfer funds to your validator address. A minimum of 1 CHT (1000000ucht) is required to start a validator.

2. Confirm your address has the funds.

```
chtd q bank balances $(chtd keys show -a <key-alias>)
```

3. Run the create-validator transaction
**Note: 1,000,000 ucht = 1 CHT, so this validator will start with 1 CHT**

```bash:
chtd tx staking create-validator \ 
--amount 1000000loki \ 
--commission-max-change-rate "0.05" \ 
--commission-max-rate "0.10" \ 
--commission-rate "0.05" \ 
--min-self-delegation "1" \ 
--details "validators write bios too" \ 
--pubkey $(chtd tendermint show-validator) \ 
--moniker $MONIKER_NAME \ 
--chain-id $CHAIN_ID \ 
--fees 2000loki \
--from <key-name>
```

To ensure your validator is active, run:
```
chtd q staking validators | grep moniker
```

### Backup critical files
```bash:
priv_validator_key.json
node_key.json
```

## Instruction for old validators

### Stop node
```bash:
systemctl stop chtd.service
```

### Install latest Chtd from source

[Install latest Chtd](#install-chtd)

### Download genesis file
```bash:
curl https://raw.githubusercontent.com/ChronicNetwork/net/main/mainnet/genesis.json > ~/.cht/config/genesis.json
```

### Clean old state

```bash:
chtd unsafe-reset-all
```

### Rerun node
```bash:
systemctl daemon-reload
systemctl start chtd.service
```
