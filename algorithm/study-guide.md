# 算法基础代码库学习文档

> 仓库：`algorithmbasic2020-master`
> 
> 使用方式：按章节学习 + 按专题复盘 + 按重点题二刷。本文档适合直接粘贴到飞书文档中继续整理。

---

# 一、这套代码库适合怎么学

这不是一个零散刷题仓库，而是一条比较完整的算法学习路径：

1. **基础结构与基本功**：排序、二分、位运算、链表、栈、队列、堆
2. **树与图**：二叉树遍历、二叉树递归套路、图的遍历与最短路/最小生成树
3. **暴力递归到动态规划**：从“会写递归”到“能抽状态、列转移、做优化”
4. **高频优化结构**：滑动窗口、单调栈、KMP、Manacher
5. **高级数据结构**：线段树、树状数组、AC 自动机、AVL、SBT、SkipList

建议每一章都按下面 5 个问题学习：

- 这章解决什么类型的问题？
- 暴力解法是什么？
- 为什么暴力会慢？
- 优化点是什么？
- 这个优化以后遇到什么题型还能复用？

---

# 二、总学习路线

## 阶段 1：基础能力
- class01：排序、二分
- class02：位运算
- class03：链表、栈、队列
- class04~06：归并、快排、堆
- class07~08：堆应用、Trie、计数/基数排序

## 阶段 2：树与图
- class09~13：链表技巧、二叉树遍历、二叉树递归套路
- class14~17：贪心、并查集、图、递归展开

## 阶段 3：动态规划主线
- class18~23：从暴力递归、记忆化搜索，到严格位置依赖 DP

## 阶段 4：高频专题
- class24~30：滑动窗口、单调栈、字符串算法、Morris

## 阶段 5：高级数据结构
- class31~37：线段树、树状数组、AC 自动机、平衡树与有序表应用

---

# 三、37 章完整学习地图

## class01｜基础排序与二分
**核心文件**
- `src/class01/Code01_SelectionSort.java`
- `src/class01/Code02_BubbleSort.java`
- `src/class01/Code03_InsertionSort.java`
- `src/class01/Code04_BSExist.java`
- `src/class01/Code05_BSNearLeft.java`
- `src/class01/Code05_BSNearRight.java`
- `src/class01/Code06_BSAwesome.java`

**本章要点**
- 熟悉选择、冒泡、插入排序的思想与边界
- 掌握二分不仅能查“是否存在”，还能找最左/最右满足条件的位置
- 理解“局部有序信息”如何支持高效搜索

**学习目标**
- 能独立手写基础排序
- 能把二分从模板题扩展到边界题

---

## class02｜位运算专题
**核心文件**
- `src/class02/Code01_Swap.java`
- `src/class02/Code02_EvenTimesOddTimes.java`
- `src/class02/Code03_KM.java`

**本章要点**
- 异或的交换律、自反性
- 找出现奇数次的数
- K 次 / M 次问题的位统计做法

**学习目标**
- 能用位运算代替额外哈希结构解决频次统计问题
- 理解按位计数的本质是“把频率问题拆成 32 个独立维度”

---

## class03｜链表、栈、队列与递归入门
**核心文件**
- `src/class03/Code01_ReverseList.java`
- `src/class03/Code02_DeleteGivenValue.java`
- `src/class03/Code03_DoubleEndsQueueToStackAndQueue.java`
- `src/class03/Code04_RingArray.java`
- `src/class03/Code05_GetMinStack.java`
- `src/class03/Code06_TwoStacksImplementQueue.java`
- `src/class03/Code07_TwoQueueImplementStack.java`
- `src/class03/Code08_GetMax.java`

**本章要点**
- 链表反转、指定值删除
- 栈和队列的实现与相互模拟
- 最小栈
- 递归求数组最大值

**学习目标**
- 理解“用一个结构模拟另一个结构”的设计思路
- 能区分数据结构接口与底层实现方式

---

## class04｜归并排序与归并扩展统计
**核心文件**
- `src/class04/Code01_MergeSort.java`
- `src/class04/Code02_SmallSum.java`
- `src/class04/Code03_ReversePair.java`
- `src/class04/Code04_BiggerThanRightTwice.java`

**本章要点**
- 归并排序模板
- 小和问题
- 逆序对问题
- “大于右侧两倍”统计

**学习目标**
- 理解归并排序不仅能排序，还能在 merge 过程中完成跨区间统计
- 形成“先分治排序，再利用有序性统计”的题感

---

## class05｜快排与区间和统计
**核心文件**
- `src/class05/Code01_CountOfRangeSum.java`
- `src/class05/Code02_PartitionAndQuickSort.java`
- `src/class05/Code03_QuickSortRecursiveAndUnrecursive.java`

