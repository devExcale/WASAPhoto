<script>
import {Post, User} from "@/utils/entities";
import {axiosConf, global} from "@/utils/global";

export default {
	name: "UserPost",
	data() {
		return {
			user: {},
		}
	},
	props: {
		post: Post,
		onDelete: Function,
	},
	computed: {
		pictureUrl() {
			return `http://localhost:3000${this.post.pictureUrl}?token=${global.token}`;
		},
		isAuthor() {
			return this.post.authorUuid === global.userUUID;
		},
	},
	methods: {
		likePost() {
			console.log('likePost');
		},
		commentPost() {
			console.log('commentPost');
		},
		async setUsername() {

			try {

				let response = await this.$axios.get(`/users/${this.userUUID}`, axiosConf.value);
				this.user = User.fromResponse(response.data);

			} catch (e) {

				console.log(e)

			}
		},
		async deletePost() {

			try {

				await this.$axios.delete(`/me/feed/${this.post.uuid}`, axiosConf.value);

				console.log(`Post ${this.post.uuid} deleted.`)

				this.onDelete(this.post);

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not delete. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not delete. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not delete. Please try again later. (?)');
				}

			}

		},
	},
	mounted() {
		this.setUsername();
	}
}
</script>

<template>
	<div class="card p-0 m-0" :key="post.uuid">

		<img :src="pictureUrl" class="card-img-top" alt="Picture here...">
		<div class="card-body">


			<RouterLink :to="`/profile/${post.authorUuid}`" class="link-primary">
				{{ user.displayName || post.authorUuid }}
			</RouterLink>
			<span class="card-title"></span>
			<span class="card-subtitle mb-2 text-body-secondary">{{ post.timestamp }}</span>
			<p class="card-text">{{ post.caption }}</p>

			<div class="btn-group">
				<button class="btn btn-primary material-symbols-rounded" @click="likePost">favorite</button>
				<button class="btn btn-primary disabled">{{ post.numLikes }}</button>
				<button class="btn btn-primary material-symbols-rounded" @click="commentPost">mode_comment</button>
				<button class="btn btn-primary disabled">{{ post.numComments }}</button>
				<button class="btn btn-danger" @click="deletePost" v-if="isAuthor">delete</button>
			</div>


		</div>

	</div>
</template>

<style scoped>

</style>
