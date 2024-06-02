import {createRouter, createWebHistory} from "vue-router";
import PostIdPage from "@/views/PostIdPage";
import MainPage from "@/views/MainPage";
//import PostPageCompositionApi from "@/pages/PostPageCompositionApi";

import AdminDishes from "@/views/AdminDishes.vue";
import AdminProducts from "@/views/AdminProducts.vue";
import AdminOrders from "@/views/AdminOrders.vue";
import AdminTags from "@/views/AdminTags.vue";

import MyShoppingCart from "@/views/MyShoppingCart.vue";
import MyOrders from "@/views/MyOrders.vue";


const routes = [
    {
        path: '/',
        name: 'mainpage',
        component: MainPage
    },  
    {
        path: '/dishdetail/:id',
        name: 'dishdetail',
        component: PostIdPage
    },
    {
        path: '/admin/orders',
        name: 'adminorders',
        component: AdminOrders
    },
    {
        path: '/admin/dishes',
        name: 'admindishes',
        component: AdminDishes
    },
    {
        path: '/admin/products',
        name: 'adminproducts',
        component: AdminProducts
    },
    {
        path: '/admin/tags',
        name: 'admintags',
        component: AdminTags
    },
    {
        path: '/mycart',
        component: MyShoppingCart
    },
    {
        path: '/myorders',
        component: MyOrders
    },


]

const router = createRouter({
    routes,
    history: createWebHistory(process.env.BASE_URL)
})

export default router;