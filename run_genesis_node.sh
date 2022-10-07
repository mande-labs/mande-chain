
## Initialize the chain
mandeNode init ${MONIKER} --chain-id ${CHAIN_ID}

# Create genesis accounts
mandeNode keys add $${VALIDATOR_NAME} --keyring-backend test
mandeNode add-genesis-account $$(mandeNode keys show $${VALIDATOR_NAME} -a --keyring-backend test) "100000000000000cred"

mandeNode keys add alice --keyring-backend test
mandeNode keys add bob --keyring-backend test
mandeNode keys add faucet1 --keyring-backend test
mandeNode keys add faucet2 --keyring-backend test
mandeNode keys add faucet3 --keyring-backend test

mandeNode add-genesis-account alice 10000000000cred > genesis_accounts.txt
mandeNode add-genesis-account bob 10000000000cred > genesis_accounts.txt
mandeNode add-genesis-account faucet1 10000000000000000cred > genesis_accounts.txt
mandeNode add-genesis-account faucet2 10000000000000000cred > genesis_accounts.txt
mandeNode add-genesis-account faucet3 10000000000000000cred > genesis_accounts.txt

## Modify genesis.json
# distribution
jq '.app_state.distribution.params.base_proposer_reward = "0.010000000000000000"' genesis.json > temp.json
jq '.app_state.distribution.params.bonus_proposer_reward = "0.040000000000000000"' genesis.json > temp.json
jq '.app_state.distribution.params.community_tax = "0.020000000000000000"' genesis.json > temp.json
jq '.app_state.distribution.params.withdraw_addr_enabled = true' genesis.json > temp.json

# staking
jq '.app_state.staking.params.bond_denom = "cred"' genesis.json > temp.json
jq '.app_state.staking.params.unbonding_time = "0.001s"' genesis.json > temp.json
jq '.app_state.staking.params.max_validators = 175' genesis.json > temp.json
jq '.app_state.staking.params.max_entries = 7' genesis.json > temp.json
jq '.app_state.staking.params.historical_entries = 10000' genesis.json > temp.json

# gov
jq '.app_state.gov.deposit_params.min_deposit = [{ "amount": "10000000", "denom": "mand" }]' genesis.json > temp.json

# mint
jq '.app_state.mint.minter.inflation = "0.070000000000000000"' genesis.json > temp.json
jq '.app_state.mint.minter.annual_provisions = "0.000000000000000000"' genesis.json > temp.json
jq '.app_state.mint.params.blocks_per_year = "4360000"' genesis.json > temp.json
jq '.app_state.mint.params.goal_bonded = "0.670000000000000000"' genesis.json > temp.json
jq '.app_state.mint.params.inflation_max = "0.200000000000000000"' genesis.json > temp.json
jq '.app_state.mint.params.inflation_min = "0.070000000000000000"' genesis.json > temp.json
jq '.app_state.mint.params.inflation_rate_change = "0.130000000000000000"' genesis.json > temp.json
jq '.app_state.mint.params.mint_denom = "mand"' genesis.json > temp.json

# crisis
jq '.app_state.crisis.constant_fee.denom = "mand"' genesis.json > temp.json

# bank
jq '.app_state.bank.params.send_enabled = [{"denom": "cred","enabled": false}]' genesis.json > temp.json

# slashing
jq '.app_state.slashing.params.downtime_jail_duration = "600s"' genesis.json > temp.json
jq '.app_state.slashing.params.min_signed_per_window = "0.050000000000000000"' genesis.json > temp.json
jq '.app_state.slashing.params.signed_blocks_window = "10000"' genesis.json > temp.json
jq '.app_state.slashing.params.slash_fraction_double_sign = "0.050000000000000000"' genesis.json > temp.json
jq '.app_state.slashing.params.slash_fraction_downtime = "0.000100000000000000"' genesis.json > temp.json

mv temp.json genesis.json


mandeNode gentx $${VALIDATOR_NAME} "1000000cred" --chain-id ${CHAIN_ID} --moniker="$${MONIKER}" --keyring-backend test
mandeNode validate-genesis
mandeNode collect-gentxs
mandeNode validate-genesis

if [[ -z "$$(grep '\[api\]' -A 5 $${APP_TOML} | grep true)" ]]; then
  sed -i '1,/enable = false/{s/enable = false/enable = true/g}' $${APP_TOML}
fi
sed -i 's/cors_allowed_origins.*/cors_allowed_origins = ["*"]/g' $${CONFIG_TOML}
sed -i 's/cors_allowed_methods.*/cors_allowed_methods = ["*"]/g' $${CONFIG_TOML}
sed -i 's/cors_allowed_headers.*/cors_allowed_headers = ["*"]/g' $${CONFIG_TOML}

mandeNode start --pruning=nothing --rpc.laddr=tcp://0.0.0.0:26657