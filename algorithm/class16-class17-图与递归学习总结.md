# class16~class17 图与递归学习总结

## 1. 这两章在讲什么

class16 和 class17 放在一起看，主题非常清楚：

- **class16：图论基础算法**
- **class17：递归、暴力递归、回溯、递归改迭代**

如果说前面的二叉树章节是在练“树形递归怎么向上返回信息”，那么这两章更像是在补两块核心能力：

1. **会表示图、会遍历图、会做图上的经典算法**
2. **会把一个问题拆成递归过程，并理解每一步的决策含义**

所以这两章不是孤立的。

它们共同训练的是：

> **如何把结构化问题拆成状态，再用合适的数据结构或递归流程去推进。**

---

## 2. 本次总结覆盖范围

### class16

- `src/class16/Node.java`
- `src/class16/Edge.java`
- `src/class16/Graph.java`
- `src/class16/GraphGenerator.java`
- `src/class16/Code01_BFS.java`
- `src/class16/Code02_DFS.java`
- `src/class16/Code03_TopologicalOrderBFS.java`
- `src/class16/Code03_TopologicalOrderDFS1.java`
- `src/class16/Code03_TopologicalOrderDFS2.java`
- `src/class16/Code03_TopologySort.java`
- `src/class16/Code04_Kruskal.java`
- `src/class16/Code05_Prim.java`
- `src/class16/Code06_Dijkstra.java`

### class17

- `src/class17/Node.java`
- `src/class17/Edge.java`
- `src/class17/Graph.java`
- `src/class17/Code01_Dijkstra.java`
- `src/class17/Code02_Hanoi.java`
- `src/class17/Code03_PrintAllSubsquences.java`
- `src/class17/Code04_PrintAllPermutations.java`
- `src/class17/Code05_ReverseStackUsingRecursive.java`

---

## 3. 先补地基：图的一些基础概念

在 class16 里，首先不是直接讲算法，而是先统一图的表示方式。这一步很重要，因为图不像数组、链表那样天然直观，必须先想清楚“图在代码里怎么存”。

### 3.1 什么是图

图由两部分组成：

- **点（Vertex / Node）**
- **边（Edge）**

例如：

```text
1 -> 2
1 -> 3
2 -> 4
```

这就是一个有向图。

### 3.2 有向图与无向图

#### 有向图
边有方向：

```text
1 -> 2
```

表示只能从 1 指向 2。

#### 无向图
边没有方向：

```text
1 -- 2
```

通常可以理解成两边都能走。

在这套代码里：

- 拓扑排序、Dijkstra 默认是 **有向图** 思路
- Kruskal、Prim 讨论的是 **无向图最小生成树**

### 3.3 带权图

如果边上还有数值，比如距离、代价、费用，这就是带权图。

例如：

```text
1 -(5)-> 2
1 -(3)-> 3
```

这里的 5、3 就是边权。

最小生成树、最短路径都是在带权图上讨论的。

### 3.4 入度与出度

对于有向图中的一个节点：

- **入度 in**：有多少条边指向它
- **出度 out**：它指向多少条边

这两个量在拓扑排序里特别重要。

### 3.5 邻接点

如果一个点能直接走到另一个点，那么后者就是它的邻接点。

代码里通常用 `nexts` 表示。

---

## 4. class16 的图结构设计

这一章先搭了一个通用图模型，这个模型是后面所有算法的基础。

### 4.1 `Node.java`

文件：`src/class16/Node.java:6`

```java
public class Node {
    public int value;
    public int in;
    public int out;
    public ArrayList<Node> nexts;
    public ArrayList<Edge> edges;
}
```

这个结构说明每个节点都维护：

- 节点值 `value`
- 入度 `in`
- 出度 `out`
- 直接相邻的下一个节点 `nexts`
- 从自己出发的边 `edges`

这是一种很常见的**邻接表表示法**。

### 4.2 `Edge.java`

文件：`src/class16/Edge.java:3`

```java
public class Edge {
    public int weight;
    public Node from;
    public Node to;
}
```

边对象维护：

- 权重 `weight`
- 起点 `from`
- 终点 `to`

### 4.3 `Graph.java`

文件：`src/class16/Graph.java:6`

```java
public class Graph {
    public HashMap<Integer, Node> nodes;
    public HashSet<Edge> edges;
}
```

整张图维护：

- 所有节点集合
- 所有边集合

### 4.4 `GraphGenerator.java`

文件：`src/class16/GraphGenerator.java:12`

它把一个 `N * 3` 的矩阵转成图：