**本章要点**
- 荷兰国旗问题
- 快速排序递归与非递归写法
- 区间和个数问题（前缀和 + 分治统计）

**学习目标**
- 掌握 partition 的三路划分思想
- 理解“子数组问题 -> 前缀和问题 -> 排序统计问题”的变形路线

---

## class06｜堆与堆排序
**核心文件**
- `src/class06/Code01_Comparator.java`
- `src/class06/Code02_Heap.java`
- `src/class06/Code03_HeapSort.java`
- `src/class06/Code04_SortArrayDistanceLessK.java`

**本章要点**
- 大根堆/小根堆
- 堆排序
- 几乎有序数组排序

**学习目标**
- 会手写 heapInsert、heapify
- 理解“如果元素离目标位置不远，就不需要全局排序”

---

## class07｜堆的应用
**核心文件**
- `src/class07/Code01_CoverMax.java`
- `src/class07/Code02_EveryStepShowBoss.java`
- `src/class07/HeapGreater.java`

**本章要点**
- 线段最大重合数
- 强化堆
- TopK/动态获奖系统类问题

**学习目标**
- 理解堆在“动态候选集维护”中的作用
- 能把几何/区间问题转成“扫描线 + 小根堆”

---

## class08｜Trie、计数排序、基数排序
**核心文件**
- `src/class08/Code01_TrieTree.java`
- `src/class08/Code02_TrieTree.java`
- `src/class08/Code03_CountSort.java`
- `src/class08/Code04_RadixSort.java`

**本章要点**
- 前缀树的插入、删除、查找、前缀统计
- 计数排序与基数排序的适用场景

**学习目标**
- 明确“桶排序类算法依赖数据范围/位数”
- 理解 Trie 的本质是共享前缀路径

---

## class09｜链表高频技巧
**核心文件**
- `src/class09/Code01_LinkedListMid.java`
- `src/class09/Code02_IsPalindromeList.java`
- `src/class09/Code03_SmallerEqualBigger.java`
- `src/class09/Code04_CopyListWithRandom.java`

**本章要点**
- 快慢指针找中点
- 判断回文链表
- 按某个值做链表分区
- 复制带随机指针的链表

**学习目标**
- 熟悉链表中的原地操作、断链、重连
- 理解空间优化通常来自“原链表结构复用”

---

## class10｜链表相交与二叉树遍历
**核心文件**
- `src/class10/Code01_FindFirstIntersectNode.java`
- `src/class10/Code02_RecursiveTraversalBT.java`
- `src/class10/Code03_UnRecursiveTraversalBT.java`

**本章要点**
- 两个链表第一个相交节点（含有环/无环讨论）
- 二叉树递归与非递归遍历

**学习目标**
- 学会先分类讨论再写逻辑
- 能把递归遍历改写成栈模拟版本

---

## class11｜二叉树按层遍历与序列化
**核心文件**
- `src/class11/Code01_LevelTraversalBT.java`
- `src/class11/Code02_SerializeAndReconstructTree.java`
- `src/class11/Code03_EncodeNaryTreeToBinaryTree.java`
- `src/class11/Code04_PrintBinaryTree.java`
- `src/class11/Code05_TreeMaxWidth.java`
- `src/class11/Code06_SuccessorNode.java`
- `src/class11/Code07_PaperFolding.java`

**本章要点**
- 层序遍历
- 先序/后序/层序序列化与反序列化
- N 叉树转二叉树
- 最大宽度、后继节点

**学习目标**
- 理解 null 标记在树序列化中的作用
- 建立 BFS/队列在树题中的通用用法

---

## class12｜二叉树递归套路（性质判断）
**核心文件**
- `src/class12/Code01_IsCBT.java`
- `src/class12/Code02_IsBST.java`
- `src/class12/Code03_IsBalanced.java`
- `src/class12/Code04_IsFull.java`
- `src/class12/Code05_MaxSubBSTSize.java`
- `src/class12/Code06_MaxDistance.java`

**本章要点**
- 完全二叉树、搜索二叉树、平衡二叉树、满二叉树判断
- 最大 BST 子树大小
- 二叉树最大距离

**学习目标**
- 学会定义“每棵子树向上返回什么信息”
- 建立树形 DP 基础

---

## class13｜二叉树递归套路（进阶）
**核心文件**
- `src/class13/Code01_IsCBT.java`
- `src/class13/Code02_MaxSubBSTHead.java`
- `src/class13/Code03_lowestAncestor.java`
- `src/class13/Code04_MaxHappy.java`
- `src/class13/Code05_LowestLexicography.java`

**本章要点**
- 最大 BST 子树头节点
- 最近公共祖先
- 派对最大快乐值
- 字典序最小拼接

