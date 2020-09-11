```


Go控制并发有三种种经典的方式，一种是通过channel通知实现并发控制 一种是WaitGroup，另外一种就是Context


1.WaitGroup 这种方式是控制多个goroutine同时完成

N个goroutine同时做完，才算是完成，先做好的就要等着其他未完成的，所有的goroutine要都全部完成才可以。
 
这是一种控制并发的方式，这种尤其适用于，好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，
只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式




2.chan通知

我们都知道一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的。

这种情况化，一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台goroutine不停的检查这个变量，如果发现被通知关闭了，就自我结束。

这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 



问题 ：

这种chan+select的方式，是比较优雅的结束一个goroutine的方式，不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？
如果这些goroutine又衍生了其他更多的goroutine怎么办呢？如果一层层的无穷尽的goroutine呢？这就非常复杂了，即使我们定义很多chan也很难解决这个问题，
因为goroutine的关系链就导致了这种场景非常复杂


3.Context

上面说的这种场景是存在的，比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。
所以我们需要一种可以跟踪goroutine的方案，才可以达到控制他们的目的，这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文



```
