##  process by not buf
1. 创建一个发送者列表和接收者列表都为空的 channel
2. 第一个协程向 channel 发送foo变量的值，第 16 行
3. channel 从池中获取一个sudog结构体变量，用于表示发送者。sudog 结构体会保持对发送者所在协程的引用，以及foo的引用
4. 发送者加入sendq队列
5. 发送者协程进入等待状态
6. 第二个协程将从 channel 中读取一个消息，第 23 行
7. channel 将sendq列表中等待状态的发送者出队列
8. chanel 使用memmove函数将发送者要发送的值进行拷贝，包装入sudog结构体，再传递给 channel 接收者的接收变量
9. 在第五步中被挂起的第一个协程将恢复运行并释放第三步中获取的sudog结构体。