**学习目标**
- 理解“树上做选择”的本质就是设计正确的信息结构
- 学会把组织树、员工树等抽象成树形 DP

---

## class14｜贪心算法与并查集引入
**核心文件**
- `src/class14/Code01_Light.java`
- `src/class14/Code02_LessMoneySplitGold.java`
- `src/class14/Code03_BestArrange.java`
- `src/class14/Code04_IPO.java`
- `src/class14/Code05_UnionFind.java`

**本章要点**
- 路灯问题
- 切金条（哈夫曼思想）
- 会议安排
- IPO 项目选择
- 并查集基础结构

**学习目标**
- 学会识别“当前看起来最优，且不会破坏整体最优”的贪心条件
- 理解并查集是高效动态连通性结构

---

## class15｜并查集应用
**核心文件**
- `src/class15/Code01_FriendCircles.java`
- `src/class15/Code02_NumberOfIslands.java`
- `src/class15/Code03_NumberOfIslandsII.java`

**本章要点**
- 朋友圈/省份数量
- 岛问题
- 动态加点岛问题

**学习目标**
- 理解并查集在网格问题、连通块问题中的价值
- 掌握路径压缩和按规模合并

---

## class16｜图算法基础
**核心文件**
- `src/class16/Code01_BFS.java`
- `src/class16/Code02_DFS.java`
- `src/class16/Code03_TopologySort.java`
- `src/class16/Code04_Kruskal.java`
- `src/class16/Code05_Prim.java`
- `src/class16/Code06_Dijkstra.java`

**本章要点**
- 图的宽搜、深搜
- 拓扑排序
- Kruskal / Prim 最小生成树
- Dijkstra 最短路

**学习目标**
- 掌握图结构抽象：点、边、邻接关系
- 明确 BFS/DFS/拓扑/MST/最短路分别解决什么问题

---

## class17｜递归展开与暴力尝试
**核心文件**
- `src/class17/Code01_Dijkstra.java`
- `src/class17/Code02_Hanoi.java`
- `src/class17/Code03_PrintAllSubsquences.java`
- `src/class17/Code04_PrintAllPermutations.java`
- `src/class17/Code05_ReverseStackUsingRecursive.java`

**本章要点**
- 汉诺塔
- 全部子序列
- 全排列
- 递归逆序栈

**学习目标**
- 建立“每一层递归做一件事”的拆解能力
- 理解子序列、排列、状态树的区别

---

## class18｜从暴力递归到记忆化
**核心文件**
- `src/class18/Code01_RobotWalk.java`
- `src/class18/Code02_CardsInLine.java`

**本章要点**
- 机器人走路
- 纸牌博弈问题

**学习目标**
- 学会从“尝试”定义递归
- 学会识别重复子问题并加缓存

---

## class19｜经典 DP：背包、转化、LCS
**核心文件**
- `src/class19/Code01_Knapsack.java`
- `src/class19/Code02_ConvertToLetterString.java`
- `src/class19/Code03_StickersToSpellWord.java`
- `src/class19/Code04_LongestCommonSubsequence.java`

**本章要点**
- 0/1 背包
- 数字字符串转字母
- 贴纸拼词
- 最长公共子序列

**学习目标**
- 理解“当前位置 + 剩余资源”是最常见 DP 状态设计方式
- 掌握从递归到二维 DP 表的转化方式

---

## class20｜动态规划进阶
**核心文件**
- `src/class20/Code01_PalindromeSubsequence.java`
- `src/class20/Code02_HorseJump.java`
- `src/class20/Code03_Coffee.java`

**本章要点**
- 最长回文子序列
- 马走日
- 咖啡机与洗杯子问题

**学习目标**
- 学会处理多个阶段串联的问题
- 理解某些题目需要先贪心定前半段，再 DP 解后半段

---

## class21｜路径与硬币模型
**核心文件**
- `src/class21/Code01_MinPathSum.java`
- `src/class21/Code02_CoinsWayEveryPaperDifferent.java`
- `src/class21/Code03_CoinsWayNoLimit.java`
- `src/class21/Code04_CoinsWaySameValueSamePapper.java`
- `src/class21/Code05_BobDie.java`

**本章要点**
- 最小路径和
- 不同硬币模型
- 生存概率问题

**学习目标**
- 能区分“每张纸币不同 / 无限张 / 同值有限张”三种模型差异
- 理解概率 DP 也是状态转移

---

## class22｜DP 优化与计数模型
**核心文件**
- `src/class22/Code01_KillMonster.java`
- `src/class22/Code02_MinCoinsNoLimit.java`
- `src/class22/Code03_SplitNumber.java`

**本章要点**
- 杀怪兽概率
- 最少硬币数（无限张）
- 整数拆分

**学习目标**
- 理解枚举行为如何通过递推式优化掉内层循环
- 学会从“方案枚举”推成“状态转移”

