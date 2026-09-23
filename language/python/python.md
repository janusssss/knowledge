# pip

-  pip install <package-name> : 安装包
- pip install --upgrade <package-name> : 升级包

# 关键字

| 关键字 | 作用                         |
| ------ | ---------------------------- |
| as     | 用于给对象或异常指定一个别名 |
| raise  | 手动引发（抛出）异常         |
| and    | 逻辑运算符and                |
| pass   | 占位语句                     |
| in     | 用于成员资格测试和迭代       |

## in

在 Python 中，`in` 关键字有多种用途，主要用于成员资格测试和迭代。下面是 `in` 的主要用法：

### 成员资格测试

`in` 可以用来检查一个值是否存在于序列（如字符串、列表、元组）或集合类型（如集合、字典的键）中。它返回一个布尔值：如果找到则为 `True`，否则为 `False`。

```python
# str 字符串
s = "hello" 
print('h' in s)  # 输出: True
print('z' in s)  # 输出: False
```

```python
# list 列表
my_list = ['apple', 'banana', 'cherry']
print('banana' in my_list)  # 输出: True
print('orange' in my_list)  # 输出: False
```

```python
# tuple 元组
my_tuple = (1, 2, 3)
print(2 in my_tuple)  # 输出: True
print(4 in my_tuple)  # 输出: False
```

```python
# conllection 集合
my_set = {'apple', 'banana', 'cherry'}
print('apple' in my_set)  # 输出: True
print('orange' in my_set)  # 输出: False
```

```python
# dirc 字典
my_dict = {'name': 'Alice', 'age': 25}
print('name' in my_dict)  # 输出: True
print('height' in my_dict)  # 输出: False
```

### 迭代（for 循环）

`in` 关键字也用于 `for` 循环中，遍历序列或其他可迭代对象中的元素。

```python
for item in ['apple', 'banana', 'cherry']:
    print(item)

for letter in 'Python':
    print(letter)

for key in {'name': 'Alice', 'age': 25}:
    print(key)

# 如果想要遍历字典的值或键值对，可以使用 .values() 或 .items()
for value in {'name': 'Alice', 'age': 25}.values():
    print(value)

for key, value in {'name': 'Alice', 'age': 25}.items():
    print(f'{key}: {value}')
```

### 在条件语句中结合 not 使用

`not in` 是 `in` 的否定形式，用于检查某个值是否**不在**序列或集合中。

 ```python
fruits = ['apple', 'banana', 'cherry']
if 'orange' not in fruits:
    print("Orange is not in the list.")
 ```

## lambda

`lambda` 关键字用于创建匿名函数，即没有名字的函数。匿名函数通常是在需要一个简单的、一次性的函数时使用，而不必正式地定义一个具有 `def` 语句的函数。`lambda` 函数可以有任意数量的参数，但只能有一个表达式。这个表达式的值是 `lambda` 函数的返回值。

### Lambda 表达式的语法

```python
lambda 参数1, 参数2, ... : 表达式
```

- `参数1, 参数2, ...`：这是传递给 `lambda` 函数的输入参数列表。
- `表达式`：这是一个单一的表达式，`lambda` 函数会计算该表达式的值并将其作为结果返回。

# 符号

| 符号 | 作用 |
| ---- | ---- |
| \|   | 或   |
|      |      |
|      |      |

# 数据类型

| 英文    | 中文   | 概述                                                 |
| ------- | ------ | ---------------------------------------------------- |
| dict    | 字典   | 键值对数据结构                                       |
| int     | 整数   | 表示整数值，没有大小限制                             |
| float   | 浮点数 | 表示带小数点的数值，遵循 IEEE 754 双精度浮点数标准   |
| complex | 复数   | 由实部和虚部组成，用于表示复数                       |
| bool    | 布尔值 | 表示真或假，只有两个可能的值：`True` 和 `False`      |
| str     | 字符串 | 用于表示文本数据，支持 Unicode 编码                  |
| list    | 列表   | 有序、可变的序列，可以包含不同类型的元素             |
| `tuple` | 元组   | 有序、不可变的序列，通常用于存储异构数据             |
| `dict`  | 字典   | 键值对集合，键必须是唯一的且不可变，值可以是任意类型 |
| set     | 集合   | 无序、不重复的元素集合，支持集合操作如并集、交集等   |

