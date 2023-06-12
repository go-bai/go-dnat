### DNAT管理工具

#### 一键安装

```bash
apt install jq curl -y
bash <(curl -Ls https://raw.githubusercontent.com/go-bai/go-dnat/master/install.sh)
```

```bash
root@dev:~# dnat -h
NAME:
   dnat - a DNAT management tool

USAGE:
   dnat [global options] command [command options] [arguments...]

COMMANDS:
   append      append a rule to the end of nat chain if it does not exist
   delete      delete a rule by id
   list, ls    list all rules
   get         get one rule by id
   sync        sync rules to local machine
   version, v  print version
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

#### 添加一个DNAT规则

```bash
dnat append -i eth0 -p 8888 -d 10.0.0.1:9999
```

#### 查看所有DNAT规则

```bash
root@dev:~# dnat ls

ID  Iface  Port  Dest           CreatedAt            
1   eth0   8888  10.0.0.1:9999  2023-06-12 05:08:45

```

##### 查看在iptables中创建了什么

```bash
root@dev:~# iptables -t nat -nvL
Chain PREROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DNAT       tcp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            tcp dpt:8888 /* go-dnat */ to:10.0.0.1:9999
    0     0 DNAT       udp  --  eth0   *       0.0.0.0/0            0.0.0.0/0            udp dpt:8888 /* go-dnat */ to:10.0.0.1:9999

Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain POSTROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination
```

#### 删除一个DNAT规则

```bash
root@dev:~# dnat delete -id 1
```