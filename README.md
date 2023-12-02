
<h1 align="center">
cwon
</h1>

A CLI for generating CRON expressions from natrual english phrases in Go. 

It works by making a GET request to [cron prompt](https://cronprompt.com/).

<h5 align="center">
🚧Warning🚧
</h5>

Cron Prompt uses a Large Language Model to generate its Cron expressions. The resulting cron expressions may be innacurate or incomplete as a result. 
It may, however, be capable of handling more complex expressions than more deterministic libraries which use Pushdown automata. 

## Building 

```
go build
```

## Usage
```
cwon New years eve
0 0 31 12 *
cwon Every minute during Halloween
0 31 10 * *
```

### Cron format

```
# ┌───────────── minute (0–59)
# │ ┌───────────── hour (0–23)
# │ │ ┌───────────── day of the month (1–31)
# │ │ │ ┌───────────── month (1–12)
# │ │ │ │ ┌───────────── day of the week (0–6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * *
```
