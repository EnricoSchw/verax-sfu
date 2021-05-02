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
    "id" : "MyClientId"
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
    "name": "MyName",
    "id": "MyClientId"
  }
}
```
#### Server -> Client
```json
{
  "type": "conference",
  "name": "joinResponse",
  "data": {
    "id": "MyClientId",
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
    "id": "RemoteClientId",
    "peerName": "Name of peer",
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
    "id": "MyClientId"
  }
}
```

#### Server -> Client

```json
{
  "type": "conference",
  "name": "leaveResponse",
  "data": {
    "id": "MyClientId",
    "room": "name"
  }
}
```

#### Server -> Broadcast Peers

```json
{
  "type": "conference",
  "name": "peerLeave",
  "data": {
    "id": "RemoteClientId",
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
  "data": {"room": "name", "to": "RemotePeerId", "id": "MyClientId"},
  "signal": {
    "desc": {"type": "offer","sdp": "v=0\r\no=- 2057321257886006919..."}
  }
}
```

### 3.2 Ice Candidates

#### Client -> Server -> Broadcast Peers

```json
{
  "type": "signal",
  "name": "trickle",
  "data": {"room": "name", "candidate": "string", "id": "ClientId"},
  "signal": {
    "ice": {"candidate": "candidate:1176663647 1 udp...", "sdpMid":"0","sdpMLineIndex":0 }
  }
}
```

## 4 System

### 4.1 Peer Disconnected

#### Server -> Broadcast Peers

Peer closed socket connection

```json
{
  "type": "system",
  "name": "peerDisconnected",
  "data": {
    "id": "RemoteClientId",
    "room": "name"
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
