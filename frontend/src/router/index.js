import {createRouter, createWebHistory} from "vue-router";
import PostIdPage from "@/views/PostIdPage";
import MainPage from "@/views/MainPage";
//import PostPageCompositionApi from "@/pages/PostPageCompositionApi";
import ListPostsPage from "@/views/ListPostsPage.vue";

import AdminDishes from "@/views/AdminDishes.vue";
import AdminProducts from "@/views/AdminProducts.vue";
import AdminOrders from "@/views/AdminOrders.vue";
import AdminTags from "@/views/AdminTags.vue";

import MyShoppingCart from "@/views/MyShoppingCart.vue";
import MyOrders from "@/views/MyOrders.vue";

import Home from '@/views/Home.vue';
import OAuthCallback from '@/views/OAuthCallback.vue';


const routes = [
    {
        path: '/mainpage',
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
    {
        path: '/posts',
        component: ListPostsPage
    },
    {
        path: '/',
        name: 'Home',
        component: Home
    },
  {
        path: '/oauth2callback',
        name: 'OAuthCallback',
        component: OAuthCallback
    }
]

const router = createRouter({
    routes,
    history: createWebHistory(process.env.BASE_URL)
})

export default router;