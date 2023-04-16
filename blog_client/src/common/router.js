import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    // 登录路由
    { path: "/login", component: () => import("../views/Login.vue") }, 
    // 添加注册
    { path: "/register", component: () => import("../views/Register.vue") },
    // 主页
    { path: "/", component: () => import("../views/MainFrame.vue") },
    { path: "/publish", component: () => import("../views/Publish.vue") },
    { path:"/myself", component: () => import("../views/Myself.vue") },
    { path:"/others", component: () => import("../views/Others.vue") },
    { path:"/detail", component: () => import("../views/Detail.vue") },
    { path:"/update", component: () => import("../views/Update.vue") },


]

const router = createRouter({
    // createWebHistory路由模式路径不带#号
    // createWebHashHistory路由模式路径带#号
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
