# pms
基于beego的进销存系统
数据库的结构参考了odoo的设计
由于公司销售的是黄金饰品，产品管理上面需要两个单位来管理，故在设计上直接使用了两个单位

数据库采用的是postgresql，版本中的数据库连接到自己的阿里云服务器，
包含了全国省市区的地址信息，初始化数据在init_xml中，包括系统管理员的帐号信息
