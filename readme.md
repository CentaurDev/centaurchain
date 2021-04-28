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

1. Initialize validators's and node's configuration files.
```
centaurchaind init [moniker] --chain-id centaurchain
```

2. Copy and replace genesis.json in config folder
```
cp /centaurchain/genesis.json ~/.centaurchain/config/genesis.json
```

3. Run the node
```
centaurchaind start --p2p.seeds "5c2861f0c73222e9a67967d45f7301fba2e8572d@64.225.103.191:26656,6577e64ac8799559f278eadb2fd51e88b63ffa58@174.138.30.107:26656,7f31fbd52e4cbb363e5b1626cbcaf30d470ad1a4@139.59.244.72:26656,9785c21dceb2e2735f7cec5d9ad74681173bf224@138.68.143.66:26656,f4bd87cb114cbfa71b0c3bb3cc453d9d77eeeafe@142.93.233.11:26656"
```

## Becoming a validator

1. Create a new key
```
centaurchaind keys add <key_name>
```

2. Send create-validator TX
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
