% Introduction to Kubernetes
% Pietro Bongiovanni
% Digital Ocean Amsterdam, the April edition

# What are we going to talk about?

* What is Kubernetes?
* Why should I care about Kubernetes?
* What are the building blocks of Kubernetes?
* How do I get started?
* Demo!
* (Q&A)

---

# whoami

Pietro Bongiovanni

Software Engineer @ Fixico

Computer Engineering and CS @ Politecnico di Torino

twitter.com/pietro_goodjohn

\#DOMeetup

https://blog.pietrobongiovanni.com

---

# What Is K8S?

*Kubernetes is a 10 ton dump truck, but it handles pretty well on the track at 321 kph.* [*](http://crunchtools.com/kubernetes-10-ton-dump-truck-handles-pretty-well-200-mph/)

---

# What is K8S?

Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications. 

([kubernetes.io](https://kubernetes.io))

---

## What is K8S?

It makes easy to bbuild applications based on microservices

---

## What is K8S?

Basically, it's cool AF.

---

# Why should you care about Kubernetes?

* It works
* It's backed by some very (very) big companies
* It is Open Source
* It has an amazing Open Source Community

---

## Why should you care about Kubernetes?

**It Works**

---

### Why should you care about Kubernetes?

![](img/pokemon-go.png)

---

## Why should you care about Kubernetes?

It's backed by big companies, yet still open source!

---

## Why should you care about Kubernetes?

![](img/react.png)

---

## Why should you care about Kubernetes?

![](img/vscode.png)

---

## Why should you care about Kubernetes?

![](img/tensorflow.png)

---

## Why should you care about Kubernetes?

![](img/kubernetes.png)

---

## Why should you care about Kubernetes?

It's a bible for open source projects and communities

---

# Let's get into the weeds of it

---

## How does it work?

![](img/architecture-diagram.png)

---

## Nodes

* Master/Slave(s) Nodes
* Physical or Virtual Machines

---

## Pods

* Basic building blocks of K8S
* Docker Container(s)
* Horizontal scalability

---

## Pod Controllers

* Deployments
* Stateful Sets
* Daemon Controllers

---

## Pod Lifecycle

* Pod failure
* Pod restarts
* Deployments

---

## Services

* Exposing Pods to HTTP traffic (internal)

---

## Ingresses

* Exposing Services to external HTTP traffic

---

## More stuff

* K8S Offer a lot more
    * Cronjobs (1.8)
    * Secret management
    * Automated rollouts and rollbacks
    * Storage orchestration

---

# Demo!

---

# Q&A?