## List

### 声明

1. 使用方括号 `[]` 来定义一个列表，并在其中用逗号分隔各个元素：
   ```python
   my_list = [element1, element2, element3]
   ```

3. 使用列表推导式（List Comprehensions）快速生成一个列表：
   ```python
   squared_numbers = [x**2 for x in range(10)]  # 创建一个包含0到9的平方数的列表
   ```

4. 使用 `list()` 构造函数来从其他可迭代对象（如字符串、元组等）创建列表：
   ```python
   list_from_string = list("hello")  # 将字符串转换为列表 ['h', 'e', 'l', 'l', 'o']
   list_from_tuple = list((1, 2, 3))  # 将元组转换为列表 [1, 2, 3]
   ```

5. 如果要创建一个特定长度的列表，并且所有元素都是相同的值，可以使用乘法：
   ```python
   repeated_elements = ['a'] * 5  # 创建一个 ['a', 'a', 'a', 'a', 'a'] 的列表
   ```

6. 使用 `range()` 函数结合 `list()` 构造函数来创建数值列表：
   ```python
   number_list = list(range(5))  # 创建一个 [0, 1, 2, 3, 4] 的列表
   ```

### 方法

#### append()

```python
# 用于在列表的末尾添加一个元素
list.append(element)
"""
list：你想要添加元素的目标列表。
element：任何类型的元素（包括数字、字符串、另一个列表等），它将被添加到列表的最后。
"""
```



## dict

### 声明

1. 使用花括号 `{}` 创建字典：

   ```python
   # 创建一个空字典
    empty_dict = {}  
   #	创建一个带有初始键值对的字典
    person_info = {	
         'name': 'Alice',
        'age': 30,
        'city': 'Beijing'
    }
   ```

2. 使用 `dict()` 构造函数创建字典：

    ```python
    empty_dict = dict()  # 创建一个空字典
    
    # 使用关键字参数创建字典
    person_info = dict(name='Alice', age=30, city='Beijing')
    
    # 使用包含键值对的可迭代对象创建字典
    list_of_tuples = [('name', 'Alice'), ('age', 30), ('city', 'Beijing')]
    person_info = dict(list_of_tuples)
    ```

3. 使用字典推导式 (Dictionary Comprehensions) 创建字典：

    ```python
    # 从已有的列表创建字典
    keys = ['name', 'age', 'city']
    values = ['Alice', 30, 'Beijing']
    person_info = {k: v for k, v in zip(keys, values)}
    ```

4. 使用 `fromkeys()` 方法创建新字典，给定默认值：

    ```python
    # 创建一个所有值都是 None 的字典
    keys = ['name', 'age', 'city']
    person_info = dict.fromkeys(keys)
    
    # 创建一个具有相同默认值的字典
    person_info_default = dict.fromkeys(keys, 'default_value')
    ```

### 遍历

1. 遍历所有键

   ```python
   person_info = {'name': 'Alice', 'age': 30, 'city': 'Beijing'}
   
   for key in person_info:
       print(key)
   ```

2. 使用 `keys()` 方法遍历所有键

   ```python
   for key in person_info.keys():
       print(key)
   ```

3. 使用 `values()` 方法遍历所有值

   ```python
   for value in person_info.values():
       print(value)
   ```

4. 使用 `items()` 方法遍历键值对

   ```python
   for key, value in person_info.items():
       print(f"{key}: {value}")
   ```

