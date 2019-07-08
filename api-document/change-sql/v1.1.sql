------------标签表--------------------
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` datetime(0) COMMENT '创建时间',
  `created_by` int COMMENT '创建人',
  `modified_on` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` int COMMENT '修改人',
  `is_deleted` tinyint(1) unsigned DEFAULT '0',
  `state` tinyint(1) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT='文章标签管理';
-------------------------------------

-------------文章表------------------
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` datetime(0),
  `created_by` int COMMENT '创建人',
  `modified_on` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `modified_by` int COMMENT '修改人',
  `is_deleted` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT='文章管理';
-------------------------------------

-------------用户表-------------------
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES (0, 'admin', '123456');

-------------------------------------
