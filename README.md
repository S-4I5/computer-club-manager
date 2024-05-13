# Computer club manager

## Description

Prototype of system for managing a computer club, capable of analyzing client commands and calculating total profit and time of occupation for each computer.
Currently suited for file input and command line output.

## Run

### Build executable file
```
go build -o task ./cmd/main.go
```
### Run it
```
task {some-filename}
```

## Deployment via docker

### Build docker image
```
docker build -t computer-club-manager:0.1.0 -f ./build/computer-club-manager.dockerfile .
```

### Run image
Via env variable `filepath` you can define input file for app.
```
docker run --name computer-club-manager -e filepath={some-file-path} computer-club-manager:0.1.0
```

## Tests

Tests are available for App struct, covering most possible cases. 

### Run tests

```
go test computer-club-manager/internal/app
```

## Workflow
To start, you need to create a configuration for the service:

### Config

First goes the number of tables with computers.
Then two dates in the format hours:minutes: the start and end time of operation.
And the last value: the cost of using a computer per hour.

Next comes a series of messages.

### Messages

Message format:
```
XX:XX {command number} {args}
```

### Command types

* ID 1. Client arrived
```
{time} 1 {client name}
```
If the client is already in the computer club, an error "YouShallNotPass" is generated.
If the client arrives during non-working hours, then "NotOpenYet".
* ID 2. Client sat at a table
```
{time} 2 {client name} {table number}
```
If the client is already sitting at a table, they can change tables.
If table is occupied (including if the client tries to move to a table they are already sitting at), an error "PlaceIsBusy" is generated.
If the client is not in the computer club, an error "ClientUnknown" is generated.
* ID 3. Client is waiting
 ```
{time} 3 {client name}
 ```
If there are free tables in the club, an error "ICanWaitNoLonger!" is generated.
If the waiting queue exceeds the total number of tables, the client leaves and event ID 11 is generated.
* ID 4. Client left
 ```
{time} 4 {client name}
 ```
If the client is not in the computer club, an error "ClientUnknown" is generated.
When the client leaves, the table they were sitting at is vacated and the first client from the waiting queue occupies it (ID 12).

All incoming messages will be duplicated, accompanied by outgoing messages.

### Outgoing messages
* ID 11. Client left
 ```
{time} 11 {client name}
```
Generated at the end of the working day for all clients remaining in the computer club, in alphabetical order of their names. Also, when a client joins the queue and the waiting queue is already full.
* ID 12. Client sat at a table
 ```
{time} 12 {client name} {table number}
 ```
Generated for the first client in the queue when any table becomes available.
* ID 13. Error
 ```
{time 13 error}
 ```
Displayed immediately after the event that caused the error. The event that caused the error is considered not executed, and has no effect on the state of the clients.

### Calculation of profit and occupation time

For each table, on a separate line, the following parameters are outputted separated by spaces: Table number, Revenue for the day, and Time it was occupied during the working day.

For each hour spent at the table, the client pays the price specified in the configuration. When paying, the time is rounded up to the nearest hour, so even if the client occupied the table for only a few minutes, they pay for a whole hour. Revenue is the sum obtained from all clients throughout the computer club's operation.