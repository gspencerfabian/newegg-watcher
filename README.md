# newegg-watcher
A Golang inventory monitoring/notification script for [newegg.com](http://newegg.com)
#### Why?
Because newegg's auto-notify is not always available for certain products. Also, this script is much faster at detecting inventory and sending out emails.

## requirements
 - Be sure that golang is installed. If not, get it [here](https://golang.org/dl/).
 - Google email address. **NOTE** you must allow less secure apps to access your account. [Learn More](https://support.google.com/accounts/answer/6010255?hl=en)

## how to install
```
git clone https://github.com/gspencerfabian/newegg-watcher.git
cd newegg-watcher
make build
```

## how to run
```
make run
```

## how to schedule (cron.d)
```
*/5 * * * * cd /home/<username>/work/src/github.com/gspencerfabian/newegg-watcher && ./newegg-watcher >> newegg-watcher.log 2>&1
```
Modify path to point to your cloned repository

## configuration
Modify config.json file 
 - items field: the newegg item number
 - email fields: sender/receiver email address.

```
{
        "items": [
                "N82E16813157746",
                "N82E16819117728",
                "N82E16835100007",
                "N82E16835181103"
        ],
        "email": {
                "receiver": {
                        "address": [
                                "jsmith@example.com",
                                "5555555555@vtext.com",
                                "5551234567@txt.att.net"
                        ]
                },
                "sender": {
                        "address": "xxxxxx@gmail.com",
                        "password": "xxxxxxxx"
                }
        }
}
```

