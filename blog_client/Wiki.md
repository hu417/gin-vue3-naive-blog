

# 简易博客社区前端搭建
## 项目准备
```bash
node        v18.12.1  `npm config set registry="https://registry.npm.taobao.org/"`
npm         8.19.2
vue3        @vue/cli 4.5.12  `npm install vue@3.2.47 -g`
vue/cli     @4.5.12          `npm install @vue/cli@4.5.12 -g`
vscode      插件：Auto Close Tag、Auto Rename Tag、Live Server
```
### 创建项目
```bash
$ npm init vite@latest blog_client
√ Select a framework: » Vue
√ Select a variant: » JavaScript

Scaffolding project in D:\study\project\golang\3-gin\go-vue-blogCommunity\blog_client\blog_client...

Done. Now run:

  cd blog_client
  npm install     // 安装依赖
  npm run dev     // 运行项目

// nodejs版本高,则需要: NODE_OPTIONS="--openssl-legacy-provider"
```

### 模块安装
```bash
// axios: 基于promise的HTTP库，用于http请求
npm install axios

// pinia: Vue的存储库，它允许您跨组件、页面共享状态
npm install pinia

// sass: CSS的开发工具，提供许多便利写法
npm install sass

// vue-router: Vue.js官方的路由插件
npm install vue-router@4

// naive-ui: Vue3的组件库,HTML+CSS样式,类似element-plus,antd
npm i -D naive-ui
npm i -D vfonts
npm i -D @vicons/ionicons5

// wangeditor: 富文本编辑器
npm install @wangeditor/editor-for-vue@next --save

```
### 修改全局格式文件
```bash
// src/style.css
# 全部注释,添加此段
body {
  background-color: #FCFAF7;
  margin: 0;
  padding: 0;
}

```
### 新建文件夹存放图片
```bash
// 该文件夹存放一些显示在页面上的图片
$ mkdir -p src/assets/image

```
### 引入基本的模块
> main.js
```bash
// 引入ui框架
// 引入pinia
// 引入路由
// 引入axios

```
> stores
// 定义存储内容
```bash
$ mkdir -p src/stores
$ touch src/stores/UserStore.js   // 存储用户token
```
> roter
```bash
// common
// views: 页面组件目录(所有页面放在这里)
$ mkdir -p src/{common,views} 
$ touch src/common/router.js  // 路由配置入口
```
> App.vue
```bash
<template>
  <router-view ></router-view>
</template>

<script setup>

</script>

<style scoped>

</style>

```

## 登录注册面

Naive UI: https://www.naiveui.com/zh-CN/os-theme/components/button

### 注册页面
```bash
// 编写注册页，获取用户的用户名、手机号和密码并传给后端
$ touch src/views/Register.vue 

```
### 保存token
```bash
$ mkdir -p src/stores
$ touch src/stores/UserStore.js   // token保存起来，以便传给其他接口

```
### 登录页面
```bash
// 编写登录页，获取用户的手机号和密码传给后端，登录成功后存储后端传过来的token
$ touch src/views/Login.vue

```
### 图片准备
```bash
// src/assets/image
```
### 修改路由
```bash
// common/router.js
```
### 测试访问
- [x] 运行server端
- [x] 运行client端