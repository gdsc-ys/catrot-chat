# catrot-chat
당근마켓 클론 코딩 서버의 채팅 기능을 담당한다.

</br>

## Tech Spec
- Framework : [Gin](https://github.com/gin-gonic/gin)

- Go Version : go 1.19.1 darwin/amd64

- DataBase : MySQL (8.0.30)

</br>

## Database Schema
### chat_room
```
CREATE TABLE `chat_room` (
  `room_id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_uid` int NOT NULL,
  `room_type` varchar(10) DEFAULT NULL,
  `prev_uid` int DEFAULT NULL,
  `prev_msg` varchar(512) DEFAULT NULL,
  `reg_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`room_id`),
  KEY `create_uid` (`create_uid`),
  KEY `room_type` (`room_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```


### chat_room_member
```
CREATE TABLE `chat_room_member` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `room_id` int DEFAULT NULL,
  `uid` int DEFAULT NULL,
  `room_name` varchar(512) NOT NULL,
  `push_state` varchar(2) NOT NULL DEFAULT 'Y',
  `reg_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_room_id_uid` (`room_id`,`uid`),
  KEY `room_id` (`room_id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

### messages
```
CREATE TABLE `messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `msg` varchar(512) NOT NULL,
  `send_uid` int NOT NULL,
  `room_id` int NOT NULL,
  `messageType` varchar(10) DEFAULT NULL,
  `reg_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `send_uid` (`send_uid`),
  KEY `room_id` (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

### unread_messages
```
CREATE TABLE `unread_messages` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `room_id` int NOT NULL,
  `uid` int NOT NULL,
  `last_mid` int DEFAULT NULL,
  `reg_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_room_id_uid` (`room_id`,`uid`),
  KEY `room_id` (`room_id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
