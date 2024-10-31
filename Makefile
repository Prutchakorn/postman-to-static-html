filename=$1

postman-to-openapi:
	p2o $(filename) > data.yaml

yaml-to-json:
	go run yaml-to-json.go -f=data.yaml > data.json
	rm data.yaml


json-to-js:
	@echo "var temp =" > ./dist/swagger.json.js && cat data.json | sed 's/^/  /' >> ./dist/swagger.json.js && echo ";" >> ./dist/swagger.json.js
	rm data.json


static-html:
	html-inline ./dist/index.html > doc.html

postman-to-static-html: postman-to-openapi yaml-to-json json-to-js static-html
