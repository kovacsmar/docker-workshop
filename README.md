# 󰡨 Docker Workshop 󰡨

Demo application containerization example with Docker

## Prerequisites

In this demo you need to have Git and Docker installed to follow the project.

###  Git

####  Linux

󰣇 Arch:

```bash
sudo pacman -S git
```

 Deb:

```bash
sudo apt isntall git-all
```

 RPM:

```bash
sudo dnf install git-all
```

####  Windows

[Installation Link](https://git-scm.com/downloads/win)

---

### 󰡨 Docker

#### 󰡨 Docker CE

Docker CE is coming with only cli option

[Docker CE Install](https://docs.docker.com/engine/install/)

#### 󰡨 Docker Desktop

[Docker Desktop on Linux](https://docs.docker.com/desktop/setup/install/linux/)
[Docker Desktop on Windows](https://docs.docker.com/desktop/setup/install/windows-install/)

Installing Docker on Windows can get tricky.
We can go over it on the workshop or beforehand.

---

### About

In this workshop we will containerize an application written in Go.

- Go down on the basics of the hows and whys on Docker
- Make a basic image
- Improve the image
- Securing the image
- Running multiple containers with orchestration
- Public repos and fun apps to run on Docker
- Optimization

##  Kubernetes

If you do not have a Kubernetes cluster you can run it on try out kind ([Install](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)) for running it locally.

> [!WARNING] kind is for testing only, **do NOT** use for production environments, read [documentation](https://kind.sigs.k8s.io/) for more information.

### Prerequisites

- `kubectl` configured against a running cluster [Install](https://kubernetes.io/docs/tasks/tools/#kubectl)

- Traefik installed as ingress controller:

```bash
  helm repo add traefik https://traefik.github.io/charts
  helm install traefik traefik/traefik -n traefik --create-namespace
```

- `docker-workshop.local` resolving to your cluster's ingress IP (via DNS or `/etc/hosts`)

### Install via Helm

Install via the public helm repo

```bash
helm install docker-workshop oci://ghcr.io/kovacsmar/docker-workshop \
  --version 0.1.0 \
  -n docker-workshop \
  --create-namespace
```

### Access

Navigate to `http://docker-workshop.local`