```text
[weight, from, to]
```

例如：

```text
[5, 0, 7]
```

表示一条从 0 到 7、边权为 5 的边。

这个类的意义在于：

> **把“输入格式”和“图结构”分离开。**

做题时经常遇到各种输入格式，但算法最好统一跑在同一种图结构上。

---

## 5. Code01_BFS：图的宽度优先遍历

文件：`src/class16/Code01_BFS.java`

### 5.1 BFS 是什么

BFS 全称 Breadth-First Search，宽度优先搜索。

核心特点：

> **一层一层往外扩。**

它先访问离起点最近的节点，再访问更远一层的节点。

### 5.2 代码核心结构

位置：`src/class16/Code01_BFS.java:10`

BFS 的关键数据结构：

- `Queue<Node>`：队列，保证先到先出
- `HashSet<Node>`：防止节点重复入队

流程：

1. 起点入队
2. 起点加入 visited 集合
3. 每次弹出队头节点 `cur`
4. 遍历它的所有邻居 `next`
5. 没见过就入队

### 5.3 这题真正要学的点

图和树不同，图可能有环：

```text
1 -> 2 -> 3 -> 1
```

所以图遍历时必须加 visited 结构，不然会死循环。

这一点是图遍历和树遍历最大的区别之一。

---

## 6. Code02_DFS：图的深度优先遍历

文件：`src/class16/Code02_DFS.java`

### 6.1 DFS 是什么

DFS 全称 Depth-First Search，深度优先搜索。

核心特点：

> **沿着一条路尽量走到底，走不动再回退。**

### 6.2 代码写法的特点

位置：`src/class16/Code02_DFS.java:8`

这份代码不是递归版，而是**手动栈模拟版**。

它使用：

- `Stack<Node>`
- `HashSet<Node>`

特别值得注意的地方：

```java
stack.push(cur);
stack.push(next);
break;
```

这段逻辑的意义是：

- 当前点 `cur` 先压回去，表示“后面还要继续处理它”
- 再压入一个新的未访问邻居 `next`
- 然后立刻 `break`，优先沿这条新路径继续往深处走

这就是 DFS 的“深度优先”本质。

### 6.3 BFS 和 DFS 的区别

#### BFS
- 用队列
- 逐层扩展
- 常用于最短层数、最少步数

#### DFS
- 用栈 / 递归
- 一路走到底再回退
- 常用于枚举、回溯、连通性判断

---

## 7. Code03 系列：拓扑排序

class16 里有 4 份拓扑排序相关代码：

- `Code03_TopologicalOrderBFS`
- `Code03_TopologicalOrderDFS1`
- `Code03_TopologicalOrderDFS2`
- `Code03_TopologySort`

它们本质都在解决一个问题：

> **有向无环图（DAG）中，找到一个满足依赖顺序的线性排列。**

### 7.1 什么是拓扑排序

例如课程依赖：

```text
先学 A 才能学 B
先学 A 才能学 C
先学 B、C 才能学 D
```

那么合法顺序可以是：

```text
A, B, C, D
```

拓扑排序常用来处理：

- 任务依赖
- 编译顺序
- 课程安排
- 有依赖关系的流程调度

前提：

> **图必须是有向无环图 DAG。**

如果有环，就不存在合法拓扑序。

---

### 7.2 `Code03_TopologicalOrderBFS`：Kahn 算法

文件：`src/class16/Code03_TopologicalOrderBFS.java`

位置：`src/class16/Code03_TopologicalOrderBFS.java:23`

核心思路：

1. 先统计每个点的入度
2. 所有入度为 0 的点先入队
3. 每次弹出一个点加入答案
4. 删除它对后继节点的影响，即让后继节点入度减 1
5. 如果某个后继节点入度变成 0，就入队

这就是最经典的拓扑排序写法。

### 7.3 `Code03_TopologySort`：基于统一图结构的拓扑排序

文件：`src/class16/Code03_TopologySort.java`

和上面 BFS 版本质一样，只不过这里直接使用了 class16 自己定义的 `Graph/Node` 结构。

位置：`src/class16/Code03_TopologySort.java:12`

这版更适合和前面的图结构类一起理解。

---

### 7.4 `Code03_TopologicalOrderDFS1`：按“最大深度”排序

文件：`src/class16/Code03_TopologicalOrderDFS1.java`

核心定义：

```java
Record(node, deep)
```

这里 `deep` 表示：

> 从当前节点出发，沿着有向边往后走，最长链长度是多少。

