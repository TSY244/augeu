# 参考

1. 日志分析

可以参考

> 一个实时的日志分析工具
>
> https://github.com/Fheidt12/Windows_Log

> 一个可以将存量的日志倒入分析的工具
>
> https://github.com/xbaogua/BaoGuaWindowsEvent

1. 综合类型的工具

> https://github.com/mir1ce/Hawkeye

> [QDoctor.exe](https://poizon.feishu.cn/file/Z5vVb0iRwoaAahxFNqEcWG9Qnyb)

# 日志分析

## 功能

包含

1. 登陆成功/失败的日志 可以检测内网密码喷洒或者是爆破记录
2. RDP 登陆日志
3. RDP 连接日志
4. 服务创建日志
5. 用户日志 用户创建的日志 用户添加到特权组的日志
6. 创建进程日志
7. MSSQL 日志
8. powershell
9. lsass 读取痕迹 通过事件id 4663 (对象访问审计)进行查看 是否使用高/异常的权限
10. Kerberos 日志读取信息 主要是防止白银票据和黄金票据 重点关注域内主机非正常的服务请求
11. DCSync 日志 防止 凭证窃取技术
12. SID 后门
13. 域内收集痕迹
14. ZeroLogon 记录 重点是防止 CVE-2020-1472
15. 系统日志
16. 应用日志
17. 安全日志

## 实现

一共有四种方式可以进行日志分析

### 通过powershell

通过powershell Get-WinEvent 的方式获取windows 日志，然后将其变成xml 格式

然后通过程序进行解析xml 进行分析

### logparser

通过外部程序调用的方式，调用logparser 进行分析

###  直接提取(可能使用这种)

https://github.com/0xrawsec/golang-evtx

### 通过 wevtapi.dll

wevtapi.dll 存在于windows 系统的动态链接库

![img](https://poizon.feishu.cn/space/api/box/stream/download/asynccode/?code=NGVmNGIzODQwMzk0YjJjMzM5ODA5YjYyYTM2NDk2YzRfalR3VXhMM1VtdDN5UXlMNjIxdG5aVktaV3B5SjRWVERfVG9rZW46QzBPV2JEaDB1bzA1WHR4VGluZGNuYU9tbm5mXzE3NDMzMzY4NjQ6MTc0MzM0MDQ2NF9WNA)

# 进程模块

## 进程信息展示

### 展示信息（中）

1. PID
2. 进程名称
3. 父进程PID
4. 父进程名称
5. 创建时间
6. 可执行文件路径
7. MD5
8. 签名信息
9. 导入资源
    1. 导入路径
    2. MD5
    3. 签名

## 进程扫描（内存查杀）

通过现有的yara 规则进行扫描

https://github.com/InQuest/yara-rules?tab=readme-ov-file

该仓库需要过滤出静态文件的yara

### 实现

https://github.com/hillu/go-yara

## 内存关键字检索

https://github.com/Fheidt12/Windows_Memory_Search

https://www.freebuf.com/sectool/408673.html

https://www.52pojie.cn/thread-1484643-1-1.html

尽可能添加正则匹配的方式

### 实现

可以通过go -> windows api 的方式进行获取其中的信息

通过打开进程句柄然后遍历字符串的形式进行检索

# 文件历史记录（暂时没实现）

## Prefetch记录

可以查看近期高频运行的程序，还可以看到程序运行的时间记录

## UserAssit话动记录

记录程序的启动次数

## Recent File记录

最近使用的文件

# 主机信息（重点）

## 用户信息

有多少个账号

## 定时任务

## 服务信息

## 启动项信息

## 镜像劫持信息

# 网络信息展示（重点）

1. 进程名
2. Pid
3. 协议
4. 本地监听地址
5. 本地监听端口
6. 远程地址
7. 远程端口
8. 状态

# “一键”功能原理

## 一键检查

1. 通过日志分析 检查是否存在爆破日志 Rdp 登陆成功 用户创建等日志
2. 通过进程扫描 通过yara检查 关键字检索
3. 是否存在不正常的程序打开次数
4. 检查是否存在不正常用户或者是否被劫持

## 一键恢复安全

我更推荐通过检查出的信息

然后在界面 选择 修改/复原的方式，而不是使用程序直接一键删除

> 比如现在扫描出了 有隐藏用户 然后该模块中，显示了修复方法，让用户点击是否需要一键删除该用户（防止工具误删除）

如果是一定需要一键恢复安全，可以先利用一个弹窗，提示后面要做的一些工作

然后再让使用者点击确认

然后再执行，并且需要把执行的过程和敏感操作使用日志文件等方式进行记录

支持一下功能

1. 删除RID 劫持的用户 可以通过操作注册表
2. 删除影子用户
3. 非法服务删除
4. 非法定时任务删除
5. 劫持检查

# 工具设计

我现在想的四个方案

## 桌面程序

通过go 的gui 库画出一个桌面程序

## 前后端

通过前端展示界面

后端在启动部分完成一部分信息的检索

前端访问，直接通过sql 差出来的结果进行返回

## agent+server 模式（采用）

在怀疑的主机上放置一个agent

通过server 命令下发进行信息返回

这样可以抽象成

agent->server->前端

## 纯命令行工具

通过csv 到处需要的数据，或者通过终端日志的形式返回结果


# 架构设计

## 整体设计

整个工具分为两个程序

第一个程序是agent

Agent 是放到可能被控的电脑上

是一个信息获取的一个锚点

第二个程序是server

提供数据给前端进行一个展示

### 功能介绍图

![img](https://poizon.feishu.cn/space/api/box/stream/download/asynccode/?code=MTk1MTZhYjVjZmUwMzc5NmE0YzRkYjM0NDA4OTMxOGRfaUxDczZ4c1R1cnh0Z0tMbVFGWUJIOUpNV2xaVlNBR1hfVG9rZW46VU0zeWJVVkdxbzB4MFh4NWhaNmNzODVXbndnXzE3NDMzMzY2NjA6MTc0MzM0MDI2MF9WNA)



### 工具交互图

![img](https://poizon.feishu.cn/space/api/box/stream/download/asynccode/?code=YmFiMmUyYWY3NGJhOWFlYTdjMGZjOTU0ODFjMDI3ZWJfYjYybndGR2Uxb3lBbWZ2OVBxUTlwTUp0M1p5ZzVGcEFfVG9rZW46RDBoMmJlNWU3bzVrTUR4T1NDbGN2UmVFbkRlXzE3NDMzMzY4MTc6MTc0MzM0MDQxN19WNA)



## Agent

使用go 源码硬写应该就行，不需要什么框架

## Server

使用gin/go swagger  提供网络服务

## 前端

由于需要下发指令，可能还需要一个websocket 端口进行一个监听

通过websocket 还可以获取到agent 端的一个日志信息的获取

## 数据库设计

使用 postgresql 作为后端数据库

## 表设计（暂定）

### 日志表

暂定将所有的日志放同一个表

但是防止数据库的数据过大

可能在插入数据的时候会设置一个deleteAt 用于软删除

暂定一个月进行软删除

并且是有两种情况下会触发日志的收集

1. 打开软件的时候主动收集一遍，并且上传
2. 用户手动收集一遍

**字段****设计**

可能需要一个hostname 或者什么作为唯一id

```sql
CREATE TABLE Event_4624 (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                source_ip TEXT,
                source_name TEXT,
                target_name TEXT,
                logon_type TEXT,
                logon_proc TEXT,
                LogonProcessName TEXT,
                AuthenticationPackageName TEXT,
                description TEXT
        );
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE Event_4625 (
                   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                source_ip TEXT,
                source_name TEXT,
                target_name TEXT,
                logon_type TEXT,
                logon_proc TEXT,
                LogonProcessName TEXT,
                AuthenticationPackageName TEXT,
                description TEXT
        );
CREATE TABLE Event_User (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                Source_name TEXT,
                Source_domain TEXT,
                Target_name TEXT,
                Target_domain TEXT,
                MemberSid TEXT,
                description TEXT
        );
CREATE TABLE Event_SIDHistory (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            create_time DATETIME,
            event_id TEXT,
                Source_Name TEXT,
                Source_Domain TEXT,
                Source_SID TEXT,
                Target_Name TEXT,
                Target_Domain TEXT,
                Target_SID TEXT,
                SidHistory TEXT,
                description TEXT
        );
CREATE TABLE Event_LsassAccess (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            create_time DATETIME,
            event_id TEXT,
                Source_Name TEXT,
                Source_Domain TEXT,
                Source_Process_Name TEXT,
                Target_Process_Name TEXT,
                description TEXT
        );
CREATE TABLE Event_CreateProcess (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            create_time DATETIME,
            event_id TEXT,
                Create_User TEXT,
                Create_User_Domain TEXT,
                NewProcessName TEXT,
                ParentProcessName TEXT,
                CommandLine TEXT,
                description TEXT
        );
CREATE TABLE Event_7045 (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                Service_Name TEXT,
                Create_Account TEXT,
                Service_Filename TEXT
        );
CREATE TABLE Event_RDPLogon (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                LoginName TEXT,
                Address TEXT,
                Domain TEXT,
                description TEXT
        );
CREATE TABLE Event_RDPConnect (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                LoginName TEXT,
                Address TEXT,
                Domain TEXT,
                description TEXT
        );
CREATE TABLE Event_PowerShell (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                Command TEXT,
                description TEXT
        );
CREATE TABLE ApplicationEvent (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                LevelDisplayName TEXT,
                description TEXT
        );
CREATE TABLE SystemEvent (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                LevelDisplayName TEXT,
                description TEXT
        );
CREATE TABLE SecurityEvent (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                create_time DATETIME,
                event_id TEXT,
                LevelDisplayName TEXT,
                description TEXT
        );
CREATE TABLE ProcessInfo (
            ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            PID TEXT,
                ProcessName TEXT,
                PPID TEXT,
                ParentName TEXT,
                UserName TEXT,
                Service TEXT,
                CreateTime DATETIME,
                Network TEXT,
                ExePath TEXT
        );
CREATE TABLE TaskInfo (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                TaskName TEXT,
                CreateUserName TEXT,
                ImagePath TEXT,
                Status TEXT,
                CreateTime TEXT,
                description TEXT
        );
CREATE TABLE ServiceInfo (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                ServiceName TEXT,
                ImagePath TEXT,
                StartType TEXT,
                Account TEXT,
                description TEXT
        );
CREATE TABLE UserInfo (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                DomainName TEXT,
                UserName TEXT,
                SID TEXT,
                Disabled TEXT,
                description TEXT
        );
CREATE TABLE StartUpInfo (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                StartUpName TEXT,
                ImagePath TEXT
                
        );
```

### 进程表（可要可不要）

可以通过websocket 的方式传输实时的数据

可以不需要存

如果需要的话下面的设计

```sql
-- 主表：进程基础信息
CREATE TABLE process_info (
    pid INTEGER PRIMARY KEY,
    process_name TEXT NOT NULL,
    parent_pid INTEGER,
    parent_name TEXT,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL,
    executable_path TEXT NOT NULL,
    file_create_time TIMESTAMP WITH TIME ZONE,
    file_modify_time TIMESTAMP WITH TIME ZONE,
    file_md5 CHAR(32),
    signature_info TEXT,
    CONSTRAINT fk_parent 
        FOREIGN KEY(parent_pid) 
        REFERENCES process_info(pid)
);

-- 子表：进程加载资源信息
CREATE TABLE process_imports (
    import_id SERIAL PRIMARY KEY,
    pid INTEGER NOT NULL,
    import_path TEXT NOT NULL,
    import_md5 CHAR(32),
    import_signature TEXT,
    CONSTRAINT fk_process
        FOREIGN KEY(pid)
        REFERENCES process_info(pid)
        ON DELETE CASCADE
);
```

并且这个数据只会存在一天，一天之后就会被软删除

### 活动痕迹

主要有三个表

1. Prefetch 记录
    1.  启动时间 可执行文件名 可执行文件md5 可执行文件路径
2. UserAssit 活动记录 文件最后一次执行时间 可执文件名字 可执行文件路径 运行次数 聚焦次数
3. Recent File 记录 文件名字 文件路径 创建时间 修改时间 目标文件路径

```sql
-- Prefetch记录表
CREATE TABLE prefetch_record (
    id SERIAL PRIMARY KEY,
    launch_time TIMESTAMP NOT NULL,
    executable_name VARCHAR(255) NOT NULL,
    executable_md5 CHAR(32) NOT NULL,
    executable_path TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- UserAssist活动记录表
CREATE TABLE user_assist_record (
    id SERIAL PRIMARY KEY,
    last_executed TIMESTAMP NOT NULL,
    executable_name VARCHAR(255) NOT NULL,
    executable_path TEXT NOT NULL,
    run_count INT DEFAULT 0 CHECK (run_count >= 0),
    focus_count INT DEFAULT 0 CHECK (focus_count >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- RecentFile记录表
CREATE TABLE recent_file_record (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    creation_time TIMESTAMP NOT NULL,
    modification_time TIMESTAMP NOT NULL,
    target_path TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 主机信息

总共有5张表

```sql
-- 账号信息表（存储用户账户状态）
CREATE TABLE account_info (
    account_name VARCHAR(255) PRIMARY KEY NOT NULL,
    is_suspicious BOOLEAN NOT NULL DEFAULT false
);

-- 定时任务表（记录计划任务配置）
CREATE TABLE scheduled_tasks (
    hostname VARCHAR(255) NOT NULL,
    task_name TEXT NOT NULL,
    next_run_time TIMESTAMP,
    mode VARCHAR(50),
    logon_status VARCHAR(50),
    last_run_time TIMESTAMP,
    last_result INTEGER,
    creator VARCHAR(255),
    task_action TEXT,
    start_condition VARCHAR(50),
    comment TEXT,
    task_status VARCHAR(50) NOT NULL,
    idle_settings VARCHAR(50),
    power_management TEXT,
    run_as_user VARCHAR(255) NOT NULL,
    delete_expired_task VARCHAR(50),
    execution_time_limit INTERVAL,
    schedule_data TEXT,
    schedule_type VARCHAR(50),
    start_time TIME,
    start_date DATE,
    end_date DATE,
    PRIMARY KEY (hostname, task_name)
);

-- 服务信息表（记录系统服务配置）
CREATE TABLE service_info (
    service_name VARCHAR(255) PRIMARY KEY,
    service_status VARCHAR(50),
    description TEXT,
    executable_path TEXT NOT NULL,
    signature_info TEXT
);

-- 启动项表（记录自启动程序）
CREATE TABLE startup_items (
    item_name VARCHAR(255) PRIMARY KEY,
    executable_path TEXT NOT NULL,
    signature_info TEXT
);

-- 镜像劫持表（记录文件关联劫持信息）
CREATE TABLE image_hijacks (
    hijack_name VARCHAR(255) PRIMARY KEY,
    hijack_path TEXT NOT NULL,
    hijack_status BOOLEAN NOT NULL DEFAULT false
);
```

**账号信息**

- 账号名称
- 是否可疑

**定时****任务**

```sql
主机名:                             DELL-G2N2B44
任务名:                             \Microsoft\Windows\WwanSvc\OobeDiscovery
下次运行时间:                       N/A
模式:                               就绪
登录状态:                           交互方式/后台方式
上次运行时间:                       2025/3/20 14:52:54
上次结果:                           0
创建者:                             N/A
要运行的任务:                       COM 处理程序
起始于:                             N/A
注释:                               N/A
计划任务状态:                       已启用
空闲时间:                           已禁用
电源管理:
作为用户运行:                       SYSTEM
删除没有计划的任务:                 已禁用
如果运行了 X 小时 X 分钟，停止任务: 01:00:00
计划:                               计划数据在此格式中不可用。
计划类型:                           未定义的
开始时间:                           N/A
开始日期:                           N/A
结束日期:                           N/A
天:                                 N/A
月:                                 N/A
重复: 每:                           N/A
重复: 截止: 时间:                   N/A
重复: 截止: 持续时间:               N/A
重复: 如果还在运行，停止:           N/A
```

**服务****信息**

```sql
服务名
服务状态
服务描述
服务可执行文件路径
签名信息
```

**启动项**

```sql
启动项名称
启动项可执行文件路径
签名信息
```

镜像劫持信息

```sql
名称
映像劫持路径
状态
```

### 网络信息展示

表设计

```sql
CREATE TABLE process_connections (
    id SERIAL PRIMARY KEY,
    process_name VARCHAR(255) NOT NULL,  -- 进程名
    pid INT NOT NULL,                    -- 进程ID
    protocol VARCHAR(10),                -- 协议 (TCP/UDP)
    local_address INET,                  -- 本地监听地址
    local_port SMALLINT,                 -- 本地监听端口 (0-65535)
    remote_address INET,                 -- 远程地址
    remote_port SMALLINT,                -- 远程端口 (0-65535)
    connection_status VARCHAR(20),       -- 连接状态
    
    -- 端口有效性检查
    CHECK (local_port BETWEEN 0 AND 65535),
    CHECK (remote_port BETWEEN 0 AND 65535),
    
    -- 可选的状态检查（根据实际需要添加）
    CHECK (connection_status IN (
        'LISTEN', 'ESTABLISHED', 'TIME_WAIT', 'CLOSE_WAIT', 
        'SYN_SENT', 'SYN_RECEIVED', 'CLOSED'))
);
```

# 具体实现

## 日志

### 日志信息的获取

重点于

```sql
C:/Windows/System32/winevt/Logs/
C:/WINDOWS/system32/config
```

然后通过解析出不同的事件id 进行匹配

#### 登陆相关事件

| 事件类型     | 事件ID | 描述                          |
| ------------ | ------ | ----------------------------- |
| 登录成功     | 4624   | 账户成功登录                  |
| 登录失败     | 4625   | 账户登录失败（密码错误等）    |
| 特殊登录成功 | 4648   | 使用显式凭据的登录（如runas） |
| 账户锁定     | 4740   | 账户因多次失败登录被锁定      |

#### Rdp

| 事件类型     | 事件ID | 描述                                                     |
| ------------ | ------ | -------------------------------------------------------- |
| 远程登录验证 | 4776   | 专门记录远程桌面服务登录尝试，包含来源计算机名及认证结果 |

#### 服务

| 事件类型     | 事件ID | 描述                                                         |
| ------------ | ------ | ------------------------------------------------------------ |
| 服务创建成功 | 7045   | 记录系统或用户成功创建新服务的事件，用于监控服务部署或第三方软件安装行为。 |
| 服务创建失败 | 7030   | 表示服务创建过程中出现错误，可能由权限不足、配置错误或系统资源问题导致。 |

#### 用户日志

| 事件类型             | 事件ID | 描述                                                         |
| -------------------- | ------ | ------------------------------------------------------------ |
| 用户账户创建         | 4720   | 当新用户账户被成功创建时触发，包含创建者、目标用户及操作时间等详细信息。 |
| 用户加入本地管理员组 | 4732   | 记录用户被添加到启用安全性的本地组（如Administrators组）的操作日志。 |
| 本地组成员变更       | 4728   | 当用户被添加到全局安全组时触发（适用于域环境中的全局组变更，如非本地操作）。 |
| 安全组属性修改       | 4737   | 当本地安全组（如Administrators组）的成员列表或属性被修改时触发。 |

#### 创建进程日志

| 事件ID | 类型名称 | 关键检测点             | 涉及权限       | 说明                         |
| ------ | -------- | ---------------------- | -------------- | ---------------------------- |
| 4688   | 进程创建 | 新进程路径、父进程路径 | 通常为当前用户 | 检查异常进程或未知来源进程。 |

#### PowerShell活动

| 事件ID | 类型名称           | 关键检测点         | 涉及权限         | 说明                         |
| ------ | ------------------ | ------------------ | ---------------- | ---------------------------- |
| 4103   | PowerShell模块日志 | 执行的脚本块内容   | 用户或管理员权限 | 检测混淆代码或敏感操作。     |
| 4104   | PowerShell脚本日志 | 完整脚本路径及参数 | 用户或管理员权限 | 检查远程下载或系统修改行为。 |
|        |                    |                    |                  |                              |

#### LSASS读取痕迹

| 事件ID | 类型名称       | 关键检测点                  | 涉及权限       | 说明                             |
| ------ | -------------- | --------------------------- | -------------- | -------------------------------- |
| 4663   | 对象访问审计   | 访问对象路径（如lsass.exe） | SYSTEM或高权限 | 检测凭证转储工具（如Mimikatz）。 |
| 10     | Sysmon进程访问 | 源进程和目标进程            | 高权限进程     | 监控非系统程序访问LSASS。        |

#### 系统日志

| 事件ID | 类型名称      | 关键检测点          | 涉及权限   | 说明               |
| ------ | ------------- | ------------------- | ---------- | ------------------ |
| 1074   | 系统关机/重启 | 触发关机/重启的用户 | 管理员权限 | 检测异常关机行为。 |

#### 应用日志

| 事件ID | 类型名称        | 关键检测点             | 涉及权限 | 说明                     |
| ------ | --------------- | ---------------------- | -------- | ------------------------ |
| 1000   | 应用程序错误    | 崩溃程序路径及错误模块 | 用户权限 | 检测恶意软件兼容性问题。 |
| 1001   | Windows错误报告 | 错误报告内容及提交数据 | 用户权限 | 分析潜在漏洞利用尝试。   |

#### 安全日志

| 事件ID | 类型名称     | 关键检测点             | 涉及权限       | 说明                          |
| ------ | ------------ | ---------------------- | -------------- | ----------------------------- |
| 4624   | 账户登录成功 | 登录类型、来源IP及账户 | 用户或系统权限 | 检测异常地理位置或时间登录。  |
| 4625   | 账户登录失败 | 失败原因及目标账户     | 无权限         | 检测暴力破解或账户枚举。      |
| 4648   | 显式凭证登录 | 使用凭据的进程及目标   | 高权限账户     | 检查横向移动或Pass-the-Hash。 |

## 进程模块

### 展示信息

可以通过windows api 的方式获取

1. PID
2. 进程名称
3. 父进程PID
4. 父进程名称
5. 创建时间
6. 可执行文件路径
7. MD5
8. 签名信息
9. 导入资源
    1. 导入路径
    2. MD5
    3. 签名

### 进程扫描

通过yara

下面有两个examples

```go
package main

import (
        "github.com/hillu/go-yara/v4"

        "bytes"
        "flag"
        "fmt"
        "log"
        "os"
        "path/filepath"
        "strconv"
        "sync"
)

func printMatches(item string, m []yara.MatchRule, err error) {
        if err != nil {
                log.Printf("%s: error: %s", item, err)
                return
        }
        if len(m) == 0 {
                log.Printf("%s: no matches", item)
                return
        }
        buf := &bytes.Buffer{}
        fmt.Fprintf(buf, "%s: [", item)
        for i, match := range m {
                if i > 0 {
                        fmt.Fprint(buf, ", ")
                }
                fmt.Fprintf(buf, "%s:%s", match.Namespace, match.Rule)
        }
        fmt.Fprint(buf, "]")
        log.Print(buf.String())
}

func main() {
        var (
                rules       rules
                vars        variables
                processScan bool
                pids        []int
                threads     int
        )
        flag.BoolVar(&processScan, "processes", false, "scan processes instead of files")
        flag.Var(&rules, "rule", "add rules in source form: [namespace:]filename")
        flag.Var(&vars, "define", "define variable referenced n ruleset")
        flag.IntVar(&threads, "threads", 1, "use specified number of threads")
        flag.Parse()

        if len(rules) == 0 {
                flag.Usage()
                log.Fatal("no rules specified")
        }

        args := flag.Args()
        if len(args) == 0 {
                flag.Usage()
                log.Fatal("no files or processes specified")
        }

        if processScan {
                for _, arg := range args {
                        if pid, err := strconv.Atoi(arg); err != nil {
                                log.Fatalf("Could not parse %s ad number", arg)
                        } else {
                                pids = append(pids, pid)
                        }
                }
        }

        c, err := yara.NewCompiler()
        if err != nil {
                log.Fatalf("Failed to initialize YARA compiler: %s", err)
        }
        for id, value := range vars {
                if err := c.DefineVariable(id, value); err != nil {
                        log.Fatal("failed to define variable '%s': %s", id, err)
                }
        }
        for _, rule := range rules {
                f, err := os.Open(rule.filename)
                if err != nil {
                        log.Fatalf("Could not open rule file %s: %s", rule.filename, err)
                }
                err = c.AddFile(f, rule.namespace)
                f.Close()
                if err != nil {
                        log.Fatalf("Could not parse rule file %s: %s", rule.filename, err)
                }
        }
        r, err := c.GetRules()
        if err != nil {
                log.Fatalf("Failed to compile rules: %s", err)
        }

        wg := sync.WaitGroup{}
        wg.Add(threads)

        if processScan {
                c := make(chan int, threads)
                for i := 0; i < threads; i++ {
                        s, _ := yara.NewScanner(r)
                        go func(c chan int, tid int) {
                                for pid := range c {
                                        var m yara.MatchRules
                                        log.Printf("<%02d> Scanning process %d...", tid, pid)
                                        err := s.SetCallback(&m).ScanProc(pid)
                                        printMatches(fmt.Sprintf("<pid %d", pid), m, err)
                                }
                                wg.Done()
                        }(c, i)
                }
                for _, pid := range pids {
                        c <- pid
                }
                close(c)
        } else {
                c := make(chan string, threads)
                for i := 0; i < threads; i++ {
                        s, _ := yara.NewScanner(r)
                        go func(c chan string, tid int) {
                                for filename := range c {
                                        var m yara.MatchRules
                                        log.Printf("<%02d> Scanning file %s... ", tid, filename)
                                        err := s.SetCallback(&m).ScanFile(filename)
                                        printMatches(filename, m, err)
                                }
                                wg.Done()
                        }(c, i)
                }
                for _, path := range args {
                        if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
                                if info.Mode().IsRegular() {
                                        c <- path
                                } else if info.Mode().IsDir() {
                                        return nil
                                } else {
                                        log.Printf("Sipping %s", path)
                                }
                                return nil
                        }); err != nil {
                                log.Printf("walk: %s: %s", path, err)
                        }
                }
                close(c)
        }
        wg.Wait()
}
```

还可以是结合windows api 获取进程

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "syscall"
    "unsafe"

    "github.com/hillu/go-yara/v4"
    "golang.org/x/sys/windows"
)

const (
    TH32CS_SNAPPROCESS = 0x00000002
    MAX_PATH           = 260
)

type PROCESSENTRY32 struct {
    Size              uint32
    Usage             uint32
    ProcessID         uint32
    DefaultHeapID     uintptr
    ModuleID          uint32
    Threads           uint32
    ParentProcessID   uint32
    PriClassBase      int32
    Flags             uint32
    ExeFile           [MAX_PATH]uint16
}

func loadYaraRules(ruleDir string) (*yara.Rules, error) {
    c, err := yara.NewCompiler()
    if err != nil {
        return nil, fmt.Errorf("创建编译器失败: %v", err)
    }

    err = filepath.Walk(ruleDir, func(path string, info os.FileInfo, err error) error {
        if err != nil || info.IsDir() {
            return nil
        }

        if filepath.Ext(path) == ".yar" {
            f, err := os.Open(path)
            if err != nil {
                return fmt.Errorf("打开规则文件失败: %v", err)
            }
            defer f.Close()

            if err := c.AddFile(f, ""); err != nil {
                return fmt.Errorf("编译规则失败: %v", err)
            }
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    rules, err := c.GetRules()
    if err != nil {
        return nil, fmt.Errorf("获取规则失败: %v", err)
    }
    return rules, nil
}

func scanProcesses(rules *yara.Rules) error {
    snapshot, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
    if err != nil {
        return fmt.Errorf("创建进程快照失败: %v", err)
    }
    defer windows.CloseHandle(snapshot)

    var entry PROCESSENTRY32
    entry.Size = uint32(unsafe.Sizeof(entry))
    
    if err = windows.Process32First(snapshot, &entry); err != nil {
        return fmt.Errorf("获取首个进程失败: %v", err)
    }

    for {
        exePath := syscall.UTF16ToString(entry.ExeFile[:])
        fmt.Printf("扫描进程 PID:%d 路径:%s\n", entry.ProcessID, exePath)

        if exePath != "" {
            m, _ := yara.NewScanner(rules)
            m.SetCallback(func(r yara.Rule) (bool, error) {
                fmt.Printf("检测到匹配规则: %s\n", r.Identifier())
                return true, nil
            }).ScanFile(exePath)
        }

        if err = windows.Process32Next(snapshot, &entry); err != nil {
            break
        }
    }
    return nil
}

func main() {
    rules, err := loadYaraRules("/rule")
    if err != nil {
        fmt.Printf("加载YARA规则失败: %v\n", err)
        return
    }

    if err := scanProcesses(rules); err != nil {
        fmt.Printf("进程扫描失败: %v\n", err)
    }
}
```

### 内存关键字

https://www.freebuf.com/sectool/408673.html

https://www.52pojie.cn/thread-1484643-1-1.html

可以通过go -> windows api 的方式进行获取其中的信息

通过打开进程句柄然后遍历字符串的形式进行检索

## 文件历史记录（暂定）

### Prefetch记录

可以查看近期高频运行的程序，还可以看到程序运行的时间记录

### UserAssit话动记录

记录程序的启动次数

### Recent File记录

最近使用的文件

## 主机信息（重点）

### 用户信息

有多少个账号

#### 方案一（可能使用这个方案）

通过读取注册表

#### 方案二

通过windows api

GetUserNameA

### 定时任务

通过windows api

https://learn.microsoft.com/zh-cn/windows/win32/taskschd/task-scheduler-interfaces

### 服务信息

winservices

https://blog.csdn.net/raoxiaoya/article/details/121658965

### 启动项信息

1. 读取windows 注册表
2. 读取启动文末人启动文件夹

### 镜像劫持信息

可以通过读取注册表，然后找到debug 键值

检查是否debug 键值对应的文件路径和文件名是否是一致的

如果发现不一致，贼很有可能有问题

但是还有一个情况就是如果恶意程序的名字一样，可能需要比较

##  网络信息展示（重点）

1. 进程名
2. Pid
3. 协议
4. 本地监听地址
5. 本地监听端口
6. 远程地址
7. 远程端口
8. 状态

下面是一个获取pid 对应 网络信息的代码，可以通过修改下面的代码实现一直监听的功能

```go
package main

import (
    "encoding/binary"
    "fmt"
    "golang.org/x/sys/windows"
    "net"
    "syscall"
    "time"
    "unsafe"
)

const (
    TCP_TABLE_OWNER_PID_ALL = 5
    AF_INET                 = 2
)

var (
    iphlpapi                  = windows.NewLazySystemDLL("iphlpapi.dll")
    kernel32                  = windows.NewLazySystemDLL("kernel32.dll")
    getExtendedTcpTable       = iphlpapi.NewProc("GetExtendedTcpTable")
    openProcess               = kernel32.NewProc("OpenProcess")
    queryFullProcessImageName = kernel32.NewProc("QueryFullProcessImageNameW")
)

type ConnectionInfo struct {
    ProcessName string
    PID         uint32
    Protocol    string
    LocalAddr   string
    LocalPort   uint16
    RemoteAddr  string
    RemotePort  uint16
    State       string
}

type MIB_TCPROW_OWNER_PID struct {
    DwState      uint32
    DwLocalAddr  uint32
    DwLocalPort  uint32
    DwRemoteAddr uint32
    DwRemotePort uint32
    DwOwningPid  uint32
}

var tcpStates = map[uint32]string{
    1:  "CLOSED",
    2:  "LISTEN",
    3:  "SYN_SENT",
    4:  "SYN_RCVD",
    5:  "ESTABLISHED",
    6:  "FIN_WAIT1",
    7:  "FIN_WAIT2",
    8:  "CLOSE_WAIT",
    9:  "CLOSING",
    10: "LAST_ACK",
    11: "TIME_WAIT",
    12: "DELETE_TCB",
}

func main() {
    monitorConnections(5 * time.Minute)
}

func monitorConnections(duration time.Duration) {
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    timeout := time.After(duration)
    fmt.Printf("%-20s %-8s %-8s %-21s %-6s %-21s %-6s %-15s\n",
       "Process", "PID", "Proto", "Local Address", "Port",
       "Remote Address", "Port", "State")

    for {
       select {
       case <-ticker.C:
          conns := getCurrentConnections()
          printConnections(conns)
       case <-timeout:
          return
       }
    }
}

func getCurrentConnections() []ConnectionInfo {
    var pTcpTable []byte
    var dwSize uint32

    // 第一次调用获取缓冲区大小
    getExtendedTcpTable.Call(
       uintptr(unsafe.Pointer(&pTcpTable)),
       uintptr(unsafe.Pointer(&dwSize)),
       0,
       syscall.AF_INET,
       TCP_TABLE_OWNER_PID_ALL,
       0,
    )

    buf := make([]byte, dwSize)
    ret, _, _ := getExtendedTcpTable.Call(
       uintptr(unsafe.Pointer(&buf[0])),
       uintptr(unsafe.Pointer(&dwSize)),
       0,
       syscall.AF_INET,
       TCP_TABLE_OWNER_PID_ALL,
       0,
    )

    if ret != 0 {
       return nil
    }

    numEntries := *(*uint32)(unsafe.Pointer(&buf[0]))
    rows := (*[1 << 20]MIB_TCPROW_OWNER_PID)(unsafe.Pointer(&buf[4]))[:numEntries]

    var conns []ConnectionInfo
    for _, row := range rows {
       conn := ConnectionInfo{
          PID:         row.DwOwningPid,
          Protocol:    "TCP",
          LocalAddr:   parseIPv4(row.DwLocalAddr),
          LocalPort:   parsePort(row.DwLocalPort),
          RemoteAddr:  parseIPv4(row.DwRemoteAddr),
          RemotePort:  parsePort(row.DwRemotePort),
          State:       tcpStates[row.DwState],
          ProcessName: getProcessName(row.DwOwningPid),
       }
       conns = append(conns, conn)
    }
    return conns
}

func printConnections(conns []ConnectionInfo) {
    for _, c := range conns {
       fmt.Printf("%-20.20s %-8d %-8s %-21s %-6d %-21s %-6d %-15s\n",
          c.ProcessName,
          c.PID,
          c.Protocol,
          c.LocalAddr,
          c.LocalPort,
          c.RemoteAddr,
          c.RemotePort,
          c.State,
       )
    }
}

func parseIPv4(addr uint32) string {
    ip := make(net.IP, 4)
    binary.LittleEndian.PutUint32(ip, addr)
    return ip.String()
}

func parsePort(port uint32) uint16 {
    return uint16(port>>8 | port<<8)
}

func getProcessName(pid uint32) string {
    const PROCESS_QUERY_LIMITED_INFORMATION = 0x1000

    hProcess, _, _ := openProcess.Call(
       uintptr(PROCESS_QUERY_LIMITED_INFORMATION),
       uintptr(0),
       uintptr(pid),
    )

    if hProcess == 0 {
       return "N/A"
    }
    defer windows.CloseHandle(windows.Handle(hProcess))

    var buf [windows.MAX_PATH]uint16
    var size = uint32(len(buf))

    ret, _, _ := queryFullProcessImageName.Call(
       hProcess,
       uintptr(0),
       uintptr(unsafe.Pointer(&buf[0])),
       uintptr(unsafe.Pointer(&size)),
    )

    if ret == 0 {
       return "N/A"
    }

    return windows.UTF16ToString(buf[:])
}
```

![img](https://poizon.feishu.cn/space/api/box/stream/download/asynccode/?code=NmY1ZDI5OTRmZjFjN2JmYzk2NjAzMzIxOWI2NTUzODVfcnI5RHNPOFAwdFVPUTgzaWpmZ3hkN0FGRWVlQVhHb0ZfVG9rZW46SWJYUWI0NWp0b1J5Sm14Rnltc2NvamRGbkxoXzE3NDMzMzY2NjA6MTc0MzM0MDI2MF9WNA)

## 权限维持检查

1. 注册表中IFEO 目录下的对应的目录是否存在debuger 值 对比这个值是否是一致的，对应的程序是否有签名
2. 自启动检查

- 注册表中信息检查

```Plain
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run      # 当前用户
HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run     # 设置了所有的用户登录都会被执行
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Policies\Explorer\Run
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\Explorer\Run
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce
HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce
HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce
HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnceSetup
HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnceSetup
```

通过对自启动的程序进行yara 静态扫描+签名检查+文件路径检查

- 自启动文件夹

```Plain
C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup
```

1. 服务（现在没具体的想法）

- 原本不是服务exe 查看是否存在包装成服务exe 的过程
- 原本就是服务exe
    - 查看服务的签名

1. 计划任务 找创建计划任务的日志
2. 用户类的 直接查找注册表names 以及对应的uid
3. Dns  解析 通过提供关键的url ,ip  对内存关键字解锁 时间22 日志检测