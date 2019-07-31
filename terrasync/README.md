# Terrasync

Terrasync is a simple go executable which continuously syncs your Terraform
config against your infrastructure.  It is intended to be used with the
[kubernetes/git-sync](https://github.com/kubernetes/git-sync) sidecar in order
to pull an up-to-date configuration from git.

## Example usage

This repo has an example under the `manifest/` directory which should get you
started using Terrasync.

First, simply build the docker image using the provided Dockerfile:

```bash
docker build -t 'terrasync:v0.1.0' .
```

Once the image has been built, go ahead and modify the kubernetes deployment
object found at `manifests/deploy.yml`.  You will likely need to modify the git
repository and branch being synced, as well as the location of the Terraform
root module using the `-dir` argument to Terrasync.

With the manifest modified for your use case, we can apply it to a cluster.
For this example we'll be using
[kind](https://github.com/kubernetes-sigs/kind):

```bash
# Create the cluster
kind create cluster

# Point kubectl at the local cluster
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"

# Sideload our docker images
docker pull k8s.gcr.io/git-sync:v3.1.1
kind load docker-image k8s.gcr.io/git-sync:v3.1.1
kind load docker-image terrasync:v0.1.0

# Apply the manifest
kubectl apply -f manifests/deploy.yml
```

At this point you should have Terrasync running in your cluster and can monitor
its progress with `kubectl logs`.