递归函数 `f(cur)` 会返回当前节点的最大深度。

排序规则：

- 深度越大，越应该排前面

为什么对？

因为如果 `A -> B`，那么 A 的后续最长路径一定比 B 更长或至少不短，所以 A 会排在 B 前面。

### 7.5 `Code03_TopologicalOrderDFS2`：按“点次”排序

文件：`src/class16/Code03_TopologicalOrderDFS2.java`

这里不是看最长深度，而是看：

> 从当前点出发，一共能到达多少个点（含自己）

代码里叫 `nodes`。

排序规则：

- 能覆盖越多后继点的节点，越排前面

这也是一种合法拓扑排序思路。

### 7.6 这三种拓扑排序的学习重点

你要理解的不是“记三份代码”，而是：

> **拓扑序的本质是找一个满足边方向约束的线性顺序。**

Kahn 算法是通过“入度”构造；
DFS 版本是通过“递归定义一个优先级”再排序。

也就是说，拓扑排序不只有一种实现思路。

---

## 8. Code04_Kruskal：最小生成树 Kruskal

文件：`src/class16/Code04_Kruskal.java`

### 8.1 什么是最小生成树 MST

针对**连通无向带权图**：

> 选一些边，把所有点连起来，且总权值最小。

注意两个关键词：

- 连通
- 无环

因为“生成树”本身必须是一棵树。

### 8.2 Kruskal 的核心贪心思想

策略非常直接：

> **按边权从小到大尝试选边，只要这条边不会形成环，就选。**

### 8.3 为什么要并查集

Kruskal 的关键问题是：

> 选一条边 `(u, v)` 之前，怎么快速判断 u 和 v 是否已经连通？

如果已经连通，再加这条边就会成环。

所以它用了并查集 Union-Find。

### 8.4 并查集实现重点

位置：`src/class16/Code04_Kruskal.java:15`

并查集维护：

- `fatherMap`：每个节点的代表父亲
- `sizeMap`：每个集合代表节点对应的集合大小

关键操作：

#### `findFather`
找代表节点，并做路径压缩。

#### `isSameSet`
判断两个节点是不是同一个集合。

#### `union`
合并两个集合，并按大小合并。

### 8.5 Kruskal 主流程

位置：`src/class16/Code04_Kruskal.java:83`

1. 所有边按权值进入小根堆
2. 每次弹出最小边
3. 如果边两端不在同一集合，就选这条边
4. 然后把两端集合合并

### 8.6 这题最值得记的本质

Kruskal 是一个非常典型的“贪心 + 并查集”模板。

记忆核心只有一句：

> **小边优先，前提是不成环。**

---

## 9. Code05_Prim：最小生成树 Prim

文件：`src/class16/Code05_Prim.java`

### 9.1 Prim 的核心思路

Prim 和 Kruskal 都是求最小生成树，但视角不同。

#### Kruskal
从**边**的角度出发，按小边排序。

#### Prim
从**点集扩张**的角度出发。

Prim 的策略是：

> 从一个点开始，每次选择“已解锁区域”通向“未解锁区域”的最小边。

### 9.2 `primMST(Graph graph)`

位置：`src/class16/Code05_Prim.java:20`

关键结构：

- 小根堆：存当前解锁边
- `nodeSet`：哪些点已经被纳入生成树
- `result`：已选择的边

流程：

1. 随机选一个起点
2. 解锁该点所有边，放入小根堆
3. 每次弹出最小边
4. 如果它指向的新点还没被纳入，就选这条边
5. 再解锁这个新点的所有边
6. 重复直到结束

### 9.3 为什么代码里有外层 for 循环

位置：`src/class16/Code05_Prim.java:31`

这是为了兼容**森林**情况，即图不一定连通。

如果图不连通，就会得到最小生成森林。

### 9.4 `prim(int[][] graph)`：邻接矩阵版

位置：`src/class16/Code05_Prim.java:58`

这版不走通用图结构，而是直接对邻接矩阵做 Prim。

它维护：

- `visit[i]`：节点是否已纳入
- `distances[i]`：当前已纳入集合到节点 i 的最小代价

这版很适合笔试题里直接套。

### 9.5 Prim 和 Kruskal 怎么区分

#### Kruskal
- 从全局最小边出发
- 需要并查集防环
- 更像边排序

#### Prim
- 从某个起点逐步扩张
- 需要小根堆维护当前解锁边
- 更像“已选点集”不断长大

---

## 10. Code06_Dijkstra：单源最短路径

