# Kubernetes Introduction Talk Material

Support material (slides, examples) for the Introduction to Kubernetes talk.

## Slides
If you want to access the slides in a better format than a markdown file that doesn't make much sense, you will need to have `pandoc` installed, follow the instructions [here](http://pandoc.org) to install it.

Once you installed it, running the following command fill create a HTML version of the slides in `slides/output`

```
$ make slides
```

## Kubernetes
If you want to try out the demo that was performed with the talk, you will need to have `minikube` and the `kubernetes-cli` (`kubectl`) applications installed on your machine. You can find the latter [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and the former [here](https://github.com/kubernetes/minikube).

Make sure you change the `path` value in `templates/00-deployment.yml` file to point it to where you cloned this repo or the application will not be deployed correctly. Once you performed all this, start minikube with `minikube start` and simply run the following command and everything will be performed for you.

```
$ make k8s
```

If you want to access your application via an URL, you will need to add the following line to your `/etc/hosts` file:


```
192.168.99.100 example.meetup.com
```

P.S. Note that `192.168.99.100` is the standard IP address for a minikube cluster. Verify that it's the correct one by running `minikube ip` before changing the `hosts` file.