---

## class23｜分组 DP 与 N 皇后
**核心文件**
- `src/class23/Code01_SplitSumClosed.java`
- `src/class23/Code02_SplitSumClosedSizeHalf.java`
- `src/class23/Code03_NQueens.java`

**本章要点**
- 集合划分接近问题
- 固定个数的集合划分
- N 皇后（普通版 + 位运算优化版）

**学习目标**
- 学会把“选择若干数让结果最接近某目标”转成 DP
- 理解位运算优化搜索的思想

---

## class24｜滑动窗口专题
**核心文件**
- `src/class24/Code01_SlidingWindowMaxArray.java`
- `src/class24/Code02_AllLessNumSubArray.java`
- `src/class24/Code03_GasStation.java`
- `src/class24/Code04_MinCoinsOnePaper.java`

**本章要点**
- 滑动窗口最大值
- 子数组 max-min <= num 个数统计
- 加油站问题
- 单纸币最少张数问题的窗口优化

**学习目标**
- 理解窗口何时右扩、何时左缩
- 学会用双端队列维护窗口内最值

---

## class25｜单调栈专题
**核心文件**
- `src/class25/Code01_MonotonousStack.java`
- `src/class25/Code02_AllTimesMinToMax.java`
- `src/class25/Code03_LargestRectangleInHistogram.java`
- `src/class25/Code04_MaximalRectangle.java`
- `src/class25/Code05_CountSubmatricesWithAllOnes.java`

**本章要点**
- 最近较小元素
- 子数组最小值乘和最大
- 柱状图最大矩形
- 二维矩阵最大全 1 矩形
- 全 1 子矩阵数量

**学习目标**
- 一眼识别“边界由最近更小/更大元素决定”的题型
- 理解单调栈的本质是延迟结算贡献

---

## class26｜单调栈延伸与矩阵快速幂类问题
**核心文件**
- `src/class26/Code01_SumOfSubarrayMinimums.java`
- `src/class26/Code02_FibonacciProblem.java`
- `src/class26/Code03_ZeroLeftOneStringNumber.java`

**本章要点**
- 子数组最小值之和
- 斐波那契问题
- 0 左边必须有 1 的字符串计数

**学习目标**
- 学会元素贡献法
- 理解线性递推可由矩阵快速幂优化

---

## class27｜KMP 与树结构匹配
**核心文件**
- `src/class27/Code01_KMP.java`
- `src/class27/Code02_TreeEqual.java`
- `src/class27/Code03_IsRotation.java`

**本章要点**
- KMP 模式匹配
- 树是否包含某结构
- 字符串旋转判断

**学习目标**
- 理解 next 数组的含义
- 明确 KMP 的核心收益是文本串指针不回退

---

## class28｜Manacher 回文串专题
**核心文件**
- `src/class28/Code01_Manacher.java`
- `src/class28/Code02_AddShortestEnd.java`

**本章要点**
- 最长回文子串
- 在字符串尾部补最短字符使其整体回文

**学习目标**
- 理解回文半径数组与镜像位置
- 掌握线性时间求回文信息的方法

---

## class29｜TopK、选择算法与蓄水池采样
**核心文件**
- `src/class29/Code01_FindMinKth.java`
- `src/class29/Code02_MaxTopK.java`
- `src/class29/Code03_ReservoirSampling.java`

**本章要点**
- 第 K 小
- TopK
- 蓄水池抽样

**学习目标**
- 理解 BFPRT/快速选择思想
- 掌握“数据流条件下的等概率采样”

---

## class30｜Morris 遍历与树高度问题
**核心文件**
- `src/class30/Code01_MorrisTraversal.java`
- `src/class30/Code05_MinHeight.java`

**本章要点**
- Morris 先序/中序/后序遍历
- 二叉树最小高度

**学习目标**
- 理解如何利用线索化思想把空间复杂度降为 O(1)
- 学会在不额外开栈的情况下遍历树

---

## class31｜线段树
**核心文件**
- `src/class31/Code01_SegmentTree.java`
- `src/class31/Code02_FallingSquares.java`

**本章要点**
- 区间加、区间更新、区间求和
- 落方块问题

**学习目标**
- 掌握懒更新（lazy propagation）
- 理解“区间问题为什么不能总是暴力下推到叶子”

---

## class32｜树状数组与 AC 自动机
**核心文件**
- `src/class32/Code01_IndexTree.java`
- `src/class32/Code02_IndexTree2D.java`
- `src/class32/Code03_AC1.java`
- `src/class32/Code04_AC2.java`

**本章要点**
- 一维/二维树状数组
- Aho-Corasick 自动机

**学习目标**
- 理解树状数组适合前缀信息维护
- 理解 AC 自动机是“Trie + fail 指针”的多模式匹配结构

