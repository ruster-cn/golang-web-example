DROP TABLE IF EXISTS t_cluster_meta;
CREATE TABLE t_cluster_meta (
                                id  INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '序号',
                                created_at DATETIME COMMENT '创建时间',
                                updated_at DATETIME COMMENT '更新时间',
                                deleted_at DATETIME COMMENT '删除时间',
                                name VARCHAR(255) NOT NULL COMMENT '集群名称',
                                introduce TEXT  COMMENT '说明',
                                region VARCHAR(255) COMMENT 'region',
                                az VARCHAR(255) COMMENT 'az',
                                department VARCHAR(255) COMMENT '所属部门',
                                environment VARCHAR(255) COMMENT '所属环境',
                                kube_conf TEXT COMMENT 'kubeconfig文件内容',
                                attribute TEXT COMMENT '属性',
                                UNIQUE KEY unique_name (name)
) COMMENT '集群元数据';