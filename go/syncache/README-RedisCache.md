# 期待的来了 redis缓存和数据库的数据一致性解决方案   

* 使用策略共**三种**:  
  * 单节点：单例模式(加读写锁🔒)
  * 主从结点：延时双删（虽然有主从延时的问题存在，但是如果使用polarDB的话，基本就解决了）、MQ异步更新缓存（这个比较难，我有限实现前面两种）
* 使用go语言实现

---

## 话不多说直接开始

### sql 体系树脚本
```sql
# 生成一个体系树🌲 每个层级都有一个父级的parent_id
# 所有的一级标签的parent_id都是1
# 这个脚本有30个一级标签 每个一级标签有5个二级标签 每个二级标签都有5个三级标签

CREATE TABLE label_tree (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT DEFAULT NULL,
    level INT NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES label_tree(id)
);

truncate label_tree;


DROP PROCEDURE IF EXISTS generate_label_tree;
DELIMITER //


CREATE PROCEDURE generate_label_tree()
BEGIN
    DECLARE i INT DEFAULT 1;
    DECLARE j INT DEFAULT 1;
    DECLARE k INT DEFAULT 1;
    DECLARE parent_id_level1 INT DEFAULT 1; -- 一级标签的 parent_id 设置为 1
    DECLARE parent_id_level2 INT DEFAULT NULL;

    -- 插入一级标签
    WHILE i <= 30 DO
        INSERT INTO label_tree (name, parent_id, level) VALUES (CONCAT('Label Level 1 - ', i), 1, 1);
        SET parent_id_level1 = LAST_INSERT_ID();

        -- 插入二级标签
        WHILE j <= 5 DO
            INSERT INTO label_tree (name, parent_id, level) VALUES (CONCAT('Label Level 2 - ', i, '.', j), parent_id_level1, 2);
            SET parent_id_level2 = LAST_INSERT_ID();

            -- 插入三级标签
            WHILE k <= 5 DO
                INSERT INTO label_tree (name, parent_id, level) VALUES (CONCAT('Label Level 3 - ', i, '.', j, '.', k), parent_id_level2, 3);
                SET k = k + 1;
END WHILE;

            SET k = 1;
            SET j = j + 1;
END WHILE;

        SET j = 1;
        SET i = i + 1;
END WHILE;
END //

DELIMITER ;


INSERT INTO label_tree (id, name, parent_id, level) VALUES (1, 'root', NULL, 0);

CALL generate_label_tree();

select * from label_tree where parent_id  =32;


```
