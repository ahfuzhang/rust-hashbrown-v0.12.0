
| 源码 | 函数 | 调用了 | 被调函数 | 说明 |
| ---- | ---- | ---- | ---- | ---- |
| hashbrown-0.12.0/src/map.rs:188 | pub struct HashMap | - | - | hashmap的结构定义 |
| | | :190 | `pub(crate) table: RawTable<(K, V), A>` | 具体的内部实现 |
| hashbrown-0.12.0/src/raw/mod.rs:367 | pub struct RawTable | | | hashtable的内部实现 |
| :1049 | new_in | | | hash表的初始化 |
|  |  | | |  |
| :440 | RawTable::new_uninitialized | | |  |
|  |  | :448 | RawTableInner::new_uninitialized() |  |
| :1063 | new_uninitialized | | | hash表的初始化 |
|  |  | :1072 | table_layout.calculate_layout_for | 计算bucket数组和ctrl数组的布局分布 |
|  |  | :1086 | do_alloc | 分配内存的函数 |
|  |  | :1091 | let ctrl = xxx | 偏移，找到ctrl数组的位置 |
| hashbrown-0.12.0/src/raw/alloc.rs:10 | pub fn do_alloc |  |  | 调用分配器的方法来分配内存<br />搞不懂为何如此复杂 |
|  |  |  |  |  |



扩容时候的分配：
| 源码 | 函数 | 调用了 | 被调函数 | 说明 |
| ---- | ---- | ---- | ---- | ---- |
| hashbrown-0.12.0\src\raw\mod.rs:1432 |  |  | self.prepare_resize() | resize_inner()中分配新的table |
|  |  | :1435 |  | 遍历旧的table |
|  |  | :1437 |  | 跳过空元素 |
|  |  | :1441 | hasher(self, i); | 重新计算key的hashcode |
|  |  | :1447 | new_table.prepare_insert_slot(hash); | 在新表中插入 |
|  |  | :1449 | ptr::copy_nonoverlapping() | 拷贝K,V |
| :1350 | RawTableInner::prepare_resize() |  |  |  |
|  |  | :1359 | RawTableInner::fallible_with_capacity() | 构造一个新的表 |
|  |  | :1374 |  | 返回guard对象 |
| :1102 | RawTableInner::fallible_with_capacity() |  |  | 通过容量来构造一个table对象 |
|  |  | :1113 | capacity_to_buckets() | 计算桶的个数 |
|  |  | :1115 | Self::new_uninitialized() | 构造table对象 |
|  |  | :1116 | result.ctrl(0).write_bytes() | memset控制数组 |
| :1063 | RawTableInner::new_uninitialized() |  |  | 静态构造方法 |
|  |  | :1072 | table_layout.calculate_layout_for() | 计算需要的总空间 |
|  |  | :1086 | do_alloc() | 执行分配 |
|  |  | :1092 |  | 返回对象 |

