# docker-credential-acr-env

A Docker credential helper for Azure Container Registry (ACR) that sources
credentials from the environment. A modern, actively maintained replacement for both
[`docker-credential-acr-env`](https://github.com/chrismellard/docker-credential-acr-env)
and Azure's own
[`acr-docker-credential-helper`](https://github.com/Azure/acr-docker-credential-helper),
both of which are now unmaintained or archived. Built on the current Azure SDK
(`azidentity` / `azcontainerregistry`) instead of the deprecated `go-autorest` stack.

## Features

- **Docker credential helper** — plug into Docker via `~/.docker/config.json`
- **go-containerregistry keychain** — drop-in `authn.Keychain` for Go programs (cosign, crane, ko, etc.)
- **Modern Azure SDK** — uses `azidentity.DefaultAzureCredential` which automatically discovers credentials
- **All Azure auth methods** — service principal, workload identity, managed identity, Azure CLI, and more
- **ACR registry detection** — recognises `azurecr.io` and `mcr.microsoft.com` endpoints

## Installation

### Binary

Download from [Releases](https://github.com/gaganhr94/docker-credential-acr/releases),
or build from source:

```bash
go install github.com/gaganhr94/docker-credential-acr@latest
```

The binary is named `docker-credential-acr-env`.

### As a Go library

```bash
go get github.com/gaganhr94/docker-credential-acr
```

## Usage

### Docker credential helper

Add to `~/.docker/config.json`:

```json
{
  "credHelpers": {
    "myregistry.azurecr.io": "acr-env"
  }
}
```

Docker (and tools like kaniko) will automatically call `docker-credential-acr-env get`
to fetch credentials.

### Go library — authn.Keychain

Use `credhelper.Keychain()` for a ready-to-use keychain that silently skips non-ACR registries:

```go
import (
    "github.com/google/go-containerregistry/pkg/authn"
    "github.com/gaganhr94/docker-credential-acr/pkg/credhelper"
)

kc := authn.NewMultiKeychain(
    authn.DefaultKeychain,
    credhelper.Keychain(),
)
```

Or use the lower-level `authn.Helper` interface:

```go
helper := credhelper.NewKeychainHelper()
kc := authn.NewKeychainFromHelper(helper)
```

### Go library — registry detection

```go
import "github.com/gaganhr94/docker-credential-acr/pkg/registry"

registry.IsACRRegistry("myregistry.azurecr.io") // true
registry.IsACRRegistry("mcr.microsoft.com")     // true
registry.IsACRRegistry("gcr.io")                // false
```

### Go library — token exchange

```go
import (
    "github.com/gaganhr94/docker-credential-acr/pkg/token"
    "github.com/gaganhr94/docker-credential-acr/pkg/registry"
)

cred, _ := token.GetCredential()
refreshToken, _ := registry.GetRegistryRefreshToken(ctx, "myregistry.azurecr.io", tenantID, cred)
```

## Authentication

This helper uses [`azidentity.NewDefaultAzureCredential`](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#NewDefaultAzureCredential),
which tries the following methods in order:

| Method | Environment Variables |
|---|---|
| Service Principal (secret) | `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, `AZURE_TENANT_ID` |
| Service Principal (certificate) | `AZURE_CLIENT_ID`, `AZURE_CLIENT_CERTIFICATE_PATH`, `AZURE_TENANT_ID` |
| Workload Identity | `AZURE_FEDERATED_TOKEN_FILE`, `AZURE_CLIENT_ID`, `AZURE_TENANT_ID` |
| Managed Identity | (automatic in Azure), or `AZURE_CLIENT_ID` for user-assigned |
| Azure CLI | (uses `az login` session) |
| Azure Developer CLI | (uses `azd auth login` session) |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute.

## License

Apache License 2.0
