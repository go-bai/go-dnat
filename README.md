### DNAT管理工具

#### 一键安装

```bash
apt/yum install jq -y
export APP_VERSION=`curl -s GET https://api.github.com/repos/go-bai/go-dnat/tags\?per_page=1 | jq -r '.[].name'`

bash <(curl -Ls https://raw.githubusercontent.com/go-bai/go-dnat/master/install.sh)
```

```bash
root@dev:~# dnat
NAME:
   dnat - a DNAT management tool

USAGE:
   dnat [global options] command [command options] [arguments...]

COMMANDS:
   append    append a rule to the end of nat chain if it does not exist
   delete    delete a rule by id
   list, ls  list all rules
   get       get one rule by id
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

#### 添加一个DNAT规则

```bash
dnat append -i eth0 -p 10022 -d 10.0.0.1:22
```

#### 查看所有DNAT规则

```bash
root@dev:~# dnat list

ID  Iface  Port   Dest              CreatedAt            
1   eth0   9001   10.0.0.1:9001     2023-05-31 21:37:35  
2   eth0   9002   10.0.0.1:9002     2023-05-31 21:43:10  
3   eth0   9003   10.0.0.1:9003     2023-05-31 21:43:14  
4   eth0   9004   10.0.0.1:9004     2023-05-31 21:43:19  
5   eth0   9005   10.0.0.1:9005     2023-05-31 21:43:23  
6   eth0   1001   192.168.3.1:1001  2023-05-31 21:43:53  
7   eth0   1002   192.168.3.1:1002  2023-05-31 21:43:57  
8   eth0   1003   192.168.3.1:1003  2023-05-31 21:44:01  
9   eth0   1004   192.168.3.1:1004  2023-05-31 21:44:05  
10  eth0   10022  10.0.0.1:22       2023-05-31 22:30:19 
```

#### 删除一个DNAT规则

```bash
root@dev:~# dnat delete -id 1
root@dev:~# dnat ls  

ID  Iface  Port   Dest              CreatedAt            
2   eth0   9002   10.0.0.1:9002     2023-05-31 21:43:10  
3   eth0   9003   10.0.0.1:9003     2023-05-31 21:43:14  
4   eth0   9004   10.0.0.1:9004     2023-05-31 21:43:19  
5   eth0   9005   10.0.0.1:9005     2023-05-31 21:43:23  
6   eth0   1001   192.168.3.1:1001  2023-05-31 21:43:53  
7   eth0   1002   192.168.3.1:1002  2023-05-31 21:43:57  
8   eth0   1003   192.168.3.1:1003  2023-05-31 21:44:01  
9   eth0   1004   192.168.3.1:1004  2023-05-31 21:44:05  
10  eth0   10022  10.0.0.1:22       2023-05-31 22:30:19

root@dev:~# iptables -t nat -L -n -v
Chain PREROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:9002 /* go-dnat */ to:10.0.0.1:9002
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:9003 /* go-dnat */ to:10.0.0.1:9003
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:9004 /* go-dnat */ to:10.0.0.1:9004
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:9005 /* go-dnat */ to:10.0.0.1:9005
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:1001 /* go-dnat */ to:192.168.3.1:1001
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:1002 /* go-dnat */ to:192.168.3.1:1002
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:1003 /* go-dnat */ to:192.168.3.1:1003
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:1004 /* go-dnat */ to:192.168.3.1:1004
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:10022 /* go-dnat */ to:10.0.0.1:22
```