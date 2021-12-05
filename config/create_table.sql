CREATE TABLE `tbl_blog_user`(
  `uid` BIGINT(20) NOT NULL,
  `user_name` VARCHAR(64) NOT NULL,
  `passwd` VARCHAR(64) NOT NULL,
  `avatar` VARCHAR(64),
  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `idx_user_name` (`user_name`) USING BTREE,
  UNIQUE KEY `idx_uid` (`uid`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `tbl_blog_category`(
    `cid` int not null AUTO_INCREMENT,
    `name` varchar(64) not null,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`cid`),
    unique key idx_cid (`cid`) using btree
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `tbl_blog_post`(
  `pid` int not null AUTO_INCREMENT,
  `title` varchar(255) not null,
  `content` varchar(10000) not null,
  `markdown` varchar(10000) not null,
  `category_id` int not null,
  `user_id` BIGINT(20) not null,
  `view_count`  int not null DEFAULT 0,
  `create_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  primary key (pid),
  unique key idx_pid (`pid`) using btree
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;