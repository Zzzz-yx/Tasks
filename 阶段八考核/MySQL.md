# MySQL

简介：MySQL是一个关系型数据库管理系统，关系数据库将数据保存在不同的表中，而不是将所有数据放在一个大仓库内，这样就增加了速度并提高了灵活性。MySQL所使用的 SQL 语言是用于访问数据库的最常用标准化语言。MySQL 软件采用了双授权政策，分为社区版和商业版，由于其体积小、速度快、总体拥有成本低，尤其是开放源码这一特点，一般中小型和大型网站的开发都选择 MySQL作为网站数据库

应用环境：与其他的大型数据库例如 Oracle、DB2、SQL Server等相比，MySQL [1] 自有它的不足之处，但是这丝毫也没有减少它受欢迎的程度。对于一般的个人使用者和中小型企业来说，MySQL提供的功能已经绰绰有余，而且由于 MySQL是开放源码软件，因此可以大大降低总体拥有成本。

Linux作为操作系统，Apache 或Nginx作为 Web 服务器，MySQL 作为数据库，PHP/Perl/Python作为服务器端脚本解释器。由于这四个软件都是免费或开放源码软件FLOSS，因此使用这种方式不用花一分钱（除开人工成本）就可以建立起一个稳定、免费的网站系统，被业界称为“LAMP“或“LNMP”组合。

## 数据表结构

示例

```
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    hire_date DATE
);
```

在上面的示例中，我们创建了一个名为"employees"的数据表，它包含以下列：

- id：整数类型，用作主键，并且使用`AUTO_INCREMENT`属性自动生成唯一的值。
- first_name：VARCHAR(50)类型，用于存储员工的名字，最大长度为50个字符。
- last_name：VARCHAR(50)类型，用于存储员工的姓氏，最大长度为50个字符。
- email：VARCHAR(100)类型，用于存储员工的电子邮件地址，最大长度为100个字符。
- hire_date：DATE类型，用于存储员工的入职日期。

## 建立关联

在MySQL中，可以使用外键来建立表与表之间的关联关系。外键是一个指向另一个表中主键的列，用于确保数据的完整性和一致性。下面是一个示例，展示如何在两个表之间建立关联：

假设我们有两个表：`orders`和`customers`。`orders`表存储订单信息，`customers`表存储客户信息。每个订单都需要关联到特定的客户。我们可以在`orders`表中添加一个`customer_id`列作为外键，指向`customers`表的主键。

首先，我们创建`customers`表：

示例

```
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100)
);
```

然后，在`orders`表中添加`customer_id`列作为外键：

```
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(50),
    order_date DATE,
    customer_id INT,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);
```

在上面的示例中，`customer_id`列被定义为外键，它引用了`customers`表的`id`列作为主键。这样，每个订单都可以通过`customer_id`与特定的客户相关联。

通过建立外键关系，MySQL会在插入或更新数据时自动验证外键的有效性，确保只有存在于`customers`表中的有效`id`才能被插入到`orders`表的`customer_id`列中。这样可以保证数据的一致性和完整性。

需要注意的是，在MySQL中，使用外键需要满足一些条件，例如表的存储引擎必须是InnoDB，且外键列和参考列必须具有相同的数据类型和长度等。

## CURD

在SQL中，CURD代表创建(Create)、读取(Read)、更新(Update)和删除(Delete)数据的操作。下面是每个操作对应的SQL查询语句示例：

创建(Create)操作：

```
INSERT INTO table_name (column1, column2, column3)
VALUES (value1, value2, value3);
```

读取(Read)操作：

```
SELECT column1, column2, column3
FROM table_name
WHERE condition;
```

更新(Update)操作：

```
UPDATE table_name
SET column1 = value1, column2 = value2
WHERE condition;
```

删除(Delete)操作：

```
DELETE FROM table_name
WHERE condition;
```

上述示例中，你需要将table_name替换为你要操作的表名，column1, column2, column3替换为相应的列名，value1, value2, value3替换为要插入或更新的值，condition替换为适当的条件。

需要注意的是，SQL查询语句的具体语法可能因数据库管理系统的不同而有所差异。上述示例是通用的SQL语法，但在实际使用时，你可能需要根据所使用的数据库管理系统的要求进行调整。另外，还可以使用其他SQL语句和子句来实现更复杂的查询和操作。