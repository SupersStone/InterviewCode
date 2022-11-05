# InterviewCode

### 目录一：contrats 是合约模块  
#####  1.开发框架：hardhat  
1. 目录结构:        
- contract : 合约部分  
- scripts : 部署脚本
- tests : 测试脚本

2. 编译命令
- npx hardhat compile

3. 部署网络
- bsctest 币安测试网

3. 部署命令
- npx hardhat run scripts/deploy.js --network bsctest
- 部署测试网合约地址：0x290908A9bd589cDC95aafef63528Ba573C7B8d9E

4. 验证合约源码
-  npx hardhat verify 0x290908A9bd589cDC95aafef63528Ba573C7B8d9E "https://hei she bei jing tu" "Toekn" "Token" --network bsctest


###### 合约空投白名单实现方法
###### js版本简单实现默克尔树：contracts/merkletreeProf
- 方法一： 简单粗暴，在合约中添加一个map来映射白名单，由管理天添加白名单到合约中，简单除暴，手续费高
- 方法二：通过默克尔树的形式验证白名单，通过离线生成默克尔树，空投用户通过请求后端拿到莫克尔证明后调用合约，实现白名单，优点手续费低，缺点是只能一次添加白名单，不能后面新增白名单
- 方法三：后端签名，合约校验签名信息， 优点是能新增百名单，不像默克尔树那样需要一次确定下来





### 目录二： go-code 是后端模块

##### 1. 开发框架： go-zero
1. 目录结构
- apis ： 对外的api接口层
- apps : 链上扫描处理
- cmd: 扫链启动
- config : 配置文件
- docker :docker 配置文件
- iternal: kafka 基础抽象等
- pkg: 通过的包

#### 2. 原理：
1. 原理是通过扫描链上解析合约事件入库数据

