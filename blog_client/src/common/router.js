import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
    // 添加注册/登录路由
    { path: "/login", component: () => import("../views/Login.vue") },
    { path: "/register", component: () => import("../views/Register.vue") },

]

const router = createRouter({
    // createWebHistory路由模式路径不带#号
    // createWebHashHistory路由模式路径带#号
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
