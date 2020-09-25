# 随机Prim算法

**Randomized Prim's algorithm（随机Prim算法)**

随机Prim算法属于打通墙壁生成迷宫的算法
RandomizedPrim

**伪代码逻辑**
1.让迷宫全是墙.
2.选一个单元格作为迷宫的通路，然后把它的邻墙放入列表
3.当列表里还有墙时
	1.从列表里随机选一个墙，如果这面墙分隔的两个单元格只有一个单元格被访问过
		1.那就从列表里移除这面墙，即把墙打通，让未访问的单元格成为迷宫的通路
		2.把这个格子的墙加入列表
	2.如果墙两面的单元格都已经被访问过，那就从列表里移除这面墙

在操作过程中，如果把墙放到列表中，比较复杂，维基里面提到了改进策略：
<https://en.wikipedia.org/wiki/Maze_generation_algorithm#Modified_version>

我们可以维护一个迷宫单元格的列表，而不是边的列表。
在这个迷宫单元格列表里面存放了未访问的单元格，（即随机可选的墙）
我们在单元格列表中随机挑选一个单元格，如果这个单元格有多面墙联系着已存在的迷宫通路，我们就随机选择一面墙打通。
这会比基于边的版本分支稍微多一点。


<https://www.cnblogs.com/WayneShao/p/5890379.html>



**链表中随机取一个**


[Linked List Random Node(返回链表中随机一个节点的值)](https://blog.csdn.net/xiangwanpeng/article/details/53136188)
[链表随机节点](https://unclegem.cn/2018/09/07/Leetcode%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0-382-%E9%93%BE%E8%A1%A8%E9%9A%8F%E6%9C%BA%E8%8A%82%E7%82%B9/)