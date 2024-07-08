import {createRouter, createWebHashHistory} from 'vue-router'
import FeedView from '@/views/FeedView.vue'
import LoginView from '@/views/LoginView.vue'
import ProfileView from '@/views/ProfileView.vue';

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/', component: FeedView},
		{path: '/home', component: FeedView},
		{path: '/profile/:user_uuid', component: ProfileView},
		{path: '/link2', component: FeedView},
		{path: '/some/:id/link', component: FeedView},
	]
})

export default router
