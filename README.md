# verax-sfu
> WebRTC Selective Forwarding Unit

# Events

### Room (Connect)
### 1. Join Room
#### Client -> Server
```json
{
  "type": "room",
  "name": "connect",
  "data": {
    "room": "name"
  }
}
```
#### Server -> Client
```json
{
  "type": "room",
  "name": "connectResponse",
  "data": {
    "room": "name",
    "id" : "ClientId"
  }
}
```

## 2. Conference

### 2.1. Join Conference

#### Client -> Server

```json
{
  "type": "conference",
  "name": "join",
  "data": {
    "room": "name",
    "id": "ClientId"
  }
}
```
#### Server -> Client
```json
{
  "type": "conference",
  "name": "joinResponse",
  "data": {
    "id": "ClientId",
    "room": "name"
  }
}
```

#### Server -> Broadcast Peers

```json
{
  "type": "conference",
  "name": "peerJoin",
  "data": {
    "id": "ClientId",
    "room": "name"
  }
}
```

### 2.2 Leave Conference

#### Client -> Server

```json
{
  "type": "conference",
  "name": "leave",
  "data": {
    "room": "name",
    "id": "ClientId"
  }
}
```

#### Server -> Client

```json
{
  "type": "conference",
  "name": "leaveResponse",
  "data": {}
}
```

#### Server -> Broadcast Peers

```json
{
  "type": "conference",
  "name": "peerLeave",
  "data": {
    "id": "ClientId",
    "room": "name"
  }
}
```

## 3. Signaling

### 3.1 Session Description

#### Client -> Server -> Peer

```json
{
  "type": "signal",
  "name": "sdp",
  "data": {
    "room": "name",
    "to": "PeerId",
    "id": "ClientId"
  },
  "signal": {
    "ice": "string",
    "desc": "string"
  }
}
```

### 3.2 Ice Candidates

#### Client -> Server -> Broadcast Peers

```json
{
  "type": "signal",
  "name": "trickle",
  "data": {
    "room": "name",
    "candidate": "string",
    "id": "ClientId"
  }
}
```

### 3.3 Renegotiation

####Server -> Broadcast Peers
```json
{
  "type": "signal",
  "name": "renegotiation",
  "data": {
    "room": "name",
    "sdp": "string",
    "id": "ClientId"
  }
}
```

## 3 Verax Signal


#### Client -> Server -> Broadcast Peers

```json
{
  "type": "signal",
  "method": "join",
  "data": {
    "sid": "sid:string",
    "uid": "uid:string",
    "offer": "sdp:string"
  }
}
```


## 4 System

### 4.1 Peer Disconnected

#### Server -> Broadcast Peers

```json
{
  "type": "system",
  "name": "peerDisconnected",
  "data": {
    "id": "ClientId",
    "room": "name"
  }
}
```

### 4.2 Keep Alive

#### Client -> Server

```json
{
  "type": "system",
  "name": "ping",
  "data": {
  }
}
```

#### Server -> Client

```json
{
  "type": "system",
  "name": "pong",
  "data": {
  }
}
```
