# try_hack_cds

Admin 账户信息

```
用户名：admin
密码：c10ud5d1e
```

需要被注入 Flag 的文件是 `./src/db/db.sqlite`（在 `docker-entrypoint.sh` 中看来是 `/app/db/db.sqlite`），用真正的 Flag 替换 `{{inject_me_with_flag}}` 即可。

向选手提供题目附件只需要将 `./src` 打包提供即可，不必担心泄露 `db.sqlite`，反正看到的 Flag 也是假的，密码是 BCrypt，弄出来也是神仙。