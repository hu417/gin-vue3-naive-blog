import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    // 登录路由
    { path: "/login", component: () => import("../views/Login.vue") }, 
    // 添加注册
    { path: "/register", component: () => import("../views/Register.vue") },
    // 主页
    { path: "/", component: () => import("../views/MainFrame.vue") },
    // 文章上传
    { path: "/publish", component: () => import("../views/Publish.vue") },
    // 个人信息
    { path:"/myself", component: () => import("../views/Myself.vue") },
    { path:"/others", component: () => import("../views/Others.vue") },
    // 文章详情
    { path:"/detail", component: () => import("../views/Detail.vue") },
    // 文章修改
    { path:"/update", component: () => import("../views/Update.vue") },


]

const router = createRouter({
    // createWebHistory路由模式路径不带#号
    // createWebHashHistory路由模式路径带#号
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
