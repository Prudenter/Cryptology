# openssl生成自签名证书步骤
＃＃ 生成的私钥可以没有密码（直接方式，不是 后去除的）

   > openssl genrsa -out server.key 1024 //不加密
   >
   > 或
   >
   > openssl genrsa -des3 -out server.key 1024  //使用des3加密

2. 使用私钥可以生成csr:

   > openssl req -new -key server.key -out server.csr,  //会引导输入信息

3. 查看csr详情:

   > openssl req -in server.csr -text，// 注意没有x509, 而是有req

4. 生成自签名证书(用自己的key给自己的csr签名）

   > openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

至此，生成了一个自签名的证书

5. 查看自签名证书

   > openssl x509 -in server.crt -text
