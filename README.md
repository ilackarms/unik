# <img src="http://i.imgur.com/npkzp8l.png" alt="Build and run unikernels" width="141" height="50">

UniK (pronounced you-neek) is a tool for compiling application sources into unikernels (lightweight bootable disk images) rather than binaries.

UniK runs and manages instances of compiled images across a variety of cloud providers as well as locally on Virtualbox.

UniK utilizes a simple docker-like command line interface, making building unikernels as easy as building containers.

UniK is built to be easily extensible, allowing (and encouraging) adding support for unikernel [compilers](docs/compilers/README.md) and cloud [providers](docs/providers/README.md). See [architecture](docs/architecture.md) for a better understanding of UniK's pluggable code design.

To learn more about the motivation behind project UniK, read our [blog](https://github.com/ilackarms/unik/wiki/UniK:-Build-and-Run-Unikernels-with-Ease) post, or [watch](https://www.youtube.com/watch?v=wcZWg3YtvnY) session ([Slides](http://www.slideshare.net/IditLevine/unik-slides)).
We also encourage you to read [this](https://github.com/ilackarms/unik/wiki/Worried-about-IoT-DDoS%3F-Think-Unikernels) blog about IoT Security.<BR><BR>
To stay up-to-date on UniK, follow us [@ProjectUnik](https://twitter.com/ProjectUniK) and join us on our [slack channel](http://project-unik.io).

---

### *Changelog*:

Due to the frequency with which features and patches are applied to UniK, we have decided to list these changes by date in our [unversioned changelog](docs/changelog.md).

---

### Documentation
- **Installation**
  - [Install UniK](docs/install.md)
  - [Configuring the daemon](docs/configure.md)
  - [Launching the InstanceListener](docs/instance_listener.md)
- **Getting Started**
  - [Pull and run a unikernel without building](docs/getting_started_pull.md) on Virtualbox with UniK
  - [Run your first Go unikernel](docs/getting_started.md) on Virtualbox with UniK
  - [Run your first Node.js unikernel](docs/getting_started_node.md) on Virtualbox with UniK
  - [Run your first Python3 unikernel](docs/getting_started_python3.md) on Virtualbox with UniK
  - [Run your first Java unikernel](docs/getting_started_java.md) on Virtualbox with UniK
  - [Run your first C++ unikernel](docs/getting_started_cpp.md) on Virtualbox with UniK
- **User Documenation**
  - Using the [command line interface](docs/cli.md)
  - Compiling [Node.js](docs/compilers/rump.md#nodejs) Applications to Unikernels
  - Compiling [Go](docs/compilers/rump.md#golang) Applications to Unikernels
  - Compiling [Java](docs/compilers/osv.md#java) Applications to Unikernels (OSv)
  - Compiling [Node.js](docs/compilers/osv.md#nodejs) Applications to Unikernels (OSv)
  - Compiling [C/C++](docs/compilers/osv.md#native) Applications to Unikernels (OSv)
  - Compiling [C/C++](docs/compilers/includeos.md) Applications to Unikernels
  - Compiling [Python3](docs/compilers/rump.md#python-3) Applications to Unikernels
- **Developer Documentation**
  - Adding [compiler](docs/compilers/README.md) support
  - Adding [provider](docs/providers/README.md) support

---

### Supported unikernel types:
* **rump**: UniK supports compiling [Python](docs/compilers/rump.md#python-3), [Node.js](docs/compilers/rump.md#nodejs) and [Go](docs/compilers/rump.md#golang) code into [rumprun](docs/compilers/rump.md) unikernels
* **OSv**: UniK supports compiling Java, Node.js, C and C++ code into [OSv](http://osv.io/) unikernels
* **IncludeOS**: UniK supports compiling C++ code into [IncludeOS](https://github.com/hioa-cs/IncludeOS) unikernels
* **MirageOS**: UniK supports compiling [OCaml](docs/compilers/mirage.md), code into [MirageOS](https://mirage.io) unikernels

*We are looking for community help to add support for more unikernel types and languages.*

### Supported providers:
* [Virtualbox](docs/providers/virtualbox.md)
* [AWS](docs/providers/aws.md)
* [Google Cloud](docs/providers/gcloud.md)
* [vSphere](docs/providers/vsphere.md)
* [QEMU](docs/providers/qemu.md)
* [UKVM](docs/providers/ukvm.md)
* [Xen](docs/providers/xen.md)
* [OpenStack](docs/providers/openstack.md)
* [Photon Controller](docs/providers/photon.md)

### Roadmap:
* dynamic volume and application arguments configuration at instance runtime (rather than at compile time)
* expanding test coverage
* better code documentation
* multi-account support per provider (i.e. multiple AWS accounts/regions, etc.)
* migrate from [martini](https://github.com/go-martini/martini) to [echo](https://github.com/labstack/echo)

UniK is still experimental! APIs and compatibility are subject to change. We are looking for community support to help identify potential bugs and compatibility issues. Please open a Github issue for any problems you may experience, and join us on our [slack channel](http://project-unik.io)

---

### Thanks

**UniK** would not be possible without the valuable open-source work of projects in the unikernel community. We would like to extend a special thank-you to [rumprun](https://github.com/rumpkernel), [deferpanic](https://github.com/deferpanic), [cloudius systems](https://github.com/cloudius-systems), [mirageos](https://mirage.io) and [includeOS](http://www.includeos.org/).

<!--(for contributors): push images: CONTAINERVER=0.1 for i in $(docker images | grep projectunik/ | awk '{print $1}'); do docker push $i:$CONTAINERVER; done-->
