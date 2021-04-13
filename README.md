# kongctl
kong 网关的命令行工具，方便在shell命令行下操作，并且支持多机器多对象批量处理，提高配置效率。命令行格式如下：
```
kongctl [global flags] [object] [command] [flags]
# global flags: --hosts ...
# object: upstream, service ...
```
## 功能
 
### [1] 支持本地或者远程执行命令
- 所有命令都可以指定操作的主机(支持ip地址或domain)，可以指定多个，格式为:`--hosts IP:PORT/Domain:PORT`。如果没有指定PORT,默认端口是8001  
- 如果没有指定主机，默认本地执行（127.0.0.1:8001)  
- 命令示例：`kongctl -H 192.168.1.1:8001 --hosts 10.196.3.203 ...`
 
### [2] Object upstream
**显示所有upstreams，仅列出部分必要信息**  
`kongctl upstream show`
 
**显示单个upstream的信息**  
`kongctl upstream show --uname {uname1, uname2 ...}`
 
**显示所有upstream的 targets 包括健康状态**  
`kongctl upstream show-targets --uname {uname1, uname2 ...}`
 
**创建upstream, 健康检查策略默认加上**  
`kongctl upstream create --uname {uname1, uname2 ...}`

**为upstream 创建targets**  
可以批量创建，前提是同一批 upstream 拥有相同的targets。  
可以指定weight, 默认值为100。  
```
kongctl \
upstream add-targets \
--uname  {uname1, uname2 ...} \
--target {target1, target2...} \
--weight {optionnal: weight} \
--hosts  {optional: host1, host2 ...}
```
**为upstream 删除targets**  

```
kongctl \
upstream del-targets \
--uname  {uname1, uname2 ...} \
--target {target1, target2...} \
--weight {optionnal: weight} \
--hosts  {optional: host1, host2 ...}
```
### [3] Object service
 
**显示所有service**  
`kongctl service show` 
 
**显示指定service**  
`kongctl service show --sname {SERVICE-NAME}` 
 
**显示service的route**  
`kongctl service show-routes --sname {SERVICE-NAME}` 
 
**显示service的plugin**  
`kongctl service show-plugins --sname {SERVICE-NAME}` 

## 用例
```
kongctl upstream show-targets  \
--uname upstream1,upstream2

kongctl upstream add-targets \
--uname upstream1,upstream2 \
--target 192.168.1.1:80,192.168.1.2:80,192.168.1.3:80

kongctl upstream del-targets \
--uname upstream1,upstream2 \
--target 192.168.1.1:80,192.168.1.2:80,192.168.1.3:80
```