---

## class33｜哈希基础
**核心文件**
- `src/class33/Hash.java`

**本章要点**
- 哈希函数与基础哈希用法

**学习目标**
- 了解哈希在字符串、集合判重、快速定位中的作用

---

## class34｜资源限制类题目说明
**核心文件**
- `src/class34/ReadMe.java`

**本章要点**
- 本章没有完整代码实现
- 主要原因是资源限制类题目输入复杂、实现量大，面试中通常重点考察思路

**学习目标**
- 学会表达解法思路、复杂度和工程折中，而不只是写代码

---

## class35｜AVL 树
**核心文件**
- `src/class35/Code01_AVLTreeMap.java`

**本章要点**
- AVL 平衡二叉搜索树
- 有序表操作

**学习目标**
- 理解旋转维持平衡的原理
- 熟悉有序表的插入、删除、查询

---

## class36｜SBT 与 SkipList
**核心文件**
- `src/class36/Code01_SizeBalancedTreeMap.java`
- `src/class36/Code02_SkipListMap.java`

**本章要点**
- Size Balanced Tree
- 跳表有序表

**学习目标**
- 理解不同有序结构的平衡策略
- 学会按秩查询、前驱后继等操作

---

## class37｜有序表综合应用
**核心文件**
- `src/class37/Code01_CountofRangeSum.java`
- `src/class37/Code02_SlidingWindowMedian.java`
- `src/class37/Code03_AddRemoveGetIndexGreat.java`
- `src/class37/Compare.java`

**本章要点**
- 区间和计数的有序表解法
- 滑动窗口中位数
- 支持按下标增删查的结构

**学习目标**
- 理解“有序 + 按秩”结构如何解决窗口中位数、范围统计等问题
- 能把抽象数据结构应用到具体题目中

---

# 四、重点题精讲（高价值必刷）

## 1）小和问题
**文件**：`src/class04/Code02_SmallSum.java`

**问题**
- 对每个位置，统计左边比它小的数字之和，累加得到总小和。

**核心思路**
- 利用归并排序的 merge 过程统计跨区间贡献。
- 当左组 `arr[p1] < arr[p2]` 时，右组从 `p2` 到结尾都比 `arr[p1]` 大，因此一次性产生贡献。

**复杂度**
- 时间：`O(N log N)`
- 空间：`O(N)`

**优化点**
- 从双重循环的 `O(N^2)`，优化为“排序时顺便统计”。

**学习价值**
- 是归并扩展题的母题，逆序对、区间和统计都与它一脉相承。

---

## 2）区间和个数
**文件**：`src/class05/Code01_CountOfRangeSum.java`

**问题**
- 求有多少个子数组的累加和落在 `[lower, upper]` 范围内。

**核心思路**
- 先转前缀和。
- 原问题变成：对每个前缀和 `sum[i]`，找之前有多少个前缀和落在 `[sum[i]-upper, sum[i]-lower]`。
- 再用归并排序在有序前缀和中统计窗口数量。

**复杂度**
- 时间：`O(N log N)`
- 空间：`O(N)`

**优化点**
- 子数组枚举 -> 前缀和差值 -> 分治排序统计。

**学习价值**
- 是“前缀和 + 有序统计”的典型题。

---

## 3）几乎有序数组排序
**文件**：`src/class06/Code04_SortArrayDistanceLessK.java`

**问题**
- 每个元素距离排好序后的最终位置不超过 `K`，要求排序。

**核心思路**
- 用大小为 `K` 的小根堆维护当前位置可能出现的最小元素。

**复杂度**
- 时间：`O(N log K)`
- 空间：`O(K)`

**优化点**
- 利用输入特征，比通用排序 `O(N log N)` 更优。

**学习价值**
- 训练“利用题目额外条件降复杂度”的能力。

---

## 4）最大线段重合数
**文件**：`src/class07/Code01_CoverMax.java`

**问题**
- 给定若干线段，求同一时刻最多有多少条线段重合。

**核心思路**
- 按起点排序。
- 用小根堆维护当前仍然覆盖当前位置的线段终点。

**复杂度**
- 时间：`O(N log N)`
- 空间：`O(N)`

**优化点**
- 把几何题转成扫描线 + 活跃区间维护。

**学习价值**
- 扫描线思想的入门题。

---

## 5）复制带随机指针的链表
**文件**：`src/class09/Code04_CopyListWithRandom.java`

**问题**
- 复制一个带 `next` 和 `rand` 指针的链表。

**核心思路**
- 方法 1：哈希表建立旧节点到新节点映射。
- 方法 2：把复制节点插到原节点后面，再利用 `old.rand.next` 找到新节点的随机指针，最后拆分链表。

