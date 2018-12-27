redis系列之sentinel
1、哨兵
sentinal,中文是哨兵
哨兵是redis集群中非常重要的一个组件，功能如下
（1）、集群监控，负责监控redis master和slave是否正常工作
（2）、消息通知，如果某个redis实例有故障，那么哨兵负责发送消息作为报警通知管理员
（3）、故障转移，如果master node挂了，自动转移到slave node上
（4）、配置中心，如果故障转移发生了，通知client客户端新的master地址
哨兵本身也是分布式的，作为一个哨兵集群去运行，相互协调工作
（1）、故障转移时，判断一个master挂掉了，需要大部分的哨兵都同意才行，涉及到了分布式选举的问题
（2）、即使部分哨兵挂掉了，哨兵集群还是能正常工作的，因为如果一个作为高可用机制重要组成部分的故障转移系统本身是单点的，那确实是很坑爹的。

2、哨兵的核心执行
（1）、哨兵至少要3个实例来保障自己的健壮性
（2）、哨兵+redis主从的部署架构，是不会保证数据零丢失的，只能保证redis集群的高可用性
（3）、对于哨兵+redis主从这种复杂的部署架构，尽量在测试环境和生产环境都进行充足的测试和演练
3、为什么redis哨兵集群只有2个节点无法正常工作
哨兵集群必须部署2个以上节点
如果哨兵集群仅仅部署了2个哨兵实例：quorum=1(即票数达到1时，就会触发failover)
+----+       +----+
| M1 |-------| R1 |
| S1  |           | S2 |
+----+       +----+
 机器1        机器2
M1：master   R1：slave
S1:sentinal1 S2:sentinal2

Configuration :quorum = 1;
master宕机，即M1进程宕机，S1和S2中只要有1个哨兵认为master宕机就可以进行切换，同时S1和S2中会选举出一个哨兵来执行故障转移
同时这个时候，需要majority,也就是大多数哨兵是运行的，2个哨兵的majority就是2（2的majority=2,3的majority=2,4的majority=2,5的majority=3，即超过哨兵数的一半，在redis中称为大多数），2个哨兵都运行着，就可以允许执行故障转移
但是如果整个M1和S1运行的机器挂了，即机器1挂了，那哨兵只有一个了，此时就没有magority来允许执行故障转移，虽然另外一台机器还有一个R1，但是故障转移不会被执行。

sentinel monitor resque 192.168.1.3 6380 4
sentinel down-after-milliseconds resque 10000
sentinel failover-timeout resque 180000
sentinel parallel-syncs resque 5
4、经典的3节点哨兵集群
+-----+
| M1  |
| S1   |
+-----+
                  |
                  |
+-----+      |     +-----+
 | R2   |---------|  R3  |
 | S2   |              |  S3  |
+-----+            +-----+
Configuration:quorum = 2 ,majority
如果M1所在机器宕机了，那么三个哨兵还剩2个，S2和S3可以一致的认为master宕机了，然后选举出一个来执行故障转移（failover），同时3个哨兵的magority是2，所以还剩下的2个哨兵运行着，就可以允许故障转移.

5、SDOWN与ODOWN的转换机制
SDOWN与ODOWN是两种失败状态
SDOWN是主观宕机，就是一个哨兵如果觉得一个master宕机了那就是主观宕机
ODOWN是客观宕机，如果有quorum个哨兵觉得一个master宕机了，那就是客观宕机
SDOWN的达成条件很简单，就是一个哨兵ping一个master，经过了down-after-milliseconds时间还没有收到回复，则认为SDOWN及了
SDOWN到ODOWN转换的条件很简单，如果在一定时间内，收到了quorum个其他哨兵也认为一个master SDOWN了，那就转换为了ODOWN了
6、SDOWN与ODOWN转换过程
当一个master被sentinel集群监控时，需要为它指定一个参数，这个参数指定了当需要判决master为不可用，并且进行failover时，所需要的sentinel数量，这里称这个参数为票数
不过，当failover主备切换真正被触发后，failover并不会马上进行，还需要sentinel中的大多数sentinel授权后才可以进行failover。
当ODOWN时，failover被触发。failover一旦被触发，尝试去进行failover的sentinel会去获得“大多数”sentinel的授权（如果票数比大多数还要大的时候，则询问更多的sentinel)
这个区别看起来很微妙，但是很容易理解和使用。例如，集群中有5个sentinel，票数被设置为2，当2个sentinel认为一个master已经不可用了以后，将会触发failover，但是，进行failover的那个sentinel必须先获得至少3个sentinel的授权才可以实行failover。
如果票数被设置为5，要达到ODOWN状态，必须所有5个sentinel都主观认为master为不可用，要进行failover，那么得获得所有5个sentinel的授权
7、sentinel.conf参数配置相关说明
在sentinel.conf配置文件中配置
sentinel monitor mymaster 127.0.0.1 6379 2(这个就是quorum)
sentinel down-after-milliseconds mymaster 60000(sentinel发送ping到master，如果在这段时间内没有返回或者返回了错误消息，则认为已经宕机了)
sentinel failover-timeout mymaster 180000（如果sentinel A推荐sentinel B去执行failover，B failover timeout，sentinel A自行再次去对同一个master执行failover，这个等待的时间是通过failover-timeout配置项去配置的）
sentinel parallel-syncs mymaster 1（在发生failover主备切换时，这个选项指定了最多可以有多少个slave同时对新的master进行同步，这个数字越小，完成failover所需的时间就越长，但是如果这个数字越大，就意味着越多的slave因为replication而不可用。可以通过将这个值设为 1 来保证每次只有一个slave处于不能处理命令请求的状态）

