# Drone vault plugin

## Building docker images

```bash
$ dep ensure
$ env GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/drone-vault github.com/jeroenrinzema/drone-vault/cmd/drone-vault
$ docker build -t jeroenrinzema/drone-vault .
```

## Available environment variables

```golang
type config struct {
	Debug          bool          `envconfig:"DEBUG"`
	Address        string        `envconfig:"SERVER_ADDRESS"`
	Secret         string        `envconfig:"SECRET_KEY"`
	VaultAddr      string        `envconfig:"VAULT_ADDR"`
	VaultRenew     time.Duration `envconfig:"VAULT_TOKEN_RENEWAL"`
	VaultTTL       time.Duration `envconfig:"VAULT_TOKEN_TTL"`
	VaultAuthType  string        `envconfig:"VAULT_AUTH_TYPE"`
	VaultAuthMount string        `envconfig:"VAULT_AUTH_MOUNT_POINT"`
	VaultKubeRole  string        `envconfig:"VAULT_KUBERNETES_ROLE"`
}
```