# Test Configs Guide

Make these changes to configs for enabling certain features for test purposes.

## app.toml
```toml
#Base Configuration
minimum-gas-prices = "0mand"

# Telemetry Configuration
enabled = true
enable-hostname = true
enable-hostname-label = true
enable-service-label = true
prometheus-retention-time = 10

# API Configuration
enable = true
swagger = true
enabled-unsafe-cors = true

# gRPC Web Configuration
enable-unsafe-cors = true
```

## config.toml
```toml
#RPC Server Configuration Options
cors_allowed_origins = ["*"]

# Instrumentation Configuration Options
prometheus = true
```

## client.toml
```toml
# Client Configuration
chain-id = "mande-test-chain-1"
keyring-backend = "test"
```

## References for more info
1. [Tendermint Configuration](https://docs.tendermint.com/master/nodes/configuration.html)