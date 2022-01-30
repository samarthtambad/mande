# First Node Setup Guide

## Create moniker
Choose a name for this node. It can be any unique name. Set it as an environment variable for ease of access.
Add the following to your `.bashrc` or `.zshrc`:
```sh
export MONIKER=<node-name-you-chose>
```

## Download the mande daemon
The daemon is the executable used to run the node, query and initiate transactions. Download and unzip the latest `manded` executable from the [releases](https://github.com/mande-labs/mande/releases) section of the repo.

## Initialize the chain
Initialize the chain using the following command:
```sh
manded init $MONIKER --chain-id mande-test-chain-1
```

Initialization creates the genesis block and and an initial chain state. You can find these files and configs in the newly created `.mande` folder.

**Note:** For the executable to be accessible from anywhere, add `manded` to the `bin` folder or add the folder to the PATH env variable.

## Create and add private key to a keyring
Create a key with the following command. Replace `my_key_name` with your own key name.
```sh
manded keys add my_key_name --keyring-backend test
```

It will generate the key and return the address and a mnemonic. Make sure to save the mnemonic.

**Note:** There are many backends supported by the keyring such as: `os`, `file`, `kwallet`, `pass` and `test`. We will use `test` for testing purpose as it does not require a password to be unlocked. 


## Add genesis account

```sh
# get your validator address from the key you created above
MY_VALIDATOR_ADDRESS=$(manded keys show my_key_name -a --keyring-backend test)

manded add-genesis-account $MY_VALIDATOR_ADDRESS 1000000000000000mand,1000000000000000cred
```

## Create a genesis transaction
```sh
manded gentx my_key_name 10000000000cred --chain-id mande-test-chain-1 --keyring-backend test --moniker=$MONIKER
```

### [Optional] Add a faucet
```sh
manded keys add faucet-key --keyring-backend test
FAUCET_ADDRESS=$(manded keys show faucet-key -a --keyring-backend test)

# create faucet account
manded add-genesis-account $FAUCET_ADDRESS 100000000000000000mand,100000000000000000cred

# TODO: how to use this account for the faucet (starport)
```

## Add the genesis transactions (gentx) to the genesis file.

```sh
manded collect-gentxs
```

**Note**: To change the default staking token (i.e. stake), manually change the `genesis.json` to your new token.

## Start the node
```sh
manded start
```

Or to start the node in the background follow [Running via Background Process](https://hub.cosmos.network/main/hub-tutorials/join-mainnet.html#running-via-background-process).

## References for more info
1. [Running a Node, API, and CLI](https://tutorials.cosmos.network/academy/3-running-a-chain/node-api-and-cli.html)
