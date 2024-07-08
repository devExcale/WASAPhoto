import {createRouter, createWebHashHistory} from 'vue-router'
import FeedView from '@/views/FeedView.vue'
import LoginView from '@/views/LoginView.vue'
import ProfileView from '@/views/ProfileView.vue';
import NewPostView from '@/views/NewPostView.vue';
import SearchView from "@/views/SearchView.vue";
import FollowedView from "@/views/FollowedView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/', redirect: '/home'},
		{path: '/home', component: FeedView},
		{path: '/search', component: SearchView},
		{path: '/profile/:user_uuid', component: ProfileView},
		{path: '/newPost', component: NewPostView},
		{path: '/follows', component: FollowedView},
	]
})

export default router
