# myDelay_queue
1.redis的有一种特殊的chanel可以监听键空间的事件-> keyevent_expired；我们可以监听键过期然后处理key val的事件
  消息全部丢失，因为redis清理过期键值对。
2.通过redis的zset的score来存放过期时间
