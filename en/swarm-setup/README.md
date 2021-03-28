---
description: Tutorial on setting up swarm node.
---

# Setting up Swarm

To run swarm you first need to install `geth` and `bzzd` which is the swarm daemon.

```go
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethersphere/swarm/cmd/swarm
```

Now we'll generate a new geth account.

```bash
$ geth account new

Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
Repeat passphrase:
Address: {0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73}
```

Export the environment variable `BZZKEY` mapping to the geth account address we just generated.

```bash
export BZZKEY=0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73
```

And now run swarm with the specified account to be our swarm account. Swarm by default will run on port `8500`.

```bash
$ swarm --bzzaccount $BZZKEY
Unlocking swarm account  [1/3]
Passphrase:
WARN [06-12|13:11:41] Starting Swarm service
```

Now that we have the swarm daemon set up and running, let's learn how to upload files to swarm in the [next section](../swarm-upload).

---

### Full code

Commands

```bash
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethersphere/swarm/cmd/swarm
geth account new
export BZZKEY=0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73
swarm --bzzaccount $BZZKEY0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73
```
