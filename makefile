.PHONY: slides

slides:
	pandoc -t revealjs -s -o ./slides/output/slides.html ./slides/slides.md -V revealjs-url=./reveal.js -V theme=white

k8s:
	kubectl apply -f ./templates/deployment.yml ./templates/service.yml ./templates/ingress.yml
