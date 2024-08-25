<script>
import {Post} from "@/utils/entities";
import {axiosConf, loadPicture, global} from "@/utils/global";

export default {
	name: "UserPost",
	data() {
		return {
			pictureSrc: global.loadingGifSrc,
		}
	},
	props: {
		post: Post,
		onDelete: Function,
	},
	inject: ['onChangeProfilePicture'],
	computed: {

		isAuthor() {
			return this.post.authorUuid === global.userUUID;
		},

		likeBtnStyle() {
			return this.post.loggedUserLiked ? 'btn-primary' : 'btn-outline-primary';
		},

	},
	methods: {

		async reloadPicture() {

			this.pictureSrc = global.loadingGifSrc;

			if (this.post.pictureUrl)
				this.pictureSrc = await loadPicture(this.post.pictureUrl);
		},

		async likePostAction() {

			if (this.loading)
				return;

			this.loading = true;

			if (this.post.loggedUserLiked)
				await this.unlikePost();
			else
				await this.likePost();

			this.loading = false;

		},

		async likePost() {

			try {

				let url = `/users/${this.post.authorUuid}/feed/${this.post.uuid}/likes/`;
				await this.$axios.put(url, {}, axiosConf.value);
				this.post.loggedUserLiked = true;
				this.post.numLikes++;

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not like post. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not like post. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not like post. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not like post. Please try again later. (?)');
				}

			}

		},

		async unlikePost() {

			try {

				let url = `/users/${this.post.authorUuid}/feed/${this.post.uuid}/likes/`;
				await this.$axios.delete(url, axiosConf.value);
				this.post.loggedUserLiked = false;
				this.post.numLikes--;

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not remove like. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 403) {
					alert('Could not remove like. You are restricted by the user.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not remove like. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not remove like. Please try again later. (?)');
				}

			}

		},

		commentPost() {
			console.log('commentPost');
		},

		async deletePost() {

			try {

				await this.$axios.delete(`/me/feed/${this.post.uuid}`, axiosConf.value);

				console.log(`Post ${this.post.uuid} deleted.`)

				if (this.onDelete)
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

		async setMainPost() {

			try {

				let data = {"post_uuid": this.post.uuid};
				await this.$axios.patch(`/me/profile_picture`, data, axiosConf.value);

				console.log(`Post ${this.post.uuid} set as main.`)

			} catch (e) {

				console.log(e)

				if (e.response && e.response.status === 400) {
					alert('Could not set profile picture. Please try again later.');
				} else if (e.response && e.response.status === 401) {
					alert('You are not logged in. Try refreshing the page.');
				} else if (e.response && e.response.status === 404) {
					alert('Could not set profile picture. The post doesn\'t exist.');
				} else if (e.response && e.response.status === 500) {
					alert('Internal Server Error')
				} else {
					alert('Could not set profile picture. Please try again later. (?)');
				}

			}

			if (this.onChangeProfilePicture)
				await this.onChangeProfilePicture(this.post.pictureUrl);

		},
	},
	mounted() {
		this.reloadPicture();
	}
}
</script>

<template>
	<div class="card p-0 m-0" :key="post.uuid">

		<img :src="pictureSrc" class="card-img-top" alt="Picture here...">
		<div class="card-body">


			<RouterLink :to="`/profile/${post.authorUuid}`" class="link-primary">
				{{ post.authorDisplayName || post.authorUsername || post.authorUuid }}
			</RouterLink>
			<span class="card-title"></span>
			<span class="card-subtitle mb-2 text-body-secondary">{{ post.timestamp }}</span>
			<p class="card-text">{{ post.caption }}</p>

			<div class="btn-group">
				<button :class="likeBtnStyle" class="btn material-symbols-rounded" @click="likePostAction">favorite</button>
				<button class="btn btn-primary disabled">{{ post.numLikes }}</button>
			</div>
			<div class="btn-group m-2">
				<button class="btn btn-primary material-symbols-rounded" @click="commentPost">mode_comment</button>
				<button class="btn btn-primary disabled">{{ post.numComments }}</button>
			</div>
			<div class="btn-group" v-if="isAuthor">
				<button class="btn btn-success" @click="setMainPost">Set as Main</button>
				<button class="btn btn-danger" @click="deletePost">Delete</button>
			</div>


		</div>

	</div>
</template>

<style scoped>

</style>
