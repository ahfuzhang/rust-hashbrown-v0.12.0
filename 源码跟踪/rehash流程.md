rehash分为三种情况：

* 删除很多，在不扩容的情况下，本地rehash
* 空间不足，增长一倍后，rehash
* 删除过多，缩容后rehash



本地rehash

| 源码 | 函数 | 调用了 | 被调函数 | 说明 |
| ---- | ---- | ------ | -------- | ---- |



删除不会导致rehash

? 查找会不会导致rehash

插入肯定导致rehash



缩容的过程：
| 源码                                | 函数                          | 调用了 | 被调函数                  | 说明                                         |
| ----------------------------------- | ----------------------------- | ------ | ------------------------- | -------------------------------------------- |
| hashbrown-0.12.0\src\map.rs:917     | HashMap::shrink_to_fit()      |        |                           |                                              |
|                                     |                               | :919   | self.table.shrink_to      |                                              |
| hashbrown-0.12.0\src\raw\mod.rs:611 | RawTable::shrink_to           |        |                           |                                              |
|                                     |                               | :637   | RawTable::resize()        |                                              |
| :702                                | RawTable::resize()            |        |                           |                                              |
|                                     |                               | :709   | self.table.resize_inner() |                                              |
| :1425                               | RawTableInner::resize_inner() |        |                           | 缩容，新分配一个hash表，然后把元素拷贝过去。 |

缩容的流程也很简单



