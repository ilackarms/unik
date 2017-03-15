# Getting Started

In this tutorial we'll be:
  1. [installing UniK](getting_started.md#installing-unik)
  2. [writing a simple HTTP Daemon in Go](getting_started.md#write-a-go-http-server)
  3. [compiling to a unikernel and launching an instance on Virtualbox](getting_started.md#compile-an-image-and-run-on-virtualbox)

### Installing UniK
#### Prerequisites
Ensure that each of the following are installed
- [Docker](http://www.docker.com/) installed and running with at least 6GB available space for building images
- [`jq`](https://stedolan.github.io/jq/)
- [`make`](https://www.gnu.org/software/make/)
- [Virtualbox](https://www.virtualbox.org/)

#### Install, configure, and launch UniK
1. Install UniK
  ```
  $ git clone https://github.com/ilackarms/unik.git
  $ cd unik
  $ make binary
  ```
  note: `make` will take quite a few minutes the first time it runs. the UniK `Makefile` is pulling all of the Docker images that bundle UniK's dependencies.

  Then, place the `unik` executable in your `$PATH` to make running UniK commands easier:
  ```
  $ mv _build/unik /usr/local/bin/
  ```

2. Configure a Host-Only Network on Virtualbox
  * Open Virtualbox
  * Open **Preferences** > **Network** > **Host-only Networks**
  * Click the green add button on the right side of the UI
  * Record the name of the new Host-Only adapter (e.g. "vboxnet0"). You will need this in your UniK configuration
  * Ensure that the Virtualbox DHCP Server is Enabled for this Host-Only Network:
    * With the Host-Only Network selected, Click the edit button (screwdriver image)
    * In the **Adapter** tab, note the IPv4 address and netmask of the adapter.
    * In the **DHCP Server** tab, check the **Enable Server** box
    * Set **Server Address** an IP on the same subnet as the Adapter IP. For example, if the adapter IP is `192.168.100.1`, make set the DHCP server IP as `192.168.100.X`, where X is a number between 2-254.
    * Set **Server Mask** to the netmask you just noted
    * Set **Upper / Lower Address Bound** to a range of IPs on the same subnet. We recommend using the range `X-254` where X is one higher than the IP you used for the DHCP server itself. E.g., if your DHCP server is `192.168.100.2`, you can set the lower and upper bounds to `192.168.100.3` and `192.168.100.254`, respectively.


3. Configure UniK daemon
  * UniK configuration files are stored in `$HOME/.unik`. Create this directory, if it is not present:
  ```
  $mkdir $HOME/.unik
  ```
  * Using a text editor, create and save the following to `$HOME/.unik/daemon-config.yaml`:
  ```yaml
  providers:
    virtualbox:
      - name: my-vbox
        adapter_type: host_only
        adapter_name: NEW_HOST_ONLY_ADAPTER
  ```
  replacing `NEW_HOST_ONLY_ADAPTER` with the name of the network adapter you created.


4. Launch UniK and automatically deploy the *Virtualbox Instance Listener*
  * Open a new terminal window/tab. This terminal will be where we leave the UniK daemon running.
  * `cd` to the `_build` directory created by `make`
  * run `unik daemon --debug` (the `--debug` flag is optional, if you want to see more verbose output)
  * UniK will compile and deploy its own 30 MB unikernel. This unikernel is the [Unik Instance Listener](./instance_listener.md). The Instance Listener uses udp broadcast to detect (the IP address) and bootstrap instances running on Virtualbox.
  * After this is finished, UniK is running and ready to accept commands.
  * Open a new terminal window and type `unik target --host localhost` to set the CLI target to the your local machine.

---

#### Write a Go HTTP server

0. Open a new terminal window, but leave the window with the daemon running. This window will be used for running UniK CLI commands.

1. Create a file `httpd.go` using a text editor. Copy and paste the following code in that file:

  ```go
  package main

  import (
      "fmt"
      "net/http"
  )

  func main() {
      http.HandleFunc("/", handler)
      http.ListenAndServe(":8080", nil)
  }

  func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "my first unikernel!")
  }
  ```

2. Try running this code with `go run http.go`. Visit [http://localhost:8080/](http://localhost:8080/) to see that the server is running.
3. We need to create a dummy `Godeps` file. This is necessary to tell the Go compiler how Go projects and their dependencies are structured. Fortunately, with this example, our project has no dependencies, and we can just fill out a simple `Godeps` file without installing [`godep`](https://github.com/tools/godep). Note: for Go projects with imported dependencies, and nested packages, you will need to install Godeps and run `GO15VENDOREXPERIMENT=1 godep save ./...` in your project. See [Compiling Go Apps with UniK](compilers/rump.md#golang) for more information.
  * To create the dummy Godeps file, create a folder named `Godeps` in the same directory as `httpd.go`. Inside, create a file named `Godeps.json` and paste the following inside:
  ```json
  {
  	"ImportPath": "my_httpd",
  	"GoVersion": "go1.6",
  	"GodepVersion": "v63",
  	"Packages": [
  		"./.."
  	],
  	"Deps": [
  		{
  			"ImportPath": "github.com/ilackarms/unik/docs/examples",
  			"Rev": "f8cc0dd435de36377eac060c93481cc9f3ae9688"
  		}
  	]
  }
  ```
  * For the purposes of this example, that matters here is `my_httpd`. It instructs the go compiler that the project should be installed from `$GOPATH/src/my_httpd`.

4. Great! Now we're ready to compile this code to a unikernel.

---

#### Compile an image and run on Virtualbox

1. run the following command from the directory where your `httpd.go` is located:
  ```
  unik build --name myImage --path ./ --base rump --language go --provider virtualbox
  ```
  This command will instruct UniK to compile the sources found in the working directory (`./`) using the `rump-go-virtualbox` compiler.
2. You can watch the output of the `build` command in the terminal window running the daemon.
3. When `build` finishes, the resulting disk image will reside at `$HOME/.unik/virtualbox/images/myImage/boot.vmdk`
4. Run an instance of this image with
  ```
  unik run --instanceName myInstance --imageName myImage
  ```
5. When the instance finishes launching, let's check its IP and see that it is running our application.
6. Run `unik instances`. The instance IP Address should be listed.
7. Direct your browser to `http://instance-ip:8080` and see that your instance is running!
8. To clean up your image and the instance you created
  ```
  unik rmi --force --image myImage
  ```