文件：`src/class16/Code06_Dijkstra.java`

### 10.1 Dijkstra 解决什么问题

前提：

> **边权不能为负。**

问题：

> 从一个源点出发，到图中其他所有可达点的最短距离是多少。

### 10.2 `dijkstra1`：朴素版

位置：`src/class16/Code06_Dijkstra.java:10`

关键结构：

- `distanceMap`：源点到每个点的当前最短距离
- `selectedNodes`：已经确定最短距离的点

核心流程：

1. 初始只有源点距离为 0
2. 从未确定的点里，选当前距离最小的点 `minNode`
3. 用 `minNode` 去更新它所有邻居的距离
4. 把 `minNode` 标记为已确定
5. 循环直到没有可扩展点

这就是 Dijkstra 的标准思想。

### 10.3 Dijkstra 为什么成立

因为边权非负。

所以一旦当前未确定点中最小的那个点被选出来，它的最短路就已经不可能再被后续改小。

这就是“贪心正确性”的来源。

### 10.4 `dijkstra2`：加强堆优化版

位置：`src/class16/Code06_Dijkstra.java:144`

朴素版瓶颈在于：

> 每次都要线性找“未确定且距离最小的点”。

所以这里手写了一个加强堆 `NodeHeap`。

#### 堆要解决的三个问题

1. 新节点第一次出现，要能加进去
2. 已经在堆里的节点，如果发现更短距离，要能更新
3. 已经弹出过、最短路已确定的节点，要能忽略

所以提供了这个非常经典的方法：

```java
addOrUpdateOrIgnore(node, distance)
```

这句几乎就是这份代码的灵魂。

### 10.5 `NodeHeap` 的设计要点

位置：`src/class16/Code06_Dijkstra.java:57`

它同时维护：

- `nodes[]`：堆数组
- `heapIndexMap`：某个节点在堆上的位置
- `distanceMap`：某个节点当前最短距离

这样就能支持：

- O(logN) 插入
- O(logN) 更新
- O(logN) 弹出最小值

### 10.6 class17 的 `Code01_Dijkstra`

文件：`src/class17/Code01_Dijkstra.java`

这一份本质上是 class16 Dijkstra 的再写一遍，思路完全一致。

可以把它理解成：

- 对 Dijkstra 的复习
- 对加强堆写法的巩固

---

## 11. class16 一句话总览

如果把 class16 所有内容压成一张图，可以这么记：

### 图的表示
- `Node`
- `Edge`
- `Graph`
- `GraphGenerator`

### 图的遍历
- BFS
- DFS

### DAG 处理
- 拓扑排序（BFS / DFS）

### 最小生成树
- Kruskal
- Prim

### 最短路径
- Dijkstra

也就是说，class16 是一整套非常标准的图论入门框架。

---

## 12. class17 的主题：暴力递归与回溯

如果说 class16 是“结构型算法”，那 class17 就是“过程型算法”。

这一章重点不是某个容器，而是：

> **如何把一个问题拆成递归决策树。**

这章代码最值得学的是三件事：

1. 明确递归函数的含义
2. 明确 base case
3. 明确每一步的选择分支

---

## 13. Code02_Hanoi：汉诺塔

文件：`src/class17/Code02_Hanoi.java`

### 13.1 汉诺塔问题是什么

有三根柱子：

- left
- mid
- right

有 N 个盘子，大盘不能压小盘。

目标：

> 把所有盘子从左边挪到右边。

### 13.2 最核心的递归思想

如果要把 `1~N` 从左移到右：

1. 先把 `1~N-1` 从左移到中
2. 再把第 `N` 个从左移到右
3. 再把 `1~N-1` 从中移到右

这三步就是整个问题的本质。

### 13.3 `hanoi1`：把动作拆成六个方向函数

位置：`src/class17/Code02_Hanoi.java:7`

这版写了：

- `leftToRight`
- `leftToMid`
- `rightToMid`
- `midToRight`
- `midToLeft`
- `rightToLeft`

优点：

- 非常直观
- 容易感受递归展开

缺点：

- 代码重复

### 13.4 `hanoi2`：抽象成统一递归函数

位置：`src/class17/Code02_Hanoi.java:73`

统一成：

```java
func(N, from, to, other)
```

含义是：

> 把 N 层盘子从 from 挪到 to，借助 other。

这是更抽象、更通用的递归写法。

### 13.5 `hanoi3`：用栈模拟递归

位置：`src/class17/Code02_Hanoi.java:105`

这版本质是：

