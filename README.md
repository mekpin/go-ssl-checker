# go-ssl-checker

a service created using golang made to be deployed in a docker environment, go-ssl-checker is built for checking ssl of your domains on daily basis and notify you via slack if there is remaining ssl live (in days) threshold reached .

output :
1. remaining days of ssl to expired
2. slack notification to a env defined slack channel webhook

how to use :
1. create manifest.yml inside /inventory/manifest.yml
2. mount your inventory folder if you had to in your deployment yml
3. hit all the endpoint in the router/router,go files
