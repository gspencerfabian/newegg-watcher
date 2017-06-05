# newegg-watcher
A Golang inventory monitoring/notification script for newegg.com

## How to install
Be sure that golang is installed. If not, get it [here](https://golang.org/dl/).

**Clone and build:**
`git clone https://github.com/gspencerfabian/newegg-watcher.git
cd newegg-watcher
make build`

**Run:**
`make run`

**Schedule to run via cron.d:**
`*/5 * * * * cd /home/<username>/work/src/github.com/gspencerfabian/newegg-watcher && ./newegg-watcher >> newegg-watcher.log 2>&1`
Modify path to point to your cloned repository

