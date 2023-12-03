
<h1 align="center">
<img src='./assets/cwon.png'>
</h1>

<p align="center">
A CLI for generating CRON expressions from natural english phrases in Go. 
</p>


It works by making a GET request to [cron prompt](https://cronprompt.com/).

<h5 align="center">
🚧Warning🚧
</h5>

Cron Prompt uses a Large Language Model to generate its Cron expressions. The resulting cron expressions may be innacurate or incomplete as a result. 
It may, however, be capable of handling more complex expressions than more deterministic libraries which use Pushdown automata. If a result seems slightly off, try to prompt the model with different wording.

Validate expressions with a tool like [Crontab](https://crontab.guru/) before use in any application.

## Building 

```
go build
```

## Usage
```
cwon at 8pm on Christmas
0 20 25 12 *
cwon During work hours
0 9-17 * * *
cwon At every minute divisible by 5 at work hours
*/5 9-17 * * *
cwon New years eve
0 0 31 12 *
cwon Every year
0 0 1 1 *
cwon At every minute divisible by 5
*/5 * * * *
cwon At 4:45 Saturday and Sunday
45 4 * * 0,6
cwon Every half hour
0,30 * * * *
cwon At 8pm during Halloween
0 20 31 10 *
cwon At 8pm during October 31st
0 20 31 10 *
cwon When McDonalds opens every day
0 6 * * *
cwon 30 minutes before 5am
30 4 * * *
cwon 30 minutes before 5th Fibbonaci Number am
30 4 * * *
```

### Example of a malformed result

```
cwon At every minute on October 31st
* 31 10 * *
```

this is incorrect, as the hours field is 31. The correct result should be * * 31 10 *.

to remedy this, I used the following prompt structure (because this is how most 'cron translators' would read out * * 31 10 *)
```
cwon At every minute at day-of-month 31 in October
* * 31 10 *
```

Sometimes, the model will generate incorrect responses but will return correct responses when re-prompted with the same prompt.

(a consequence of the probabalistic nature of LLMs)

```
cwon At every minute at day-of-month 31 in October
* 31 10 * *
cwon At every minute at day-of-month 31 in October
* * 31 10 *
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
