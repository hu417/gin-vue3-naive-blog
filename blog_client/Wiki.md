

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
用到的组件为表单Form表单: https://www.naiveui.com/zh-CN/os-theme/components/form

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


## 顶栏组件
顶栏组件用到:
- 头像组件头像[ Avatar - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/avatar)
- 按钮组件按钮[ Button - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/button)
顶栏组件，可以跳转至主页、个人信息页、登录页以及发布文章页
### 添加路由
```bash
$ touch src/views/{MainFrame.vue,Myself.vue,Others.vue,Publish.vue,Update.vue,Detail.vue}

// src/common/router.js

```
### 顶栏按钮
```bash
// 编写顶栏，顶栏渲染时向后端接口/user获取头像，若用户已登录，将成功获取用户头像，否则可跳转至登录页
$ touch src/components/TopBar.vue

```
### 修改mian.js
```bash
// 修改main.js，添加拦截器传token，即每个页面都向后端传token，无论后端需不需要

```

## 个人信息页
个人信息页用到的组件有:
- 卡片卡片[ Card - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/card)
- 模态框模态框[ Modal - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/modal)
- 图标图标[ Icon - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/icon)
> 用户点击头像可进入用户信息页，登录用户查看自身与他人的信息页渲染有所不同，自身的个人信息页有修改信息按键，而他人的个人信息页有关注按键
### 修改Myself.vue
```bash
// src/views/Myself.vue
```
### 修改Others.vue
```bash
// src/views/Others.vue
```
### 测试访问
```bash
http://localhost:5173/#/myself
http://localhost:5173/#/others
```

## 主页
用户在主页的输入框文本输入[ Input - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/input)输入关键词，通过选择器弹出选择[ Popselect - Naive UI](https://www.naiveui.com/zh-CN/os-theme/components/popselect)选择分类，通过分页器分页[ Pagination - Naive UI分页](https://www.naiveui.com/zh-CN/os-theme/components/pagination)

### 修改MainFrame.vue
```bash
// src/views/MainFrame.vue
```
### 测试访问
// http://localhost:5173/#/

## 富文本编辑组件
富文本: https://www.wangeditor.com/v5/for-frame.html
```bash
$ touch src/components/RichTextEditor.vue

```
## 文章发布修改页
文章发布与修改页类似，不同的是修改页要先获取原文章数据再将其渲染
### 修改Publish.vue
```bash
// src/views/Publish.vue
```
### 修改Updata.vue
```bash
// src/views/Update.vue
```
