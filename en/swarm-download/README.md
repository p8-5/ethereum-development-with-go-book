---
description: Tutorial on downloadig files from swarm in Go.
---

# Downloading Files from Swarm

In the [previous section](../swarm-upload) we uploaded a hello.txt file to swarm and in return we got a manifest hash.

```go
manifestHash := "0x5439291bD121C69EccAc5542E32751401Ac251FF"
```

Let's inspect the manifest by downloading it first by calling `DownloadManfest`.

```go
manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
if err != nil {
  log.Fatal(err)
}
```

We can iterate over the manifest entries and see what the content-type, size, and content hash are.

```go
for _, entry := range manifest.Entries {
  fmt.Println(entry.Hash)        // 0xe92A52398E068941D9aC03E001e14aF636bcB2F3
  fmt.Println(entry.ContentType) // text/plain; charset=utf-8
  fmt.Println(entry.Path)        // ""
}
```

If you're familiar with swarm urls, they're in the format `bzz:/<hash>/<path>`, so in order to download the file we specify the manifest hash and path. The path in this case is an empty string. We pass this data to the `Download` function and get back a file object.

```go
file, err := client.Download(manifestHash, "0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664")
if err != nil {
  log.Fatal(err)
}
```

We may now read and print the contents of the returned file reader.

```go
content, err := ioutil.ReadAll(file)
if err != nil {
  log.Fatal(err)
}

fmt.Println(string(content)) // hello world
```

As expected, it logs *hello world* which what our original file contained.

---

### Full code

Commands

```bash
geth account new
export BZZKEY=0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73
swarm --bzzaccount $BZZKEY
```

[swarm_download.go](https://github.com/Browser-Coin/ethereum-development-with-go-book/blob/master/code/swarm_download.go)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	bzzclient "github.com/ethersphere/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")
	manifestHash := "0x380347b99285a3c7fEE2489A0A6EF9cf018589F1"
	manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isEncrypted) // false

	for _, entry := range manifest.Entries {
		fmt.Println(entry.Hash)        // 0xe92A52398E068941D9aC03E001e14aF636bcB2F3
		fmt.Println(entry.ContentType) // text/plain; charset=utf-8
		fmt.Println(entry.Size)        // 12
		fmt.Println(entry.Path)        // "0x304Cd3750060E18c54eCa2716C6AC5f9c180ed73"
	}

	file, err := client.Download(manifestHash, "")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // hello world
}
```
