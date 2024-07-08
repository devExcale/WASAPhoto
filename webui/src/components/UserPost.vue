<script>
import {Post, User} from "@/utils/entities";
import {axiosConf, global} from "@/utils/global";

export default {
	name: "UserPost",
	props: {
		post: Post,
	},
	computed: {
		pictureUrl() {
			return `http://localhost:3000${this.post.pictureUrl}?token=${global.token}`;
		}
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
		}
	},
}
</script>

<template>
	<div class="card p-0 m-0" :key="post.uuid">

		<img :src="pictureUrl" class="card-img-top" alt="Picture here...">
		<div class="card-body">


			<RouterLink :to="`/profile/${post.authorUuid}`" class="link-primary">
				{{ post.authorUuid }}
			</RouterLink>
			<span class="card-title"></span>
			<span class="card-subtitle mb-2 text-body-secondary">{{ post.timestamp }}</span>
			<p class="card-text">{{ post.caption }}</p>

			<nav class="nav nav-pills">
				<!-- TODO: change material symbols to feather sprite -->
				<button class="nav-link material-symbols-rounded" @click.prevent="">favorite</button>
				<span class="">{{ post.numLikes }}</span>
				<button class="nav-link material-symbols-rounded" @click.prevent="">mode_comment</button>
				<span class="">{{ post.numComments }}</span>
			</nav>

		</div>

	</div>
</template>

<style scoped>

</style>
