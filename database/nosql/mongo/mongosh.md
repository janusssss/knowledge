### 概述

​	`mongosh` 是MongoDB 官方推出的新一代交互式 Shell 工具，用于替代旧版的 `mongo` Shell。它专为开发者设计，提供了更友好、更强大的命令行体验，支持现代 JavaScript 语法和丰富的扩展功能。

### Run  Commands

| command    | 作用                                   |
| ---------- | -------------------------------------- |
| `db`       | 显示当前数据库                         |
| `show dbs` | 数据库列表                             |
| `use <db>` | 切换数据库或创建（插入数据时才真创建） |

### insert

| 语句                       | 作用     |
| -------------------------- | -------- |
| db.collection.insertOne()  | 单条插入 |
| db.collection.insertMany() | 批量插入 |

```javascript
db.movies.insertOne(
  {
    title: "The Favourite",
    genres: [ "Drama", "History" ],
    runtime: 121,
    rated: "R",
    year: 2018,
    directors: [ "Yorgos Lanthimos" ],
    cast: [ "Olivia Colman", "Emma Stone", "Rachel Weisz" ],
    type: "movie"
  }
)

db.movies.insertMany([
   {
      title: "Jurassic World: Fallen Kingdom",
      genres: [ "Action", "Sci-Fi" ],
      runtime: 130,
      rated: "PG-13",
      year: 2018,
      directors: [ "J. A. Bayona" ],
      cast: [ "Chris Pratt", "Bryce Dallas Howard", "Rafe Spall" ],
      type: "movie"
    },
    {
      title: "Tag",
      genres: [ "Comedy", "Action" ],
      runtime: 105,
      rated: "R",
      year: 2018,
      directors: [ "Jeff Tomsic" ],
      cast: [ "Annabelle Wallis", "Jeremy Renner", "Jon Hamm" ],
      type: "movie"
    }
])