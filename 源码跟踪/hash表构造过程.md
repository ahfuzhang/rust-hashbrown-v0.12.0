
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