5. 列表推导式或其他推导式

   ```python
   # 创建一个新的列表，其中包含原字典中每个键值对组成的字符串
   pairs = [f"{key}={value}" for key, value in person_info.items()]
   print(pairs)  # 输出: ['name=Alice', 'age=30', 'city=Beijing']
   ```

   注意事项

   - 如果在遍历字典的过程中尝试添加或删除键（除了改变现有键的值），会引发运行时错误。

   - 如果需要在遍历时修改字典，一种常见做法是先创建字典的副本再进行修改，或者收集要添加或移除的键，在遍历完成后执行这些操作。




# 语法

| 字符 | 作用                                             |
| ---- | ------------------------------------------------ |
| del  | 用于删除对象、变量、列表中的元素或字典中的键值对 |
|      |                                                  |
|      |                                                  |

## Type Hints

```python
name: type = value
# 为变量、函数参数和返回值指定类型
```



# 配置文件

setup.cfg

setup.py 可以读取 setup.cfg 文件中的配置，并将其传递给 setuptools.setup() 函数。

# 标准包

## collections

### 类

#### defaultdict

```bash
一种特殊字典类型，它继承自内置的 dict 类。
defaultdict 在访问一个不存在的键时不会抛出 KeyError，而是自动创建一个默认值。
```



### 函数

#### Counter()

```python
# Counter 可以接受任何可迭代对象（如列表、字符串、元组等），并自动计算其中每个元素出现的次数
counter = Counter('abracadabra')
print(counter)  
# 输出: Counter({'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1})
```

# 内置函数

## main函数

```python
if __name__ == "__main__":
    # 主程序代码写在这里
```

## 三元表达式

```python
value_if_true if condition else value_if_false
```

```python
x = 10
y = 20

# 使用三元表达式选择较大的数
max_value = x if x > y else y
print(max_value)  # 输出: 20
```

## sorted()

```py
sorted(iterable, *, key=None, reverse=False)
# 对可迭代对象（如列表、元组、字符串等）进行排序，返回一个新的已排序列表
numbers = [5, 2, 9, 1, 5, 6]
sorted_numbers = sorted(numbers)
print(sorted_numbers)  
# 输出: [1, 2, 5, 5, 6, 9]
```

## enumerate()

用于将一个可迭代对象（如列表、元组、字符串等）组合为一个索引序列，常用于在 `for` 循环中获取元素的索引及其对应的值。它返回的是一个枚举对象，该对象是一个迭代器，可以生成由索引和值组成的元组对。

```python
enumerate(iterable, start=0)
# iterable：需要遍历的对象，例如列表、元组、字符串等。
# start：索引起始位置，默认是 `0`，也可以指定其他整数值作为起始索引。
```

```python
# 基本用法
for index, value in enumerate(['apple', 'banana', 'cherry']):
    print(f'Index {index}: {value}')
"""
Index 0: apple
Index 1: banana
Index 2: cherry
"""
```

```python
# 指定起始索引
for index, value in enumerate(['apple', 'banana', 'cherry'], start=1):
    print(f'Index {index}: {value}')
"""
Index 1: apple
Index 2: banana
Index 3: cherry
"""
```

```python
print(list(enumerate(['apple', 'banana', 'cherry'])))
# [(0, 'apple'), (1, 'banana'), (2, 'cherry')]
```

## str()

用于将对象转换为字符串表示，可以接受任何类型的对象，并返回该对象的字符串形式。

```bash
# 将整数转换为字符串
number = 12345
string_number = str(number)
print(string_number)  # 输出: "12345"

# 将浮点数转换为字符串
pi = 3.14159
string_pi = str(pi)
print(string_pi)  # 输出: "3.14159"

# 将布尔值转换为字符串
boolean_value = True
string_boolean = str(boolean_value)
print(string_boolean)  # 输出: "True"
```

## dir()

```python
def dir(__o: object = ...) -> List[str]
# 返回一个列表，包含指定对象的属性和方法名。如果没有参数传递给 dir()，它将列出当前本地作用域内的名称
```

## type()

```py
def __init__(self, __o: object) -> None
# 获取对象的类型
```

`__import__`

```python
__import__(name, globals=None, locals=None, fromlist=(), level=0)
# 动态导入模块
```

