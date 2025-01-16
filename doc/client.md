首先是客户端client模块的实现

分为三个模块
- cmd：命令解析模块，解析命令行参数，作为项目的部署入口
- cui：ui界面模块，用于端到端调试
- sdk：客户端逻辑层，封装客户端行为

技术选型上采用的是[cobra](https://github.com/spf13/cobra)这一个组件来用为命令解析，而[gocui](https://github.com/jroimartin/gocui)这个是用来作为暂时的ui交互层，先实现初始的一个界面，后续会有替换

客户端整体的流程就是：
通过cobra解析命令cmd，执行客户端的逻辑
进入到ui界面 等待用户输入信息
用户输入信息 *cui*调用`sendMsg()` 经过sdk执行发送消息的逻辑，消息通过发送端关联的channel，发送到接收端关联的channel 调用`receiveMsg()`返回消息到用户输出

- [x] 实现一个简单的复读机，即输入与输出相同 此时还不能进行多客户端之间的交互
