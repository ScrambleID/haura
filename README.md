# Haura 

Haura is pure go simple concurrent lock-free in-memory queue.

## Features
- Currently support API (`enqueue`, `dequeue`, `stats`)
- Communication beetween client over HTTP Protocol

## Installation

    go get -u github.com/vural/haura

## Running

It's so simple!

```
# Default 0.0.0.0:8080
haura
# Or Specific Address
haura -listen :5353
```

## Usage
```
# Enqueue
curl -H "Content-Type: application/json" -X POST -d '{"data":"my_data"}' http://localhost:8080/enqueue
# Dequeue
curl -X GET http://localhost:8080/dequeue
# Stats
curl -X GET http://localhost:8080/stats
```

## Example Stats
```javascript
{
    "go_version": "go1.9.2",
    "host": "0.0.0.0:8080",
    "start_time": 1515840014,
    "total_enqueue": 100,
    "total_dequeue": 50,
    "queue_count": 50
}
```

Cheers !
