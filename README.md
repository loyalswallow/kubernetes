# Kubernetes with Object Metadata Tagging
This repository is a fork of the official [Kubernetes repository](https://github.com/kubernetes/kubernetes).

As a side project, I have added support for tagging object metadata and extended the kubectl CLI with a tag command to simplify adding or deleting tags.  

I worked on this feature hoping to implement something similar to AWS tagging in Kubernetes, but without the side effects associated with labels and annotations.

For detailed instructions, see the whole flow in my [blog posts](https://medium.com/@loyalswallow21/tagging-kubernetes-objects-9774500cb214)