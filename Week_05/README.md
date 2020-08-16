# [Week5] - 学习笔记

## 树

树是一种层状结构的数据结构, 可以自然地以层状的方式存储信息.

### AVL 树

#### AVL 树的定义

AVL 树是一种平衡二叉树. 平衡二叉树的递归定义如下:
> 1. 左右子树的高度差为 $[-1, 0, 1]$
> 
> 2. 其每一个子树均为平衡二叉树

- **平衡因子**: 某个结点的右子树的 **高度** 减去左子树的 **高度** 得到的差值.
- **AVL 树**: 所有结点的平衡因子的绝对值都不超过 1 的二叉树

#### AVL 树的平衡化操作

二叉树的平衡化操作有 4 种:

1. 左旋
2. 右旋
3. 左右旋
4. 右左旋

其中 **左旋** 和 **右旋** 为基本操作.

1. 左旋操作

    ![左旋](https://pic1.zhimg.com/v2-591b2b8b5dece5da4936d823a07a3f39_b.gif)

2. 右旋操作

    ![右旋](https://pic4.zhimg.com/v2-772eb4f59b7f5ba6da045f220a506d04_b.gif)

当出现以下四种情况时, 可以采用对应的平衡化操作:

1. 左-左型: 右旋
2. 右-右型: 左旋
3. 左-右型: 先左旋, 再右旋
4. 右-左型: 先右旋, 再左旋

### 红黑树

红黑树是一种 **近似平衡** 的二叉搜索树, 它能够确保任何一个结点的左右子树的 **高度差小于两倍**. 

#### 红黑树的特点
1. 每个结点要么是红色, 要么是黑色
2. 根节点是黑色
3. 每个叶子节点 (NIL 结点, 空结点) 是黑色的
4. 不能有相邻接的两个红色结点
5. 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点

### AVL 树与红黑树的对比

1. AVL trees provide **faster lookups** than Red Black Trees because they are **more strictlybalanced**.

2. Red Black Trees provide **faster insertion and removal** operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.

3. AVL trees store balance **factors or heights** with eachnode, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of infomation per node.

4. Red Black Trees are used in most of the **language libraries**. **like map, multimap, multisetin C++** whereas AVL trees are used in **databases** where faster retrievals are required.


## 位运算

### 位运算符

|含义|运算符|示例|
|:-:|:-:|:-:|
|左移|`<<`|$0011 \Rightarrow 0110$|
|右移|`>>`|$0110 \Rightarrow 0011$|

|含义|运算符|示例|
|:-:|:-:|:-:|
|按位或|`|`|$0011 \| 1011 \Rightarrow 1011$|
|按位与|`&`|$0011 & 1011 \Rightarrow 0011$|
|按位取反|`~`|$~0011 \Rightarrow 1100$|
|按位异或|`^`|$0011 ^ 1011 \Rightarrow 1000$|

### XOR-异或的特点
1. `x ^ 0 = x`
2. `x ^ 1s = ~x`
3. `x ^ (~x) = 1s`
4. `x ^ x = 0`
5. `c = a ^ b => a ^ c = b, b^ c = a`
6. `a ^ b ^ c = a ^ ( b ^ c ) = ( a ^ b ) ^ c`

### 指定位置的位运算

1. 将 `x` 最右边的 `n` 位清零: `x & (~0 << n)`
2. 获取 `x` 的第 `n` 位的值(0, 1): `(x >> n) & 1`
3. 获取 `x` 的第 `n` 位的幂值: `x & (1 << n)`
4. 仅将第 `n` 位置为 `1`: `x | (1 << n)`
5. 仅将第 `n` 位置为 `0`: `x & (~(1 << n))`
6. 将 `x` 最高位至第 `n` 位(含)清零: `x & ((1 << n) - 1)`

### 实战位运算要点
- 判断奇偶
    1. `x % 2 == 1` --> `x & 1 == 1`
    2. `x % 2 == 0` --> `x & 1 == 0`
- `x >> 1` --> `x / 2`
- `x = x & (x - 1)`: 清零最低位的 `1`
- `x & -x`: 得到最低位的 `1`
- `x & -x ==> 0`