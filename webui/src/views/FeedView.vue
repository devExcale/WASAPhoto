<script>
import UserPost from "@/components/UserPost.vue";
import {Post} from "@/utils/entities";
import {axiosConf} from "@/utils/global";

export default {
	components: {UserPost},
	data: function () {
		return {
			errormsg: null,
			loading: true,
			some_data: null,
			posts: [],
		}
	},
	props: {
		userUUID: String,
	},
	methods: {

		async loadFeed() {

			this.loading = true;

			try {

				let url = (this.userUUID) ? `/users/${this.userUUID}/feed/` : `/me/feed/`;
				let response = await this.$axios.get(url, axiosConf.value);
				this.posts = response.data.map((post) => Post.fromResponse(post));

				console.log(response)
				console.log(this.posts)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not get feed. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error');
				} else {
					alert('Could not get feed. Please try again later.');
				}

				return;
			}

			this.loading = false;
		},

		onPostDelete(post) {
			this.posts = this.posts.filter((p) => p.uuid !== post.uuid);
		},

	},
	mounted() {
		this.loadFeed()
	}
}
</script>

<template>
	<div>
		<div class="container-fluid" v-if="loading">
			<p>Loading...</p>
		</div>
		<div v-else
			 class="container-fluid">
			<div v-for="post in posts" :key="post.uuid" class="row m-4 mb-3">
				<UserPost :post="post" :on-delete="onPostDelete"/>
			</div>
			<div v-if="posts.length === 0" class="row m-4 mb-3">
				<p class="h3 text-center">No posts here :/</p>
			</div>
		</div>
	</div>
</template>

<style>
</style>
