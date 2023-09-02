CREATE TABLE user (
      id bigint AUTO_INCREMENT,
      user_name varchar(255) NULL COMMENT '用户名',
      password varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
      gender char(10) NOT NULL DEFAULT 'male' COMMENT '性别 male|female|unknown',
      phone varchar(11) NOT NULL DEFAULT '' COMMENT '手机号码',
      email varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
      nick_name varchar(100) NULL DEFAULT '' COMMENT '昵称',
      google_secret varchar(100) NULL DEFAULT '' COMMENT 'google密钥',
      last_login_time timestamp NULL COMMENT '最后登录时间',
      register_time timestamp NULL COMMENT '注册时间',
      update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      UNIQUE phone_index (phone),
      UNIQUE email_index (email),
      UNIQUE user_name_index (user_name),
      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user用户表';