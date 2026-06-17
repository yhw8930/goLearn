# CLAUDE.md

本文件供 Claude Code / AI 快速理解本仓库。**先读这里,再读代码。**

## 这是什么

`goLearn` 是一个**个人 Go 学习仓库**(`module goLearn`, Go 1.22),不是一个统一的应用。
它由 ~20 个**互相独立的主题目录**组成,每个目录是一个或多个独立的学习单元,
大多数子目录是自带 `main` 的可运行 demo,少量是带 `_test.go` 的练习。

- 共 ~198 个 `.go` 文件,~70 个 `package main` 可运行目录,7 个测试文件。
- 各目录之间基本**没有依赖关系**,可单独运行/阅读。
- 第三方依赖见 `go.mod`:gin、gorm、fyne、olivere/elastic、validator.v9 等。

## 目录地图(最重要)

| 目录 | 主题 | 说明 |
|------|------|------|
| `basic/` | 语言基础 | 分支/循环/函数/正则/atomic 等小例子,每个子目录一个 `main` |
| `chapter1/` | 《Go 程序设计语言》第 1 章 | echo/fetch/lissajous/dup 等经典示例 |
| `consts/` | 常量与时间格式化 | |
| `channel/` | 并发:channel | `done`(关闭通知)、`select`(多路复用) |
| `container/` | 容器类型 | arrays/slices/maps/strings,含「最长不重复子串」练习 |
| `functional/` | 函数式 | 闭包、adder、斐波那契生成器 |
| `errhandling/` | 错误处理 | defer/recover、文件服务器的 errWrapper 模式(含测试) |
| `crawler/` | **并发爬虫**(较完整) | engine/fetcher/scheduler/persist 分层,珍爱网爬虫示例 |
| `retriever/` | 接口与 mock | real vs mock 两套实现演示面向接口编程 |
| `queue/` | 队列 + 泛型 | 含测试 |
| `tree/` | 二叉树 | 节点定义与遍历 |
| `maze/` | BFS 走迷宫 | |
| `stdlib/` | 标准库示例 | time/file/strconv 等 |
| `gin/` | **Gin Web 框架**(15 个子目录) | 路由/参数绑定/中间件/校验/优雅关闭,每个子目录独立 `main` |
| `gorm/` | **GORM ORM** | 连接数据库 + CRUD + quickstart(需数据库,见下) |
| `http/` | net/http 基础 | |
| `algorithm/lession/` | **算法课 class01~class24** | 左程云风格算法练习,每个 class 一个独立 `main` |
| `test/`, `trae/` | 杂项/草稿 | 类型实验、排序等临时练习,**非正式模块** |

> 根目录的 `ip.yaml` / `ip2.yaml` / `ip3.yaml` 是临时数据文件,非源码。

## 怎么运行 / 测试

每个子目录是独立单元,**进到对应目录运行**:

```bash
# 运行某个 demo(以 gin 中间件为例)
go run ./gin/middleware

# 运行算法某一课
go run ./algorithm/lession/class24

# 跑某个目录的测试
go test ./queue/...
go test ./errhandling/...

# 全仓库构建 / 测试(注意:gin/gorm 等需要联网或数据库,可能失败)
go build ./...
go test ./...
```

**注意事项**
- `gorm/` 下的示例需要可用的数据库连接,直接运行可能因连不上而报错——属正常,改连接串后再跑。
- `gin/other_autotis`(autotls)涉及 HTTPS/证书,本地运行需相应配置。
- 同一目录下若有多个 `package main` 文件,`go run` 整个目录即可;单文件运行用 `go run path/to/file.go`。

## 给 AI 的工作约定

- 这是学习/练习代码,**风格偏教学、注释多、不追求生产级健壮性**——改动时保持示例的清晰直观,不要过度工程化。
- 新增练习时,**沿用现有约定**:按主题放进对应目录,可运行 demo 用 `package main`,带练习答案的用 `_test.go`。
- 修改某个 demo 前,确认它是否独立——通常不会影响其他目录。
- 提交信息沿用现有风格(如 `feat: class 24`)。
