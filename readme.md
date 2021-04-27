# Centaur Chain

**Centaur Chain** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Requirements
- [Starport](https://docs.starport.network/intro/install.html)

## Installation

1. Clone the repo
```
git clone https://github.com/CentaurDev/centaurchain.git && cd centaurchain
```

2. Build with Starport
```
starport build
```

## Running a node

1 Initialize validators's and node's configuration files.
```
centaurchaind init [moniker] --chain-id centaurchain
```

2. Copy and replace genesis.json in config folder
```
cp /centaurchain/genesis.json ~/.centaurchain/config/genesis.json
```

3. Run the node (Get seeds from genesis.json)
```
centaurchaind start --p2p.seeds [ID@host:port]
```

## Becoming a validator

1. Create a new key
```
centaurchaind keys add <key_name>
```

3. Send create-validator TX
```
centaurchaind tx staking create-validator \
  --amount=2500000000000000000000000cntr \
  --pubkey=$(centaurchaind tendermint show-validator) \
  --moniker=[moniker] \
  --chain-id=centaurchain \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name>
```
