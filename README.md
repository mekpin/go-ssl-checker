# go-ssl-checker

a service created using golang made to be deployed in a docker environment, go-ssl-checker is built for checking ssl expiration date of your managed domains and run using cron every night. if there is a domain with less than the threshold days, it will notify you via slack (configured using slack webhook)

output :
1. remaining days of your listed manifests.yml ssl to expired
2. slack notification to a env defined slack channel webhook

how to use :
1. Create manifest.yml with a list of your managed domains (see manifest/example.yml) and put it on /inventory/manifest.yml
2. Mount your inventory folder if you had to in your deployment yml
3. you can test to hit all the endpoint in the router/router.go files, its all a harmless get API method
4. run it in your instance for daily reminder

list of ENV
PORT > envDefault:"3300"`
THRESHOLD > envDefault:"30"`
CRON_INTERVAL > envDefault:"daily"`
SLACK_WEBHOOK > envDefault:""`
ENABLE_CRON > envDefault:"false"`
TIME_FORMAT > envDefault:"15:04:05 MST 02 Jan 2006"`
TIME_FORMAT > envDefault:"Asia/Jakarta"`
MANIFEST_PATH > envDefault:"inventory/manifestSSL.yaml"`