**复杂度**
- 时间：`O(N)`
- 空间：`O(N)` 或 `O(1)`（不算返回链表本身）

**优化点**
- 原地穿插法用结构信息替代外部映射表。

**学习价值**
- 很经典的“空间优化来自结构复用”。

---

## 6）最大 BST 子树大小
**文件**：`src/class12/Code05_MaxSubBSTSize.java`

**问题**
- 在一棵普通二叉树中，找到最大的 BST 子树大小。

**核心思路**
- 每个节点向上返回：
  - 子树最小值
  - 子树最大值
  - 子树节点数
  - 子树内最大 BST 大小
- 父节点根据左右子树信息判断：整棵树是否能成为 BST。

**复杂度**
- 时间：`O(N)`
- 空间：`O(H)`

**优化点**
- 不做重复遍历，一次后序递归收集全部关键信息。

**学习价值**
- 树形 DP 代表题。

---

## 7）IPO 项目选择
**文件**：`src/class14/Code04_IPO.java`

**问题**
- 给定启动资金、项目利润和最多可做项目数，求最终最大资金。

**核心思路**
- 小根堆按成本维护“未解锁项目”。
- 大根堆按利润维护“已解锁项目”。
- 每轮把当前资金能做的项目放入利润堆，再选利润最大者。

**复杂度**
- 时间：`O(N log N + K log N)`
- 空间：`O(N)`

**优化点**
- 用两个堆把“能不能做”和“做哪个最好”分开处理。

**学习价值**
- 双堆贪心典型题。

---

## 8）朋友圈 / 省份数量
**文件**：`src/class15/Code01_FriendCircles.java`

**问题**
- 给定相邻矩阵，求有多少个连通分量。

**核心思路**
- 并查集把相互认识的人合并。
- 最后看集合数量。

**复杂度**
- 时间：`O(N^2)`
- 空间：`O(N)`

**优化点**
- 路径压缩 + 按规模合并。

**学习价值**
- 并查集最标准应用。

---

## 9）Dijkstra 最短路
**文件**：`src/class16/Code06_Dijkstra.java`

**问题**
- 求单源最短路径。

**核心思路**
- 非负边权条件下，每轮确定当前未锁定点中距离最小者。
- 可用堆优化选点过程。

**复杂度**
- 普通版：`O(N^2 + M)`
- 堆优化版：`O((N + M) log N)`

**优化点**
- 用堆减少每轮线性寻找最小距离点的代价。

**学习价值**
- 图最短路核心算法之一。

---

## 10）机器人走路
**文件**：`src/class18/Code01_RobotWalk.java`

**问题**
- 从 `start` 出发走 `K` 步到 `aim` 的方法数。

**核心思路**
- 递归状态：当前位置 + 剩余步数。
- 再做记忆化，最后改成 DP 表。

**复杂度**
- 记忆化 / DP：`O(NK)`
- 空间：`O(NK)`

**优化点**
- 从暴力递归转为可复用状态。

**学习价值**
- 动态规划入门母题。

---

## 11）0/1 背包
**文件**：`src/class19/Code01_Knapsack.java`

**问题**
- 每个物品只能拿 0 次或 1 次，在容量限制下求最大价值。

**核心思路**
- 状态通常定义为：来到第几个物品、还剩多少容量。
- 决策是“拿 / 不拿”。

**复杂度**
- 时间：`O(N * bag)`
- 空间：`O(N * bag)`

**优化点**
- 用表结构消除重复递归。

**学习价值**
- 二维 DP 最经典模板。

---

## 12）贴纸拼词
**文件**：`src/class19/Code03_StickersToSpellWord.java`

**问题**
- 给定若干贴纸，问拼出目标字符串最少需要几张。

**核心思路**
- 先把贴纸变成字符频次数组。
- 递归处理剩余目标串，并用缓存记住中间状态。
- 通过“贴纸必须覆盖目标首字符”做剪枝。

**复杂度**
- 理论上指数级，缓存后大幅优化

**优化点**
- 字符串问题转换为词频向量问题。

**学习价值**
- 很典型的“状态压缩 + 剪枝 + 记忆化”题。

---

## 13）咖啡机与洗杯子
**文件**：`src/class20/Code03_Coffee.java`

**问题**
- 多台咖啡机做咖啡，喝完后的杯子可以洗或自然挥发，求全部杯子变干净的最短时间。

**核心思路**
- 先用堆贪心安排每个人最早喝到咖啡的时间。
- 再对每个杯子做“机洗 / 自然干”的 DP 决策。

**复杂度**
- 前半段堆：`O(N log M)`
- 后半段 DP：与洗杯机时间状态范围相关

**优化点**
- 多阶段问题拆分处理：前半贪心，后半 DP。

**学习价值**
- 训练复杂问题拆阶段的能力。

