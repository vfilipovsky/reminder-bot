# WIP: Reminder bot

## TODO:
Parse message

### Getting started

1. `touch reminder.db`
2. `touch .env` put BOT_TOKEN=your-token
3. `make` or `go run main.go`

### Rules

```
Message must starts with !remindme
then you write your reminder text and 
ending it with your desired time

y  - years
mo - months
w  - weeks
d  - days
h  - hours
m  - minutes
```

### Examples

```
!remindme pay for the phone in 2h 30m
!remindme buy a coffee tomorrow at 14:00
!remindme wife birthday on September 30
!remindme play football on Monday at 18:00
```