8、哨兵和slave集群的自动发现机制
哨兵的互相发现通过redis的pub/sub系统实现，每个哨兵都会网_sentinel_:hello这个channel里发送消息，由于其他哨兵监听了这个channel的消息，所以其他哨兵就能感知到其他哨兵的存在了
每隔2秒钟，每个哨兵都会网自己监控的某个master+slaves对应的_sentinel_:hello这个channel里发送消息，内容是自己的ip host 和runid还有对这个master监控的配置
每个哨兵也会去监听自己监控的master+slaves对应的_sentinel_:hello的channel，然后感知同样监控这个master+slaves的其他的哨兵。
每个哨兵还会跟其他哨兵交换自己监控master的配置信息。相互进行监控配置的同步

9、slave配置的自动纠正
哨兵会负责自动纠正slave的一些配置，比如slave要成为潜在的master，哨兵会确保slave复制现有master的数据；
如果slave连接了一个错误的master，如failover后，那么哨兵会确保slave连接到正确的master，即新的master
10、slave到master的选举算法
如果一个master被确认ODOWN了，而且majority哨兵都允许切换，那么某个哨兵就会进行主备切换，此时要选举一个slave，算法如下
（1)跟master断开连接时长，如果过长（超过down-after-milliseconds的10倍）就不适合做新的master，一般这样的slave会被排除掉
（2）slave优先级，每个slave都会有个优先级，在redis.conf中配置，数字越小，优先级越高
（3）复制offset，offset跟原来的master越接近，数据相差最少，就越适合做新的master
（4）run id，run id越小就越容易选上
11、数据复制丢失及脑裂问题
1、两种数据丢失情况
主备切换过程中，可能导致数据丢失问题
（1）异步复制导致数据丢失问题
 由于数据从master到slave是异步复制的，当部分数据还没有复制到slave时，master宕机了，导致数据丢失。
（2)、脑裂导致数据丢失问题
脑裂：是指master与其他slave、哨兵等由于网络原因断开了，但master实际上是还运行着的，客户端继续连着这个旧的master，而此时其他slave又选了新的master，也就在此时旧的master又连接到了网络，此时出现了两个master，这就称为脑裂。
此时虽然有了新的master，但客户端还没有切换过来，客户端还继续向旧的master写数据，当旧的master重新连接到了网络，而此时会把旧的master转换为slave（原来版本更低，有更高版本的时候低版本的master会变成slave），此时旧的master数据会被清空，并从新的master中复制数据，导致这段时间内客户端写的数据丢失。
2、解决异步复制和脑裂导致的数据丢失问题
由于redis不能保证数据零丢失，但可以尽可能少的丢失数据，正对异步复制和脑裂导致数据丢失解决方案如下：
配置两个参数（在redis.conf中配置）
min-slaves-to-write 1
min-slaves-max-lag 10
要求至少有1个slave，数据复制和同步的延迟不能超过10秒
如果说所有的slave数据复制和同步都延迟超过了10秒，这个时候，master就不会在接受客户端的数据了，这保证了最多会丢失10秒钟的数据。
（1）、减少异步复制的数据丢失
有了min-slaves-max-lag这个配置，就可以确保说，一旦slave复制数据和 ack延时太多，就认为由于master宕机了导致数据丢失太多了，那么就拒绝谢请求，这样可以把宕机后数据丢失降低到可控范围内
（2）、减少脑裂数据丢失
如果一个master出现了脑裂，与其他slave丢失了，以上两个参数配置可以保证说如果给定数量的slave发送消息，在10秒内不能收到ack消息，那么就直接拒绝客户端写请求，也就避免了数据丢失的情况，上面的配置就确保了，如果跟任何一个slave丢了连接，在10秒内没有slave给自己发送ack消息，就拒绝新的请求，因此在脑裂的时候最多丢失10秒钟的数据

12、sentinel启动及状态查看
启动sentinel

./redis-sentinel  ../sentinel.conf

如果是在生产环境中，需要配置后台启动，并且指定日志文件，在sentinel.conf中加入以下信息

daemonize yes 

logfile /var/log/redis/sentinel/sentinel.log

日志会打印哨兵监控的master并且会发现连接master的slave，哨兵之间也会相互发现

查看sentinel状态

进入sentinel客户端

./redis-cli -h 10.202.15.198 -p 26379

然后输入

sentinel master mymaster(配置文件中对应的名字)

sentinel slaves mymaster

sentinel sentinels mymaster

sentinel get-master-addr-by-name mymaster
--------------------- 
作者：Simba_hua 
来源：CSDN 
原文：https://blog.csdn.net/simba_1986/article/details/77531497 
版权声明：本文为博主原创文章，转载请附上博文链接！
