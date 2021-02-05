# redisgo
redigo二次封装  分布式锁  
使用list结构(双向链表)实现消息队列时，如何进行消息ACK；当队列为空时，lpop和rpop会一直空轮训，消耗资源；所以引入阻塞读blpop和brpop（b代表blocking），阻塞读在队列没有数据的时候进入休眠状态，


