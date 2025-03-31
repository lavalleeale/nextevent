# nextevent

A simple command-line tool that shows the next event from an ICS calendar file occurring today.

## Why?

This was developed to add a module to my status bar that could show the next event in my calendar. I wanted to be able to see at a glance what was coming up next without having to open a calendar app.

## Installation

```bash
go install github.com/lavalleeale/nextevent@latest
```

Or clone and build manually:

```bash
git clone https://github.com/lavalleeale/nextevent.git
cd nextevent
go build
```

## Usage

```bash
./nextevent calendar.ics
```

The program will:

1. Read the specified ICS file
2. Find all events happening today (including recurring events)
3. Show the next upcoming event and how long until it starts
4. If an event started in the last 5 minutes, it will show how long ago it started

## Example Output

```
Daily Standup in 45m
```

or

```
Team Meeting 2m ago
```

## Dependencies

- [gocal](https://github.com/apognu/gocal) - For parsing ICS files
