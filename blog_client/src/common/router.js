import { createRouter, createWebHashHistory } from "vue-router";

let routes = [
]

const router = createRouter({
    // createWebHistory路由模式路径不带#号
    // createWebHashHistory路由模式路径带#号
    history: createWebHashHistory(),
    routes,
});

export { router, routes }
