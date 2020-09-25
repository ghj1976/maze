

迷宫生成算法是处于这样一个场景：

* 一个row行，col列的网格地图，一开始默认所有网格四周的墙是封闭的
* 要求在网格地图边缘，也就是网格的边上打通2面墙
* 所有网格都至少保证网格周围至少有一堵墙打通
* 所有网格都能通过打通的墙能形成一条通路



这三种算法分别适合不同的迷宫情况，

* 深度优先适合于那种主线支线明显的游戏（如RPG），
* 递归分割则适合转角较少的游戏（也许是FPS和ACT）
* prim，似乎适合最标准的迷宫游戏（因为很难走）。


**参考研究的文章**

[三种迷宫生成算法概述](https://www.jianshu.com/p/f643b0a0b887)

[三大迷宫生成算法 (Maze generation algorithm) -- 深度优先，随机Prim，递归分割](https://blog.csdn.net/juzihongle1/article/details/73135920)

[游戏常用算法：四种迷宫生成算法](https://www.gameres.com/829184.html)


[Maze generation algorithm](https://en.wikipedia.org/wiki/Maze_generation_algorithm)