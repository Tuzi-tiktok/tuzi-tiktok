-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`
(
    `id`         BIGINT   NOT NULL AUTO_INCREMENT COMMENT '消息id',
    `uid`        BIGINT   NOT NULL COMMENT '用户id',
    `vid`        BIGINT   NOT NULL COMMENT '视频id',
    `content`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX        `uid`(`uid`) USING BTREE,
    INDEX        `vid`(`vid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`
(
    `id`         BIGINT   NOT NULL AUTO_INCREMENT,
    `uid`        BIGINT   NOT NULL,
    `vid`        BIGINT   NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`
(
    `id`           BIGINT   NOT NULL AUTO_INCREMENT,
    `to_user_id`   BIGINT   NOT NULL COMMENT '消息接收者id',
    `form_user_id` BIGINT NULL DEFAULT NULL COMMENT '消息发送者id',
    `content`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '消息内容',
    `created_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime NULL DEFAULT NULL,
    `deleted_at`   datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`
(
    `id`           BIGINT   NOT NULL AUTO_INCREMENT,
    `follower_id`  BIGINT   NOT NULL COMMENT '关注者id',
    `following_id` BIGINT   NOT NULL COMMENT '被关注者id',
    `created_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime NULL DEFAULT NULL,
    `deleted_at`   datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX          `follower_id`(`follower_id`) USING BTREE,
    INDEX          `following_id`(`following_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`               BIGINT                                                        NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `username`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
    `password`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `follow_count`     BIGINT                                                        NOT NULL DEFAULT 0 COMMENT '关注数',
    `follower_count`   BIGINT                                                        NOT NULL DEFAULT 0 COMMENT '粉丝数',
    `avatar`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像地址',
    `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户个人页顶部大图',
    `signature`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人签名',
    `created_at`       datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       datetime NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`       datetime NULL DEFAULT NULL COMMENT '删除时间 支持软删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`
(
    `id`             BIGINT                                                        NOT NULL AUTO_INCREMENT COMMENT '视频id',
    `author_id`      BIGINT                                                        NOT NULL COMMENT '上传用户id',
    `title`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
    `cover_url`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面url',
    `play_url`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放url',
    `favorite_count` BIGINT                                                        NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comment_count`  BIGINT                                                        NOT NULL DEFAULT 0 COMMENT '评论数',
    `created_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime NULL DEFAULT NULL,
    `deleted_at`     datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX            `author_id`(`author_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