---

## 14）N 皇后
**文件**：`src/class23/Code03_NQueens.java`

**问题**
- 求 N 皇后的合法摆放数量。

**核心思路**
- 普通版：回溯 + 冲突判断。
- 优化版：用位信息维护列和对角线限制。

**复杂度**
- 指数级搜索
- 位运算版常数更优

**优化点**
- 冲突判断从扫描历史记录，优化为位操作。

**学习价值**
- 回溯优化代表题。

---

## 15）滑动窗口最大值
**文件**：`src/class24/Code01_SlidingWindowMaxArray.java`

**问题**
- 返回每个窗口大小为 `w` 的窗口最大值。

**核心思路**
- 用双端队列维护候选下标，保证队列内值单调递减。

**复杂度**
- 时间：`O(N)`
- 空间：`O(W)`

**优化点**
- 每个元素最多进出队一次。

**学习价值**
- 滑动窗口模板题。

---

## 16）柱状图最大矩形
**文件**：`src/class25/Code03_LargestRectangleInHistogram.java`

**问题**
- 给定柱状图高度，求最大矩形面积。

**核心思路**
- 对每根柱子，找左右两侧第一个比它小的位置。
- 该柱子作为高时的最大宽度就确定了。

**复杂度**
- 时间：`O(N)`
- 空间：`O(N)`

**优化点**
- 单调栈替代双向暴力查边界。

**学习价值**
- 单调栈最经典母题之一。

---

## 17）子数组最小值之和
**文件**：`src/class26/Code01_SumOfSubarrayMinimums.java`

**问题**
- 所有子数组最小值的总和。

**核心思路**
- 对每个元素计算：有多少个子数组把它当最小值。
- 贡献 = 元素值 × 左扩方案数 × 右扩方案数。

**复杂度**
- 时间：`O(N)`
- 空间：`O(N)`

**优化点**
- 从枚举所有子数组，转成枚举每个元素的贡献。

**学习价值**
- 贡献法 + 单调栈代表题。

---

## 18）KMP
**文件**：`src/class27/Code01_KMP.java`

**问题**
- 字符串匹配，找模式串第一次出现位置。

**核心思路**
- 预处理 next 数组。
- 失配时模式串跳转，文本串不回退。

**复杂度**
- 时间：`O(N + M)`
- 空间：`O(M)`

**优化点**
- 避免暴力匹配的重复比较。

**学习价值**
- 字符串匹配核心算法。

---

## 19）Manacher
**文件**：`src/class28/Code01_Manacher.java`

**问题**
- 最长回文子串长度。

**核心思路**
- 插入分隔符统一奇偶回文。
- 利用中心对称与最右回文边界快速复用信息。

**复杂度**
- 时间：`O(N)`
- 空间：`O(N)`

**优化点**
- 把大量重复扩展合并为线性处理。

**学习价值**
- 回文问题高阶算法。

---

## 20）线段树
**文件**：`src/class31/Code01_SegmentTree.java`

**问题**
- 支持区间加、区间更新、区间查询。

**核心思路**
- 线段树分治维护区间信息。
- 用 lazy 标记推迟对子节点的更新。

**复杂度**
- 建树：`O(N)`
- 单次操作：`O(log N)`

**优化点**
- 不把区间修改立即下发到底，避免重复更新。

**学习价值**
- 区间数据结构核心内容。

---

## 21）AC 自动机
**文件**：`src/class32/Code04_AC2.java`

**问题**
- 多模式串匹配，判断哪些敏感词在文本中出现。

**核心思路**
- 先构建 Trie。
- 再用 BFS 构建 fail 指针。
- 扫描文本时像 KMP 一样失配跳转。

**复杂度**
- 时间：`O(模式串总长度 + 文本长度 + 匹配结果数)`
- 空间：`O(模式串总长度)`

**优化点**
- 多模式串共享前缀，失配共享回退逻辑。

**学习价值**
- Trie 与 KMP 思想结合的经典结构。

---

## 22）滑动窗口中位数
**文件**：`src/class37/Code02_SlidingWindowMedian.java`

**问题**
- 求每个长度为 `k` 的滑动窗口中位数。

**核心思路**
- 用支持按秩查询的有序结构维护窗口元素。
- 插入一个、删除一个，再按排名取中位数。

**复杂度**
- 时间：`O(N log K)`
- 空间：`O(K)`

**优化点**
- 普通堆难以删除窗口外元素，有序表更适合。

**学习价值**
- 有序表应用题代表作。

---

# 五、专题复盘方法

## 1. 归并扩展专题
重点代码：
- `src/class04/Code02_SmallSum.java`
- `src/class04/Code03_ReversePair.java`
- `src/class05/Code01_CountOfRangeSum.java`