> 把系统递归调用栈，手动改写成用户自己维护的栈。

这对理解“递归到底在干嘛”非常有帮助。

### 13.6 这题真正的价值

汉诺塔不只是练题，它是在训练：

- 怎么定义递归函数含义
- 怎么找 base case
- 怎么把大问题拆成同类子问题
- 递归和栈之间的等价关系

---

## 14. Code03_PrintAllSubsquences：打印全部子序列

文件：`src/class17/Code03_PrintAllSubsquences.java`

### 14.1 什么是子序列

子序列不是子串。

子序列允许跳着选，但必须保持相对顺序。

例如字符串 `abc` 的子序列包括：

- `""`
- `a`
- `b`
- `c`
- `ab`
- `ac`
- `bc`
- `abc`

### 14.2 递归本质

对于每个位置字符，都有两种选择：

- 要
- 不要

所以形成一棵二叉决策树。

### 14.3 `process1` 的含义

位置：`src/class17/Code03_PrintAllSubsquences.java:24`

```java
process1(str, index, ans, path)
```

含义：

- `str` 固定不变
- 当前来到 `index`
- `path` 表示前面已经做出的选择
- 把从 `index` 往后所有可能形成的子序列加入 `ans`

核心分支：

1. 不要当前字符
2. 要当前字符

### 14.4 去重版本 `subsNoRepeat`

位置：`src/class17/Code03_PrintAllSubsquences.java:35`

如果字符串中有重复字符，普通枚举会产生重复子序列。

这里用 `HashSet<String>` 去重。

### 14.5 这题要学什么

这道题是暴力递归最标准的模板之一。

记忆重点：

> **每个位置做选择，路径记录之前决定，index 决定当前来到哪。**

---

## 15. Code04_PrintAllPermutations：打印全部排列

文件：`src/class17/Code04_PrintAllPermutations.java`

这题比子序列更进一步。

### 15.1 排列和子序列的区别

#### 子序列
- 顺序不打乱
- 只决定“要不要”

#### 排列
- 要重排顺序
- 决定“当前位置放谁”

### 15.2 `permutation1`：rest + path 写法

位置：`src/class17/Code04_PrintAllPermutations.java:8`

思路：

- `rest`：还没使用的字符
- `path`：已经拼出来的前缀

每次从 `rest` 中拿一个字符加入 `path`，继续递归。

这是最容易理解的排列递归写法。

### 15.3 `permutation2`：原地交换写法

位置：`src/class17/Code04_PrintAllPermutations.java:37`

定义：

```java
g1(str, index, ans)
```

含义：

> 当前要决定 `index` 位置放谁。

所以枚举 `i = index..end`：

- 把 `i` 位置字符交换到 `index`
- 递归决定下一个位置
- 再交换回来恢复现场

这就是典型回溯。

### 15.4 `permutation3`：去重排列

位置：`src/class17/Code04_PrintAllPermutations.java:59`

当字符串有重复字符时，比如 `acc`，普通交换会得到重复答案。

所以在每一层递归，使用：

```java
boolean[] visited = new boolean[256];
```

表示：

> 这一层里，同一种字符只尝试一次。

这样就能完成剪枝去重。

### 15.5 这题的价值

这题是回溯法的入门核心题。

它训练你理解：

- 决策位置是什么
- 可选集合是什么
- 如何恢复现场
- 如何做同层去重

---

## 16. Code05_ReverseStackUsingRecursive：递归逆序栈

文件：`src/class17/Code05_ReverseStackUsingRecursive.java`

### 16.1 题目特点

要求：

> 只能使用递归，不能申请额外数据结构，完成栈逆序。

这题很经典，因为它训练的是：

> **如何利用递归调用栈本身来完成额外存储。**

### 16.2 核心函数 `f(stack)`

位置：`src/class17/Code05_ReverseStackUsingRecursive.java:19`

它的含义不是“逆序”，而是：

> **移除并返回栈底元素，同时保持上面元素原顺序不变。**

例如栈顶到栈底是：

```text
5 4 3 2 1
```

调用 `f(stack)` 会拿到 `1`，并把剩余栈恢复成：

```text
5 4 3 2
```

### 16.3 `reverse(stack)` 的过程

位置：`src/class17/Code05_ReverseStackUsingRecursive.java:7`

步骤：

1. 先用 `f` 取出栈底元素
2. 递归逆序剩下的栈
3. 再把原栈底元素压回栈顶

这样层层展开后，就完成了整个逆序。

### 16.4 这题的真正难点

不是代码长短，而是要先想到一个“子过程”：

