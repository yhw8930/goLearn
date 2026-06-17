# goLearn

个人 Go 语言学习仓库。Go 1.22,`module goLearn`。

仓库按**主题**组织,每个目录是一个独立的学习单元,大多可单独运行。
内容覆盖:语言基础 → 并发 → 标准库 → Web(Gin)/ORM(GORM) → 算法练习。

## 目录索引

| 目录 | 内容 |
|------|------|
| [`basic/`](basic) | 语言基础:分支、循环、函数、正则、atomic |
| [`chapter1/`](chapter1) | 《Go 程序设计语言》第 1 章示例 |
| [`consts/`](consts) | 常量与时间格式化 |
| [`channel/`](channel) | 并发:channel、select、done 通知 |
| [`container/`](container) | 数组 / 切片 / map / 字符串,含最长不重复子串 |
| [`functional/`](functional) | 函数式:闭包、adder、斐波那契生成器 |
| [`errhandling/`](errhandling) | 错误处理:defer、recover、errWrapper 模式 |
| [`crawler/`](crawler) | 并发爬虫(engine / fetcher / scheduler / persist 分层) |
| [`retriever/`](retriever) | 面向接口编程:real 与 mock 两套实现 |
| [`queue/`](queue) | 队列(含泛型与测试) |
| [`tree/`](tree) | 二叉树定义与遍历 |
| [`maze/`](maze) | BFS 走迷宫 |
| [`stdlib/`](stdlib) | 标准库:time、file、strconv 等 |
| [`http/`](http) | net/http 基础 |
| [`gin/`](gin) | Gin Web 框架:路由、参数绑定、中间件、校验、优雅关闭 |
| [`gorm/`](gorm) | GORM:数据库连接与 CRUD(需数据库) |
| [`algorithm/lession/`](algorithm/lession) | 算法练习 class01 ~ class24 |

## 运行

每个子目录独立运行:

```bash
go run ./gin/middleware            # 运行某个 demo
go run ./algorithm/lession/class24 # 运行算法某一课
go test ./queue/...                # 跑测试
```

> `gorm/` 示例需要可用的数据库连接;`gin/other_autotis` 涉及 HTTPS 证书,本地运行需额外配置。

---

> 项目结构与 AI 协作约定详见 [CLAUDE.md](CLAUDE.md)。
