package main

import (
	"fmt"
	"sync"
	"time"
)

/*
接下来，看一个百米赛跑开始时的例子，来学习Cond的使用方法。10个运动员进入赛场之后需要先做拉伸活动活动筋骨，向观众和粉丝招手致敬，在自己的赛道上做好准备；等所有的运动员都准好之后，裁判员才会打响发令枪。

每个运动员做好准备之后，将ready加一，表明自己做好准备了，同时调用Broadcast方法通知裁判员。因为裁判员只有一个，所以这里可以直接替换成Signal方法调用。

裁判员会等待运动员都准备好。虽然每个运动员准备好之后都唤醒了裁判员，但是裁判员被唤醒之后需要检查等待条件是否满足（运动员都准备好了）。可以看到，裁判员被唤醒之后一定要检查等待条件，如果条件不满足还是要继续等待。
*/
func main() {
	var condition = sync.NewCond(&sync.Mutex{})
	var readyNumber = 0

	for i := 0; i < 10; i++ {
		go func(seq int) {
			time.Sleep(1000)

			// 加锁更改临界变量
			condition.L.Lock()
			readyNumber++
			condition.L.Unlock()
			fmt.Printf("运动员 %v 已就绪\n", seq)

			// 广播唤醒所有的等待者，此时就会进入到 condition 的判断 line 36: for readyNumber != 10
			condition.Broadcast()
		}(i)
	}

	condition.L.Lock()
	for readyNumber != 10 {
		condition.Wait()
	}
	condition.L.Unlock()

	// 所有的运动员就绪
	fmt.Println("均已就绪！准备开始")
}