> 如何在不借助额外栈的情况下拿到底部元素？

这道题非常适合训练“递归函数职责拆分”的能力。

---

## 17. class17 一句话总览

如果把 class17 内容压缩成几类：

### 递归模型理解
- 汉诺塔

### 二叉决策树枚举
- 全部子序列

### 回溯与去重
- 全部排列

### 递归替代额外数据结构
- 递归逆序栈

再加上 `Code01_Dijkstra` 是对上一章图算法的巩固。

---

## 18. 这两章最重要的通用方法

### 18.1 图算法通用方法

遇到图题先问自己：

1. 图怎么表示？邻接表还是邻接矩阵？
2. 图有没有环？
3. 是有向图还是无向图？
4. 有没有权重？
5. 是遍历问题、依赖排序问题、最小生成树问题，还是最短路问题？

### 18.2 递归题通用方法

遇到递归题先问自己：

1. 递归函数的含义是什么？
2. 参数里哪些是“状态”，哪些是“固定输入”？
3. base case 是什么？
4. 当前层有哪些选择分支？
5. 是否需要恢复现场？
6. 是否需要剪枝去重？

这两套思维方式，是 class16~17 最核心的收获。

---

## 19. 容易混淆的点

### 19.1 BFS vs DFS

- BFS：一层层扩，适合最少步数
- DFS：一条路走到底，适合枚举和回溯

### 19.2 Kruskal vs Prim

- Kruskal：按边排序，小边优先，不成环就选
- Prim：从点集逐步扩张，每次选跨边界最小边

### 19.3 子序列 vs 排列

- 子序列：要不要当前字符
- 排列：当前位置放哪个字符

### 19.4 递归 vs 回溯

- 递归：一种函数调用方式
- 回溯：递归中的一种搜索策略，强调“试、撤销、换分支再试”

---

## 20. 建议的学习顺序

建议按下面顺序复习：

1. 图结构 `Node / Edge / Graph / GraphGenerator`
2. BFS
3. DFS
4. 拓扑排序
5. Kruskal
6. Prim
7. Dijkstra
8. 汉诺塔
9. 子序列
10. 排列
11. 递归逆序栈

原因是：

- 先把图结构和基础遍历打牢
- 再学图上的高级算法
- 再切到递归与回溯
- 最后做递归技巧题巩固

---

## 21. 一句话总结每份代码

### `Code01_BFS`

> 队列 + visited，实现图的按层遍历。

### `Code02_DFS`

> 栈 + visited，模拟深度优先一路走到底。

### `Code03_TopologicalOrderBFS`

> 基于入度和队列的经典拓扑排序。

### `Code03_TopologicalOrderDFS1`

> 按“最长后继深度”定义优先级来做拓扑排序。

### `Code03_TopologicalOrderDFS2`

> 按“后继可达点总数”定义优先级来做拓扑排序。

### `Code03_TopologySort`

> 基于统一图结构实现 Kahn 拓扑排序。

### `Code04_Kruskal`

> 小边优先，并查集防环，求最小生成树。

### `Code05_Prim`

> 从点集逐步扩张，每次选跨边界最小边，求最小生成树。

### `Code06_Dijkstra`

> 单源最短路，朴素版 + 加强堆优化版。

### `Code02_Hanoi`

> 经典递归拆解问题，也展示了递归与手动栈模拟的关系。

### `Code03_PrintAllSubsquences`

> 每个字符“要 / 不要”形成二叉决策树，枚举所有子序列。

### `Code04_PrintAllPermutations`

> 通过“当前位置放谁”进行回溯，枚举全部排列并处理去重。

### `Code05_ReverseStackUsingRecursive`

> 利用递归调用栈拿到底元素，再完成栈整体逆序。

---

## 22. 最终总结

class16~class17 的核心，不只是会写几道题，而是掌握两种大能力：

### 第一种：图算法能力

- 会建图
- 会遍历图
- 会处理依赖关系
- 会求最小生成树
- 会求最短路径

### 第二种：递归建模能力

- 会定义递归函数含义
- 会写 base case
- 会列出决策分支
- 会做回溯恢复现场
- 会把递归改成显式栈

如果要把这两章压成一句话，就是：

> **class16 训练你在“图结构”上做算法，class17 训练你在“决策过程”上做递归。**

如果你把这两章真正吃透，后面很多算法题都会变得顺手，因为你已经掌握了两个特别通用的工具箱：

- 一个叫“图论建模”
- 一个叫“递归回溯建模”
