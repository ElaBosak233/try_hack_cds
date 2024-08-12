# try_hack_cds

本题目是以 [cloudsdale](https://github.com/elabosak233/cloudsdale) 的某个有漏洞的版本为基础修改的。与题目最接近的一次 Commit 是 [1100b98](https://github.com/ElaBosak233/cloudsdale/tree/1100b9824acdef88695ab0da57df69661c63cc20)，当然，你现在看到的还在迭代中的 cloudsdale 已经把这个问题修复了。

为了方便出题，我屏蔽了所有与容器控制相关的内容，所以不需要向题目镜像提供 Docker 守护进程的套接字。

Flag 具体的注入方法写在了 `./service/docker-entrypoint.sh` 里面，从环境变量中提取 Flag 后使用 SQL 语句进行注入。

已存放在 `db.sqlite` 中的 admin 账户信息，若有需要，注意修改（修改方式也是在 `docker-entrypoint.sh` 中用 SQL 修改）：用户名：`admin`，密码：`c10ud5d1e`。

向选手提供题目附件只需要将本仓库中的 `./src` 打包提供即可，不必担心泄露 `db.sqlite`，反正看到的 Flag 也是假的，密码用的 BCrypt，弄出来也是神仙。

如果你想先自己试一试，按照下面的步骤走：

```bash
git clone https://github.com/ElaBosak233/try_hack_cds.git
```

```bash
cd try_hack_cds
```

```bash
docker compose -f ./docker/docker-compose.yml up
```

访问 8888 端口即可，默认注入的 Flag 是 `flag{a63b4d37-7681-4850-b6a7-0d7109febb19}`。