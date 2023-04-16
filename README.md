



## git 操作
git clone https://github.com/hu417/gin-vue3-naive-blog.git
git init
git config --global user.name "***"
git config --global user.email ****@qq.com

// ssl认证关闭
git config --global http.sslVerify "false"
git config --global credential.helper manager

// 提交项目
git add .
git commit -m "fix: blog项目-前端开发
1、注册/登录页面开发
" 
git tag -a v0.2 -m "版本v0.2"
git push -u origin main --tags


