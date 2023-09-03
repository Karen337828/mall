CREATE TABLE message (
    id        bigint AUTO_INCREMENT,
    msg_type  int NULL COMMENT '信息类型 1-邮箱 2-短信',
    scene_no  varchar(10)  NOT NULL DEFAULT '' COMMENT '场景码',
    user_name varchar(255) NULL COMMENT '接收人用户名',
    phone     varchar(11)  NOT NULL DEFAULT '' COMMENT '接收人手机号码',
    email     varchar(50)  NOT NULL DEFAULT '' COMMENT '接收人邮箱',
    title     varchar(200) NOT NULL DEFAULT '' COMMENT '标题',
    content   TINYTEXT NULL COMMENT '发送内容',
    status    int NULL COMMENT '发送状态 0-待发送 1-发送成功 2-发送失败 3-发送中',
    create_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'message信息表';