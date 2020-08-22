---
title: Community & Contributing
weight: 7
---

We would love to hear from you! Here are some places you can find us.

## Mailing list

Our mailing list is
[metallb-users@googlegroups.com](https://groups.google.com/forum/#!forum/metallb-users). It's
for discussions around MetalLB usage, community support, and developer
discussion (although for the latter we mostly use Github directly).

## Slack

For a more interactive experience, we have the [#metallb slack channel
on k8s.slack.com](https://kubernetes.slack.com/messages/metallb/). If
you're not already logged into the Kubernetes slack organization,
you'll need to [request an invite](http://slack.k8s.io/) before you
can join.

Development of MetalLB is discussed in the [#metallb-dev slack channel
](https://kubernetes.slack.com/messages/metallb-dev/).

## IRC

If you prefer a more classic chat experience, we're also on `#metallb`
on the Freenode IRC network. You can use Freenode's [web
client](http://webchat.freenode.net?randomnick=1&channels=%23metallb&uio=d4)
if you don't already have an IRC client.

## Issue Tracker

Use the [GitHub issue
tracker](https://github.com/google/metallb/issues) to file bugs and
features request. If you need support, please send your questions to
the metallb-users mailing list rather than filing a GitHub issue.

# Contributing

We welcome contributions to MetalLB! Here's some information to get
you started.

## Code of Conduct

This project is released with a [Contributor Code of Conduct]({{%
relref "code-of-conduct.md" %}}). By participating in this project you
agree to abide by its terms.

## Code changes

Before you make significant code changes, please open an issue to
discuss your plans. This will minimize the amount of review required
for pull requests.

All submissions require review. We use GitHub pull requests for this
purpose. Consult [GitHub
Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Code organization

MetalLB's code is divided between a number of binaries, and some
supporting libraries. The libraries live in the `internal` directory,
and each binary has its own top-level directory. Here's what we
currently have, relative to the top-level directory:

- `controller` is the cluster-wide MetalLB controller, in charge of
  IP assignment.
- `speaker` is the per-node daemon that advertises services with
  assigned IPs using various advertising strategies.
- `internal/k8s` contains the bowels of the logic to talk to the
  Kubernetes apiserver to get and modify service information. It
  allows most of the rest of the MetalLB code to be ignorant of the
  Kubernetes client library, other than the objects (Service,
  ConfigMap...) that they manipulate.
- `internal/config` parses and validates the MetalLB configmap.
- `internal/allocator` is the IP address manager. Given pools from the
  MetalLB configmap, it can allocate addresses on demand.
- `internal/bgp` is a _very_ stripped down implementation of BGP. It
  speaks just enough of the protocol to keep peering sessions up, and
  to push routes to the peer.
- `internal/layer2` is an implementation of an ARP and NDP responder.
- `internal/logging` is a logging shim that redirects both
  Kubernetes's `klog` and Go's standard library `log` output to
  go-kit's structured logger, which is what MetalLB itself uses for
  logging.
- `internal/version` just burns version numbers and git commit
  information into compiled binaries, so that MetalLB can print its
  build information.

In addition to code, there's deployment configuration and
documentation:

- `manifests` contains a variety of Kubernetes manifests. The most
  important one is `manifests/metallb.yaml`, which specifies how to
  deploy MetalLB onto a cluster.
- `website` contains the website for MetalLB. The `website/content`
  subdirectory is where all the pages live, in Markdown format.

## Required software

To develop MetalLB, you'll need a couple of pieces of software:

- [git](https://git-scm.com), the version control system
- The [Go](https://golang.org) programming language (notably the `go`
  tool)
- [Docker](https://www.docker.com/docker-community), the container
  running system
- [kind](https://github.com/kubernetes-sigs/kind), a lightweight Kubernetes cluster running in Docker
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), the Kubernetes commandline interface
- [Invoke](https://www.pyinvoke.org) to drive the build system

>NOTE: The development environment was tested with **kind `v0.7.0`**. Older
>versions may not work since there have been breaking changes between minor
>versions.

## Building the code

Start by fetching the MetalLB repository, with `git clone
https://github.com/danderson/metallb`.

From there, you can use Invoke to build Docker images, push them to
registries, and so forth. `inv -l` lists the available tasks.

To build and deploy MetalLB to a local development environment using a kind
cluster, run `inv dev-env`.

For development, fork
the [github repository](https://github.com/google/metallb), and add
your fork as a remote in `$GOPATH/src/go.universe.tf/metallb`, with
`git remote add fork git@github.com:<your-github-user>/metallb.git`.

## The website

The website at https://metallb.universe.tf is pinned to the latest
released version, so that users who don't care about ongoing
development see documentation that is consistent with the released
code.

However, there is a version of the website synced to the latest main
branch
at
[https://main--metallb.netlify.com](https://main--metallb.netlify.com). Similarly,
every branch has a published website at `<branch
name>--metallb.netlify.com`. So if you want to view the documentation
for the 0.2 version, regardless of what the currently released version
is, you can
visit
[https://v0.2--metallb.netlify.com](https://v0.2--metallb.netlify.com).

When editing the website, you can preview your changes locally by
installing [Hugo](https://gohugo.io/) and running `hugo server` from
the `website` directory.