复盘重点：
- 为什么这些题都能在 merge 阶段统计？
- 统计的到底是“跨左右组的什么关系”？

## 2. 树形 DP 专题
重点代码：
- `src/class12/Code05_MaxSubBSTSize.java`
- `src/class12/Code06_MaxDistance.java`
- `src/class13/Code04_MaxHappy.java`

复盘重点：
- 一个节点需要向上返回哪些信息？
- 父节点如何只靠左右子树信息完成决策？

## 3. 贪心专题
重点代码：
- `src/class14/Code01_Light.java`
- `src/class14/Code02_LessMoneySplitGold.java`
- `src/class14/Code03_BestArrange.java`
- `src/class14/Code04_IPO.java`

复盘重点：
- 贪心的局部最优是什么？
- 为什么这个局部最优不会影响整体最优？

## 4. 动态规划专题
重点代码：
- `src/class18/Code01_RobotWalk.java`
- `src/class19/Code01_Knapsack.java`
- `src/class20/Code03_Coffee.java`
- `src/class22/Code02_MinCoinsNoLimit.java`

复盘重点：
- 状态是什么？
- 转移从哪来？
- 边界条件如何写？
- 能不能压缩空间？

## 5. 单调栈专题
重点代码：
- `src/class25/Code01_MonotonousStack.java`
- `src/class25/Code03_LargestRectangleInHistogram.java`
- `src/class26/Code01_SumOfSubarrayMinimums.java`

复盘重点：
- 栈里维护的单调性是什么？
- 元素在何时结算贡献？
- 左右边界分别是谁？

## 6. 字符串算法专题
重点代码：
- `src/class27/Code01_KMP.java`
- `src/class28/Code01_Manacher.java`
- `src/class32/Code04_AC2.java`

复盘重点：
- KMP：next 数组是什么
- Manacher：回文半径是什么
- AC 自动机：fail 指针是什么

---

# 六、推荐刷题优先级

## 第一优先级：必须掌握
- `src/class04/Code02_SmallSum.java`
- `src/class05/Code01_CountOfRangeSum.java`
- `src/class06/Code04_SortArrayDistanceLessK.java`
- `src/class09/Code04_CopyListWithRandom.java`
- `src/class12/Code05_MaxSubBSTSize.java`
- `src/class14/Code04_IPO.java`
- `src/class15/Code01_FriendCircles.java`
- `src/class16/Code06_Dijkstra.java`
- `src/class18/Code01_RobotWalk.java`
- `src/class19/Code01_Knapsack.java`
- `src/class23/Code03_NQueens.java`
- `src/class25/Code03_LargestRectangleInHistogram.java`
- `src/class27/Code01_KMP.java`
- `src/class31/Code01_SegmentTree.java`

## 第二优先级：进阶拔高
- `src/class19/Code03_StickersToSpellWord.java`
- `src/class20/Code03_Coffee.java`
- `src/class24/Code02_AllLessNumSubArray.java`
- `src/class26/Code01_SumOfSubarrayMinimums.java`
- `src/class28/Code01_Manacher.java`
- `src/class32/Code04_AC2.java`
- `src/class37/Code02_SlidingWindowMedian.java`

## 第三优先级：结构理解
- `src/class35/Code01_AVLTreeMap.java`
- `src/class36/Code01_SizeBalancedTreeMap.java`
- `src/class36/Code02_SkipListMap.java`

---

# 七、实际学习建议

## 第一轮：建立全局地图
- 快速看完 37 章目录
- 每章只抓“这章教什么套路”
- 不求每题一次吃透

## 第二轮：按专题打通
- 归并扩展
- 树形 DP
- 贪心
- 动态规划
- 单调栈
- 字符串算法
- 高级数据结构

## 第三轮：默写模板
建议至少能手写：
- 归并排序
- 快速排序
- 堆排序
- 并查集
- BFS / DFS / 拓扑排序
- Dijkstra
- 0/1 背包
- 滑动窗口最大值
- 单调栈模板
- KMP
- 线段树

## 第四轮：面试视角复盘
每道题复盘以下内容：
- 暴力怎么做
- 为什么慢
- 优化思路
- 边界条件
- 时间复杂度 / 空间复杂度
- 有没有同类题

---

# 八、结论

这套代码库最大的价值不是“题多”，而是**路线清晰、递进明确、模板完整**。

如果你的目标是面试算法学习，建议：
- 把前 25 章作为必学核心
- 把 27、28、31、32、35~37 作为进阶加分项
- 每学完一个专题，都要总结成自己的“识题信号 + 模板解法 + 易错点”

如果后续继续扩展这份文档，最值得补的内容是：
1. 每章加入 1~2 道“面试问法”
2. 每个专题加入“模板伪代码”
3. 给重点题补充“常见错误点”
